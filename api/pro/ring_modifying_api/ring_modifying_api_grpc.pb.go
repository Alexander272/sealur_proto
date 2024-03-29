// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pro/ring_modifying_api.proto

package ring_modifying_api

import (
	context "context"
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RingModifyingService_GetAll_FullMethodName = "/ring_modifying_api.RingModifyingService/GetAll"
	RingModifyingService_Create_FullMethodName = "/ring_modifying_api.RingModifyingService/Create"
	RingModifyingService_Update_FullMethodName = "/ring_modifying_api.RingModifyingService/Update"
	RingModifyingService_Delete_FullMethodName = "/ring_modifying_api.RingModifyingService/Delete"
)

// RingModifyingServiceClient is the client API for RingModifyingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RingModifyingServiceClient interface {
	GetAll(ctx context.Context, in *GetRingModifying, opts ...grpc.CallOption) (*RingModifying, error)
	Create(ctx context.Context, in *CreateRingModifying, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdateRingModifying, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeleteRingModifying, opts ...grpc.CallOption) (*response_model.Response, error)
}

type ringModifyingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRingModifyingServiceClient(cc grpc.ClientConnInterface) RingModifyingServiceClient {
	return &ringModifyingServiceClient{cc}
}

func (c *ringModifyingServiceClient) GetAll(ctx context.Context, in *GetRingModifying, opts ...grpc.CallOption) (*RingModifying, error) {
	out := new(RingModifying)
	err := c.cc.Invoke(ctx, RingModifyingService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ringModifyingServiceClient) Create(ctx context.Context, in *CreateRingModifying, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, RingModifyingService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ringModifyingServiceClient) Update(ctx context.Context, in *UpdateRingModifying, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, RingModifyingService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ringModifyingServiceClient) Delete(ctx context.Context, in *DeleteRingModifying, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, RingModifyingService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RingModifyingServiceServer is the server API for RingModifyingService service.
// All implementations must embed UnimplementedRingModifyingServiceServer
// for forward compatibility
type RingModifyingServiceServer interface {
	GetAll(context.Context, *GetRingModifying) (*RingModifying, error)
	Create(context.Context, *CreateRingModifying) (*response_model.Response, error)
	Update(context.Context, *UpdateRingModifying) (*response_model.Response, error)
	Delete(context.Context, *DeleteRingModifying) (*response_model.Response, error)
	mustEmbedUnimplementedRingModifyingServiceServer()
}

// UnimplementedRingModifyingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRingModifyingServiceServer struct {
}

func (UnimplementedRingModifyingServiceServer) GetAll(context.Context, *GetRingModifying) (*RingModifying, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedRingModifyingServiceServer) Create(context.Context, *CreateRingModifying) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRingModifyingServiceServer) Update(context.Context, *UpdateRingModifying) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRingModifyingServiceServer) Delete(context.Context, *DeleteRingModifying) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRingModifyingServiceServer) mustEmbedUnimplementedRingModifyingServiceServer() {}

// UnsafeRingModifyingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RingModifyingServiceServer will
// result in compilation errors.
type UnsafeRingModifyingServiceServer interface {
	mustEmbedUnimplementedRingModifyingServiceServer()
}

func RegisterRingModifyingServiceServer(s grpc.ServiceRegistrar, srv RingModifyingServiceServer) {
	s.RegisterService(&RingModifyingService_ServiceDesc, srv)
}

func _RingModifyingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRingModifying)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RingModifyingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RingModifyingService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RingModifyingServiceServer).GetAll(ctx, req.(*GetRingModifying))
	}
	return interceptor(ctx, in, info, handler)
}

func _RingModifyingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRingModifying)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RingModifyingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RingModifyingService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RingModifyingServiceServer).Create(ctx, req.(*CreateRingModifying))
	}
	return interceptor(ctx, in, info, handler)
}

func _RingModifyingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRingModifying)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RingModifyingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RingModifyingService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RingModifyingServiceServer).Update(ctx, req.(*UpdateRingModifying))
	}
	return interceptor(ctx, in, info, handler)
}

func _RingModifyingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRingModifying)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RingModifyingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RingModifyingService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RingModifyingServiceServer).Delete(ctx, req.(*DeleteRingModifying))
	}
	return interceptor(ctx, in, info, handler)
}

// RingModifyingService_ServiceDesc is the grpc.ServiceDesc for RingModifyingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RingModifyingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ring_modifying_api.RingModifyingService",
	HandlerType: (*RingModifyingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _RingModifyingService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RingModifyingService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RingModifyingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RingModifyingService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/ring_modifying_api.proto",
}
