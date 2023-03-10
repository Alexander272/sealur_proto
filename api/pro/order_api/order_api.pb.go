// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pro/order_api.proto

package order_api

import (
	order_model "github.com/Alexander272/sealur_proto/api/pro/models/order_model"
	position_model "github.com/Alexander272/sealur_proto/api/pro/models/position_model"
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrder) Reset() {
	*x = GetOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrder) ProtoMessage() {}

func (x *GetOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrder.ProtoReflect.Descriptor instead.
func (*GetOrder) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetAllOrders) Reset() {
	*x = GetAllOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrders) ProtoMessage() {}

func (x *GetAllOrders) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrders.ProtoReflect.Descriptor instead.
func (*GetAllOrders) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllOrders) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string                         `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Count     string                         `protobuf:"bytes,3,opt,name=count,proto3" json:"count,omitempty"`
	Positions []*position_model.FullPosition `protobuf:"bytes,4,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *CreateOrder) Reset() {
	*x = CreateOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrder) ProtoMessage() {}

func (x *CreateOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrder.ProtoReflect.Descriptor instead.
func (*CreateOrder) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrder) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrder) GetCount() string {
	if x != nil {
		return x.Count
	}
	return ""
}

func (x *CreateOrder) GetPositions() []*position_model.FullPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

type DeleteOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteOrder) Reset() {
	*x = DeleteOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrder) ProtoMessage() {}

func (x *DeleteOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrder.ProtoReflect.Descriptor instead.
func (*DeleteOrder) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*order_model.Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{4}
}

func (x *Orders) GetOrders() []*order_model.Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Positions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*position_model.FullPosition `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *Positions) Reset() {
	*x = Positions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Positions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Positions) ProtoMessage() {}

func (x *Positions) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Positions.ProtoReflect.Descriptor instead.
func (*Positions) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{5}
}

func (x *Positions) GetPositions() []*position_model.FullPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

type OrderNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *OrderNumber) Reset() {
	*x = OrderNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderNumber) ProtoMessage() {}

func (x *OrderNumber) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderNumber.ProtoReflect.Descriptor instead.
func (*OrderNumber) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{6}
}

func (x *OrderNumber) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_pro_order_api_proto protoreflect.FileDescriptor

var file_pro_order_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1d,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x47, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3a, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x25, 0x0a, 0x0b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x32, 0xa8, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x1a, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46,
	0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x11, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x36,
	0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x16,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13,
	0x5a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_order_api_proto_rawDescOnce sync.Once
	file_pro_order_api_proto_rawDescData = file_pro_order_api_proto_rawDesc
)

func file_pro_order_api_proto_rawDescGZIP() []byte {
	file_pro_order_api_proto_rawDescOnce.Do(func() {
		file_pro_order_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_order_api_proto_rawDescData)
	})
	return file_pro_order_api_proto_rawDescData
}

var file_pro_order_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pro_order_api_proto_goTypes = []interface{}{
	(*GetOrder)(nil),                    // 0: order_api.GetOrder
	(*GetAllOrders)(nil),                // 1: order_api.GetAllOrders
	(*CreateOrder)(nil),                 // 2: order_api.CreateOrder
	(*DeleteOrder)(nil),                 // 3: order_api.DeleteOrder
	(*Orders)(nil),                      // 4: order_api.Orders
	(*Positions)(nil),                   // 5: order_api.Positions
	(*OrderNumber)(nil),                 // 6: order_api.OrderNumber
	(*position_model.FullPosition)(nil), // 7: position_model.FullPosition
	(*order_model.Order)(nil),           // 8: order_model.Order
	(*order_model.FullOrder)(nil),       // 9: order_model.FullOrder
	(*response_model.Response)(nil),     // 10: response_model.Response
}
var file_pro_order_api_proto_depIdxs = []int32{
	7,  // 0: order_api.CreateOrder.positions:type_name -> position_model.FullPosition
	8,  // 1: order_api.Orders.orders:type_name -> order_model.Order
	7,  // 2: order_api.Positions.positions:type_name -> position_model.FullPosition
	0,  // 3: order_api.OrderService.Get:input_type -> order_api.GetOrder
	1,  // 4: order_api.OrderService.GetAll:input_type -> order_api.GetAllOrders
	2,  // 5: order_api.OrderService.Save:input_type -> order_api.CreateOrder
	2,  // 6: order_api.OrderService.Create:input_type -> order_api.CreateOrder
	3,  // 7: order_api.OrderService.Delete:input_type -> order_api.DeleteOrder
	9,  // 8: order_api.OrderService.Get:output_type -> order_model.FullOrder
	4,  // 9: order_api.OrderService.GetAll:output_type -> order_api.Orders
	6,  // 10: order_api.OrderService.Save:output_type -> order_api.OrderNumber
	10, // 11: order_api.OrderService.Create:output_type -> response_model.Response
	10, // 12: order_api.OrderService.Delete:output_type -> response_model.Response
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pro_order_api_proto_init() }
func file_pro_order_api_proto_init() {
	if File_pro_order_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_order_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pro_order_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOrders); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pro_order_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pro_order_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pro_order_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pro_order_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Positions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pro_order_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderNumber); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pro_order_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_order_api_proto_goTypes,
		DependencyIndexes: file_pro_order_api_proto_depIdxs,
		MessageInfos:      file_pro_order_api_proto_msgTypes,
	}.Build()
	File_pro_order_api_proto = out.File
	file_pro_order_api_proto_rawDesc = nil
	file_pro_order_api_proto_goTypes = nil
	file_pro_order_api_proto_depIdxs = nil
}
