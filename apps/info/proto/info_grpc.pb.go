// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: info.proto

package proto

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

// InfoClient is the client API for Info service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type infoClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoClient(cc grpc.ClientConnInterface) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/Info/info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServer is the server API for Info service.
// All implementations must embed UnimplementedInfoServer
// for forward compatibility
type InfoServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	mustEmbedUnimplementedInfoServer()
}

// UnimplementedInfoServer must be embedded to have forward compatible implementations.
type UnimplementedInfoServer struct {
}

func (UnimplementedInfoServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedInfoServer) mustEmbedUnimplementedInfoServer() {}

// UnsafeInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServer will
// result in compilation errors.
type UnsafeInfoServer interface {
	mustEmbedUnimplementedInfoServer()
}

func RegisterInfoServer(s grpc.ServiceRegistrar, srv InfoServer) {
	s.RegisterService(&Info_ServiceDesc, srv)
}

func _Info_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Info/info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Info_ServiceDesc is the grpc.ServiceDesc for Info service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Info_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "info",
			Handler:    _Info_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}
