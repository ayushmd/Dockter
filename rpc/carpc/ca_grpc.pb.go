// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: rpc/carpc/ca.proto

package carpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CAService_WhoAmI_FullMethodName             = "/CAService/WhoAmI"
	CAService_HealthMetrics_FullMethodName      = "/CAService/HealthMetrics"
	CAService_GenerateSSHKeyPair_FullMethodName = "/CAService/GenerateSSHKeyPair"
)

// CAServiceClient is the client API for CAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CAServiceClient interface {
	WhoAmI(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CAWhoAmIResponse, error)
	HealthMetrics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CAHealthMetricResponse, error)
	GenerateSSHKeyPair(ctx context.Context, in *KeyPairRequest, opts ...grpc.CallOption) (*KeyPairs, error)
}

type cAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCAServiceClient(cc grpc.ClientConnInterface) CAServiceClient {
	return &cAServiceClient{cc}
}

func (c *cAServiceClient) WhoAmI(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CAWhoAmIResponse, error) {
	out := new(CAWhoAmIResponse)
	err := c.cc.Invoke(ctx, CAService_WhoAmI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAServiceClient) HealthMetrics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CAHealthMetricResponse, error) {
	out := new(CAHealthMetricResponse)
	err := c.cc.Invoke(ctx, CAService_HealthMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAServiceClient) GenerateSSHKeyPair(ctx context.Context, in *KeyPairRequest, opts ...grpc.CallOption) (*KeyPairs, error) {
	out := new(KeyPairs)
	err := c.cc.Invoke(ctx, CAService_GenerateSSHKeyPair_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CAServiceServer is the server API for CAService service.
// All implementations must embed UnimplementedCAServiceServer
// for forward compatibility
type CAServiceServer interface {
	WhoAmI(context.Context, *emptypb.Empty) (*CAWhoAmIResponse, error)
	HealthMetrics(context.Context, *emptypb.Empty) (*CAHealthMetricResponse, error)
	GenerateSSHKeyPair(context.Context, *KeyPairRequest) (*KeyPairs, error)
	mustEmbedUnimplementedCAServiceServer()
}

// UnimplementedCAServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCAServiceServer struct {
}

func (UnimplementedCAServiceServer) WhoAmI(context.Context, *emptypb.Empty) (*CAWhoAmIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedCAServiceServer) HealthMetrics(context.Context, *emptypb.Empty) (*CAHealthMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthMetrics not implemented")
}
func (UnimplementedCAServiceServer) GenerateSSHKeyPair(context.Context, *KeyPairRequest) (*KeyPairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSSHKeyPair not implemented")
}
func (UnimplementedCAServiceServer) mustEmbedUnimplementedCAServiceServer() {}

// UnsafeCAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CAServiceServer will
// result in compilation errors.
type UnsafeCAServiceServer interface {
	mustEmbedUnimplementedCAServiceServer()
}

func RegisterCAServiceServer(s grpc.ServiceRegistrar, srv CAServiceServer) {
	s.RegisterService(&CAService_ServiceDesc, srv)
}

func _CAService_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServiceServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CAService_WhoAmI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServiceServer).WhoAmI(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAService_HealthMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServiceServer).HealthMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CAService_HealthMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServiceServer).HealthMetrics(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAService_GenerateSSHKeyPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServiceServer).GenerateSSHKeyPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CAService_GenerateSSHKeyPair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServiceServer).GenerateSSHKeyPair(ctx, req.(*KeyPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CAService_ServiceDesc is the grpc.ServiceDesc for CAService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CAService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CAService",
	HandlerType: (*CAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhoAmI",
			Handler:    _CAService_WhoAmI_Handler,
		},
		{
			MethodName: "HealthMetrics",
			Handler:    _CAService_HealthMetrics_Handler,
		},
		{
			MethodName: "GenerateSSHKeyPair",
			Handler:    _CAService_GenerateSSHKeyPair_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/carpc/ca.proto",
}
