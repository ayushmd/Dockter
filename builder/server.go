package builder

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/builderrpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BuilderServer struct {
	builderrpc.UnimplementedBuilderServiceServer
}

func (w *BuilderServer) WhoAmI(ctx context.Context, in *emptypb.Empty) (*builderrpc.BWhoAmIResponse, error) {
	return &builderrpc.BWhoAmIResponse{
		State: "BUILDER",
	}, nil
}

func (w *BuilderServer) BuildHealthMetrics(ctx context.Context, in *emptypb.Empty) (*builderrpc.BuildHealthResponse, error) {
	// cpuUsage, memUsage, diskUsage, err := internal.HealthMetrics()
	basedHealth, err := internal.HealthMetricsBased()
	if err != nil {
		return &builderrpc.BuildHealthResponse{}, err
	}
	return &builderrpc.BuildHealthResponse{
		CpuPercent:       basedHealth.CpuPercent,
		MemUsage:         basedHealth.MemUsage,
		TotalMem:         basedHealth.TotalMem,
		MemUsedPercent:   float32(basedHealth.MemUsedPercent),
		DiskUsage:        basedHealth.DiskUsage,
		TotalDisk:        basedHealth.TotalDisk,
		DiskUsagePercent: float32(basedHealth.DiskUsagePercent),
	}, nil
}

func (w *BuilderServer) BuildRaw(ctx context.Context, in *builderrpc.BuildRawRequest) (*builderrpc.BuildRawResponse, error) {
	image, basedMetrics, err := Builder_.BuildRaw(
		in.GetName(),
		in.GetGitlink(),
		in.GetBranch(),
		in.GetBuildCmd(),
		in.GetStartCmd(),
		in.GetRuntimeEnv(),
		in.GetRunningPort(),
		in.GetEnvVars(),
	)
	if err != nil {
		return &builderrpc.BuildRawResponse{}, err
	}
	return &builderrpc.BuildRawResponse{
		Success:     true,
		Name:        in.GetName(),
		ImageName:   image,
		RunningPort: in.GetRunningPort(),
		BasedMetrics: &builderrpc.ContainerStats{
			CpuPercent: basedMetrics.CpuPercent,
			MemUsage:   basedMetrics.MemUsage,
			DiskUsage:  basedMetrics.DiskUsage,
		},
	}, nil
}

func NewBuilderServer(port int) {
	var waitgrp sync.WaitGroup
	waitgrp.Add(1)
	Builder_.Port = port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic("builder server not started")
	}
	serverRegistrar := grpc.NewServer()
	service := &BuilderServer{}
	builderrpc.RegisterBuilderServiceServer(serverRegistrar, service)
	go func() {
		defer waitgrp.Done()
		err = serverRegistrar.Serve(lis)
		if err != nil {
			log.Fatalf("impossible to serve: %s", err)
		}
	}()
	fmt.Println("builder server started on port ", port)
	waitgrp.Wait()
}
