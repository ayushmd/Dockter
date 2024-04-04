package master

import (
	"context"
	"strings"

	"github.com/ayush18023/Load_balancer_Fyp/rpc/masterrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MasterGrpc struct {
	masterrpc.UnimplementedMasterServiceServer
}

func (m *MasterGrpc) WhoAmI(ctx context.Context, in *emptypb.Empty) (*masterrpc.MWhoAmIResponse, error) {
	return &masterrpc.MWhoAmIResponse{
		State: "MASTER",
	}, nil
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
		strings.Split(p.Addr.String(), ":")[0]+in.GetUrl(),
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
