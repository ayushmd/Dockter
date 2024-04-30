package master

import (
	"context"
	"strings"

	"github.com/ayush18023/Load_balancer_Fyp/internal"
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
		HealthStats := &masterrpc.BasedMetrics{
			CpuPercent:       server.Stats.CpuPercent,
			MemUsage:         server.Stats.MemUsage,
			TotalMem:         server.Stats.TotalMem,
			MemUsedPercent:   float32(server.Stats.MemUsedPercent),
			DiskUsage:        server.Stats.DiskUsage,
			TotalDisk:        server.Stats.TotalDisk,
			DiskUsagePercent: float32(server.Stats.MemUsedPercent),
		}
		servpool = append(servpool, &masterrpc.JoinServer{
			Url:   server.URL.Host,
			State: server.State,
			Stats: HealthStats,
		})
	}
	p, _ := peer.FromContext(ctx)
	Master_.Join(
		strings.Split(p.Addr.String(), ":")[0]+in.GetUrl(),
		in.GetState(),
		internal.ContainerBasedMetric{},
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
