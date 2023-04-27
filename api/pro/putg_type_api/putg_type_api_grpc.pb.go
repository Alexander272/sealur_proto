// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/putg_type_api.proto

package putg_type_api

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

// PutgTypeServiceClient is the client API for PutgTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PutgTypeServiceClient interface {
	Get(ctx context.Context, in *GetPutgType, opts ...grpc.CallOption) (*PutgType, error)
	Create(ctx context.Context, in *CreatePutgType, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdatePutgType, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeletePutgType, opts ...grpc.CallOption) (*response_model.Response, error)
}

type putgTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPutgTypeServiceClient(cc grpc.ClientConnInterface) PutgTypeServiceClient {
	return &putgTypeServiceClient{cc}
}

func (c *putgTypeServiceClient) Get(ctx context.Context, in *GetPutgType, opts ...grpc.CallOption) (*PutgType, error) {
	out := new(PutgType)
	err := c.cc.Invoke(ctx, "/putg_type_api.PutgTypeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgTypeServiceClient) Create(ctx context.Context, in *CreatePutgType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_type_api.PutgTypeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgTypeServiceClient) Update(ctx context.Context, in *UpdatePutgType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_type_api.PutgTypeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgTypeServiceClient) Delete(ctx context.Context, in *DeletePutgType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_type_api.PutgTypeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutgTypeServiceServer is the server API for PutgTypeService service.
// All implementations must embed UnimplementedPutgTypeServiceServer
// for forward compatibility
type PutgTypeServiceServer interface {
	Get(context.Context, *GetPutgType) (*PutgType, error)
	Create(context.Context, *CreatePutgType) (*response_model.Response, error)
	Update(context.Context, *UpdatePutgType) (*response_model.Response, error)
	Delete(context.Context, *DeletePutgType) (*response_model.Response, error)
	mustEmbedUnimplementedPutgTypeServiceServer()
}

// UnimplementedPutgTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPutgTypeServiceServer struct {
}

func (UnimplementedPutgTypeServiceServer) Get(context.Context, *GetPutgType) (*PutgType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPutgTypeServiceServer) Create(context.Context, *CreatePutgType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPutgTypeServiceServer) Update(context.Context, *UpdatePutgType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPutgTypeServiceServer) Delete(context.Context, *DeletePutgType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPutgTypeServiceServer) mustEmbedUnimplementedPutgTypeServiceServer() {}

// UnsafePutgTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PutgTypeServiceServer will
// result in compilation errors.
type UnsafePutgTypeServiceServer interface {
	mustEmbedUnimplementedPutgTypeServiceServer()
}

func RegisterPutgTypeServiceServer(s grpc.ServiceRegistrar, srv PutgTypeServiceServer) {
	s.RegisterService(&PutgTypeService_ServiceDesc, srv)
}

func _PutgTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_type_api.PutgTypeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgTypeServiceServer).Get(ctx, req.(*GetPutgType))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgTypeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePutgType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgTypeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_type_api.PutgTypeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgTypeServiceServer).Create(ctx, req.(*CreatePutgType))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgTypeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePutgType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgTypeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_type_api.PutgTypeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgTypeServiceServer).Update(ctx, req.(*UpdatePutgType))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgTypeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePutgType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgTypeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_type_api.PutgTypeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgTypeServiceServer).Delete(ctx, req.(*DeletePutgType))
	}
	return interceptor(ctx, in, info, handler)
}

// PutgTypeService_ServiceDesc is the grpc.ServiceDesc for PutgTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PutgTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "putg_type_api.PutgTypeService",
	HandlerType: (*PutgTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PutgTypeService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PutgTypeService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PutgTypeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PutgTypeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/putg_type_api.proto",
}
