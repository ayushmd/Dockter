package worker

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/masterrpc"
	"github.com/docker/docker/api/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// type LocalTask struct {
// 	Subdomain   string //a unique id
// 	Hostport    string
// 	Runningport string
// 	ImageName   string
// 	ContainerID string
// }

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
	// cacheDns   *lru.Cache[string, LocalTask]
}

var Worker_ *Worker = &Worker{}

func (w *Worker) AddTask(id, imageName, runningPort string, hasSSH bool) (string, string, error) {
	// repoimageName := internal.GetKey("DOCKER_HUB_REPO_NAME") + imageName
	log.Println("Adding Task ", id, imageName, runningPort)
	doc := internal.Dockter{}
	doc.Init()
	defer doc.Close()
	err := doc.PullFromRegistery(imageName)
	if err != nil {
		return "", "", err
	}
	port, err := internal.GetFreePort()
	if err != nil {
		panic("port not found")
	}
	strPort := fmt.Sprintf("%d", port)
	fmt.Println(strPort)
	exposedPorts := []string{fmt.Sprintf("%s:%s/tcp", strPort, runningPort)}
	if hasSSH {
		port, err = internal.GetFreePort()
		if err != nil {
			panic("port not found")
		}
		fmt.Println("SSH PORT: ", port)
		exposedPorts = append(exposedPorts, fmt.Sprintf("%d:22/tcp", port))
	}
	containerID, err := doc.RunContainer(
		imageName,
		"",
		exposedPorts,
	)
	// builder.RunContainer(imageName, strPort, runningPort)
	if err != nil {
		return "", "", err
	}
	// w.cacheDns.Add(id, LocalTask{
	// 	Subdomain:   id,
	// 	ContainerID: containerID,
	// 	ImageName:   imageName,
	// 	Runningport: runningPort,
	// 	Hostport:    strPort,
	// })
	return strPort, containerID, nil
}

func (w *Worker) GetTasks() []types.Container {
	doc := internal.Dockter{}
	doc.Init()
	defer doc.Close()
	return doc.ListContainers()
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
	// CpuUsage, MemUsage, DiskUsage, err := internal.HealthMetrics()
	basedMetrics, err := internal.HealthMetricsBased()
	if err != nil {
		panic(err)
	}
	HealthStats := &masterrpc.BasedMetrics{
		CpuPercent:       basedMetrics.CpuPercent,
		MemUsage:         basedMetrics.MemUsage,
		TotalMem:         basedMetrics.TotalMem,
		MemUsedPercent:   float32(basedMetrics.MemUsedPercent),
		DiskUsage:        basedMetrics.DiskUsage,
		TotalDisk:        basedMetrics.TotalDisk,
		DiskUsagePercent: float32(basedMetrics.DiskUsagePercent),
	}
	// fmt.Printf("The cpu usage is %f", CpuUsage)
	if err != nil {
		panic("Health")
	}
	joinresp, err := master.Join(
		context.Background(),
		&masterrpc.JoinServer{
			Url:   myurl,
			State: "WORKER",
			Stats: HealthStats,
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
	log.Printf("Joined Master Server(%s)\n", masterurl)
	// do get response and hydrate Serverpool
}

func (w *Worker) Obliterate(name, containerID, imageName string) error {
	doc := internal.NewDockter()
	defer doc.Close()
	if err := doc.TrashContainer(containerID); err != nil {
		return err
	}
	if err := doc.ClearImages(imageName); err != nil {
		return err
	}
	return nil
}
