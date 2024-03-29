// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/models/order_model.proto

package order_model

import (
	position_model "github.com/Alexander272/sealur_proto/api/pro/models/position_model"
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

type OrderStatus int32

const (
	OrderStatus_new    OrderStatus = 0
	OrderStatus_work   OrderStatus = 1
	OrderStatus_finish OrderStatus = 2
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "new",
		1: "work",
		2: "finish",
	}
	OrderStatus_value = map[string]int32{
		"new":    0,
		"work":   1,
		"finish": 2,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pro_models_order_model_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_pro_models_order_model_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_pro_models_order_model_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date          string                     `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	CountPosition int64                      `protobuf:"varint,3,opt,name=countPosition,proto3" json:"countPosition,omitempty"`
	Number        int64                      `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	Info          string                     `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	Positions     []*position_model.Position `protobuf:"bytes,6,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_order_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_order_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_pro_models_order_model_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Order) GetCountPosition() int64 {
	if x != nil {
		return x.CountPosition
	}
	return 0
}

func (x *Order) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Order) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *Order) GetPositions() []*position_model.Position {
	if x != nil {
		return x.Positions
	}
	return nil
}

type FullOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date          string                         `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	CountPosition int64                          `protobuf:"varint,3,opt,name=countPosition,proto3" json:"countPosition,omitempty"`
	Number        int64                          `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	Positions     []*position_model.FullPosition `protobuf:"bytes,5,rep,name=positions,proto3" json:"positions,omitempty"`
	UserId        string                         `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
	Info          string                         `protobuf:"bytes,7,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *FullOrder) Reset() {
	*x = FullOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_order_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullOrder) ProtoMessage() {}

func (x *FullOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_order_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullOrder.ProtoReflect.Descriptor instead.
func (*FullOrder) Descriptor() ([]byte, []int) {
	return file_pro_models_order_model_proto_rawDescGZIP(), []int{1}
}

func (x *FullOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FullOrder) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *FullOrder) GetCountPosition() int64 {
	if x != nil {
		return x.CountPosition
	}
	return 0
}

func (x *FullOrder) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *FullOrder) GetPositions() []*position_model.FullPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

func (x *FullOrder) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FullOrder) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type CurrentOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number    int64                           `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Positions []*position_model.OrderPosition `protobuf:"bytes,3,rep,name=positions,proto3" json:"positions,omitempty"`
	UserId    string                          `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Info      string                          `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CurrentOrder) Reset() {
	*x = CurrentOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_order_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentOrder) ProtoMessage() {}

func (x *CurrentOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_order_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentOrder.ProtoReflect.Descriptor instead.
func (*CurrentOrder) Descriptor() ([]byte, []int) {
	return file_pro_models_order_model_proto_rawDescGZIP(), []int{2}
}

func (x *CurrentOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CurrentOrder) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *CurrentOrder) GetPositions() []*position_model.OrderPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

func (x *CurrentOrder) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CurrentOrder) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type ManagerOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date          string      `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	CountPosition int64       `protobuf:"varint,3,opt,name=countPosition,proto3" json:"countPosition,omitempty"`
	Number        int64       `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	Status        OrderStatus `protobuf:"varint,5,opt,name=status,proto3,enum=order_model.OrderStatus" json:"status,omitempty"`
	UserId        string      `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
	Company       string      `protobuf:"bytes,7,opt,name=company,proto3" json:"company,omitempty"`
	ManagerId     string      `protobuf:"bytes,8,opt,name=managerId,proto3" json:"managerId,omitempty"`
}

func (x *ManagerOrder) Reset() {
	*x = ManagerOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_order_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerOrder) ProtoMessage() {}

func (x *ManagerOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_order_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerOrder.ProtoReflect.Descriptor instead.
func (*ManagerOrder) Descriptor() ([]byte, []int) {
	return file_pro_models_order_model_proto_rawDescGZIP(), []int{3}
}

func (x *ManagerOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ManagerOrder) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ManagerOrder) GetCountPosition() int64 {
	if x != nil {
		return x.CountPosition
	}
	return 0
}

func (x *ManagerOrder) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ManagerOrder) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_new
}

func (x *ManagerOrder) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ManagerOrder) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *ManagerOrder) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

var File_pro_models_order_model_proto protoreflect.FileDescriptor

var file_pro_models_order_model_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x70, 0x72, 0x6f,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x09,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd5, 0x01, 0x0a, 0x09, 0x46, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x9f, 0x01, 0x0a,
	0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xf2,
	0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x49, 0x64, 0x2a, 0x2c, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x77,
	0x6f, 0x72, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x10,
	0x02, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x32, 0x37, 0x32, 0x2f, 0x73, 0x65, 0x61,
	0x6c, 0x75, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_models_order_model_proto_rawDescOnce sync.Once
	file_pro_models_order_model_proto_rawDescData = file_pro_models_order_model_proto_rawDesc
)

func file_pro_models_order_model_proto_rawDescGZIP() []byte {
	file_pro_models_order_model_proto_rawDescOnce.Do(func() {
		file_pro_models_order_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_models_order_model_proto_rawDescData)
	})
	return file_pro_models_order_model_proto_rawDescData
}

var file_pro_models_order_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pro_models_order_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pro_models_order_model_proto_goTypes = []interface{}{
	(OrderStatus)(0),                     // 0: order_model.OrderStatus
	(*Order)(nil),                        // 1: order_model.Order
	(*FullOrder)(nil),                    // 2: order_model.FullOrder
	(*CurrentOrder)(nil),                 // 3: order_model.CurrentOrder
	(*ManagerOrder)(nil),                 // 4: order_model.ManagerOrder
	(*position_model.Position)(nil),      // 5: position_model.Position
	(*position_model.FullPosition)(nil),  // 6: position_model.FullPosition
	(*position_model.OrderPosition)(nil), // 7: position_model.OrderPosition
}
var file_pro_models_order_model_proto_depIdxs = []int32{
	5, // 0: order_model.Order.positions:type_name -> position_model.Position
	6, // 1: order_model.FullOrder.positions:type_name -> position_model.FullPosition
	7, // 2: order_model.CurrentOrder.positions:type_name -> position_model.OrderPosition
	0, // 3: order_model.ManagerOrder.status:type_name -> order_model.OrderStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pro_models_order_model_proto_init() }
func file_pro_models_order_model_proto_init() {
	if File_pro_models_order_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_models_order_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_pro_models_order_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullOrder); i {
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
		file_pro_models_order_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentOrder); i {
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
		file_pro_models_order_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerOrder); i {
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
			RawDescriptor: file_pro_models_order_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pro_models_order_model_proto_goTypes,
		DependencyIndexes: file_pro_models_order_model_proto_depIdxs,
		EnumInfos:         file_pro_models_order_model_proto_enumTypes,
		MessageInfos:      file_pro_models_order_model_proto_msgTypes,
	}.Build()
	File_pro_models_order_model_proto = out.File
	file_pro_models_order_model_proto_rawDesc = nil
	file_pro_models_order_model_proto_goTypes = nil
	file_pro_models_order_model_proto_depIdxs = nil
}
