// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/ring_construction_api.proto

package ring_construction_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	ring_construction_model "github.com/Alexander272/sealur_proto/api/pro/models/ring_construction_model"
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

type GetRingConstructions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRingConstructions) Reset() {
	*x = GetRingConstructions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_construction_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRingConstructions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRingConstructions) ProtoMessage() {}

func (x *GetRingConstructions) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_construction_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRingConstructions.ProtoReflect.Descriptor instead.
func (*GetRingConstructions) Descriptor() ([]byte, []int) {
	return file_pro_ring_construction_api_proto_rawDescGZIP(), []int{0}
}

type CreateRingConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId      string `protobuf:"bytes,1,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Count       int64  `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CreateRingConstruction) Reset() {
	*x = CreateRingConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_construction_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRingConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRingConstruction) ProtoMessage() {}

func (x *CreateRingConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_construction_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRingConstruction.ProtoReflect.Descriptor instead.
func (*CreateRingConstruction) Descriptor() ([]byte, []int) {
	return file_pro_ring_construction_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRingConstruction) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *CreateRingConstruction) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRingConstruction) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateRingConstruction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRingConstruction) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateRingConstruction) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateRingConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId      string `protobuf:"bytes,2,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Code        string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Image       string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Count       int64  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UpdateRingConstruction) Reset() {
	*x = UpdateRingConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_construction_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRingConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRingConstruction) ProtoMessage() {}

func (x *UpdateRingConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_construction_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRingConstruction.ProtoReflect.Descriptor instead.
func (*UpdateRingConstruction) Descriptor() ([]byte, []int) {
	return file_pro_ring_construction_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRingConstruction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRingConstruction) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *UpdateRingConstruction) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateRingConstruction) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRingConstruction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateRingConstruction) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UpdateRingConstruction) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeleteRingConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRingConstruction) Reset() {
	*x = DeleteRingConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_construction_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRingConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRingConstruction) ProtoMessage() {}

func (x *DeleteRingConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_construction_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRingConstruction.ProtoReflect.Descriptor instead.
func (*DeleteRingConstruction) Descriptor() ([]byte, []int) {
	return file_pro_ring_construction_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRingConstruction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RingConstructions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repeated ring_construction_model.RingConstruction ringConstruction = 1;
	RingConstruction *ring_construction_model.RingConstructionMap `protobuf:"bytes,1,opt,name=ringConstruction,proto3" json:"ringConstruction,omitempty"`
}

func (x *RingConstructions) Reset() {
	*x = RingConstructions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_construction_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RingConstructions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RingConstructions) ProtoMessage() {}

func (x *RingConstructions) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_construction_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RingConstructions.ProtoReflect.Descriptor instead.
func (*RingConstructions) Descriptor() ([]byte, []int) {
	return file_pro_ring_construction_api_proto_rawDescGZIP(), []int{4}
}

func (x *RingConstructions) GetRingConstruction() *ring_construction_model.RingConstructionMap {
	if x != nil {
		return x.RingConstruction
	}
	return nil
}

var File_pro_ring_construction_api_proto protoreflect.FileDescriptor

var file_pro_ring_construction_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x11, 0x52,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x58, 0x0a, 0x10, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x52, 0x10, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xee, 0x02, 0x0a, 0x12, 0x53,
	0x6e, 0x70, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x2b, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x28, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x51, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x2d, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x2d, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_ring_construction_api_proto_rawDescOnce sync.Once
	file_pro_ring_construction_api_proto_rawDescData = file_pro_ring_construction_api_proto_rawDesc
)

func file_pro_ring_construction_api_proto_rawDescGZIP() []byte {
	file_pro_ring_construction_api_proto_rawDescOnce.Do(func() {
		file_pro_ring_construction_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_ring_construction_api_proto_rawDescData)
	})
	return file_pro_ring_construction_api_proto_rawDescData
}

var file_pro_ring_construction_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_ring_construction_api_proto_goTypes = []interface{}{
	(*GetRingConstructions)(nil),                        // 0: ring_construction_api.GetRingConstructions
	(*CreateRingConstruction)(nil),                      // 1: ring_construction_api.CreateRingConstruction
	(*UpdateRingConstruction)(nil),                      // 2: ring_construction_api.UpdateRingConstruction
	(*DeleteRingConstruction)(nil),                      // 3: ring_construction_api.DeleteRingConstruction
	(*RingConstructions)(nil),                           // 4: ring_construction_api.RingConstructions
	(*ring_construction_model.RingConstructionMap)(nil), // 5: ring_construction_model.RingConstructionMap
	(*response_model.Response)(nil),                     // 6: response_model.Response
}
var file_pro_ring_construction_api_proto_depIdxs = []int32{
	5, // 0: ring_construction_api.RingConstructions.ringConstruction:type_name -> ring_construction_model.RingConstructionMap
	0, // 1: ring_construction_api.SnpMaterialService.GetAll:input_type -> ring_construction_api.GetRingConstructions
	1, // 2: ring_construction_api.SnpMaterialService.Create:input_type -> ring_construction_api.CreateRingConstruction
	2, // 3: ring_construction_api.SnpMaterialService.Update:input_type -> ring_construction_api.UpdateRingConstruction
	3, // 4: ring_construction_api.SnpMaterialService.Delete:input_type -> ring_construction_api.DeleteRingConstruction
	4, // 5: ring_construction_api.SnpMaterialService.GetAll:output_type -> ring_construction_api.RingConstructions
	6, // 6: ring_construction_api.SnpMaterialService.Create:output_type -> response_model.Response
	6, // 7: ring_construction_api.SnpMaterialService.Update:output_type -> response_model.Response
	6, // 8: ring_construction_api.SnpMaterialService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_ring_construction_api_proto_init() }
func file_pro_ring_construction_api_proto_init() {
	if File_pro_ring_construction_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_ring_construction_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRingConstructions); i {
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
		file_pro_ring_construction_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRingConstruction); i {
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
		file_pro_ring_construction_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRingConstruction); i {
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
		file_pro_ring_construction_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRingConstruction); i {
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
		file_pro_ring_construction_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingConstructions); i {
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
			RawDescriptor: file_pro_ring_construction_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_ring_construction_api_proto_goTypes,
		DependencyIndexes: file_pro_ring_construction_api_proto_depIdxs,
		MessageInfos:      file_pro_ring_construction_api_proto_msgTypes,
	}.Build()
	File_pro_ring_construction_api_proto = out.File
	file_pro_ring_construction_api_proto_rawDesc = nil
	file_pro_ring_construction_api_proto_goTypes = nil
	file_pro_ring_construction_api_proto_depIdxs = nil
}
