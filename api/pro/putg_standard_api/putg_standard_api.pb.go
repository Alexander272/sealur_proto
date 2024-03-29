// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/putg_standard_api.proto

package putg_standard_api

import (
	putg_standard_model "github.com/Alexander272/sealur_proto/api/pro/models/putg_standard_model"
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

type GetPutgStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPutgStandard) Reset() {
	*x = GetPutgStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_standard_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPutgStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPutgStandard) ProtoMessage() {}

func (x *GetPutgStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_standard_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPutgStandard.ProtoReflect.Descriptor instead.
func (*GetPutgStandard) Descriptor() ([]byte, []int) {
	return file_pro_putg_standard_api_proto_rawDescGZIP(), []int{0}
}

type CreatePutgStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlangeStandardId string `protobuf:"bytes,1,opt,name=flangeStandardId,proto3" json:"flangeStandardId,omitempty"`
	DnTitle          string `protobuf:"bytes,2,opt,name=dnTitle,proto3" json:"dnTitle,omitempty"`
	PnTitle          string `protobuf:"bytes,3,opt,name=pnTitle,proto3" json:"pnTitle,omitempty"`
	Count            int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CreatePutgStandard) Reset() {
	*x = CreatePutgStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_standard_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePutgStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePutgStandard) ProtoMessage() {}

func (x *CreatePutgStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_standard_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePutgStandard.ProtoReflect.Descriptor instead.
func (*CreatePutgStandard) Descriptor() ([]byte, []int) {
	return file_pro_putg_standard_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePutgStandard) GetFlangeStandardId() string {
	if x != nil {
		return x.FlangeStandardId
	}
	return ""
}

func (x *CreatePutgStandard) GetDnTitle() string {
	if x != nil {
		return x.DnTitle
	}
	return ""
}

func (x *CreatePutgStandard) GetPnTitle() string {
	if x != nil {
		return x.PnTitle
	}
	return ""
}

func (x *CreatePutgStandard) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdatePutgStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FlangeStandardId string `protobuf:"bytes,2,opt,name=flangeStandardId,proto3" json:"flangeStandardId,omitempty"`
	DnTitle          string `protobuf:"bytes,3,opt,name=dnTitle,proto3" json:"dnTitle,omitempty"`
	PnTitle          string `protobuf:"bytes,4,opt,name=pnTitle,proto3" json:"pnTitle,omitempty"`
	Count            int64  `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UpdatePutgStandard) Reset() {
	*x = UpdatePutgStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_standard_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePutgStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePutgStandard) ProtoMessage() {}

func (x *UpdatePutgStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_standard_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePutgStandard.ProtoReflect.Descriptor instead.
func (*UpdatePutgStandard) Descriptor() ([]byte, []int) {
	return file_pro_putg_standard_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePutgStandard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePutgStandard) GetFlangeStandardId() string {
	if x != nil {
		return x.FlangeStandardId
	}
	return ""
}

func (x *UpdatePutgStandard) GetDnTitle() string {
	if x != nil {
		return x.DnTitle
	}
	return ""
}

func (x *UpdatePutgStandard) GetPnTitle() string {
	if x != nil {
		return x.PnTitle
	}
	return ""
}

func (x *UpdatePutgStandard) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeletePutgStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePutgStandard) Reset() {
	*x = DeletePutgStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_standard_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePutgStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePutgStandard) ProtoMessage() {}

func (x *DeletePutgStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_standard_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePutgStandard.ProtoReflect.Descriptor instead.
func (*DeletePutgStandard) Descriptor() ([]byte, []int) {
	return file_pro_putg_standard_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePutgStandard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PutgStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Standards []*putg_standard_model.PutgStandard `protobuf:"bytes,1,rep,name=standards,proto3" json:"standards,omitempty"`
}

func (x *PutgStandard) Reset() {
	*x = PutgStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_standard_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutgStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutgStandard) ProtoMessage() {}

func (x *PutgStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_standard_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutgStandard.ProtoReflect.Descriptor instead.
func (*PutgStandard) Descriptor() ([]byte, []int) {
	return file_pro_putg_standard_api_proto_rawDescGZIP(), []int{4}
}

func (x *PutgStandard) GetStandards() []*putg_standard_model.PutgStandard {
	if x != nil {
		return x.Standards
	}
	return nil
}

var File_pro_putg_standard_api_proto protoreflect.FileDescriptor

var file_pro_putg_standard_api_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70,
	0x75, 0x74, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69,
	0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75,
	0x74, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x6c, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6e, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6e, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a,
	0x0a, 0x10, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6e,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6e, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75,
	0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x0c, 0x50, 0x75,
	0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x3f, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x75, 0x74, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x50, 0x75, 0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x32, 0xc2, 0x02, 0x0a, 0x13,
	0x50, 0x75, 0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x75, 0x74,
	0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x1a, 0x1f,
	0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x75, 0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12,
	0x49, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x70, 0x75, 0x74, 0x67,
	0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x75, 0x74, 0x67, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x1a, 0x18, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x25, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67,
	0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_putg_standard_api_proto_rawDescOnce sync.Once
	file_pro_putg_standard_api_proto_rawDescData = file_pro_putg_standard_api_proto_rawDesc
)

func file_pro_putg_standard_api_proto_rawDescGZIP() []byte {
	file_pro_putg_standard_api_proto_rawDescOnce.Do(func() {
		file_pro_putg_standard_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_putg_standard_api_proto_rawDescData)
	})
	return file_pro_putg_standard_api_proto_rawDescData
}

var file_pro_putg_standard_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_putg_standard_api_proto_goTypes = []interface{}{
	(*GetPutgStandard)(nil),                  // 0: putg_standard_api.GetPutgStandard
	(*CreatePutgStandard)(nil),               // 1: putg_standard_api.CreatePutgStandard
	(*UpdatePutgStandard)(nil),               // 2: putg_standard_api.UpdatePutgStandard
	(*DeletePutgStandard)(nil),               // 3: putg_standard_api.DeletePutgStandard
	(*PutgStandard)(nil),                     // 4: putg_standard_api.PutgStandard
	(*putg_standard_model.PutgStandard)(nil), // 5: putg_standard_model.PutgStandard
	(*response_model.Response)(nil),          // 6: response_model.Response
}
var file_pro_putg_standard_api_proto_depIdxs = []int32{
	5, // 0: putg_standard_api.PutgStandard.standards:type_name -> putg_standard_model.PutgStandard
	0, // 1: putg_standard_api.PutgStandardService.Get:input_type -> putg_standard_api.GetPutgStandard
	1, // 2: putg_standard_api.PutgStandardService.Create:input_type -> putg_standard_api.CreatePutgStandard
	2, // 3: putg_standard_api.PutgStandardService.Update:input_type -> putg_standard_api.UpdatePutgStandard
	3, // 4: putg_standard_api.PutgStandardService.Delete:input_type -> putg_standard_api.DeletePutgStandard
	4, // 5: putg_standard_api.PutgStandardService.Get:output_type -> putg_standard_api.PutgStandard
	6, // 6: putg_standard_api.PutgStandardService.Create:output_type -> response_model.Response
	6, // 7: putg_standard_api.PutgStandardService.Update:output_type -> response_model.Response
	6, // 8: putg_standard_api.PutgStandardService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_putg_standard_api_proto_init() }
func file_pro_putg_standard_api_proto_init() {
	if File_pro_putg_standard_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_putg_standard_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPutgStandard); i {
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
		file_pro_putg_standard_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePutgStandard); i {
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
		file_pro_putg_standard_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePutgStandard); i {
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
		file_pro_putg_standard_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePutgStandard); i {
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
		file_pro_putg_standard_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutgStandard); i {
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
			RawDescriptor: file_pro_putg_standard_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_putg_standard_api_proto_goTypes,
		DependencyIndexes: file_pro_putg_standard_api_proto_depIdxs,
		MessageInfos:      file_pro_putg_standard_api_proto_msgTypes,
	}.Build()
	File_pro_putg_standard_api_proto = out.File
	file_pro_putg_standard_api_proto_rawDesc = nil
	file_pro_putg_standard_api_proto_goTypes = nil
	file_pro_putg_standard_api_proto_depIdxs = nil
}
