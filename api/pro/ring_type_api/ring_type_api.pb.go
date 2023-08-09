// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/ring_type_api.proto

package ring_type_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	ring_type_model "github.com/Alexander272/sealur_proto/api/pro/models/ring_type_model"
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

type GetRingTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRingTypes) Reset() {
	*x = GetRingTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_type_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRingTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRingTypes) ProtoMessage() {}

func (x *GetRingTypes) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_type_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRingTypes.ProtoReflect.Descriptor instead.
func (*GetRingTypes) Descriptor() ([]byte, []int) {
	return file_pro_ring_type_api_proto_rawDescGZIP(), []int{0}
}

type CreateRingType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code          string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Description   string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	HasRotaryPlug bool   `protobuf:"varint,4,opt,name=hasRotaryPlug,proto3" json:"hasRotaryPlug,omitempty"`
	HasDensity    bool   `protobuf:"varint,5,opt,name=hasDensity,proto3" json:"hasDensity,omitempty"`
	HasThickness  bool   `protobuf:"varint,6,opt,name=hasThickness,proto3" json:"hasThickness,omitempty"`
	MaterialType  string `protobuf:"bytes,7,opt,name=materialType,proto3" json:"materialType,omitempty"`
}

func (x *CreateRingType) Reset() {
	*x = CreateRingType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_type_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRingType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRingType) ProtoMessage() {}

func (x *CreateRingType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_type_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRingType.ProtoReflect.Descriptor instead.
func (*CreateRingType) Descriptor() ([]byte, []int) {
	return file_pro_ring_type_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRingType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRingType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateRingType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRingType) GetHasRotaryPlug() bool {
	if x != nil {
		return x.HasRotaryPlug
	}
	return false
}

func (x *CreateRingType) GetHasDensity() bool {
	if x != nil {
		return x.HasDensity
	}
	return false
}

func (x *CreateRingType) GetHasThickness() bool {
	if x != nil {
		return x.HasThickness
	}
	return false
}

func (x *CreateRingType) GetMaterialType() string {
	if x != nil {
		return x.MaterialType
	}
	return ""
}

type UpdateRingType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code          string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description   string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	HasRotaryPlug bool   `protobuf:"varint,5,opt,name=hasRotaryPlug,proto3" json:"hasRotaryPlug,omitempty"`
	HasDensity    bool   `protobuf:"varint,6,opt,name=hasDensity,proto3" json:"hasDensity,omitempty"`
	HasThickness  bool   `protobuf:"varint,7,opt,name=hasThickness,proto3" json:"hasThickness,omitempty"`
	MaterialType  string `protobuf:"bytes,8,opt,name=materialType,proto3" json:"materialType,omitempty"`
}

func (x *UpdateRingType) Reset() {
	*x = UpdateRingType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_type_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRingType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRingType) ProtoMessage() {}

func (x *UpdateRingType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_type_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRingType.ProtoReflect.Descriptor instead.
func (*UpdateRingType) Descriptor() ([]byte, []int) {
	return file_pro_ring_type_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRingType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRingType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateRingType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRingType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateRingType) GetHasRotaryPlug() bool {
	if x != nil {
		return x.HasRotaryPlug
	}
	return false
}

func (x *UpdateRingType) GetHasDensity() bool {
	if x != nil {
		return x.HasDensity
	}
	return false
}

func (x *UpdateRingType) GetHasThickness() bool {
	if x != nil {
		return x.HasThickness
	}
	return false
}

func (x *UpdateRingType) GetMaterialType() string {
	if x != nil {
		return x.MaterialType
	}
	return ""
}

type DeleteRingType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRingType) Reset() {
	*x = DeleteRingType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_type_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRingType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRingType) ProtoMessage() {}

func (x *DeleteRingType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_type_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRingType.ProtoReflect.Descriptor instead.
func (*DeleteRingType) Descriptor() ([]byte, []int) {
	return file_pro_ring_type_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRingType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RingTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RingTypes []*ring_type_model.RingType `protobuf:"bytes,1,rep,name=ringTypes,proto3" json:"ringTypes,omitempty"`
}

func (x *RingTypes) Reset() {
	*x = RingTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_type_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RingTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RingTypes) ProtoMessage() {}

func (x *RingTypes) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_type_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RingTypes.ProtoReflect.Descriptor instead.
func (*RingTypes) Descriptor() ([]byte, []int) {
	return file_pro_ring_type_api_proto_rawDescGZIP(), []int{4}
}

func (x *RingTypes) GetRingTypes() []*ring_type_model.RingType {
	if x != nil {
		return x.RingTypes
	}
	return nil
}

var File_pro_ring_type_api_proto protoreflect.FileDescriptor

var file_pro_ring_type_api_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0xea, 0x01, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61,
	0x73, 0x52, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x50, 0x6c, 0x75, 0x67,
	0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79,
	0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x54, 0x68, 0x69, 0x63, 0x6b, 0x6e, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x54, 0x68, 0x69, 0x63, 0x6b,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x52, 0x6f,
	0x74, 0x61, 0x72, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x68, 0x61, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x12, 0x1e, 0x0a,
	0x0a, 0x68, 0x61, 0x73, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a,
	0x0c, 0x68, 0x61, 0x73, 0x54, 0x68, 0x69, 0x63, 0x6b, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x54, 0x68, 0x69, 0x63, 0x6b, 0x6e, 0x65, 0x73,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x09, 0x52, 0x69, 0x6e, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x32, 0x9e, 0x02,
	0x0a, 0x12, 0x53, 0x6e, 0x70, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1b,
	0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x18, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x6e, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17,
	0x5a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_ring_type_api_proto_rawDescOnce sync.Once
	file_pro_ring_type_api_proto_rawDescData = file_pro_ring_type_api_proto_rawDesc
)

func file_pro_ring_type_api_proto_rawDescGZIP() []byte {
	file_pro_ring_type_api_proto_rawDescOnce.Do(func() {
		file_pro_ring_type_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_ring_type_api_proto_rawDescData)
	})
	return file_pro_ring_type_api_proto_rawDescData
}

var file_pro_ring_type_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_ring_type_api_proto_goTypes = []interface{}{
	(*GetRingTypes)(nil),             // 0: ring_type_api.GetRingTypes
	(*CreateRingType)(nil),           // 1: ring_type_api.CreateRingType
	(*UpdateRingType)(nil),           // 2: ring_type_api.UpdateRingType
	(*DeleteRingType)(nil),           // 3: ring_type_api.DeleteRingType
	(*RingTypes)(nil),                // 4: ring_type_api.RingTypes
	(*ring_type_model.RingType)(nil), // 5: ring_type_model.RingType
	(*response_model.Response)(nil),  // 6: response_model.Response
}
var file_pro_ring_type_api_proto_depIdxs = []int32{
	5, // 0: ring_type_api.RingTypes.ringTypes:type_name -> ring_type_model.RingType
	0, // 1: ring_type_api.SnpMaterialService.GetAll:input_type -> ring_type_api.GetRingTypes
	1, // 2: ring_type_api.SnpMaterialService.Create:input_type -> ring_type_api.CreateRingType
	2, // 3: ring_type_api.SnpMaterialService.Update:input_type -> ring_type_api.UpdateRingType
	3, // 4: ring_type_api.SnpMaterialService.Delete:input_type -> ring_type_api.DeleteRingType
	4, // 5: ring_type_api.SnpMaterialService.GetAll:output_type -> ring_type_api.RingTypes
	6, // 6: ring_type_api.SnpMaterialService.Create:output_type -> response_model.Response
	6, // 7: ring_type_api.SnpMaterialService.Update:output_type -> response_model.Response
	6, // 8: ring_type_api.SnpMaterialService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_ring_type_api_proto_init() }
func file_pro_ring_type_api_proto_init() {
	if File_pro_ring_type_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_ring_type_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRingTypes); i {
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
		file_pro_ring_type_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRingType); i {
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
		file_pro_ring_type_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRingType); i {
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
		file_pro_ring_type_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRingType); i {
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
		file_pro_ring_type_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingTypes); i {
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
			RawDescriptor: file_pro_ring_type_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_ring_type_api_proto_goTypes,
		DependencyIndexes: file_pro_ring_type_api_proto_depIdxs,
		MessageInfos:      file_pro_ring_type_api_proto_msgTypes,
	}.Build()
	File_pro_ring_type_api_proto = out.File
	file_pro_ring_type_api_proto_rawDesc = nil
	file_pro_ring_type_api_proto_goTypes = nil
	file_pro_ring_type_api_proto_depIdxs = nil
}
