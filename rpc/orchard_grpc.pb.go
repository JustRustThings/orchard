// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: orchard.proto

package rpc

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

// ControllerClient is the client API for Controller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControllerClient interface {
	// message bus between the controller and a worker
	Watch(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Controller_WatchClient, error)
	// single purpose method when a port forward is requested and running
	// session information is passed in the requests metadata
	PortForward(ctx context.Context, opts ...grpc.CallOption) (Controller_PortForwardClient, error)
}

type controllerClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerClient(cc grpc.ClientConnInterface) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) Watch(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Controller_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Controller_ServiceDesc.Streams[0], "/Controller/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Controller_WatchClient interface {
	Recv() (*WatchInstruction, error)
	grpc.ClientStream
}

type controllerWatchClient struct {
	grpc.ClientStream
}

func (x *controllerWatchClient) Recv() (*WatchInstruction, error) {
	m := new(WatchInstruction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controllerClient) PortForward(ctx context.Context, opts ...grpc.CallOption) (Controller_PortForwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &Controller_ServiceDesc.Streams[1], "/Controller/PortForward", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerPortForwardClient{stream}
	return x, nil
}

type Controller_PortForwardClient interface {
	Send(*PortForwardData) error
	Recv() (*PortForwardData, error)
	grpc.ClientStream
}

type controllerPortForwardClient struct {
	grpc.ClientStream
}

func (x *controllerPortForwardClient) Send(m *PortForwardData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controllerPortForwardClient) Recv() (*PortForwardData, error) {
	m := new(PortForwardData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControllerServer is the server API for Controller service.
// All implementations must embed UnimplementedControllerServer
// for forward compatibility
type ControllerServer interface {
	// message bus between the controller and a worker
	Watch(*emptypb.Empty, Controller_WatchServer) error
	// single purpose method when a port forward is requested and running
	// session information is passed in the requests metadata
	PortForward(Controller_PortForwardServer) error
	mustEmbedUnimplementedControllerServer()
}

// UnimplementedControllerServer must be embedded to have forward compatible implementations.
type UnimplementedControllerServer struct {
}

func (UnimplementedControllerServer) Watch(*emptypb.Empty, Controller_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedControllerServer) PortForward(Controller_PortForwardServer) error {
	return status.Errorf(codes.Unimplemented, "method PortForward not implemented")
}
func (UnimplementedControllerServer) mustEmbedUnimplementedControllerServer() {}

// UnsafeControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerServer will
// result in compilation errors.
type UnsafeControllerServer interface {
	mustEmbedUnimplementedControllerServer()
}

func RegisterControllerServer(s grpc.ServiceRegistrar, srv ControllerServer) {
	s.RegisterService(&Controller_ServiceDesc, srv)
}

func _Controller_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerServer).Watch(m, &controllerWatchServer{stream})
}

type Controller_WatchServer interface {
	Send(*WatchInstruction) error
	grpc.ServerStream
}

type controllerWatchServer struct {
	grpc.ServerStream
}

func (x *controllerWatchServer) Send(m *WatchInstruction) error {
	return x.ServerStream.SendMsg(m)
}

func _Controller_PortForward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControllerServer).PortForward(&controllerPortForwardServer{stream})
}

type Controller_PortForwardServer interface {
	Send(*PortForwardData) error
	Recv() (*PortForwardData, error)
	grpc.ServerStream
}

type controllerPortForwardServer struct {
	grpc.ServerStream
}

func (x *controllerPortForwardServer) Send(m *PortForwardData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controllerPortForwardServer) Recv() (*PortForwardData, error) {
	m := new(PortForwardData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Controller_ServiceDesc is the grpc.ServiceDesc for Controller service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Controller_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Controller_Watch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PortForward",
			Handler:       _Controller_PortForward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orchard.proto",
}
