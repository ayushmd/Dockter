package master

import (
	"context"

	"github.com/ayush18023/Load_balancer_Fyp/rpc/masterrpc"
	"google.golang.org/grpc"
)

type MasterGrpc struct {
	masterrpc.UnimplementedMasterServiceServer
}

func (m *MasterGrpc) Join(ctx context.Context, in *masterrpc.JoinRequest) (*masterrpc.JoinResponse, error) {
	// URL:            urlparsed,
	// 	State:          peerState,
	// 	CurrentConnect: 0,
	// 	CpuUsage:       CpuUsage,
	// 	MemUsage:       MemUsage,
	// 	DiskUsage:      DiskUsage,
	// 	IsAlive:        true,
	Master_.Join(
		in.GetUrl(),
		in.GetState(),
		float64(in.GetCpuUsage()),
		float64(in.GetMemUsage()),
		float64(in.GetDiskUsage()),
	)
	return &masterrpc.JoinResponse{
		Success: true,
	}, nil
}

func NewMasterGrpcInstance() *grpc.Server {
	serverRegistrar := grpc.NewServer()
	service := &MasterGrpc{}
	masterrpc.RegisterMasterServiceServer(serverRegistrar, service)
	return serverRegistrar
}
