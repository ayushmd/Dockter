package master

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/builderrpc"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/workerrpc"
	"github.com/google/uuid"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Backend struct {
	URL             url.URL                       `json:"url"`
	State           string                        `json:"state"`
	IsAlive         bool                          `json:"isalive"`
	CurrentConnect  int                           `json:"connections"`
	Stats           internal.ContainerBasedMetric `json:"stats"`
	Rtt             time.Duration                 `json:"rtt"`
	NumofContainers int                           `json:"numofContainers"`
	mux             sync.RWMutex
	//Algo part
}

const TYPE_WEBSERVICE = "Web Service"
const TYPE_STATIC = "Static"

type Task struct {
	Subdomain   string //a unique id
	URL         url.URL
	Hostport    string
	Runningport string
	ImageName   string
	ContainerID string
	Status      string
	Type        string
}

type TaskRawRequest struct {
	Name        string            `json:"name"`
	Gitlink     string            `json:"gitLink"`
	Branch      string            `json:"branch"`
	BuildCmd    string            `json:"buildCmd"`
	StartCmd    string            `json:"startCmd"`
	RuntimeEnv  string            `json:"runtimeEnv"`
	RunningPort string            `json:"runningPort"`
	EnvVars     map[string]string `json:"envVars"`
} // BUILD DEPLOY

type TaskFileRequest struct {
	Name    string            `json:"name"`
	Gitlink string            `json:"gitLink"`
	Branch  string            `json:"branch"`
	EnvVars map[string]string `json:"envVars"`
} // BUILD DEPLOY

type TaskImageRequest struct {
	Name        string `json:"name"`
	DockerImage string `json:"dockerImage"`
	RunningPort string `json:"runningPort"`
	BasedMetric internal.ContainerBasedMetric
} // DEPLOY

func (b *Backend) AddConn() {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.CurrentConnect += 1
}

func (b *Backend) ResConn() {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.CurrentConnect -= 1
}

// type TaskConfig struct {
// 	Name string `json:"name"`

// 	Gitlink       string `json:"gitLink"`
// 	Branch        string `json:"branch"`
// 	HasDockerFile bool   `json:"hasDockerFile"`

// 	BuildCmd   string `json:"buildCmd"`
// 	StartCmd   string `json:"startCmd"`
// 	RuntimeEnv string `json:"runtimeEnv"`

// 	EnvVars map[string]string `json:"envVars"`

// 	HasDockerImage bool   `json:"hasDockerImage"`
// 	DockerImage    string `json:"dockerImage"`
// 	RunningPort    string `json:"runningPort"`
// }

type status int

const (
	RUNNING status = iota
	STARTED
	EXITED
)

type Master struct {
	kwriter    *internal.KafkaWriter
	ServerPool []*Backend
	dnsStatus  status
	dbDns      *sql.DB
	cacheDns   *lru.Cache[string, Task]
}

var Master_ *Master = &Master{}

var kacp = keepalive.ClientParameters{
	Timeout:             2 * time.Second, // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,            // send pings even without active streams
}

func (m *Master) AddDnsRecord(task Task) error {
	query := fmt.Sprintf("INSERT INTO dns (Subdomain, HostIp, HostPort, RunningPort, ImageName, ContainerID, Status) VALUES ('%s','%s','%s','%s','%s','%s','%s');",
		task.Subdomain, task.URL.Host, task.Hostport, task.Runningport, task.ImageName, task.ContainerID, task.Status)
	_, err := m.dbDns.Exec(query)
	return err
}

func (m *Master) AddBuilding(sudomain, runningPort string) error {
	query := fmt.Sprintf("INSERT INTO dns (Subdomain, RunningPort, Status, Type) VALUES ('%s','%s','%s','%s');", sudomain, runningPort, "Building", TYPE_WEBSERVICE)
	_, err := m.dbDns.Exec(query)
	return err
}

func (m *Master) UpdateDnsStatus(domain, status string) error {
	query := fmt.Sprintf("UPDATE dns SET Status = '%s' WHERE Subdomain = '%s';", status, domain)
	_, err := m.dbDns.Exec(query)
	return err
}

func (m *Master) DeployDnsRecord(domain, runningPort string) error {
	row := m.PingRecord(domain)
	var subdomain string
	err := row.Scan(&subdomain)
	if err == sql.ErrNoRows {
		query := fmt.Sprintf("INSERT INTO dns (Subdomain, RunningPort, Status, Type) VALUES ('%s','%s','%s','%s');", domain, runningPort, "Deploying", TYPE_WEBSERVICE)
		_, err := m.dbDns.Exec(query)
		return err
	} else {
		query := fmt.Sprintf("UPDATE dns SET Status = '%s' WHERE Subdomain = '%s';", "Deploying", domain)
		_, err := m.dbDns.Exec(query)
		return err
	}
}

func (m *Master) UpdateDeployRecord(task Task) error {
	query := fmt.Sprintf("UPDATE dns SET HostIp = '%s', HostPort = '%s', ImageName = '%s', ContainerID = '%s', Status = 'Deployed' WHERE Subdomain = '%s';",
		task.URL.Host, task.Hostport, task.ImageName, task.ContainerID, task.Subdomain)
	_, err := m.dbDns.Exec(query)
	return err
}

func (m *Master) GetDnsRecord(id string) *sql.Row {
	query := fmt.Sprintf("SELECT * FROM dns WHERE Subdomain='%s'", id)
	row := m.dbDns.QueryRow(query)
	return row
}

func (m *Master) PingRecord(id string) *sql.Row {
	query := fmt.Sprintf("SELECT Subdomain FROM dns WHERE Subdomain='%s'", id)
	row := m.dbDns.QueryRow(query)
	return row
}

func (m *Master) GetDnsStatus(id string) *sql.Row {
	query := fmt.Sprintf("SELECT Status FROM dns WHERE Subdomain='%s'", id)
	row := m.dbDns.QueryRow(query)
	return row
}

func (m *Master) DeleteDnsRecord(id string) error {
	_, err := m.dbDns.Exec("DELETE FROM dns WHERE Subdomain='?'", id)
	return err
}

func (m *Master) DeployStaticRecord(domain, id string) error {
	query := fmt.Sprintf("INSERT INTO dns (Subdomain, ContainerID, Status, Type) VALUES ('%s','%s','%s','%s');", domain, id, "Deployed", TYPE_STATIC)
	_, err := m.dbDns.Exec(query)
	return err
}

type NoRecordError struct{}
type TerminationError struct{}
type NoMetricsError struct{}

func (m *NoRecordError) Error() string {
	return "No record found"
}
func (m *TerminationError) Error() string {
	return "Task not terminated"
}
func (m *NoMetricsError) Error() string {
	return "Metrics not found"
}

func (m *Master) GetRecord(name string) (*Task, error) {
	task := &Task{}
	ctask, ok := m.cacheDns.Get(name)
	if ok {
		task = &ctask
	} else if m.cacheDns != nil {
		row := Master_.GetDnsRecord(name)
		var HostIp string
		err := row.Scan(&task.Subdomain, &HostIp, &task.Hostport, &task.Runningport, &task.ImageName, &task.ContainerID, &task.Status, &task.Type)
		if err != nil {
			return nil, err
		}
		task.URL = url.URL{
			Host: HostIp,
		}
	} else {
		return nil, &NoRecordError{}
	}
	return task, nil
}

func (m *Master) RemoveRecord(name string) {
	m.cacheDns.Remove(name)
	m.DeleteDnsRecord(name)
}

func (m *Master) RetryDed(serv *Backend) {
	for retry_count := 0; retry_count < 3; retry_count++ {
		if m.PoolServ(serv) != nil {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (m *Master) PoolServ(serv *Backend) error {

	conn, err := grpc.Dial(
		serv.URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	)
	if err != nil {
		// log.Fatalf("did not connect: %v", err)
		return err
	}
	defer conn.Close()

	if serv.State == "WORKER" {
		w := workerrpc.NewWorkerServiceClient(conn)
		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()
		res, err := w.HealthMetrics(ctx, &emptypb.Empty{})
		end := time.Now()
		if err != nil {
			fmt.Printf("Server %s didnt respond\n", serv.URL.Host)
			serv.IsAlive = false
			return err
		}
		// metricData, err := json.Marshal(res)
		//log.Println("Pooled server ", serv.URL.Host, string(metricData))
		serv.Stats.CpuPercent = res.CpuPercent
		serv.Stats.MemUsage = res.MemUsage
		serv.Stats.TotalMem = res.TotalMem
		serv.Stats.MemUsedPercent = float64(res.MemUsedPercent)
		serv.Stats.DiskUsage = res.DiskUsage
		serv.Stats.TotalDisk = res.TotalDisk
		serv.Stats.DiskUsagePercent = float64(res.DiskUsagePercent)
		// serv.CpuUsage = float64(res.CpuUsage)
		// serv.MemUsage = float64(res.MemUsage)
		// serv.DiskUsage = float64(res.DiskUsage)
		serv.Rtt = end.Sub(start)
		serv.IsAlive = true
	} else {
		b := builderrpc.NewBuilderServiceClient(conn)
		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()
		res, err := b.BuildHealthMetrics(ctx, &emptypb.Empty{})
		end := time.Now()
		if err != nil {
			fmt.Printf("Server %s didnt respond\n", serv.URL.Host)
			serv.IsAlive = false
			return err
		}
		// metricData, err := json.Marshal(res)
		//log.Println("Pooled server ", serv.URL.Host, string(metricData))
		serv.Stats.CpuPercent = res.CpuPercent
		serv.Stats.MemUsage = res.MemUsage
		serv.Stats.TotalMem = res.TotalMem
		serv.Stats.MemUsedPercent = float64(res.MemUsedPercent)
		serv.Stats.DiskUsage = res.DiskUsage
		serv.Stats.TotalDisk = res.TotalDisk
		serv.Stats.DiskUsagePercent = float64(res.DiskUsagePercent)
		serv.Rtt = end.Sub(start)
		serv.IsAlive = true
	}
	return nil
}

func (m *Master) Pool() {
	//log.Println("Pooling started")
	var waitgrp *sync.WaitGroup = &sync.WaitGroup{}
	for _, serv := range m.ServerPool {
		if serv.IsAlive {
			waitgrp.Add(1)
			go func(serv *Backend) {
				err := m.PoolServ(serv)
				if err != nil {
					go m.RetryDed(serv)
				}
				waitgrp.Done()
			}(serv)
		}
	}
	waitgrp.Wait()
}

func (m *Master) HasJoined(peerurl string) int {
	for i, serv := range m.ServerPool {
		if serv.URL.Host == peerurl {
			return i
		}
	}
	return -1
}

func Handshake(peerurl string) {
	for retries := 0; retries < 3; retries++ {
		if Master_.dnsStatus != STARTED {
			err := HandshakePolicy(CLEAR_INACTIVE, peerurl)
			if err != nil {
				fmt.Print(err)
			}
			break
		} else if Master_.dnsStatus == RUNNING {
			time.Sleep(2 * time.Second)
		}
	}
}

func (m *Master) Join(peerurl, peerState string, HealthStats internal.ContainerBasedMetric) {
	if i := m.HasJoined(peerurl); i == -1 {
		log.Printf("%s(%s) joined\n", peerState, peerurl)
		urlparsed := url.URL{
			Host: peerurl,
		}
		m.ServerPool = append(m.ServerPool, &Backend{
			URL:             urlparsed,
			State:           peerState,
			CurrentConnect:  0,
			IsAlive:         true,
			Stats:           HealthStats,
			NumofContainers: 0,
		})
	} else {
		m.ServerPool[i].IsAlive = true
	}
	go Handshake(peerurl)
}

func (m *Master) GetServerPoolHandler() ([]byte, error) {
	marshalData, err := json.Marshal(m.ServerPool)
	if err != nil {
		return []byte(""), err
	}
	return marshalData, nil
}

func (m *Master) AddTask(request string, marshTask []byte) {
	//log.Println("Add Task:", request, string(marshTask))
	m.kwriter.Write(
		[]byte(request),
		marshTask,
	)
}

func (m *Master) BuildRaw(message kafka.Message) {
	var configs TaskRawRequest
	err := json.Unmarshal(message.Value, &configs)
	if err != nil {
		panic("Couldnt unmarshal")
	}
	if len(m.ServerPool) == 0 {
		return
	}
	configs.Name = strings.ToLower(configs.Name)
	go func() {
		err := m.AddBuilding(configs.Name, configs.RunningPort)
		if err != nil {
			fmt.Println("Sqlite error ", err)
		}
	}()
	backend := MasterPlanAlgo(m.ServerPool, "BUILDER")
	conn, err := grpc.Dial(
		backend.URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic("Connection failed")
	}
	builder := builderrpc.NewBuilderServiceClient(conn)
	backend.AddConn()
	buildRawResponse, err := builder.BuildRaw(
		context.Background(),
		&builderrpc.BuildRawRequest{
			Name:        configs.Name,
			Gitlink:     configs.Gitlink,
			Branch:      configs.Branch,
			BuildCmd:    configs.BuildCmd,
			StartCmd:    configs.StartCmd,
			RuntimeEnv:  configs.RuntimeEnv,
			RunningPort: configs.RunningPort,
			EnvVars:     configs.EnvVars,
		},
	)
	backend.ResConn()
	if err != nil {
		panic(err)
	}
	sendDeploy, err := json.Marshal(TaskImageRequest{
		Name:        buildRawResponse.GetName(),
		DockerImage: buildRawResponse.GetImageName(),
		RunningPort: buildRawResponse.GetRunningPort(),
		BasedMetric: internal.ContainerBasedMetric{
			CpuPercent: buildRawResponse.BasedMetrics.CpuPercent,
			MemUsage:   buildRawResponse.BasedMetrics.MemUsage,
			DiskUsage:  buildRawResponse.BasedMetrics.DiskUsage,
		},
	})
	if err != nil {
		panic("Couldnt marshal 2")
	}
	m.AddTask("DEPLOY", sendDeploy)
}

func (m *Master) BuildFile(message kafka.Message) {
	var configs TaskFileRequest
	err := json.Unmarshal(message.Value, &configs)
	if err != nil {
		panic("Couldnt unmarshal")
	}
	backend := MasterPlanAlgo(m.ServerPool, "BUILDER")
	conn, err := grpc.Dial(
		backend.URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic("Connection failed")
	}
	builder := builderrpc.NewBuilderServiceClient(conn)
	backend.AddConn()
	buildRawResponse, err := builder.BuildSpec(
		context.Background(),
		&builderrpc.BuildSpecRequest{
			Name:    configs.Name,
			Gitlink: configs.Gitlink,
			Branch:  configs.Branch,
			EnvVars: configs.EnvVars,
		},
	)
	backend.ResConn()
	if err != nil {
		panic("Couldnt marshal 3")
	}
	sendDeploy, err := json.Marshal(TaskImageRequest{
		Name:        buildRawResponse.GetName(),
		DockerImage: buildRawResponse.GetImageName(),
		RunningPort: buildRawResponse.GetRunningPort(),
	})
	if err != nil {
		panic("Couldnt marshal 4")
	}
	m.AddTask("DEPLOY", sendDeploy)
}

func (m *Master) Deploy(message kafka.Message) {
	startDep := time.Now()
	var configs TaskImageRequest
	err := json.Unmarshal(message.Value, &configs)
	if err != nil {
		panic("Couldnt unmarshal")
	}
	if len(m.ServerPool) == 0 {
		return
	}
	go func() {
		err := m.DeployDnsRecord(configs.Name, configs.RunningPort)
		if err != nil {
			fmt.Println("Sqlite error ", err)
		}
	}()
	backend := MasterPlanAlgo(m.ServerPool, "WORKER")
	// backendJson, err := json.Marshal(backend)
	//log.Println("Selected Backend:", backend.URL.Host)
	conn, err := grpc.Dial(
		backend.URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic("Connection failed")
	}
	worker := workerrpc.NewWorkerServiceClient(conn)
	backend.AddConn()
	buildRawResponse, err := worker.AddTask(
		context.Background(),
		&workerrpc.Task{
			Name:        configs.Name,
			ImageName:   configs.DockerImage,
			RunningPort: configs.RunningPort,
		},
	)
	backend.ResConn()
	if err != nil {
		panic(err)
	}
	backend.NumofContainers += 1
	endDep := time.Now()
	log.Printf("%s-%s - %s:%d\n", backend.URL.Host, configs.DockerImage, configs.Name, endDep.Sub(startDep))
	task := Task{
		Subdomain:   configs.Name,
		URL:         backend.URL,
		Runningport: configs.RunningPort,
		ImageName:   configs.DockerImage,
		Hostport:    buildRawResponse.GetHostPort(),
		ContainerID: buildRawResponse.GetContainerID(),
		Status:      "Deployed",
		Type:        TYPE_WEBSERVICE,
	}
	go func() {
		err := m.UpdateDeployRecord(task)
		if err != nil {
			fmt.Println("Sqlite error ", err)
		}
	}()
	// if m.dbDns != nil {
	// 	err = m.AddDnsRecord(task)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	m.cacheDns.Add(configs.Name, task)
}

func (m *Master) TerminateTask(name string) error {
	task, err := m.GetRecord(name)
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(
		task.URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	worker := workerrpc.NewWorkerServiceClient(conn)
	resp, err := worker.TerminateTask(context.Background(), &workerrpc.TerminateTaskRequest{
		Name:        task.Subdomain,
		ContainerID: task.ContainerID,
		ImageName:   task.ImageName,
	})
	if !resp.Success {
		return &TerminationError{}
	}
	m.RemoveRecord(task.Subdomain)
	return nil
}

func (m *Master) TaskMetrics(name string) (*internal.ContainerBasedMetric, error) {
	task, err := m.GetRecord(name)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(
		task.URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	worker := workerrpc.NewWorkerServiceClient(conn)
	resp, err := worker.GetTaskMetrics(context.Background(), &workerrpc.MetricRequest{
		ContainerID: task.ContainerID,
	})
	if err != nil {
		return nil, &NoMetricsError{}
	}
	return &internal.ContainerBasedMetric{
		CpuPercent: resp.CpuPercent,
		MemUsage:   resp.MemUsage,
		DiskUsage:  resp.DiskUsage,
	}, nil
}

func GetFilePathAndMeta(fhsHeader textproto.MIMEHeader) (string, string) {
	var filename string
	contentDisp := fhsHeader.Get("Content-Disposition")
	keys := strings.Split(contentDisp, ";")
	filename = strings.Split(keys[2], "=")[1]
	return strings.ReplaceAll(filename, "\"", ""), fhsHeader.Get("Content-Type")
}

func GetS3Path(id, path string) string {
	oldName := strings.Split(path, "/")
	oldName[0] = id
	return strings.Join(oldName, "/")
}

func NewS3() *s3.S3 {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApSouth1RegionID),
	}))
	// creds := stscreds.NewCredentials(sess, "arn:aws:s3:::66049c07d9e8546699fe0872fd32d8f6")
	return s3.New(sess, &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ID"),
			os.Getenv("AWS_SECRET"),
			"",
		),
	})
}

func (m *Master) DeployStatic(domain string, files []*multipart.FileHeader) error {
	id := uuid.New().String()
	for _, fhs := range files {
		file, err := fhs.Open()
		if err != nil {
			return err
		}
		filePath, metaType := GetFilePathAndMeta(fhs.Header)
		s3path := GetS3Path(id, filePath)
		bucket := NewS3()
		_, err = bucket.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(os.Getenv("BUCKET_NAME")),
			Key:         aws.String(s3path),
			Body:        file,
			ContentType: &metaType,
		})
		fmt.Println(s3path)
		if err != nil {
			return err
		}
		// handleFile(w, file, filename)
		file.Close()
	}
	m.cacheDns.Add(domain, Task{
		Subdomain:   domain,
		ContainerID: id,
		Status:      "Deployed",
		Type:        TYPE_STATIC,
	})
	m.DeployStaticRecord(domain, id)
	return nil
}

func (m *Master) KafkaHandler(message kafka.Message) {
	key := string(message.Key)
	switch key {
	case "BUILDRAW":
		m.BuildRaw(message)
	case "BUILDWITHFILE":
		m.BuildFile(message)
	case "DEPLOY":
		m.Deploy(message)
	}
}

func (m *Master) KafkaError() {
	panic("A kafka error occured")
}
