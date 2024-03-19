package master

import (
	"context"

	"github.com/ayush18023/Load_balancer_Fyp/rpc/masterrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type MasterGrpc struct {
	masterrpc.UnimplementedMasterServiceServer
}

func (m *MasterGrpc) Join(ctx context.Context, in *masterrpc.JoinServer) (*masterrpc.JoinResponse, error) {
	var servpool []*masterrpc.JoinServer
	for _, server := range Master_.ServerPool {
		servpool = append(servpool, &masterrpc.JoinServer{
			Url:       server.URL.Host,
			State:     server.State,
			CpuUsage:  float32(server.CpuUsage),
			MemUsage:  float32(server.MemUsage),
			DiskUsage: float32(server.DiskUsage),
		})
	}
	p, _ := peer.FromContext(ctx)
	Master_.Join(
		p.Addr.String(),
		in.GetState(),
		float64(in.GetCpuUsage()),
		float64(in.GetMemUsage()),
		float64(in.GetDiskUsage()),
	)
	return &masterrpc.JoinResponse{
		ServerPool: servpool,
	}, nil
}

func NewMasterGrpcInstance() *grpc.Server {
	serverRegistrar := grpc.NewServer()
	service := &MasterGrpc{}
	masterrpc.RegisterMasterServiceServer(serverRegistrar, service)
	return serverRegistrar
}
