// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/ring_density_api.proto

package ring_density_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	ring_density_model "github.com/Alexander272/sealur_proto/api/pro/models/ring_density_model"
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

type GetRingDensity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRingDensity) Reset() {
	*x = GetRingDensity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_density_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRingDensity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRingDensity) ProtoMessage() {}

func (x *GetRingDensity) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_density_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRingDensity.ProtoReflect.Descriptor instead.
func (*GetRingDensity) Descriptor() ([]byte, []int) {
	return file_pro_ring_density_api_proto_rawDescGZIP(), []int{0}
}

type CreateRingDensity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId        string `protobuf:"bytes,1,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code          string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description   string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	HasRotaryPlug bool   `protobuf:"varint,5,opt,name=hasRotaryPlug,proto3" json:"hasRotaryPlug,omitempty"`
}

func (x *CreateRingDensity) Reset() {
	*x = CreateRingDensity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_density_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRingDensity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRingDensity) ProtoMessage() {}

func (x *CreateRingDensity) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_density_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRingDensity.ProtoReflect.Descriptor instead.
func (*CreateRingDensity) Descriptor() ([]byte, []int) {
	return file_pro_ring_density_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRingDensity) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *CreateRingDensity) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRingDensity) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateRingDensity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRingDensity) GetHasRotaryPlug() bool {
	if x != nil {
		return x.HasRotaryPlug
	}
	return false
}

type UpdateRingDensity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId        string `protobuf:"bytes,2,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Code          string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	HasRotaryPlug bool   `protobuf:"varint,6,opt,name=hasRotaryPlug,proto3" json:"hasRotaryPlug,omitempty"`
}

func (x *UpdateRingDensity) Reset() {
	*x = UpdateRingDensity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_density_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRingDensity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRingDensity) ProtoMessage() {}

func (x *UpdateRingDensity) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_density_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRingDensity.ProtoReflect.Descriptor instead.
func (*UpdateRingDensity) Descriptor() ([]byte, []int) {
	return file_pro_ring_density_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRingDensity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRingDensity) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *UpdateRingDensity) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateRingDensity) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRingDensity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateRingDensity) GetHasRotaryPlug() bool {
	if x != nil {
		return x.HasRotaryPlug
	}
	return false
}

type DeleteRingDensity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRingDensity) Reset() {
	*x = DeleteRingDensity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_density_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRingDensity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRingDensity) ProtoMessage() {}

func (x *DeleteRingDensity) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_density_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRingDensity.ProtoReflect.Descriptor instead.
func (*DeleteRingDensity) Descriptor() ([]byte, []int) {
	return file_pro_ring_density_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRingDensity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RingDensity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repeated ring_density_model.RingDensity density = 1;
	Density *ring_density_model.RingDensityMap `protobuf:"bytes,1,opt,name=density,proto3" json:"density,omitempty"`
}

func (x *RingDensity) Reset() {
	*x = RingDensity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_density_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RingDensity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RingDensity) ProtoMessage() {}

func (x *RingDensity) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_density_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RingDensity.ProtoReflect.Descriptor instead.
func (*RingDensity) Descriptor() ([]byte, []int) {
	return file_pro_ring_density_api_proto_rawDescGZIP(), []int{4}
}

func (x *RingDensity) GetDensity() *ring_density_model.RingDensityMap {
	if x != nil {
		return x.Density
	}
	return nil
}

var File_pro_ring_density_api_proto protoreflect.FileDescriptor

var file_pro_ring_density_api_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f,
	0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x50, 0x6c, 0x75,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x52, 0x6f, 0x74, 0x61,
	0x72, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x22, 0xad, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x50, 0x6c, 0x75,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x52, 0x6f, 0x74, 0x61,
	0x72, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x0b, 0x52,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x07, 0x64, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x52, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x70, 0x52,
	0x07, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x32, 0xba, 0x02, 0x0a, 0x12, 0x52, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x49, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x79, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_ring_density_api_proto_rawDescOnce sync.Once
	file_pro_ring_density_api_proto_rawDescData = file_pro_ring_density_api_proto_rawDesc
)

func file_pro_ring_density_api_proto_rawDescGZIP() []byte {
	file_pro_ring_density_api_proto_rawDescOnce.Do(func() {
		file_pro_ring_density_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_ring_density_api_proto_rawDescData)
	})
	return file_pro_ring_density_api_proto_rawDescData
}

var file_pro_ring_density_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_ring_density_api_proto_goTypes = []interface{}{
	(*GetRingDensity)(nil),                    // 0: ring_density_api.GetRingDensity
	(*CreateRingDensity)(nil),                 // 1: ring_density_api.CreateRingDensity
	(*UpdateRingDensity)(nil),                 // 2: ring_density_api.UpdateRingDensity
	(*DeleteRingDensity)(nil),                 // 3: ring_density_api.DeleteRingDensity
	(*RingDensity)(nil),                       // 4: ring_density_api.RingDensity
	(*ring_density_model.RingDensityMap)(nil), // 5: ring_density_model.RingDensityMap
	(*response_model.Response)(nil),           // 6: response_model.Response
}
var file_pro_ring_density_api_proto_depIdxs = []int32{
	5, // 0: ring_density_api.RingDensity.density:type_name -> ring_density_model.RingDensityMap
	0, // 1: ring_density_api.RingDensityService.GetAll:input_type -> ring_density_api.GetRingDensity
	1, // 2: ring_density_api.RingDensityService.Create:input_type -> ring_density_api.CreateRingDensity
	2, // 3: ring_density_api.RingDensityService.Update:input_type -> ring_density_api.UpdateRingDensity
	3, // 4: ring_density_api.RingDensityService.Delete:input_type -> ring_density_api.DeleteRingDensity
	4, // 5: ring_density_api.RingDensityService.GetAll:output_type -> ring_density_api.RingDensity
	6, // 6: ring_density_api.RingDensityService.Create:output_type -> response_model.Response
	6, // 7: ring_density_api.RingDensityService.Update:output_type -> response_model.Response
	6, // 8: ring_density_api.RingDensityService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_ring_density_api_proto_init() }
func file_pro_ring_density_api_proto_init() {
	if File_pro_ring_density_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_ring_density_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRingDensity); i {
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
		file_pro_ring_density_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRingDensity); i {
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
		file_pro_ring_density_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRingDensity); i {
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
		file_pro_ring_density_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRingDensity); i {
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
		file_pro_ring_density_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingDensity); i {
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
			RawDescriptor: file_pro_ring_density_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_ring_density_api_proto_goTypes,
		DependencyIndexes: file_pro_ring_density_api_proto_depIdxs,
		MessageInfos:      file_pro_ring_density_api_proto_msgTypes,
	}.Build()
	File_pro_ring_density_api_proto = out.File
	file_pro_ring_density_api_proto_rawDesc = nil
	file_pro_ring_density_api_proto_goTypes = nil
	file_pro_ring_density_api_proto_depIdxs = nil
}
