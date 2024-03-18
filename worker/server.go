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
	cpuUsage, memUsage, diskUsage, err := internal.HealthMetrics()
	if err != nil {
		return &workerrpc.HealthMetricResponse{}, err
	}
	return &workerrpc.HealthMetricResponse{
		CpuUsage:  float32(cpuUsage),
		MemUsage:  float32(memUsage),
		DiskUsage: float32(diskUsage),
	}, nil
}

func (w *WorkerServer) AddTask(ctx context.Context, in *workerrpc.Task) (*workerrpc.AddTaskResponse, error) {
	port, err := Worker_.AddTask(
		in.GetName(),
		in.GetImageName(),
		in.GetRunningPort(),
	)
	if err != nil {
		return &workerrpc.AddTaskResponse{}, err
	}
	return &workerrpc.AddTaskResponse{
		HostPort: port,
	}, nil
}

func NewWorkerServer(port int) {
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
