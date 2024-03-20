package master

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/builderrpc"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/workerrpc"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Backend struct {
	URL            url.URL `json:"url"`
	State          string  `json:"state"`
	IsAlive        bool    `json:"isalive"`
	CurrentConnect int     `json:"connections"`
	CpuUsage       float64 `json:"cpu"`
	MemUsage       float64 `json:"mem"`
	DiskUsage      float64 `json:"disk"`
	// mux            sync.RWMutex
	//Algo part
}

type Task struct {
	Subdomain   string //a unique id
	URL         url.URL
	Hostport    string
	Runningport string
	ImageName   string
}

type TaskRawRequest struct {
	Name       string            `json:"name"`
	Gitlink    string            `json:"gitLink"`
	Branch     string            `json:"branch"`
	BuildCmd   string            `json:"buildCmd"`
	StartCmd   string            `json:"startCmd"`
	RuntimeEnv string            `json:"runtimeEnv"`
	EnvVars    map[string]string `json:"envVars"`
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
} // DEPLOY

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

type Master struct {
	kwriter    *internal.KafkaWriter
	ServerPool []*Backend
	cacheDns   *lru.Cache[string, Task]
}

var Master_ *Master

var kacp = keepalive.ClientParameters{
	Timeout:             2 * time.Second, // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,            // send pings even without active streams
}

func (m *Master) PoolServ(waitgrp *sync.WaitGroup, serv *Backend) {
	fmt.Println("Pooling server ", serv.URL.Host)
	defer waitgrp.Done()
	conn, err := grpc.Dial(
		serv.URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	if serv.State == "WORKER" {
		w := workerrpc.NewWorkerServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()
		res, err := w.HealthMetrics(ctx, &emptypb.Empty{})
		if err != nil {
			fmt.Printf("Server %s didnt respond", serv.URL.Host)
			serv.IsAlive = false
			fmt.Println(err)
			return
		}
		serv.CpuUsage = float64(res.CpuUsage)
		serv.MemUsage = float64(res.MemUsage)
		serv.DiskUsage = float64(res.DiskUsage)
		serv.IsAlive = true
	} else {
		b := builderrpc.NewBuilderServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()
		res, err := b.BuildHealthMetrics(ctx, &emptypb.Empty{})
		if err != nil {
			fmt.Printf("Server %s didnt respond", serv.URL.Host)
			serv.IsAlive = false
			fmt.Println(err)
			return
		}
		serv.CpuUsage = float64(res.CpuUsage)
		serv.MemUsage = float64(res.MemUsage)
		serv.DiskUsage = float64(res.DiskUsage)
		serv.IsAlive = true
	}
}

func (m *Master) Pool() {
	fmt.Println("Pooling started")
	var waitgrp *sync.WaitGroup = &sync.WaitGroup{}
	for _, serv := range m.ServerPool {
		waitgrp.Add(1)
		go m.PoolServ(waitgrp, serv)
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
func (m *Master) Join(peerurl, peerState string, CpuUsage, MemUsage, DiskUsage float64) {
	if i := m.HasJoined(peerurl); i == -1 {
		fmt.Println("This is called by ", peerurl)
		urlparsed := url.URL{
			Host: peerurl,
		}
		m.ServerPool = append(m.ServerPool, &Backend{
			URL:            urlparsed,
			State:          peerState,
			CurrentConnect: 0,
			CpuUsage:       CpuUsage,
			MemUsage:       MemUsage,
			DiskUsage:      DiskUsage,
			IsAlive:        true,
		})
	} else {
		m.ServerPool[i].IsAlive = true
	}
}

func (m *Master) GetServerPoolHandler() ([]byte, error) {
	marshalData, err := json.Marshal(m.ServerPool)
	if err != nil {
		return []byte(""), err
	}
	return marshalData, nil
}

func (m *Master) AddTask(request string, marshTask []byte) {
	m.kwriter.Write(
		[]byte(request),
		marshTask,
	)
	// if tconfig.HasDockerImage {
	// 	newtask := TaskConfig{
	// 		Name:        tconfig.Name,
	// 		DockerImage: tconfig.DockerImage,
	// 		RunningPort: tconfig.RunningPort,
	// 	}
	// 	enc, _ := json.Marshal(newtask)
	// 	m.KafkaManager.Write(
	// 		[]byte("DEPLOY"),
	// 		enc,
	// 	)
	// } else {
	// 	if tconfig.HasDockerFile {
	// 		buildtask := TaskConfig{
	// 			Gitlink: tconfig.Gitlink,
	// 			Branch:  tconfig.Branch,
	// 			EnvVars: tconfig.EnvVars,
	// 		}
	// 		enc, _ := json.Marshal(buildtask)
	// 		m.KafkaManager.Write(
	// 			[]byte("BUILDFILE"),
	// 			enc,
	// 		)
	// 	} else {
	// 		buildtask := TaskConfig{
	// 			Gitlink:    tconfig.Gitlink,
	// 			Branch:     tconfig.Branch,
	// 			BuildCmd:   tconfig.BuildCmd,
	// 			StartCmd:   tconfig.StartCmd,
	// 			RuntimeEnv: tconfig.RuntimeEnv,
	// 			EnvVars:    tconfig.EnvVars,
	// 		}
	// 		enc, _ := json.Marshal(buildtask)
	// 		m.KafkaManager.Write(
	// 			[]byte("BUILDRAW"),
	// 			enc,
	// 		)
	// 	}
	// }
}

func (m *Master) BuildRaw(message kafka.Message) {
	var configs TaskRawRequest
	err := json.Unmarshal(message.Value, &configs)
	if err != nil {
		panic("Couldnt unmarshal")
	}
	conn, err := grpc.Dial(
		MasterPlanAlgo(m.ServerPool, "BUILDER").URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic("Connection failed")
	}
	builder := builderrpc.NewBuilderServiceClient(conn)
	buildRawResponse, err := builder.BuildRaw(
		context.Background(),
		&builderrpc.BuildRawRequest{
			Name:       configs.Name,
			Gitlink:    configs.Gitlink,
			Branch:     configs.Branch,
			BuildCmd:   configs.BuildCmd,
			StartCmd:   configs.StartCmd,
			RuntimeEnv: configs.RuntimeEnv,
			EnvVars:    configs.EnvVars,
		},
	)
	if err != nil {
		panic(err)
	}
	sendDeploy, err := json.Marshal(TaskImageRequest{
		Name:        buildRawResponse.GetName(),
		DockerImage: buildRawResponse.GetImageName(),
		RunningPort: buildRawResponse.GetRunningPort(),
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
	conn, err := grpc.Dial(
		MasterPlanAlgo(m.ServerPool, "BUILDER").URL.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic("Connection failed")
	}
	builder := builderrpc.NewBuilderServiceClient(conn)
	buildRawResponse, err := builder.BuildSpec(
		context.Background(),
		&builderrpc.BuildSpecRequest{
			Name:    configs.Name,
			Gitlink: configs.Gitlink,
			Branch:  configs.Branch,
			EnvVars: configs.EnvVars,
		},
	)
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
	var configs TaskImageRequest
	err := json.Unmarshal(message.Value, &configs)
	if err != nil {
		panic("Couldnt unmarshal")
	}
	servUrl := MasterPlanAlgo(m.ServerPool, "WORKER").URL
	conn, err := grpc.Dial(
		servUrl.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic("Connection failed")
	}
	fmt.Println("The recieved stuff ", configs.Name, configs.DockerImage, configs.RunningPort)
	worker := workerrpc.NewWorkerServiceClient(conn)
	buildRawResponse, err := worker.AddTask(
		context.Background(),
		&workerrpc.Task{
			Name:        configs.Name,
			ImageName:   configs.DockerImage,
			RunningPort: configs.RunningPort,
		},
	)
	if err != nil {
		panic(err)
	}
	m.cacheDns.Add(configs.Name, Task{
		Subdomain:   configs.Name,
		URL:         servUrl,
		Runningport: configs.RunningPort,
		ImageName:   configs.DockerImage,
		Hostport:    buildRawResponse.GetHostPort(),
	})
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
