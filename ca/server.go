package ca

import (
	"context"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
	"github.com/ayush18023/Load_balancer_Fyp/rpc/carpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CAServer struct {
	carpc.UnimplementedCAServiceServer
}

func (c *CAServer) WhoAmI(ctx context.Context, in *emptypb.Empty) (*carpc.CAWhoAmIResponse, error) {
	return &carpc.CAWhoAmIResponse{
		State: "CA",
	}, nil
}

func (c *CAServer) HealthMetrics(ctx context.Context, in *emptypb.Empty) (*carpc.CAHealthMetricResponse, error) {
	// cpuUsage, memUsage, diskUsage, err := internal.HealthMetrics()
	basedHealth, err := internal.HealthMetricsBased()
	if err != nil {
		return &carpc.CAHealthMetricResponse{}, err
	}
	return &carpc.CAHealthMetricResponse{
		CpuPercent:       basedHealth.CpuPercent,
		MemUsage:         basedHealth.MemUsage,
		TotalMem:         basedHealth.TotalMem,
		MemUsedPercent:   float32(basedHealth.MemUsedPercent),
		DiskUsage:        basedHealth.DiskUsage,
		TotalDisk:        basedHealth.TotalDisk,
		DiskUsagePercent: float32(basedHealth.DiskUsagePercent),
	}, nil
}

func (c *CAServer) GenerateSSHKeyPair(ctx context.Context, in *carpc.KeyPairRequest) (*carpc.KeyPairs, error) {
	pkey, err := GenerateSSHKeyPairs(in.Algorithm, in.Keyname)
	if err != nil {
		return &carpc.KeyPairs{}, err
	}
	return &carpc.KeyPairs{
		PrivateKey: pkey,
	}, nil
}

func NewCAInstance() *grpc.Server {
	serverRegistrar := grpc.NewServer()
	service := &CAServer{}
	carpc.RegisterCAServiceServer(serverRegistrar, service)
	return serverRegistrar
}
