package master

import (
	"github.com/ayush18023/Load_balancer_Fyp/rpc/masterrpc"
	"google.golang.org/grpc"
)

type MasterGrpc struct {
	masterrpc.UnimplementedMasterServiceServer
}

func NewMasterGrpcInstance() *grpc.Server {
	serverRegistrar := grpc.NewServer()
	service := &MasterGrpc{}
	masterrpc.RegisterMasterServiceServer(serverRegistrar, service)
	return serverRegistrar
}
