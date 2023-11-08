// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/tasks/v1/tasks.proto

package tasksv1

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
	TasksWriterService_CreateTask_FullMethodName = "/api.tasks.v1.TasksWriterService/CreateTask"
)

// TasksWriterServiceClient is the client API for TasksWriterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TasksWriterServiceClient interface {
	// CreateTask creates a Task.
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error)
}

type tasksWriterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksWriterServiceClient(cc grpc.ClientConnInterface) TasksWriterServiceClient {
	return &tasksWriterServiceClient{cc}
}

func (c *tasksWriterServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, TasksWriterService_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksWriterServiceServer is the server API for TasksWriterService service.
// All implementations must embed UnimplementedTasksWriterServiceServer
// for forward compatibility
type TasksWriterServiceServer interface {
	// CreateTask creates a Task.
	CreateTask(context.Context, *CreateTaskRequest) (*Task, error)
	mustEmbedUnimplementedTasksWriterServiceServer()
}

// UnimplementedTasksWriterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTasksWriterServiceServer struct {
}

func (UnimplementedTasksWriterServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTasksWriterServiceServer) mustEmbedUnimplementedTasksWriterServiceServer() {}

// UnsafeTasksWriterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksWriterServiceServer will
// result in compilation errors.
type UnsafeTasksWriterServiceServer interface {
	mustEmbedUnimplementedTasksWriterServiceServer()
}

func RegisterTasksWriterServiceServer(s grpc.ServiceRegistrar, srv TasksWriterServiceServer) {
	s.RegisterService(&TasksWriterService_ServiceDesc, srv)
}

func _TasksWriterService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksWriterServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksWriterService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksWriterServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TasksWriterService_ServiceDesc is the grpc.ServiceDesc for TasksWriterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TasksWriterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.tasks.v1.TasksWriterService",
	HandlerType: (*TasksWriterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TasksWriterService_CreateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tasks/v1/tasks.proto",
}
