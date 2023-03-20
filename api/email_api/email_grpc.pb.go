// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: email_api/email.proto

package email_api

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

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Interview --------------------------------------------------
	SendInterview(ctx context.Context, opts ...grpc.CallOption) (EmailService_SendInterviewClient, error)
	// Order --------------------------------------------------
	SendOrder(ctx context.Context, opts ...grpc.CallOption) (EmailService_SendOrderClient, error)
	SendNotification(ctx context.Context, in *NotificationData, opts ...grpc.CallOption) (*SuccessResponse, error)
	// Confirm --------------------------------------------------
	SendConfirm(ctx context.Context, in *ConfirmUserRequestOld, opts ...grpc.CallOption) (*SuccessResponse, error)
	ConfirmUser(ctx context.Context, in *ConfirmUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	// Reject --------------------------------------------------
	SendReject(ctx context.Context, in *RejectUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	// Join --------------------------------------------------
	SendJoin(ctx context.Context, in *JoinUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	// Blocked ------------------------------------------------
	SendBlocked(ctx context.Context, in *BlockedUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	// TEST ---------------------------------------------------
	SendTest(ctx context.Context, in *SendTestRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/email_api.EmailService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendInterview(ctx context.Context, opts ...grpc.CallOption) (EmailService_SendInterviewClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmailService_ServiceDesc.Streams[0], "/email_api.EmailService/SendInterview", opts...)
	if err != nil {
		return nil, err
	}
	x := &emailServiceSendInterviewClient{stream}
	return x, nil
}

type EmailService_SendInterviewClient interface {
	Send(*SendInterviewRequest) error
	CloseAndRecv() (*SuccessResponse, error)
	grpc.ClientStream
}

type emailServiceSendInterviewClient struct {
	grpc.ClientStream
}

func (x *emailServiceSendInterviewClient) Send(m *SendInterviewRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *emailServiceSendInterviewClient) CloseAndRecv() (*SuccessResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SuccessResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *emailServiceClient) SendOrder(ctx context.Context, opts ...grpc.CallOption) (EmailService_SendOrderClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmailService_ServiceDesc.Streams[1], "/email_api.EmailService/SendOrder", opts...)
	if err != nil {
		return nil, err
	}
	x := &emailServiceSendOrderClient{stream}
	return x, nil
}

type EmailService_SendOrderClient interface {
	Send(*SendOrderRequest) error
	CloseAndRecv() (*SuccessResponse, error)
	grpc.ClientStream
}

type emailServiceSendOrderClient struct {
	grpc.ClientStream
}

func (x *emailServiceSendOrderClient) Send(m *SendOrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *emailServiceSendOrderClient) CloseAndRecv() (*SuccessResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SuccessResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *emailServiceClient) SendNotification(ctx context.Context, in *NotificationData, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/email_api.EmailService/SendNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendConfirm(ctx context.Context, in *ConfirmUserRequestOld, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/email_api.EmailService/SendConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) ConfirmUser(ctx context.Context, in *ConfirmUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/email_api.EmailService/ConfirmUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendReject(ctx context.Context, in *RejectUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/email_api.EmailService/SendReject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendJoin(ctx context.Context, in *JoinUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/email_api.EmailService/SendJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendBlocked(ctx context.Context, in *BlockedUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/email_api.EmailService/SendBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendTest(ctx context.Context, in *SendTestRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/email_api.EmailService/SendTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
// All implementations must embed UnimplementedEmailServiceServer
// for forward compatibility
type EmailServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Interview --------------------------------------------------
	SendInterview(EmailService_SendInterviewServer) error
	// Order --------------------------------------------------
	SendOrder(EmailService_SendOrderServer) error
	SendNotification(context.Context, *NotificationData) (*SuccessResponse, error)
	// Confirm --------------------------------------------------
	SendConfirm(context.Context, *ConfirmUserRequestOld) (*SuccessResponse, error)
	ConfirmUser(context.Context, *ConfirmUserRequest) (*SuccessResponse, error)
	// Reject --------------------------------------------------
	SendReject(context.Context, *RejectUserRequest) (*SuccessResponse, error)
	// Join --------------------------------------------------
	SendJoin(context.Context, *JoinUserRequest) (*SuccessResponse, error)
	// Blocked ------------------------------------------------
	SendBlocked(context.Context, *BlockedUserRequest) (*SuccessResponse, error)
	// TEST ---------------------------------------------------
	SendTest(context.Context, *SendTestRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedEmailServiceServer()
}

// UnimplementedEmailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (UnimplementedEmailServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedEmailServiceServer) SendInterview(EmailService_SendInterviewServer) error {
	return status.Errorf(codes.Unimplemented, "method SendInterview not implemented")
}
func (UnimplementedEmailServiceServer) SendOrder(EmailService_SendOrderServer) error {
	return status.Errorf(codes.Unimplemented, "method SendOrder not implemented")
}
func (UnimplementedEmailServiceServer) SendNotification(context.Context, *NotificationData) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (UnimplementedEmailServiceServer) SendConfirm(context.Context, *ConfirmUserRequestOld) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendConfirm not implemented")
}
func (UnimplementedEmailServiceServer) ConfirmUser(context.Context, *ConfirmUserRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmUser not implemented")
}
func (UnimplementedEmailServiceServer) SendReject(context.Context, *RejectUserRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReject not implemented")
}
func (UnimplementedEmailServiceServer) SendJoin(context.Context, *JoinUserRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendJoin not implemented")
}
func (UnimplementedEmailServiceServer) SendBlocked(context.Context, *BlockedUserRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBlocked not implemented")
}
func (UnimplementedEmailServiceServer) SendTest(context.Context, *SendTestRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTest not implemented")
}
func (UnimplementedEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {}

// UnsafeEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceServer will
// result in compilation errors.
type UnsafeEmailServiceServer interface {
	mustEmbedUnimplementedEmailServiceServer()
}

func RegisterEmailServiceServer(s grpc.ServiceRegistrar, srv EmailServiceServer) {
	s.RegisterService(&EmailService_ServiceDesc, srv)
}

func _EmailService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_api.EmailService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendInterview_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmailServiceServer).SendInterview(&emailServiceSendInterviewServer{stream})
}

type EmailService_SendInterviewServer interface {
	SendAndClose(*SuccessResponse) error
	Recv() (*SendInterviewRequest, error)
	grpc.ServerStream
}

type emailServiceSendInterviewServer struct {
	grpc.ServerStream
}

func (x *emailServiceSendInterviewServer) SendAndClose(m *SuccessResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *emailServiceSendInterviewServer) Recv() (*SendInterviewRequest, error) {
	m := new(SendInterviewRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EmailService_SendOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmailServiceServer).SendOrder(&emailServiceSendOrderServer{stream})
}

type EmailService_SendOrderServer interface {
	SendAndClose(*SuccessResponse) error
	Recv() (*SendOrderRequest, error)
	grpc.ServerStream
}

type emailServiceSendOrderServer struct {
	grpc.ServerStream
}

func (x *emailServiceSendOrderServer) SendAndClose(m *SuccessResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *emailServiceSendOrderServer) Recv() (*SendOrderRequest, error) {
	m := new(SendOrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EmailService_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_api.EmailService/SendNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendNotification(ctx, req.(*NotificationData))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmUserRequestOld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_api.EmailService/SendConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendConfirm(ctx, req.(*ConfirmUserRequestOld))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_ConfirmUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).ConfirmUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_api.EmailService/ConfirmUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).ConfirmUser(ctx, req.(*ConfirmUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendReject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendReject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_api.EmailService/SendReject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendReject(ctx, req.(*RejectUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_api.EmailService/SendJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendJoin(ctx, req.(*JoinUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockedUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_api.EmailService/SendBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendBlocked(ctx, req.(*BlockedUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_api.EmailService/SendTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendTest(ctx, req.(*SendTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailService_ServiceDesc is the grpc.ServiceDesc for EmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "email_api.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _EmailService_Ping_Handler,
		},
		{
			MethodName: "SendNotification",
			Handler:    _EmailService_SendNotification_Handler,
		},
		{
			MethodName: "SendConfirm",
			Handler:    _EmailService_SendConfirm_Handler,
		},
		{
			MethodName: "ConfirmUser",
			Handler:    _EmailService_ConfirmUser_Handler,
		},
		{
			MethodName: "SendReject",
			Handler:    _EmailService_SendReject_Handler,
		},
		{
			MethodName: "SendJoin",
			Handler:    _EmailService_SendJoin_Handler,
		},
		{
			MethodName: "SendBlocked",
			Handler:    _EmailService_SendBlocked_Handler,
		},
		{
			MethodName: "SendTest",
			Handler:    _EmailService_SendTest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendInterview",
			Handler:       _EmailService_SendInterview_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SendOrder",
			Handler:       _EmailService_SendOrder_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "email_api/email.proto",
}
