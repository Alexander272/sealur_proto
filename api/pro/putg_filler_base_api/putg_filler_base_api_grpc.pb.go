// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/putg_filler_base_api.proto

package putg_filler_base_api

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

// PutgBaseFillerServiceClient is the client API for PutgBaseFillerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PutgBaseFillerServiceClient interface {
	Get(ctx context.Context, in *GetPutgBaseFiller, opts ...grpc.CallOption) (*PutgBaseFiller, error)
	Create(ctx context.Context, in *CreatePutgBaseFiller, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdatePutgBaseFiller, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeletePutgBaseFiller, opts ...grpc.CallOption) (*response_model.Response, error)
}

type putgBaseFillerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPutgBaseFillerServiceClient(cc grpc.ClientConnInterface) PutgBaseFillerServiceClient {
	return &putgBaseFillerServiceClient{cc}
}

func (c *putgBaseFillerServiceClient) Get(ctx context.Context, in *GetPutgBaseFiller, opts ...grpc.CallOption) (*PutgBaseFiller, error) {
	out := new(PutgBaseFiller)
	err := c.cc.Invoke(ctx, "/putg_filler_base_api.PutgBaseFillerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgBaseFillerServiceClient) Create(ctx context.Context, in *CreatePutgBaseFiller, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_filler_base_api.PutgBaseFillerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgBaseFillerServiceClient) Update(ctx context.Context, in *UpdatePutgBaseFiller, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_filler_base_api.PutgBaseFillerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgBaseFillerServiceClient) Delete(ctx context.Context, in *DeletePutgBaseFiller, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_filler_base_api.PutgBaseFillerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutgBaseFillerServiceServer is the server API for PutgBaseFillerService service.
// All implementations must embed UnimplementedPutgBaseFillerServiceServer
// for forward compatibility
type PutgBaseFillerServiceServer interface {
	Get(context.Context, *GetPutgBaseFiller) (*PutgBaseFiller, error)
	Create(context.Context, *CreatePutgBaseFiller) (*response_model.Response, error)
	Update(context.Context, *UpdatePutgBaseFiller) (*response_model.Response, error)
	Delete(context.Context, *DeletePutgBaseFiller) (*response_model.Response, error)
	mustEmbedUnimplementedPutgBaseFillerServiceServer()
}

// UnimplementedPutgBaseFillerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPutgBaseFillerServiceServer struct {
}

func (UnimplementedPutgBaseFillerServiceServer) Get(context.Context, *GetPutgBaseFiller) (*PutgBaseFiller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPutgBaseFillerServiceServer) Create(context.Context, *CreatePutgBaseFiller) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPutgBaseFillerServiceServer) Update(context.Context, *UpdatePutgBaseFiller) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPutgBaseFillerServiceServer) Delete(context.Context, *DeletePutgBaseFiller) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPutgBaseFillerServiceServer) mustEmbedUnimplementedPutgBaseFillerServiceServer() {}

// UnsafePutgBaseFillerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PutgBaseFillerServiceServer will
// result in compilation errors.
type UnsafePutgBaseFillerServiceServer interface {
	mustEmbedUnimplementedPutgBaseFillerServiceServer()
}

func RegisterPutgBaseFillerServiceServer(s grpc.ServiceRegistrar, srv PutgBaseFillerServiceServer) {
	s.RegisterService(&PutgBaseFillerService_ServiceDesc, srv)
}

func _PutgBaseFillerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgBaseFiller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgBaseFillerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_filler_base_api.PutgBaseFillerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgBaseFillerServiceServer).Get(ctx, req.(*GetPutgBaseFiller))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgBaseFillerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePutgBaseFiller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgBaseFillerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_filler_base_api.PutgBaseFillerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgBaseFillerServiceServer).Create(ctx, req.(*CreatePutgBaseFiller))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgBaseFillerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePutgBaseFiller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgBaseFillerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_filler_base_api.PutgBaseFillerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgBaseFillerServiceServer).Update(ctx, req.(*UpdatePutgBaseFiller))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgBaseFillerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePutgBaseFiller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgBaseFillerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_filler_base_api.PutgBaseFillerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgBaseFillerServiceServer).Delete(ctx, req.(*DeletePutgBaseFiller))
	}
	return interceptor(ctx, in, info, handler)
}

// PutgBaseFillerService_ServiceDesc is the grpc.ServiceDesc for PutgBaseFillerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PutgBaseFillerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "putg_filler_base_api.PutgBaseFillerService",
	HandlerType: (*PutgBaseFillerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PutgBaseFillerService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PutgBaseFillerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PutgBaseFillerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PutgBaseFillerService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/putg_filler_base_api.proto",
}