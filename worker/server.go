package worker

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/workerrpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WorkerServer struct {
	workerrpc.UnimplementedWorkerServiceServer
}

func (w *WorkerServer) WhoAmI(ctx context.Context, in *emptypb.Empty) (*workerrpc.WWhoAmIResponse, error) {
	return &workerrpc.WWhoAmIResponse{
		State: "WORKER",
	}, nil
}

func (w *WorkerServer) HealthScore(ctx context.Context, in *emptypb.Empty) (*workerrpc.HealthScoreResponse, error) {
	return &workerrpc.HealthScoreResponse{
		Score: 10.0,
	}, nil
}

func (w *WorkerServer) HealthMetrics(ctx context.Context, in *emptypb.Empty) (*workerrpc.HealthMetricResponse, error) {
	// cpuUsage, memUsage, diskUsage, err := internal.HealthMetrics()
	basedHealth, err := internal.HealthMetricsBased()
	if err != nil {
		return &workerrpc.HealthMetricResponse{}, err
	}
	return &workerrpc.HealthMetricResponse{
		CpuPercent:       basedHealth.CpuPercent,
		MemUsage:         basedHealth.MemUsage,
		TotalMem:         basedHealth.TotalMem,
		MemUsedPercent:   float32(basedHealth.MemUsedPercent),
		DiskUsage:        basedHealth.DiskUsage,
		TotalDisk:        basedHealth.TotalDisk,
		DiskUsagePercent: float32(basedHealth.DiskUsagePercent),
	}, nil
}

func (w *WorkerServer) GetTasks(ctx context.Context, in *emptypb.Empty) (*workerrpc.RunningTasks, error) {
	containers := Worker_.GetTasks()
	reqContainers := []*workerrpc.RunningTask{}
	for _, container := range containers {
		var respPorts []*workerrpc.Port
		for _, port := range container.Ports {
			respPorts = append(respPorts, &workerrpc.Port{
				IP:          port.IP,
				RunningPort: int32(port.PrivatePort),
				HostPort:    int32(port.PublicPort),
				Type:        port.Type,
			})
		}
		reqContainers = append(reqContainers, &workerrpc.RunningTask{
			ContainerID: container.ID,
			ImageName:   container.Image,
			Ports:       respPorts,
			State:       container.State,
		})
	}
	return &workerrpc.RunningTasks{
		Containers: reqContainers,
	}, nil
}

func (w *WorkerServer) AddTask(ctx context.Context, in *workerrpc.Task) (*workerrpc.AddTaskResponse, error) {
	port, containerID, err := Worker_.AddTask(
		in.GetName(),
		in.GetImageName(),
		in.GetRunningPort(),
	)
	if err != nil {
		return &workerrpc.AddTaskResponse{}, err
	}
	return &workerrpc.AddTaskResponse{
		HostPort:    port,
		ContainerID: containerID,
	}, nil
}

func NewWorkerServer(port int) {
	// cache, err := lru.New[string, LocalTask](128)
	// if err != nil {
	// 	log.Fatal("Master server not started")
	// }
	// Worker_.cacheDns = cache
	var waitgrp sync.WaitGroup
	waitgrp.Add(1)
	Worker_.Port = port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic("Worker server not started")
	}
	serverRegistrar := grpc.NewServer()
	service := &WorkerServer{}
	workerrpc.RegisterWorkerServiceServer(serverRegistrar, service)
	go func() {
		defer waitgrp.Done()
		err = serverRegistrar.Serve(lis)
		if err != nil {
			log.Fatalf("impossible to serve: %s", err)
		}
	}()
	fmt.Println("Worker server started on port ", port)
	waitgrp.Wait()
}
