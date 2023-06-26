// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/services/task_svc/v1/task_template_svc.proto

package task_svc

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
	TaskTemplateService_CreateTaskTemplate_FullMethodName        = "/proto.services.task_svc.v1.TaskTemplateService/CreateTaskTemplate"
	TaskTemplateService_GetAllTaskTemplatesByWard_FullMethodName = "/proto.services.task_svc.v1.TaskTemplateService/GetAllTaskTemplatesByWard"
)

// TaskTemplateServiceClient is the client API for TaskTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskTemplateServiceClient interface {
	CreateTaskTemplate(ctx context.Context, in *CreateTaskTemplateRequest, opts ...grpc.CallOption) (*CreateTaskTemplateResponse, error)
	GetAllTaskTemplatesByWard(ctx context.Context, in *GetAllTaskTemplatesByWardRequest, opts ...grpc.CallOption) (*GetAllTaskTemplatesByWardResponse, error)
}

type taskTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskTemplateServiceClient(cc grpc.ClientConnInterface) TaskTemplateServiceClient {
	return &taskTemplateServiceClient{cc}
}

func (c *taskTemplateServiceClient) CreateTaskTemplate(ctx context.Context, in *CreateTaskTemplateRequest, opts ...grpc.CallOption) (*CreateTaskTemplateResponse, error) {
	out := new(CreateTaskTemplateResponse)
	err := c.cc.Invoke(ctx, TaskTemplateService_CreateTaskTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskTemplateServiceClient) GetAllTaskTemplatesByWard(ctx context.Context, in *GetAllTaskTemplatesByWardRequest, opts ...grpc.CallOption) (*GetAllTaskTemplatesByWardResponse, error) {
	out := new(GetAllTaskTemplatesByWardResponse)
	err := c.cc.Invoke(ctx, TaskTemplateService_GetAllTaskTemplatesByWard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskTemplateServiceServer is the server API for TaskTemplateService service.
// All implementations must embed UnimplementedTaskTemplateServiceServer
// for forward compatibility
type TaskTemplateServiceServer interface {
	CreateTaskTemplate(context.Context, *CreateTaskTemplateRequest) (*CreateTaskTemplateResponse, error)
	GetAllTaskTemplatesByWard(context.Context, *GetAllTaskTemplatesByWardRequest) (*GetAllTaskTemplatesByWardResponse, error)
	mustEmbedUnimplementedTaskTemplateServiceServer()
}

// UnimplementedTaskTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskTemplateServiceServer struct {
}

func (UnimplementedTaskTemplateServiceServer) CreateTaskTemplate(context.Context, *CreateTaskTemplateRequest) (*CreateTaskTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskTemplate not implemented")
}
func (UnimplementedTaskTemplateServiceServer) GetAllTaskTemplatesByWard(context.Context, *GetAllTaskTemplatesByWardRequest) (*GetAllTaskTemplatesByWardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTaskTemplatesByWard not implemented")
}
func (UnimplementedTaskTemplateServiceServer) mustEmbedUnimplementedTaskTemplateServiceServer() {}

// UnsafeTaskTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskTemplateServiceServer will
// result in compilation errors.
type UnsafeTaskTemplateServiceServer interface {
	mustEmbedUnimplementedTaskTemplateServiceServer()
}

func RegisterTaskTemplateServiceServer(s grpc.ServiceRegistrar, srv TaskTemplateServiceServer) {
	s.RegisterService(&TaskTemplateService_ServiceDesc, srv)
}

func _TaskTemplateService_CreateTaskTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskTemplateServiceServer).CreateTaskTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskTemplateService_CreateTaskTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskTemplateServiceServer).CreateTaskTemplate(ctx, req.(*CreateTaskTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskTemplateService_GetAllTaskTemplatesByWard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTaskTemplatesByWardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskTemplateServiceServer).GetAllTaskTemplatesByWard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskTemplateService_GetAllTaskTemplatesByWard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskTemplateServiceServer).GetAllTaskTemplatesByWard(ctx, req.(*GetAllTaskTemplatesByWardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskTemplateService_ServiceDesc is the grpc.ServiceDesc for TaskTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.services.task_svc.v1.TaskTemplateService",
	HandlerType: (*TaskTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTaskTemplate",
			Handler:    _TaskTemplateService_CreateTaskTemplate_Handler,
		},
		{
			MethodName: "GetAllTaskTemplatesByWard",
			Handler:    _TaskTemplateService_GetAllTaskTemplatesByWard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services/task_svc/v1/task_template_svc.proto",
}
