// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/peerService.proto

package proto

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

// PeerProtoClient is the client API for PeerProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerProtoClient interface {
	CheckConnection(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	LeaveNetwork(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type peerProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerProtoClient(cc grpc.ClientConnInterface) PeerProtoClient {
	return &peerProtoClient{cc}
}

func (c *peerProtoClient) CheckConnection(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/proto.PeerProto/CheckConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerProtoClient) LeaveNetwork(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.PeerProto/LeaveNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerProtoClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.PeerProto/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerProtoServer is the server API for PeerProto service.
// All implementations must embed UnimplementedPeerProtoServer
// for forward compatibility
type PeerProtoServer interface {
	CheckConnection(context.Context, *ConnectRequest) (*ConnectResponse, error)
	LeaveNetwork(context.Context, *LeaveRequest) (*emptypb.Empty, error)
	Vote(context.Context, *VoteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPeerProtoServer()
}

// UnimplementedPeerProtoServer must be embedded to have forward compatible implementations.
type UnimplementedPeerProtoServer struct {
}

func (UnimplementedPeerProtoServer) CheckConnection(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckConnection not implemented")
}
func (UnimplementedPeerProtoServer) LeaveNetwork(context.Context, *LeaveRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveNetwork not implemented")
}
func (UnimplementedPeerProtoServer) Vote(context.Context, *VoteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedPeerProtoServer) mustEmbedUnimplementedPeerProtoServer() {}

// UnsafePeerProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerProtoServer will
// result in compilation errors.
type UnsafePeerProtoServer interface {
	mustEmbedUnimplementedPeerProtoServer()
}

func RegisterPeerProtoServer(s grpc.ServiceRegistrar, srv PeerProtoServer) {
	s.RegisterService(&PeerProto_ServiceDesc, srv)
}

func _PeerProto_CheckConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerProtoServer).CheckConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PeerProto/CheckConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerProtoServer).CheckConnection(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerProto_LeaveNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerProtoServer).LeaveNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PeerProto/LeaveNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerProtoServer).LeaveNetwork(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerProto_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerProtoServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PeerProto/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerProtoServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PeerProto_ServiceDesc is the grpc.ServiceDesc for PeerProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeerProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PeerProto",
	HandlerType: (*PeerProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckConnection",
			Handler:    _PeerProto_CheckConnection_Handler,
		},
		{
			MethodName: "LeaveNetwork",
			Handler:    _PeerProto_LeaveNetwork_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _PeerProto_Vote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/peerService.proto",
}
