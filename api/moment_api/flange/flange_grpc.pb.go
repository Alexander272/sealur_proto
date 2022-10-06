// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: moment_api/flange.proto

package flange

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

// FlangeServiceClient is the client API for FlangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlangeServiceClient interface {
	GetBolts(ctx context.Context, in *GetBoltsRequest, opts ...grpc.CallOption) (*BoltsResponse, error)
	GetAllBolts(ctx context.Context, in *GetBoltsRequest, opts ...grpc.CallOption) (*BoltsResponse, error)
	CreateBolt(ctx context.Context, in *CreateBoltRequest, opts ...grpc.CallOption) (*Response, error)
	CreateBolts(ctx context.Context, in *CreateBoltsRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateBolt(ctx context.Context, in *UpdateBoltRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteBolt(ctx context.Context, in *DeleteBoltRequest, opts ...grpc.CallOption) (*Response, error)
	GetTypeFlange(ctx context.Context, in *GetTypeFlangeRequest, opts ...grpc.CallOption) (*TypeFlangeResponse, error)
	CreateTypeFlange(ctx context.Context, in *CreateTypeFlangeRequest, opts ...grpc.CallOption) (*IdResponse, error)
	UpdateTypeFlange(ctx context.Context, in *UpdateTypeFlangeRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteTypeFlange(ctx context.Context, in *DeleteTypeFlangeRequest, opts ...grpc.CallOption) (*Response, error)
	GetStandarts(ctx context.Context, in *GetStandartsRequest, opts ...grpc.CallOption) (*StandartsResponse, error)
	GetStandartsWithSize(ctx context.Context, in *GetStandartsRequest, opts ...grpc.CallOption) (*StandartsWithSizeResponse, error)
	CreateStandart(ctx context.Context, in *CreateStandartRequest, opts ...grpc.CallOption) (*IdResponse, error)
	UpdateStandart(ctx context.Context, in *UpdateStandartRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteStandart(ctx context.Context, in *DeleteStandartRequest, opts ...grpc.CallOption) (*Response, error)
	CreateFlangeSize(ctx context.Context, in *CreateFlangeSizeRequest, opts ...grpc.CallOption) (*Response, error)
	CreateFlangeSizes(ctx context.Context, in *CreateFlangeSizesRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateFlangeSize(ctx context.Context, in *UpdateFlangeSizeRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteFlangeSize(ctx context.Context, in *DeleteFlangeSizeRequest, opts ...grpc.CallOption) (*Response, error)
	GetBasisFlangeSize(ctx context.Context, in *GetBasisFlangeSizeRequest, opts ...grpc.CallOption) (*BasisFlangeSizeResponse, error)
	GetFlangeSize(ctx context.Context, in *GetFullFlangeSizeRequest, opts ...grpc.CallOption) (*FullFlangeSizeResponse, error)
}

type flangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlangeServiceClient(cc grpc.ClientConnInterface) FlangeServiceClient {
	return &flangeServiceClient{cc}
}

func (c *flangeServiceClient) GetBolts(ctx context.Context, in *GetBoltsRequest, opts ...grpc.CallOption) (*BoltsResponse, error) {
	out := new(BoltsResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/GetBolts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) GetAllBolts(ctx context.Context, in *GetBoltsRequest, opts ...grpc.CallOption) (*BoltsResponse, error) {
	out := new(BoltsResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/GetAllBolts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) CreateBolt(ctx context.Context, in *CreateBoltRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/CreateBolt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) CreateBolts(ctx context.Context, in *CreateBoltsRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/CreateBolts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) UpdateBolt(ctx context.Context, in *UpdateBoltRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/UpdateBolt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) DeleteBolt(ctx context.Context, in *DeleteBoltRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/DeleteBolt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) GetTypeFlange(ctx context.Context, in *GetTypeFlangeRequest, opts ...grpc.CallOption) (*TypeFlangeResponse, error) {
	out := new(TypeFlangeResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/GetTypeFlange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) CreateTypeFlange(ctx context.Context, in *CreateTypeFlangeRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/CreateTypeFlange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) UpdateTypeFlange(ctx context.Context, in *UpdateTypeFlangeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/UpdateTypeFlange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) DeleteTypeFlange(ctx context.Context, in *DeleteTypeFlangeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/DeleteTypeFlange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) GetStandarts(ctx context.Context, in *GetStandartsRequest, opts ...grpc.CallOption) (*StandartsResponse, error) {
	out := new(StandartsResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/GetStandarts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) GetStandartsWithSize(ctx context.Context, in *GetStandartsRequest, opts ...grpc.CallOption) (*StandartsWithSizeResponse, error) {
	out := new(StandartsWithSizeResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/GetStandartsWithSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) CreateStandart(ctx context.Context, in *CreateStandartRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/CreateStandart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) UpdateStandart(ctx context.Context, in *UpdateStandartRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/UpdateStandart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) DeleteStandart(ctx context.Context, in *DeleteStandartRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/DeleteStandart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) CreateFlangeSize(ctx context.Context, in *CreateFlangeSizeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/CreateFlangeSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) CreateFlangeSizes(ctx context.Context, in *CreateFlangeSizesRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/CreateFlangeSizes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) UpdateFlangeSize(ctx context.Context, in *UpdateFlangeSizeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/UpdateFlangeSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) DeleteFlangeSize(ctx context.Context, in *DeleteFlangeSizeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/DeleteFlangeSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) GetBasisFlangeSize(ctx context.Context, in *GetBasisFlangeSizeRequest, opts ...grpc.CallOption) (*BasisFlangeSizeResponse, error) {
	out := new(BasisFlangeSizeResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/GetBasisFlangeSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeServiceClient) GetFlangeSize(ctx context.Context, in *GetFullFlangeSizeRequest, opts ...grpc.CallOption) (*FullFlangeSizeResponse, error) {
	out := new(FullFlangeSizeResponse)
	err := c.cc.Invoke(ctx, "/moment_api.flange.FlangeService/GetFlangeSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlangeServiceServer is the server API for FlangeService service.
// All implementations must embed UnimplementedFlangeServiceServer
// for forward compatibility
type FlangeServiceServer interface {
	GetBolts(context.Context, *GetBoltsRequest) (*BoltsResponse, error)
	GetAllBolts(context.Context, *GetBoltsRequest) (*BoltsResponse, error)
	CreateBolt(context.Context, *CreateBoltRequest) (*Response, error)
	CreateBolts(context.Context, *CreateBoltsRequest) (*Response, error)
	UpdateBolt(context.Context, *UpdateBoltRequest) (*Response, error)
	DeleteBolt(context.Context, *DeleteBoltRequest) (*Response, error)
	GetTypeFlange(context.Context, *GetTypeFlangeRequest) (*TypeFlangeResponse, error)
	CreateTypeFlange(context.Context, *CreateTypeFlangeRequest) (*IdResponse, error)
	UpdateTypeFlange(context.Context, *UpdateTypeFlangeRequest) (*Response, error)
	DeleteTypeFlange(context.Context, *DeleteTypeFlangeRequest) (*Response, error)
	GetStandarts(context.Context, *GetStandartsRequest) (*StandartsResponse, error)
	GetStandartsWithSize(context.Context, *GetStandartsRequest) (*StandartsWithSizeResponse, error)
	CreateStandart(context.Context, *CreateStandartRequest) (*IdResponse, error)
	UpdateStandart(context.Context, *UpdateStandartRequest) (*Response, error)
	DeleteStandart(context.Context, *DeleteStandartRequest) (*Response, error)
	CreateFlangeSize(context.Context, *CreateFlangeSizeRequest) (*Response, error)
	CreateFlangeSizes(context.Context, *CreateFlangeSizesRequest) (*Response, error)
	UpdateFlangeSize(context.Context, *UpdateFlangeSizeRequest) (*Response, error)
	DeleteFlangeSize(context.Context, *DeleteFlangeSizeRequest) (*Response, error)
	GetBasisFlangeSize(context.Context, *GetBasisFlangeSizeRequest) (*BasisFlangeSizeResponse, error)
	GetFlangeSize(context.Context, *GetFullFlangeSizeRequest) (*FullFlangeSizeResponse, error)
	mustEmbedUnimplementedFlangeServiceServer()
}

// UnimplementedFlangeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlangeServiceServer struct {
}

func (UnimplementedFlangeServiceServer) GetBolts(context.Context, *GetBoltsRequest) (*BoltsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBolts not implemented")
}
func (UnimplementedFlangeServiceServer) GetAllBolts(context.Context, *GetBoltsRequest) (*BoltsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBolts not implemented")
}
func (UnimplementedFlangeServiceServer) CreateBolt(context.Context, *CreateBoltRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBolt not implemented")
}
func (UnimplementedFlangeServiceServer) CreateBolts(context.Context, *CreateBoltsRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBolts not implemented")
}
func (UnimplementedFlangeServiceServer) UpdateBolt(context.Context, *UpdateBoltRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBolt not implemented")
}
func (UnimplementedFlangeServiceServer) DeleteBolt(context.Context, *DeleteBoltRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBolt not implemented")
}
func (UnimplementedFlangeServiceServer) GetTypeFlange(context.Context, *GetTypeFlangeRequest) (*TypeFlangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypeFlange not implemented")
}
func (UnimplementedFlangeServiceServer) CreateTypeFlange(context.Context, *CreateTypeFlangeRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTypeFlange not implemented")
}
func (UnimplementedFlangeServiceServer) UpdateTypeFlange(context.Context, *UpdateTypeFlangeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTypeFlange not implemented")
}
func (UnimplementedFlangeServiceServer) DeleteTypeFlange(context.Context, *DeleteTypeFlangeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTypeFlange not implemented")
}
func (UnimplementedFlangeServiceServer) GetStandarts(context.Context, *GetStandartsRequest) (*StandartsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStandarts not implemented")
}
func (UnimplementedFlangeServiceServer) GetStandartsWithSize(context.Context, *GetStandartsRequest) (*StandartsWithSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStandartsWithSize not implemented")
}
func (UnimplementedFlangeServiceServer) CreateStandart(context.Context, *CreateStandartRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStandart not implemented")
}
func (UnimplementedFlangeServiceServer) UpdateStandart(context.Context, *UpdateStandartRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStandart not implemented")
}
func (UnimplementedFlangeServiceServer) DeleteStandart(context.Context, *DeleteStandartRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStandart not implemented")
}
func (UnimplementedFlangeServiceServer) CreateFlangeSize(context.Context, *CreateFlangeSizeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlangeSize not implemented")
}
func (UnimplementedFlangeServiceServer) CreateFlangeSizes(context.Context, *CreateFlangeSizesRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlangeSizes not implemented")
}
func (UnimplementedFlangeServiceServer) UpdateFlangeSize(context.Context, *UpdateFlangeSizeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlangeSize not implemented")
}
func (UnimplementedFlangeServiceServer) DeleteFlangeSize(context.Context, *DeleteFlangeSizeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFlangeSize not implemented")
}
func (UnimplementedFlangeServiceServer) GetBasisFlangeSize(context.Context, *GetBasisFlangeSizeRequest) (*BasisFlangeSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasisFlangeSize not implemented")
}
func (UnimplementedFlangeServiceServer) GetFlangeSize(context.Context, *GetFullFlangeSizeRequest) (*FullFlangeSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlangeSize not implemented")
}
func (UnimplementedFlangeServiceServer) mustEmbedUnimplementedFlangeServiceServer() {}

// UnsafeFlangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlangeServiceServer will
// result in compilation errors.
type UnsafeFlangeServiceServer interface {
	mustEmbedUnimplementedFlangeServiceServer()
}

func RegisterFlangeServiceServer(s grpc.ServiceRegistrar, srv FlangeServiceServer) {
	s.RegisterService(&FlangeService_ServiceDesc, srv)
}

func _FlangeService_GetBolts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoltsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).GetBolts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/GetBolts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).GetBolts(ctx, req.(*GetBoltsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_GetAllBolts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoltsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).GetAllBolts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/GetAllBolts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).GetAllBolts(ctx, req.(*GetBoltsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_CreateBolt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBoltRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).CreateBolt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/CreateBolt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).CreateBolt(ctx, req.(*CreateBoltRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_CreateBolts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBoltsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).CreateBolts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/CreateBolts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).CreateBolts(ctx, req.(*CreateBoltsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_UpdateBolt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBoltRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).UpdateBolt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/UpdateBolt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).UpdateBolt(ctx, req.(*UpdateBoltRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_DeleteBolt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBoltRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).DeleteBolt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/DeleteBolt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).DeleteBolt(ctx, req.(*DeleteBoltRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_GetTypeFlange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTypeFlangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).GetTypeFlange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/GetTypeFlange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).GetTypeFlange(ctx, req.(*GetTypeFlangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_CreateTypeFlange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTypeFlangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).CreateTypeFlange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/CreateTypeFlange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).CreateTypeFlange(ctx, req.(*CreateTypeFlangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_UpdateTypeFlange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTypeFlangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).UpdateTypeFlange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/UpdateTypeFlange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).UpdateTypeFlange(ctx, req.(*UpdateTypeFlangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_DeleteTypeFlange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTypeFlangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).DeleteTypeFlange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/DeleteTypeFlange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).DeleteTypeFlange(ctx, req.(*DeleteTypeFlangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_GetStandarts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStandartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).GetStandarts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/GetStandarts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).GetStandarts(ctx, req.(*GetStandartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_GetStandartsWithSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStandartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).GetStandartsWithSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/GetStandartsWithSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).GetStandartsWithSize(ctx, req.(*GetStandartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_CreateStandart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStandartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).CreateStandart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/CreateStandart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).CreateStandart(ctx, req.(*CreateStandartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_UpdateStandart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStandartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).UpdateStandart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/UpdateStandart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).UpdateStandart(ctx, req.(*UpdateStandartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_DeleteStandart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStandartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).DeleteStandart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/DeleteStandart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).DeleteStandart(ctx, req.(*DeleteStandartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_CreateFlangeSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlangeSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).CreateFlangeSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/CreateFlangeSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).CreateFlangeSize(ctx, req.(*CreateFlangeSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_CreateFlangeSizes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlangeSizesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).CreateFlangeSizes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/CreateFlangeSizes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).CreateFlangeSizes(ctx, req.(*CreateFlangeSizesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_UpdateFlangeSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlangeSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).UpdateFlangeSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/UpdateFlangeSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).UpdateFlangeSize(ctx, req.(*UpdateFlangeSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_DeleteFlangeSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFlangeSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).DeleteFlangeSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/DeleteFlangeSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).DeleteFlangeSize(ctx, req.(*DeleteFlangeSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_GetBasisFlangeSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasisFlangeSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).GetBasisFlangeSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/GetBasisFlangeSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).GetBasisFlangeSize(ctx, req.(*GetBasisFlangeSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeService_GetFlangeSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullFlangeSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeServiceServer).GetFlangeSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment_api.flange.FlangeService/GetFlangeSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeServiceServer).GetFlangeSize(ctx, req.(*GetFullFlangeSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlangeService_ServiceDesc is the grpc.ServiceDesc for FlangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moment_api.flange.FlangeService",
	HandlerType: (*FlangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBolts",
			Handler:    _FlangeService_GetBolts_Handler,
		},
		{
			MethodName: "GetAllBolts",
			Handler:    _FlangeService_GetAllBolts_Handler,
		},
		{
			MethodName: "CreateBolt",
			Handler:    _FlangeService_CreateBolt_Handler,
		},
		{
			MethodName: "CreateBolts",
			Handler:    _FlangeService_CreateBolts_Handler,
		},
		{
			MethodName: "UpdateBolt",
			Handler:    _FlangeService_UpdateBolt_Handler,
		},
		{
			MethodName: "DeleteBolt",
			Handler:    _FlangeService_DeleteBolt_Handler,
		},
		{
			MethodName: "GetTypeFlange",
			Handler:    _FlangeService_GetTypeFlange_Handler,
		},
		{
			MethodName: "CreateTypeFlange",
			Handler:    _FlangeService_CreateTypeFlange_Handler,
		},
		{
			MethodName: "UpdateTypeFlange",
			Handler:    _FlangeService_UpdateTypeFlange_Handler,
		},
		{
			MethodName: "DeleteTypeFlange",
			Handler:    _FlangeService_DeleteTypeFlange_Handler,
		},
		{
			MethodName: "GetStandarts",
			Handler:    _FlangeService_GetStandarts_Handler,
		},
		{
			MethodName: "GetStandartsWithSize",
			Handler:    _FlangeService_GetStandartsWithSize_Handler,
		},
		{
			MethodName: "CreateStandart",
			Handler:    _FlangeService_CreateStandart_Handler,
		},
		{
			MethodName: "UpdateStandart",
			Handler:    _FlangeService_UpdateStandart_Handler,
		},
		{
			MethodName: "DeleteStandart",
			Handler:    _FlangeService_DeleteStandart_Handler,
		},
		{
			MethodName: "CreateFlangeSize",
			Handler:    _FlangeService_CreateFlangeSize_Handler,
		},
		{
			MethodName: "CreateFlangeSizes",
			Handler:    _FlangeService_CreateFlangeSizes_Handler,
		},
		{
			MethodName: "UpdateFlangeSize",
			Handler:    _FlangeService_UpdateFlangeSize_Handler,
		},
		{
			MethodName: "DeleteFlangeSize",
			Handler:    _FlangeService_DeleteFlangeSize_Handler,
		},
		{
			MethodName: "GetBasisFlangeSize",
			Handler:    _FlangeService_GetBasisFlangeSize_Handler,
		},
		{
			MethodName: "GetFlangeSize",
			Handler:    _FlangeService_GetFlangeSize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moment_api/flange.proto",
}
