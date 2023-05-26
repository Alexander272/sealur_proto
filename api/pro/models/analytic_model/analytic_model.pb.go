// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/models/analytic_model.proto

package analytic_model

import (
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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Manager string     `protobuf:"bytes,2,opt,name=manager,proto3" json:"manager,omitempty"`
	Clients []*Clients `protobuf:"bytes,3,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_analytic_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_analytic_model_proto_msgTypes[0]
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
	return file_pro_models_analytic_model_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetManager() string {
	if x != nil {
		return x.Manager
	}
	return ""
}

func (x *Order) GetClients() []*Clients {
	if x != nil {
		return x.Clients
	}
	return nil
}

type OrderCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company    string `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	OrderCount int64  `protobuf:"varint,4,opt,name=orderCount,proto3" json:"orderCount,omitempty"`
}

func (x *OrderCount) Reset() {
	*x = OrderCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_analytic_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCount) ProtoMessage() {}

func (x *OrderCount) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_analytic_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCount.ProtoReflect.Descriptor instead.
func (*OrderCount) Descriptor() ([]byte, []int) {
	return file_pro_models_analytic_model_proto_rawDescGZIP(), []int{1}
}

func (x *OrderCount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderCount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderCount) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *OrderCount) GetOrderCount() int64 {
	if x != nil {
		return x.OrderCount
	}
	return 0
}

type Clients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OrdersCount      int64  `protobuf:"varint,3,opt,name=ordersCount,proto3" json:"ordersCount,omitempty"`
	PositionCount    int64  `protobuf:"varint,4,opt,name=positionCount,proto3" json:"positionCount,omitempty"`
	SnpPositionCount int64  `protobuf:"varint,5,opt,name=snpPositionCount,proto3" json:"snpPositionCount,omitempty"`
}

func (x *Clients) Reset() {
	*x = Clients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_analytic_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clients) ProtoMessage() {}

func (x *Clients) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_analytic_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clients.ProtoReflect.Descriptor instead.
func (*Clients) Descriptor() ([]byte, []int) {
	return file_pro_models_analytic_model_proto_rawDescGZIP(), []int{2}
}

func (x *Clients) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Clients) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Clients) GetOrdersCount() int64 {
	if x != nil {
		return x.OrdersCount
	}
	return 0
}

func (x *Clients) GetPositionCount() int64 {
	if x != nil {
		return x.PositionCount
	}
	return 0
}

func (x *Clients) GetSnpPositionCount() int64 {
	if x != nil {
		return x.SnpPositionCount
	}
	return 0
}

type FullOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Manager string    `protobuf:"bytes,2,opt,name=manager,proto3" json:"manager,omitempty"`
	Clients []*Client `protobuf:"bytes,3,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *FullOrder) Reset() {
	*x = FullOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_analytic_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullOrder) ProtoMessage() {}

func (x *FullOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_analytic_model_proto_msgTypes[3]
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
	return file_pro_models_analytic_model_proto_rawDescGZIP(), []int{3}
}

func (x *FullOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FullOrder) GetManager() string {
	if x != nil {
		return x.Manager
	}
	return ""
}

func (x *FullOrder) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Company string         `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	Name    string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Orders  []*ClientOrder `protobuf:"bytes,4,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_analytic_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_analytic_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_pro_models_analytic_model_proto_rawDescGZIP(), []int{4}
}

func (x *Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Client) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Client) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Client) GetOrders() []*ClientOrder {
	if x != nil {
		return x.Orders
	}
	return nil
}

type ClientOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Date   string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ClientOrder) Reset() {
	*x = ClientOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_analytic_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientOrder) ProtoMessage() {}

func (x *ClientOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_analytic_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientOrder.ProtoReflect.Descriptor instead.
func (*ClientOrder) Descriptor() ([]byte, []int) {
	return file_pro_models_analytic_model_proto_rawDescGZIP(), []int{5}
}

func (x *ClientOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientOrder) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *ClientOrder) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ClientOrder) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_pro_models_analytic_model_proto protoreflect.FileDescriptor

var file_pro_models_analytic_model_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x22, 0x64, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x07,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6a, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x73,
	0x6e, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x6e, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x67, 0x0a, 0x09, 0x46, 0x75, 0x6c, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x7b, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x61, 0x0a,
	0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x32, 0x37, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x6c,
	0x75, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_models_analytic_model_proto_rawDescOnce sync.Once
	file_pro_models_analytic_model_proto_rawDescData = file_pro_models_analytic_model_proto_rawDesc
)

func file_pro_models_analytic_model_proto_rawDescGZIP() []byte {
	file_pro_models_analytic_model_proto_rawDescOnce.Do(func() {
		file_pro_models_analytic_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_models_analytic_model_proto_rawDescData)
	})
	return file_pro_models_analytic_model_proto_rawDescData
}

var file_pro_models_analytic_model_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pro_models_analytic_model_proto_goTypes = []interface{}{
	(*Order)(nil),       // 0: analytic_model.Order
	(*OrderCount)(nil),  // 1: analytic_model.OrderCount
	(*Clients)(nil),     // 2: analytic_model.Clients
	(*FullOrder)(nil),   // 3: analytic_model.FullOrder
	(*Client)(nil),      // 4: analytic_model.Client
	(*ClientOrder)(nil), // 5: analytic_model.ClientOrder
}
var file_pro_models_analytic_model_proto_depIdxs = []int32{
	2, // 0: analytic_model.Order.clients:type_name -> analytic_model.Clients
	4, // 1: analytic_model.FullOrder.clients:type_name -> analytic_model.Client
	5, // 2: analytic_model.Client.orders:type_name -> analytic_model.ClientOrder
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pro_models_analytic_model_proto_init() }
func file_pro_models_analytic_model_proto_init() {
	if File_pro_models_analytic_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_models_analytic_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pro_models_analytic_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCount); i {
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
		file_pro_models_analytic_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clients); i {
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
		file_pro_models_analytic_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pro_models_analytic_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_pro_models_analytic_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientOrder); i {
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
			RawDescriptor: file_pro_models_analytic_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pro_models_analytic_model_proto_goTypes,
		DependencyIndexes: file_pro_models_analytic_model_proto_depIdxs,
		MessageInfos:      file_pro_models_analytic_model_proto_msgTypes,
	}.Build()
	File_pro_models_analytic_model_proto = out.File
	file_pro_models_analytic_model_proto_rawDesc = nil
	file_pro_models_analytic_model_proto_goTypes = nil
	file_pro_models_analytic_model_proto_depIdxs = nil
}
