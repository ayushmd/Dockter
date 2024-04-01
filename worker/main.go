package worker

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/masterrpc"
	lru "github.com/hashicorp/golang-lru/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LocalTask struct {
	Subdomain   string //a unique id
	Hostport    string
	Runningport string
	ImageName   string
	ContainerID string
}

type BackendCopy struct {
	URL   url.URL `json:"url"`
	State string  `json:"state"`
	// mux            sync.RWMutex
	//Algo part
}

type Worker struct {
	Port       int
	Master     url.URL
	ServerPool []*BackendCopy
	cacheDns   *lru.Cache[string, LocalTask]
}

var Worker_ *Worker = &Worker{}

func (w *Worker) AddTask(id, imageName, runningPort string) (string, error) {
	// repoimageName := internal.GetKey("DOCKER_HUB_REPO_NAME") + imageName
	doc := internal.Dockter{}
	doc.Init()
	defer doc.Close()
	err := doc.PullFromRegistery(imageName)
	if err != nil {
		return "", err
	}
	port, err := internal.GetFreePort()
	if err != nil {
		panic("port not found")
	}
	strPort := fmt.Sprintf("%d", port)
	fmt.Println(strPort)
	containerID, err := doc.RunContainer(
		imageName,
		"",
		[]string{fmt.Sprintf("%s:%s/tcp", strPort, runningPort)},
	)
	// builder.RunContainer(imageName, strPort, runningPort)
	if err != nil {
		return "", err
	}
	w.cacheDns.Add(id, LocalTask{
		Subdomain:   id,
		ContainerID: containerID,
		ImageName:   imageName,
		Runningport: runningPort,
		Hostport:    strPort,
	})
	return strPort, nil
}

func (w *Worker) JoinMaster(masterurl string) {
	w.Master = url.URL{
		Host: masterurl,
	}
	// fmt.Println("here is ")
	conn, err := grpc.Dial(
		masterurl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	// fmt.Println("here is executed")
	if err != nil {
		panic("Connection failed")
	}
	master := masterrpc.NewMasterServiceClient(conn)
	// ip, err := internal.GetIP()
	// if err != nil {
	// 	panic("Error ip")
	// }
	myurl := fmt.Sprintf(":%d", w.Port)
	// fmt.Println("Till here ", myurl)
	CpuUsage, MemUsage, DiskUsage, err := internal.HealthMetrics()
	// fmt.Printf("The cpu usage is %f", CpuUsage)
	if err != nil {
		panic("Health")
	}
	joinresp, err := master.Join(
		context.Background(),
		&masterrpc.JoinServer{
			Url:       myurl,
			State:     "WORKER",
			CpuUsage:  float32(CpuUsage),
			MemUsage:  float32(MemUsage),
			DiskUsage: float32(DiskUsage),
		},
	)
	if err != nil {
		panic("Join error")
	}
	for _, servInPool := range joinresp.GetServerPool() {
		w.ServerPool = append(w.ServerPool, &BackendCopy{
			URL: url.URL{
				Host: servInPool.Url,
			},
			State: servInPool.State,
		})
	}
	fmt.Printf("Joined Master(%s)\n", masterurl)
	// do get response and hydrate Serverpool
}
