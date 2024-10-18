// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: announce.proto

package announce

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	jsonfeed "within.website/x/proto/external/jsonfeed"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Announce_Announce_FullMethodName = "/within.website.x.mimi.announce.Announce/Announce"
)

// AnnounceClient is the client API for Announce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnnounceClient interface {
	Announce(ctx context.Context, in *jsonfeed.Item, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type announceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnounceClient(cc grpc.ClientConnInterface) AnnounceClient {
	return &announceClient{cc}
}

func (c *announceClient) Announce(ctx context.Context, in *jsonfeed.Item, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Announce_Announce_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnounceServer is the server API for Announce service.
// All implementations must embed UnimplementedAnnounceServer
// for forward compatibility.
type AnnounceServer interface {
	Announce(context.Context, *jsonfeed.Item) (*emptypb.Empty, error)
	mustEmbedUnimplementedAnnounceServer()
}

// UnimplementedAnnounceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnnounceServer struct{}

func (UnimplementedAnnounceServer) Announce(context.Context, *jsonfeed.Item) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Announce not implemented")
}
func (UnimplementedAnnounceServer) mustEmbedUnimplementedAnnounceServer() {}
func (UnimplementedAnnounceServer) testEmbeddedByValue()                  {}

// UnsafeAnnounceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnnounceServer will
// result in compilation errors.
type UnsafeAnnounceServer interface {
	mustEmbedUnimplementedAnnounceServer()
}

func RegisterAnnounceServer(s grpc.ServiceRegistrar, srv AnnounceServer) {
	// If the following call pancis, it indicates UnimplementedAnnounceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Announce_ServiceDesc, srv)
}

func _Announce_Announce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(jsonfeed.Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnounceServer).Announce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Announce_Announce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnounceServer).Announce(ctx, req.(*jsonfeed.Item))
	}
	return interceptor(ctx, in, info, handler)
}

// Announce_ServiceDesc is the grpc.ServiceDesc for Announce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Announce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "within.website.x.mimi.announce.Announce",
	HandlerType: (*AnnounceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Announce",
			Handler:    _Announce_Announce_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "announce.proto",
}

const (
	Post_Post_FullMethodName = "/within.website.x.mimi.announce.Post/Post"
)

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostClient interface {
	Post(ctx context.Context, in *StatusUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type postClient struct {
	cc grpc.ClientConnInterface
}

func NewPostClient(cc grpc.ClientConnInterface) PostClient {
	return &postClient{cc}
}

func (c *postClient) Post(ctx context.Context, in *StatusUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Post_Post_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
// All implementations must embed UnimplementedPostServer
// for forward compatibility.
type PostServer interface {
	Post(context.Context, *StatusUpdate) (*emptypb.Empty, error)
	mustEmbedUnimplementedPostServer()
}

// UnimplementedPostServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostServer struct{}

func (UnimplementedPostServer) Post(context.Context, *StatusUpdate) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedPostServer) mustEmbedUnimplementedPostServer() {}
func (UnimplementedPostServer) testEmbeddedByValue()              {}

// UnsafePostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServer will
// result in compilation errors.
type UnsafePostServer interface {
	mustEmbedUnimplementedPostServer()
}

func RegisterPostServer(s grpc.ServiceRegistrar, srv PostServer) {
	// If the following call pancis, it indicates UnimplementedPostServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Post_ServiceDesc, srv)
}

func _Post_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_Post_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Post(ctx, req.(*StatusUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// Post_ServiceDesc is the grpc.ServiceDesc for Post service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Post_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "within.website.x.mimi.announce.Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Post_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "announce.proto",
}
