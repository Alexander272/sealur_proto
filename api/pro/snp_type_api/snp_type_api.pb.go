// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pro/snp_type_api.proto

package snp_type_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	snp_type_model "github.com/Alexander272/sealur_proto/api/pro/models/snp_type_model"
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

type GetSnpTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StandardId string `protobuf:"bytes,1,opt,name=standardId,proto3" json:"standardId,omitempty"`
}

func (x *GetSnpTypes) Reset() {
	*x = GetSnpTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_type_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnpTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnpTypes) ProtoMessage() {}

func (x *GetSnpTypes) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_type_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnpTypes.ProtoReflect.Descriptor instead.
func (*GetSnpTypes) Descriptor() ([]byte, []int) {
	return file_pro_snp_type_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetSnpTypes) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

type CreateSnpType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	FlangeTypeId string `protobuf:"bytes,2,opt,name=flangeTypeId,proto3" json:"flangeTypeId,omitempty"`
	Code         string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	StandardId   string `protobuf:"bytes,4,opt,name=standardId,proto3" json:"standardId,omitempty"`
	Description  string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	HasD4        bool   `protobuf:"varint,6,opt,name=hasD4,proto3" json:"hasD4,omitempty"`
	HasD3        bool   `protobuf:"varint,7,opt,name=hasD3,proto3" json:"hasD3,omitempty"`
	HasD2        bool   `protobuf:"varint,8,opt,name=hasD2,proto3" json:"hasD2,omitempty"`
	HasD1        bool   `protobuf:"varint,9,opt,name=hasD1,proto3" json:"hasD1,omitempty"`
}

func (x *CreateSnpType) Reset() {
	*x = CreateSnpType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_type_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnpType) ProtoMessage() {}

func (x *CreateSnpType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_type_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnpType.ProtoReflect.Descriptor instead.
func (*CreateSnpType) Descriptor() ([]byte, []int) {
	return file_pro_snp_type_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSnpType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateSnpType) GetFlangeTypeId() string {
	if x != nil {
		return x.FlangeTypeId
	}
	return ""
}

func (x *CreateSnpType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateSnpType) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

func (x *CreateSnpType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateSnpType) GetHasD4() bool {
	if x != nil {
		return x.HasD4
	}
	return false
}

func (x *CreateSnpType) GetHasD3() bool {
	if x != nil {
		return x.HasD3
	}
	return false
}

func (x *CreateSnpType) GetHasD2() bool {
	if x != nil {
		return x.HasD2
	}
	return false
}

func (x *CreateSnpType) GetHasD1() bool {
	if x != nil {
		return x.HasD1
	}
	return false
}

type CreateSeveralSnpType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnpTypes []*CreateSnpType `protobuf:"bytes,1,rep,name=snpTypes,proto3" json:"snpTypes,omitempty"`
}

func (x *CreateSeveralSnpType) Reset() {
	*x = CreateSeveralSnpType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_type_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeveralSnpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeveralSnpType) ProtoMessage() {}

func (x *CreateSeveralSnpType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_type_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeveralSnpType.ProtoReflect.Descriptor instead.
func (*CreateSeveralSnpType) Descriptor() ([]byte, []int) {
	return file_pro_snp_type_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSeveralSnpType) GetSnpTypes() []*CreateSnpType {
	if x != nil {
		return x.SnpTypes
	}
	return nil
}

type UpdateSnpType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	FlangeTypeId string `protobuf:"bytes,3,opt,name=flangeTypeId,proto3" json:"flangeTypeId,omitempty"`
	Code         string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	StandardId   string `protobuf:"bytes,5,opt,name=standardId,proto3" json:"standardId,omitempty"`
	Description  string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	HasD4        bool   `protobuf:"varint,7,opt,name=hasD4,proto3" json:"hasD4,omitempty"`
	HasD3        bool   `protobuf:"varint,8,opt,name=hasD3,proto3" json:"hasD3,omitempty"`
	HasD2        bool   `protobuf:"varint,9,opt,name=hasD2,proto3" json:"hasD2,omitempty"`
	HasD1        bool   `protobuf:"varint,10,opt,name=hasD1,proto3" json:"hasD1,omitempty"`
}

func (x *UpdateSnpType) Reset() {
	*x = UpdateSnpType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_type_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSnpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnpType) ProtoMessage() {}

func (x *UpdateSnpType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_type_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnpType.ProtoReflect.Descriptor instead.
func (*UpdateSnpType) Descriptor() ([]byte, []int) {
	return file_pro_snp_type_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSnpType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSnpType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateSnpType) GetFlangeTypeId() string {
	if x != nil {
		return x.FlangeTypeId
	}
	return ""
}

func (x *UpdateSnpType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateSnpType) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

func (x *UpdateSnpType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateSnpType) GetHasD4() bool {
	if x != nil {
		return x.HasD4
	}
	return false
}

func (x *UpdateSnpType) GetHasD3() bool {
	if x != nil {
		return x.HasD3
	}
	return false
}

func (x *UpdateSnpType) GetHasD2() bool {
	if x != nil {
		return x.HasD2
	}
	return false
}

func (x *UpdateSnpType) GetHasD1() bool {
	if x != nil {
		return x.HasD1
	}
	return false
}

type DeleteSnpType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSnpType) Reset() {
	*x = DeleteSnpType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_type_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnpType) ProtoMessage() {}

func (x *DeleteSnpType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_type_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnpType.ProtoReflect.Descriptor instead.
func (*DeleteSnpType) Descriptor() ([]byte, []int) {
	return file_pro_snp_type_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSnpType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SnpTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnpTypes []*snp_type_model.SnpType `protobuf:"bytes,1,rep,name=snpTypes,proto3" json:"snpTypes,omitempty"`
}

func (x *SnpTypes) Reset() {
	*x = SnpTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_type_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnpTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnpTypes) ProtoMessage() {}

func (x *SnpTypes) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_type_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnpTypes.ProtoReflect.Descriptor instead.
func (*SnpTypes) Descriptor() ([]byte, []int) {
	return file_pro_snp_type_api_proto_rawDescGZIP(), []int{5}
}

func (x *SnpTypes) GetSnpTypes() []*snp_type_model.SnpType {
	if x != nil {
		return x.SnpTypes
	}
	return nil
}

var File_pro_snp_type_api_proto protoreflect.FileDescriptor

var file_pro_snp_type_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53,
	0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0xf7, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x73,
	0x44, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x61, 0x73, 0x44, 0x34, 0x12,
	0x14, 0x0a, 0x05, 0x68, 0x61, 0x73, 0x44, 0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x68, 0x61, 0x73, 0x44, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x73, 0x44, 0x32, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x61, 0x73, 0x44, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x68,
	0x61, 0x73, 0x44, 0x31, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x61, 0x73, 0x44,
	0x31, 0x22, 0x4f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x53, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x6e, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6e,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x6e, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x22, 0x87, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6c,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x73, 0x44, 0x34, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x61, 0x73, 0x44, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61,
	0x73, 0x44, 0x33, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x61, 0x73, 0x44, 0x33,
	0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x73, 0x44, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x68, 0x61, 0x73, 0x44, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x73, 0x44, 0x31, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x61, 0x73, 0x44, 0x31, 0x22, 0x1f, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a,
	0x08, 0x53, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x6e, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6e,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x6e, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x73, 0x32, 0xdc,
	0x02, 0x0a, 0x0e, 0x53, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x22, 0x2e,
	0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x6e, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a,
	0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_snp_type_api_proto_rawDescOnce sync.Once
	file_pro_snp_type_api_proto_rawDescData = file_pro_snp_type_api_proto_rawDesc
)

func file_pro_snp_type_api_proto_rawDescGZIP() []byte {
	file_pro_snp_type_api_proto_rawDescOnce.Do(func() {
		file_pro_snp_type_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_snp_type_api_proto_rawDescData)
	})
	return file_pro_snp_type_api_proto_rawDescData
}

var file_pro_snp_type_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pro_snp_type_api_proto_goTypes = []interface{}{
	(*GetSnpTypes)(nil),             // 0: snp_type_api.GetSnpTypes
	(*CreateSnpType)(nil),           // 1: snp_type_api.CreateSnpType
	(*CreateSeveralSnpType)(nil),    // 2: snp_type_api.CreateSeveralSnpType
	(*UpdateSnpType)(nil),           // 3: snp_type_api.UpdateSnpType
	(*DeleteSnpType)(nil),           // 4: snp_type_api.DeleteSnpType
	(*SnpTypes)(nil),                // 5: snp_type_api.SnpTypes
	(*snp_type_model.SnpType)(nil),  // 6: snp_type_model.SnpType
	(*response_model.Response)(nil), // 7: response_model.Response
}
var file_pro_snp_type_api_proto_depIdxs = []int32{
	1, // 0: snp_type_api.CreateSeveralSnpType.snpTypes:type_name -> snp_type_api.CreateSnpType
	6, // 1: snp_type_api.SnpTypes.snpTypes:type_name -> snp_type_model.SnpType
	0, // 2: snp_type_api.SnpTypeService.Get:input_type -> snp_type_api.GetSnpTypes
	1, // 3: snp_type_api.SnpTypeService.Create:input_type -> snp_type_api.CreateSnpType
	2, // 4: snp_type_api.SnpTypeService.CreateSeveral:input_type -> snp_type_api.CreateSeveralSnpType
	3, // 5: snp_type_api.SnpTypeService.Update:input_type -> snp_type_api.UpdateSnpType
	4, // 6: snp_type_api.SnpTypeService.Delete:input_type -> snp_type_api.DeleteSnpType
	5, // 7: snp_type_api.SnpTypeService.Get:output_type -> snp_type_api.SnpTypes
	7, // 8: snp_type_api.SnpTypeService.Create:output_type -> response_model.Response
	7, // 9: snp_type_api.SnpTypeService.CreateSeveral:output_type -> response_model.Response
	7, // 10: snp_type_api.SnpTypeService.Update:output_type -> response_model.Response
	7, // 11: snp_type_api.SnpTypeService.Delete:output_type -> response_model.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pro_snp_type_api_proto_init() }
func file_pro_snp_type_api_proto_init() {
	if File_pro_snp_type_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_snp_type_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnpTypes); i {
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
		file_pro_snp_type_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnpType); i {
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
		file_pro_snp_type_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeveralSnpType); i {
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
		file_pro_snp_type_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSnpType); i {
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
		file_pro_snp_type_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnpType); i {
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
		file_pro_snp_type_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnpTypes); i {
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
			RawDescriptor: file_pro_snp_type_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_snp_type_api_proto_goTypes,
		DependencyIndexes: file_pro_snp_type_api_proto_depIdxs,
		MessageInfos:      file_pro_snp_type_api_proto_msgTypes,
	}.Build()
	File_pro_snp_type_api_proto = out.File
	file_pro_snp_type_api_proto_rawDesc = nil
	file_pro_snp_type_api_proto_goTypes = nil
	file_pro_snp_type_api_proto_depIdxs = nil
}
