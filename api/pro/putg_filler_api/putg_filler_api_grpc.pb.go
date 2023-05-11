// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/putg_filler_api.proto

package putg_filler_api

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

// PutgFillerServiceClient is the client API for PutgFillerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PutgFillerServiceClient interface {
	Get(ctx context.Context, in *GetPutgFiller, opts ...grpc.CallOption) (*PutgFiller, error)
	Create(ctx context.Context, in *CreatePutgFiller, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdatePutgFiller, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeletePutgFiller, opts ...grpc.CallOption) (*response_model.Response, error)
}

type putgFillerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPutgFillerServiceClient(cc grpc.ClientConnInterface) PutgFillerServiceClient {
	return &putgFillerServiceClient{cc}
}

func (c *putgFillerServiceClient) Get(ctx context.Context, in *GetPutgFiller, opts ...grpc.CallOption) (*PutgFiller, error) {
	out := new(PutgFiller)
	err := c.cc.Invoke(ctx, "/putg_filler_api.PutgFillerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgFillerServiceClient) Create(ctx context.Context, in *CreatePutgFiller, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_filler_api.PutgFillerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgFillerServiceClient) Update(ctx context.Context, in *UpdatePutgFiller, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_filler_api.PutgFillerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgFillerServiceClient) Delete(ctx context.Context, in *DeletePutgFiller, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_filler_api.PutgFillerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutgFillerServiceServer is the server API for PutgFillerService service.
// All implementations must embed UnimplementedPutgFillerServiceServer
// for forward compatibility
type PutgFillerServiceServer interface {
	Get(context.Context, *GetPutgFiller) (*PutgFiller, error)
	Create(context.Context, *CreatePutgFiller) (*response_model.Response, error)
	Update(context.Context, *UpdatePutgFiller) (*response_model.Response, error)
	Delete(context.Context, *DeletePutgFiller) (*response_model.Response, error)
	mustEmbedUnimplementedPutgFillerServiceServer()
}

// UnimplementedPutgFillerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPutgFillerServiceServer struct {
}

func (UnimplementedPutgFillerServiceServer) Get(context.Context, *GetPutgFiller) (*PutgFiller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPutgFillerServiceServer) Create(context.Context, *CreatePutgFiller) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPutgFillerServiceServer) Update(context.Context, *UpdatePutgFiller) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPutgFillerServiceServer) Delete(context.Context, *DeletePutgFiller) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPutgFillerServiceServer) mustEmbedUnimplementedPutgFillerServiceServer() {}

// UnsafePutgFillerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PutgFillerServiceServer will
// result in compilation errors.
type UnsafePutgFillerServiceServer interface {
	mustEmbedUnimplementedPutgFillerServiceServer()
}

func RegisterPutgFillerServiceServer(s grpc.ServiceRegistrar, srv PutgFillerServiceServer) {
	s.RegisterService(&PutgFillerService_ServiceDesc, srv)
}

func _PutgFillerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgFiller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgFillerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_filler_api.PutgFillerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgFillerServiceServer).Get(ctx, req.(*GetPutgFiller))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgFillerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePutgFiller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgFillerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_filler_api.PutgFillerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgFillerServiceServer).Create(ctx, req.(*CreatePutgFiller))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgFillerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePutgFiller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgFillerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_filler_api.PutgFillerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgFillerServiceServer).Update(ctx, req.(*UpdatePutgFiller))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgFillerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePutgFiller)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgFillerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_filler_api.PutgFillerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgFillerServiceServer).Delete(ctx, req.(*DeletePutgFiller))
	}
	return interceptor(ctx, in, info, handler)
}

// PutgFillerService_ServiceDesc is the grpc.ServiceDesc for PutgFillerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PutgFillerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "putg_filler_api.PutgFillerService",
	HandlerType: (*PutgFillerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PutgFillerService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PutgFillerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PutgFillerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PutgFillerService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/putg_filler_api.proto",
}