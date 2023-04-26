// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/putg_flange_type_api.proto

package putg_flange_type_api

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

// PutgFlangeTypeServiceClient is the client API for PutgFlangeTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PutgFlangeTypeServiceClient interface {
	Get(ctx context.Context, in *GetPutgFlangeType, opts ...grpc.CallOption) (*PutgFlangeType, error)
	Create(ctx context.Context, in *CreatePutgFlangeType, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdatePutgFlangeType, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeletePutgFlangeType, opts ...grpc.CallOption) (*response_model.Response, error)
}

type putgFlangeTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPutgFlangeTypeServiceClient(cc grpc.ClientConnInterface) PutgFlangeTypeServiceClient {
	return &putgFlangeTypeServiceClient{cc}
}

func (c *putgFlangeTypeServiceClient) Get(ctx context.Context, in *GetPutgFlangeType, opts ...grpc.CallOption) (*PutgFlangeType, error) {
	out := new(PutgFlangeType)
	err := c.cc.Invoke(ctx, "/putg_flange_type_api.PutgFlangeTypeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgFlangeTypeServiceClient) Create(ctx context.Context, in *CreatePutgFlangeType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_flange_type_api.PutgFlangeTypeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgFlangeTypeServiceClient) Update(ctx context.Context, in *UpdatePutgFlangeType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_flange_type_api.PutgFlangeTypeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgFlangeTypeServiceClient) Delete(ctx context.Context, in *DeletePutgFlangeType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_flange_type_api.PutgFlangeTypeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutgFlangeTypeServiceServer is the server API for PutgFlangeTypeService service.
// All implementations must embed UnimplementedPutgFlangeTypeServiceServer
// for forward compatibility
type PutgFlangeTypeServiceServer interface {
	Get(context.Context, *GetPutgFlangeType) (*PutgFlangeType, error)
	Create(context.Context, *CreatePutgFlangeType) (*response_model.Response, error)
	Update(context.Context, *UpdatePutgFlangeType) (*response_model.Response, error)
	Delete(context.Context, *DeletePutgFlangeType) (*response_model.Response, error)
	mustEmbedUnimplementedPutgFlangeTypeServiceServer()
}

// UnimplementedPutgFlangeTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPutgFlangeTypeServiceServer struct {
}

func (UnimplementedPutgFlangeTypeServiceServer) Get(context.Context, *GetPutgFlangeType) (*PutgFlangeType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPutgFlangeTypeServiceServer) Create(context.Context, *CreatePutgFlangeType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPutgFlangeTypeServiceServer) Update(context.Context, *UpdatePutgFlangeType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPutgFlangeTypeServiceServer) Delete(context.Context, *DeletePutgFlangeType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPutgFlangeTypeServiceServer) mustEmbedUnimplementedPutgFlangeTypeServiceServer() {}

// UnsafePutgFlangeTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PutgFlangeTypeServiceServer will
// result in compilation errors.
type UnsafePutgFlangeTypeServiceServer interface {
	mustEmbedUnimplementedPutgFlangeTypeServiceServer()
}

func RegisterPutgFlangeTypeServiceServer(s grpc.ServiceRegistrar, srv PutgFlangeTypeServiceServer) {
	s.RegisterService(&PutgFlangeTypeService_ServiceDesc, srv)
}

func _PutgFlangeTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgFlangeTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_flange_type_api.PutgFlangeTypeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgFlangeTypeServiceServer).Get(ctx, req.(*GetPutgFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgFlangeTypeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePutgFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgFlangeTypeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_flange_type_api.PutgFlangeTypeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgFlangeTypeServiceServer).Create(ctx, req.(*CreatePutgFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgFlangeTypeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePutgFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgFlangeTypeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_flange_type_api.PutgFlangeTypeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgFlangeTypeServiceServer).Update(ctx, req.(*UpdatePutgFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgFlangeTypeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePutgFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgFlangeTypeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_flange_type_api.PutgFlangeTypeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgFlangeTypeServiceServer).Delete(ctx, req.(*DeletePutgFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

// PutgFlangeTypeService_ServiceDesc is the grpc.ServiceDesc for PutgFlangeTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PutgFlangeTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "putg_flange_type_api.PutgFlangeTypeService",
	HandlerType: (*PutgFlangeTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PutgFlangeTypeService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PutgFlangeTypeService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PutgFlangeTypeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PutgFlangeTypeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/putg_flange_type_api.proto",
}
