// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/position_api.proto

package position_api

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

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	Create(ctx context.Context, in *CreatePosition, opts ...grpc.CallOption) (*response_model.IdResponse, error)
	Update(ctx context.Context, in *UpdatePosition, opts ...grpc.CallOption) (*response_model.Response, error)
	Copy(ctx context.Context, in *CopyPosition, opts ...grpc.CallOption) (*CopyDrawing, error)
	Delete(ctx context.Context, in *DeletePosition, opts ...grpc.CallOption) (*response_model.Response, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) Create(ctx context.Context, in *CreatePosition, opts ...grpc.CallOption) (*response_model.IdResponse, error) {
	out := new(response_model.IdResponse)
	err := c.cc.Invoke(ctx, "/position_api.PositionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Update(ctx context.Context, in *UpdatePosition, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/position_api.PositionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Copy(ctx context.Context, in *CopyPosition, opts ...grpc.CallOption) (*CopyDrawing, error) {
	out := new(CopyDrawing)
	err := c.cc.Invoke(ctx, "/position_api.PositionService/Copy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Delete(ctx context.Context, in *DeletePosition, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/position_api.PositionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations must embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	Create(context.Context, *CreatePosition) (*response_model.IdResponse, error)
	Update(context.Context, *UpdatePosition) (*response_model.Response, error)
	Copy(context.Context, *CopyPosition) (*CopyDrawing, error)
	Delete(context.Context, *DeletePosition) (*response_model.Response, error)
	mustEmbedUnimplementedPositionServiceServer()
}

// UnimplementedPositionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) Create(context.Context, *CreatePosition) (*response_model.IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPositionServiceServer) Update(context.Context, *UpdatePosition) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPositionServiceServer) Copy(context.Context, *CopyPosition) (*CopyDrawing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
func (UnimplementedPositionServiceServer) Delete(context.Context, *DeletePosition) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPositionServiceServer) mustEmbedUnimplementedPositionServiceServer() {}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_api.PositionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Create(ctx, req.(*CreatePosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_api.PositionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Update(ctx, req.(*UpdatePosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_api.PositionService/Copy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Copy(ctx, req.(*CopyPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_api.PositionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Delete(ctx, req.(*DeletePosition))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "position_api.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PositionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PositionService_Update_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _PositionService_Copy_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PositionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/position_api.proto",
}
