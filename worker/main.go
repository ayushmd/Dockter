package worker

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/masterrpc"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
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

type Worker struct {
	Port       int
	Master     url.URL
	ServerPool []*Backend
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
	portBindings := nat.PortMap{
		nat.Port(runningPort): []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: strPort,
			},
		},
	}
	containerID, err := doc.RunContainer(
		imageName,
		&container.HostConfig{
			NetworkMode:  "host",
			PortBindings: portBindings,
		},
	)
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
	fmt.Println("here is ")
	conn, err := grpc.Dial(
		masterurl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	fmt.Println("here is executed")
	if err != nil {
		panic("Connection failed")
	}
	master := masterrpc.NewMasterServiceClient(conn)
	ip, err := internal.GetIP()
	if err != nil {
		panic("Error ip")
	}
	myurl := ip + fmt.Sprintf("%d", w.Port)
	fmt.Println("Till here ", myurl)
	CpuUsage, MemUsage, DiskUsage, err := internal.HealthMetrics()
	if err != nil {
		panic("Health")
	}
	_, err = master.Join(
		context.Background(),
		&masterrpc.JoinRequest{
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
	// do get response and hydrate Serverpool
}
