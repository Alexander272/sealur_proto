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
	user_model "github.com/Alexander272/sealur_proto/api/user/models/user_model"
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

type GetCurrentOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetCurrentOrder) Reset() {
	*x = GetCurrentOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentOrder) ProtoMessage() {}

func (x *GetCurrentOrder) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCurrentOrder.ProtoReflect.Descriptor instead.
func (*GetCurrentOrder) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrentOrder) GetUserId() string {
	if x != nil {
		return x.UserId
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
		mi := &file_pro_order_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrders) ProtoMessage() {}

func (x *GetAllOrders) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllOrders.ProtoReflect.Descriptor instead.
func (*GetAllOrders) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllOrders) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetManagerOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerId string `protobuf:"bytes,1,opt,name=managerId,proto3" json:"managerId,omitempty"`
}

func (x *GetManagerOrders) Reset() {
	*x = GetManagerOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManagerOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManagerOrders) ProtoMessage() {}

func (x *GetManagerOrders) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetManagerOrders.ProtoReflect.Descriptor instead.
func (*GetManagerOrders) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetManagerOrders) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

type GetAllManagerOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllManagerOrders) Reset() {
	*x = GetAllManagerOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllManagerOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllManagerOrders) ProtoMessage() {}

func (x *GetAllManagerOrders) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllManagerOrders.ProtoReflect.Descriptor instead.
func (*GetAllManagerOrders) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{4}
}

type CreateOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string                         `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Count     int64                          `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Positions []*position_model.FullPosition `protobuf:"bytes,4,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *CreateOrder) Reset() {
	*x = CreateOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrder) ProtoMessage() {}

func (x *CreateOrder) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateOrder.ProtoReflect.Descriptor instead.
func (*CreateOrder) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{5}
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

func (x *CreateOrder) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
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
		mi := &file_pro_order_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrder) ProtoMessage() {}

func (x *DeleteOrder) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteOrder.ProtoReflect.Descriptor instead.
func (*DeleteOrder) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CopyOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	FromId   string `protobuf:"bytes,2,opt,name=fromId,proto3" json:"fromId,omitempty"`
	Count    int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CopyOrder) Reset() {
	*x = CopyOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyOrder) ProtoMessage() {}

func (x *CopyOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyOrder.ProtoReflect.Descriptor instead.
func (*CopyOrder) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{7}
}

func (x *CopyOrder) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *CopyOrder) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *CopyOrder) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  order_model.OrderStatus `protobuf:"varint,1,opt,name=status,proto3,enum=order_model.OrderStatus" json:"status,omitempty"`
	OrderId string                  `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Date    string                  `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{8}
}

func (x *Status) GetStatus() order_model.OrderStatus {
	if x != nil {
		return x.Status
	}
	return order_model.OrderStatus(0)
}

func (x *Status) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Status) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *user_model.User          `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Order *order_model.CurrentOrder `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[9]
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
	return file_pro_order_api_proto_rawDescGZIP(), []int{9}
}

func (x *Order) GetUser() *user_model.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Order) GetOrder() *order_model.CurrentOrder {
	if x != nil {
		return x.Order
	}
	return nil
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
		mi := &file_pro_order_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[10]
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
	return file_pro_order_api_proto_rawDescGZIP(), []int{10}
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

	Positions []*position_model.Position `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *Positions) Reset() {
	*x = Positions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Positions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Positions) ProtoMessage() {}

func (x *Positions) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[11]
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
	return file_pro_order_api_proto_rawDescGZIP(), []int{11}
}

func (x *Positions) GetPositions() []*position_model.Position {
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
		mi := &file_pro_order_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderNumber) ProtoMessage() {}

func (x *OrderNumber) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[12]
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
	return file_pro_order_api_proto_rawDescGZIP(), []int{12}
}

func (x *OrderNumber) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ManagerOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*order_model.ManagerOrder `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ManagerOrders) Reset() {
	*x = ManagerOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_order_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerOrders) ProtoMessage() {}

func (x *ManagerOrders) ProtoReflect() protoreflect.Message {
	mi := &file_pro_order_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerOrders.ProtoReflect.Descriptor instead.
func (*ManagerOrders) Descriptor() ([]byte, []int) {
	return file_pro_order_api_proto_rawDescGZIP(), []int{13}
}

func (x *ManagerOrders) GetOrders() []*order_model.ManagerOrder {
	if x != nil {
		return x.Orders
	}
	return nil
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
	0x1a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x1d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x55, 0x0a, 0x09, 0x43, 0x6f, 0x70, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x5e, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x2f, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x34, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x43, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x42, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0xa5, 0x05, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a,
	0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x11, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1b, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x6e,
	0x12, 0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x1a, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x53, 0x61,
	0x76, 0x65, 0x12, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x04, 0x43, 0x6f, 0x70, 0x79, 0x12, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
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

var file_pro_order_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pro_order_api_proto_goTypes = []interface{}{
	(*GetOrder)(nil),                    // 0: order_api.GetOrder
	(*GetCurrentOrder)(nil),             // 1: order_api.GetCurrentOrder
	(*GetAllOrders)(nil),                // 2: order_api.GetAllOrders
	(*GetManagerOrders)(nil),            // 3: order_api.GetManagerOrders
	(*GetAllManagerOrders)(nil),         // 4: order_api.GetAllManagerOrders
	(*CreateOrder)(nil),                 // 5: order_api.CreateOrder
	(*DeleteOrder)(nil),                 // 6: order_api.DeleteOrder
	(*CopyOrder)(nil),                   // 7: order_api.CopyOrder
	(*Status)(nil),                      // 8: order_api.Status
	(*Order)(nil),                       // 9: order_api.Order
	(*Orders)(nil),                      // 10: order_api.Orders
	(*Positions)(nil),                   // 11: order_api.Positions
	(*OrderNumber)(nil),                 // 12: order_api.OrderNumber
	(*ManagerOrders)(nil),               // 13: order_api.ManagerOrders
	(*position_model.FullPosition)(nil), // 14: position_model.FullPosition
	(order_model.OrderStatus)(0),        // 15: order_model.OrderStatus
	(*user_model.User)(nil),             // 16: user_model.User
	(*order_model.CurrentOrder)(nil),    // 17: order_model.CurrentOrder
	(*order_model.Order)(nil),           // 18: order_model.Order
	(*position_model.Position)(nil),     // 19: position_model.Position
	(*order_model.ManagerOrder)(nil),    // 20: order_model.ManagerOrder
	(*response_model.FileResponse)(nil), // 21: response_model.FileResponse
	(*response_model.IdResponse)(nil),   // 22: response_model.IdResponse
	(*response_model.Response)(nil),     // 23: response_model.Response
}
var file_pro_order_api_proto_depIdxs = []int32{
	14, // 0: order_api.CreateOrder.positions:type_name -> position_model.FullPosition
	15, // 1: order_api.Status.status:type_name -> order_model.OrderStatus
	16, // 2: order_api.Order.user:type_name -> user_model.User
	17, // 3: order_api.Order.order:type_name -> order_model.CurrentOrder
	18, // 4: order_api.Orders.orders:type_name -> order_model.Order
	19, // 5: order_api.Positions.positions:type_name -> position_model.Position
	20, // 6: order_api.ManagerOrders.orders:type_name -> order_model.ManagerOrder
	0,  // 7: order_api.OrderService.Get:input_type -> order_api.GetOrder
	1,  // 8: order_api.OrderService.GetCurrent:input_type -> order_api.GetCurrentOrder
	2,  // 9: order_api.OrderService.GetAll:input_type -> order_api.GetAllOrders
	0,  // 10: order_api.OrderService.GetFile:input_type -> order_api.GetOrder
	3,  // 11: order_api.OrderService.GetOpen:input_type -> order_api.GetManagerOrders
	4,  // 12: order_api.OrderService.GetAllOpen:input_type -> order_api.GetAllManagerOrders
	5,  // 13: order_api.OrderService.Save:input_type -> order_api.CreateOrder
	5,  // 14: order_api.OrderService.Create:input_type -> order_api.CreateOrder
	7,  // 15: order_api.OrderService.Copy:input_type -> order_api.CopyOrder
	8,  // 16: order_api.OrderService.SetStatus:input_type -> order_api.Status
	6,  // 17: order_api.OrderService.Delete:input_type -> order_api.DeleteOrder
	9,  // 18: order_api.OrderService.Get:output_type -> order_api.Order
	17, // 19: order_api.OrderService.GetCurrent:output_type -> order_model.CurrentOrder
	10, // 20: order_api.OrderService.GetAll:output_type -> order_api.Orders
	21, // 21: order_api.OrderService.GetFile:output_type -> response_model.FileResponse
	13, // 22: order_api.OrderService.GetOpen:output_type -> order_api.ManagerOrders
	13, // 23: order_api.OrderService.GetAllOpen:output_type -> order_api.ManagerOrders
	12, // 24: order_api.OrderService.Save:output_type -> order_api.OrderNumber
	22, // 25: order_api.OrderService.Create:output_type -> response_model.IdResponse
	23, // 26: order_api.OrderService.Copy:output_type -> response_model.Response
	23, // 27: order_api.OrderService.SetStatus:output_type -> response_model.Response
	23, // 28: order_api.OrderService.Delete:output_type -> response_model.Response
	18, // [18:29] is the sub-list for method output_type
	7,  // [7:18] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
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
			switch v := v.(*GetCurrentOrder); i {
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
		file_pro_order_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManagerOrders); i {
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
			switch v := v.(*GetAllManagerOrders); i {
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
		file_pro_order_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_pro_order_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyOrder); i {
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
		file_pro_order_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_pro_order_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_pro_order_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
		file_pro_order_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
		file_pro_order_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
		file_pro_order_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerOrders); i {
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
			NumMessages:   14,
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
