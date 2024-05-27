// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: scheduling.proto

package scheduling

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Scheduling_ParseEmail_FullMethodName = "/within.website.x.mimi.scheduling.Scheduling/ParseEmail"
)

// SchedulingClient is the client API for Scheduling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulingClient interface {
	ParseEmail(ctx context.Context, in *ParseReq, opts ...grpc.CallOption) (*ParseResp, error)
}

type schedulingClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulingClient(cc grpc.ClientConnInterface) SchedulingClient {
	return &schedulingClient{cc}
}

func (c *schedulingClient) ParseEmail(ctx context.Context, in *ParseReq, opts ...grpc.CallOption) (*ParseResp, error) {
	out := new(ParseResp)
	err := c.cc.Invoke(ctx, Scheduling_ParseEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulingServer is the server API for Scheduling service.
// All implementations must embed UnimplementedSchedulingServer
// for forward compatibility
type SchedulingServer interface {
	ParseEmail(context.Context, *ParseReq) (*ParseResp, error)
	mustEmbedUnimplementedSchedulingServer()
}

// UnimplementedSchedulingServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulingServer struct {
}

func (UnimplementedSchedulingServer) ParseEmail(context.Context, *ParseReq) (*ParseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseEmail not implemented")
}
func (UnimplementedSchedulingServer) mustEmbedUnimplementedSchedulingServer() {}

// UnsafeSchedulingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulingServer will
// result in compilation errors.
type UnsafeSchedulingServer interface {
	mustEmbedUnimplementedSchedulingServer()
}

func RegisterSchedulingServer(s grpc.ServiceRegistrar, srv SchedulingServer) {
	s.RegisterService(&Scheduling_ServiceDesc, srv)
}

func _Scheduling_ParseEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulingServer).ParseEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduling_ParseEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulingServer).ParseEmail(ctx, req.(*ParseReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Scheduling_ServiceDesc is the grpc.ServiceDesc for Scheduling service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scheduling_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "within.website.x.mimi.scheduling.Scheduling",
	HandlerType: (*SchedulingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseEmail",
			Handler:    _Scheduling_ParseEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduling.proto",
}
