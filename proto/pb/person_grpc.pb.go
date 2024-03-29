// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// LifeClient is the client API for Life service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LifeClient interface {
	Work(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkResponse, error)
	Study(ctx context.Context, in *StudyRequest, opts ...grpc.CallOption) (*StudyResponse, error)
}

type lifeClient struct {
	cc grpc.ClientConnInterface
}

func NewLifeClient(cc grpc.ClientConnInterface) LifeClient {
	return &lifeClient{cc}
}

func (c *lifeClient) Work(ctx context.Context, in *WorkRequest, opts ...grpc.CallOption) (*WorkResponse, error) {
	out := new(WorkResponse)
	err := c.cc.Invoke(ctx, "/pb.Life/work", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifeClient) Study(ctx context.Context, in *StudyRequest, opts ...grpc.CallOption) (*StudyResponse, error) {
	out := new(StudyResponse)
	err := c.cc.Invoke(ctx, "/pb.Life/study", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LifeServer is the server API for Life service.
// All implementations must embed UnimplementedLifeServer
// for forward compatibility
type LifeServer interface {
	Work(context.Context, *WorkRequest) (*WorkResponse, error)
	Study(context.Context, *StudyRequest) (*StudyResponse, error)
	mustEmbedUnimplementedLifeServer()
}

// UnimplementedLifeServer must be embedded to have forward compatible implementations.
type UnimplementedLifeServer struct {
}

func (UnimplementedLifeServer) Work(context.Context, *WorkRequest) (*WorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Work not implemented")
}
func (UnimplementedLifeServer) Study(context.Context, *StudyRequest) (*StudyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Study not implemented")
}
func (UnimplementedLifeServer) mustEmbedUnimplementedLifeServer() {}

// UnsafeLifeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LifeServer will
// result in compilation errors.
type UnsafeLifeServer interface {
	mustEmbedUnimplementedLifeServer()
}

func RegisterLifeServer(s grpc.ServiceRegistrar, srv LifeServer) {
	s.RegisterService(&Life_ServiceDesc, srv)
}

func _Life_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifeServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Life/work",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifeServer).Work(ctx, req.(*WorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Life_Study_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifeServer).Study(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Life/study",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifeServer).Study(ctx, req.(*StudyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Life_ServiceDesc is the grpc.ServiceDesc for Life service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Life_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Life",
	HandlerType: (*LifeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "work",
			Handler:    _Life_Work_Handler,
		},
		{
			MethodName: "study",
			Handler:    _Life_Study_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}
