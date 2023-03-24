// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pro/snp_filler_api.proto

package snp_filler_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	snp_filler_model "github.com/Alexander272/sealur_proto/api/pro/models/snp_filler_model"
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

type GetSnpFillers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string temperatureId = 1;
	StandardId string `protobuf:"bytes,1,opt,name=standardId,proto3" json:"standardId,omitempty"`
}

func (x *GetSnpFillers) Reset() {
	*x = GetSnpFillers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_filler_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnpFillers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnpFillers) ProtoMessage() {}

func (x *GetSnpFillers) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_filler_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnpFillers.ProtoReflect.Descriptor instead.
func (*GetSnpFillers) Descriptor() ([]byte, []int) {
	return file_pro_snp_filler_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetSnpFillers) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

type CreateSnpFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	AnotherTitle string `protobuf:"bytes,2,opt,name=anotherTitle,proto3" json:"anotherTitle,omitempty"`
	Code         string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Designation  string `protobuf:"bytes,5,opt,name=designation,proto3" json:"designation,omitempty"`
}

func (x *CreateSnpFiller) Reset() {
	*x = CreateSnpFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_filler_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnpFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnpFiller) ProtoMessage() {}

func (x *CreateSnpFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_filler_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnpFiller.ProtoReflect.Descriptor instead.
func (*CreateSnpFiller) Descriptor() ([]byte, []int) {
	return file_pro_snp_filler_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSnpFiller) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateSnpFiller) GetAnotherTitle() string {
	if x != nil {
		return x.AnotherTitle
	}
	return ""
}

func (x *CreateSnpFiller) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateSnpFiller) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateSnpFiller) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

type CreateSeveralSnpFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnpFillers []*CreateSnpFiller `protobuf:"bytes,1,rep,name=snpFillers,proto3" json:"snpFillers,omitempty"`
}

func (x *CreateSeveralSnpFiller) Reset() {
	*x = CreateSeveralSnpFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_filler_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeveralSnpFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeveralSnpFiller) ProtoMessage() {}

func (x *CreateSeveralSnpFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_filler_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeveralSnpFiller.ProtoReflect.Descriptor instead.
func (*CreateSeveralSnpFiller) Descriptor() ([]byte, []int) {
	return file_pro_snp_filler_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSeveralSnpFiller) GetSnpFillers() []*CreateSnpFiller {
	if x != nil {
		return x.SnpFillers
	}
	return nil
}

type UpdateSnpFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	AnotherTitle string `protobuf:"bytes,3,opt,name=anotherTitle,proto3" json:"anotherTitle,omitempty"`
	Code         string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Description  string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Designation  string `protobuf:"bytes,6,opt,name=designation,proto3" json:"designation,omitempty"`
}

func (x *UpdateSnpFiller) Reset() {
	*x = UpdateSnpFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_filler_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSnpFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnpFiller) ProtoMessage() {}

func (x *UpdateSnpFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_filler_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnpFiller.ProtoReflect.Descriptor instead.
func (*UpdateSnpFiller) Descriptor() ([]byte, []int) {
	return file_pro_snp_filler_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSnpFiller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSnpFiller) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateSnpFiller) GetAnotherTitle() string {
	if x != nil {
		return x.AnotherTitle
	}
	return ""
}

func (x *UpdateSnpFiller) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateSnpFiller) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateSnpFiller) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

type DeleteSnpFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSnpFiller) Reset() {
	*x = DeleteSnpFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_filler_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnpFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnpFiller) ProtoMessage() {}

func (x *DeleteSnpFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_filler_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnpFiller.ProtoReflect.Descriptor instead.
func (*DeleteSnpFiller) Descriptor() ([]byte, []int) {
	return file_pro_snp_filler_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSnpFiller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SnpFillers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnpFillers []*snp_filler_model.SnpFiller `protobuf:"bytes,1,rep,name=snpFillers,proto3" json:"snpFillers,omitempty"`
}

func (x *SnpFillers) Reset() {
	*x = SnpFillers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_filler_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnpFillers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnpFillers) ProtoMessage() {}

func (x *SnpFillers) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_filler_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnpFillers.ProtoReflect.Descriptor instead.
func (*SnpFillers) Descriptor() ([]byte, []int) {
	return file_pro_snp_filler_api_proto_rawDescGZIP(), []int{5}
}

func (x *SnpFillers) GetSnpFillers() []*snp_filler_model.SnpFiller {
	if x != nil {
		return x.SnpFillers
	}
	return nil
}

var File_pro_snp_filler_api_proto protoreflect.FileDescriptor

var file_pro_snp_filler_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x6e, 0x70, 0x5f,
	0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22,
	0xa3, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6e, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x3f, 0x0a, 0x0a, 0x73, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x46, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x73,
	0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x46, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6e,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0a, 0x53, 0x6e, 0x70,
	0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x6e, 0x70, 0x46, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6e,
	0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53,
	0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x6e, 0x70, 0x46, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x73, 0x32, 0xf6, 0x02, 0x0a, 0x10, 0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x1a,
	0x1a, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70,
	0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x51, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61,
	0x6c, 0x12, 0x26, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c,
	0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e,
	0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x46, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a,
	0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_snp_filler_api_proto_rawDescOnce sync.Once
	file_pro_snp_filler_api_proto_rawDescData = file_pro_snp_filler_api_proto_rawDesc
)

func file_pro_snp_filler_api_proto_rawDescGZIP() []byte {
	file_pro_snp_filler_api_proto_rawDescOnce.Do(func() {
		file_pro_snp_filler_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_snp_filler_api_proto_rawDescData)
	})
	return file_pro_snp_filler_api_proto_rawDescData
}

var file_pro_snp_filler_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pro_snp_filler_api_proto_goTypes = []interface{}{
	(*GetSnpFillers)(nil),              // 0: snp_filler_api.GetSnpFillers
	(*CreateSnpFiller)(nil),            // 1: snp_filler_api.CreateSnpFiller
	(*CreateSeveralSnpFiller)(nil),     // 2: snp_filler_api.CreateSeveralSnpFiller
	(*UpdateSnpFiller)(nil),            // 3: snp_filler_api.UpdateSnpFiller
	(*DeleteSnpFiller)(nil),            // 4: snp_filler_api.DeleteSnpFiller
	(*SnpFillers)(nil),                 // 5: snp_filler_api.SnpFillers
	(*snp_filler_model.SnpFiller)(nil), // 6: snp_filler_model.SnpFiller
	(*response_model.Response)(nil),    // 7: response_model.Response
}
var file_pro_snp_filler_api_proto_depIdxs = []int32{
	1, // 0: snp_filler_api.CreateSeveralSnpFiller.snpFillers:type_name -> snp_filler_api.CreateSnpFiller
	6, // 1: snp_filler_api.SnpFillers.snpFillers:type_name -> snp_filler_model.SnpFiller
	0, // 2: snp_filler_api.SnpFillerService.Get:input_type -> snp_filler_api.GetSnpFillers
	1, // 3: snp_filler_api.SnpFillerService.Create:input_type -> snp_filler_api.CreateSnpFiller
	2, // 4: snp_filler_api.SnpFillerService.CreateSeveral:input_type -> snp_filler_api.CreateSeveralSnpFiller
	3, // 5: snp_filler_api.SnpFillerService.Update:input_type -> snp_filler_api.UpdateSnpFiller
	4, // 6: snp_filler_api.SnpFillerService.Delete:input_type -> snp_filler_api.DeleteSnpFiller
	5, // 7: snp_filler_api.SnpFillerService.Get:output_type -> snp_filler_api.SnpFillers
	7, // 8: snp_filler_api.SnpFillerService.Create:output_type -> response_model.Response
	7, // 9: snp_filler_api.SnpFillerService.CreateSeveral:output_type -> response_model.Response
	7, // 10: snp_filler_api.SnpFillerService.Update:output_type -> response_model.Response
	7, // 11: snp_filler_api.SnpFillerService.Delete:output_type -> response_model.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pro_snp_filler_api_proto_init() }
func file_pro_snp_filler_api_proto_init() {
	if File_pro_snp_filler_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_snp_filler_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnpFillers); i {
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
		file_pro_snp_filler_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnpFiller); i {
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
		file_pro_snp_filler_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeveralSnpFiller); i {
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
		file_pro_snp_filler_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSnpFiller); i {
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
		file_pro_snp_filler_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnpFiller); i {
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
		file_pro_snp_filler_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnpFillers); i {
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
			RawDescriptor: file_pro_snp_filler_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_snp_filler_api_proto_goTypes,
		DependencyIndexes: file_pro_snp_filler_api_proto_depIdxs,
		MessageInfos:      file_pro_snp_filler_api_proto_msgTypes,
	}.Build()
	File_pro_snp_filler_api_proto = out.File
	file_pro_snp_filler_api_proto_rawDesc = nil
	file_pro_snp_filler_api_proto_goTypes = nil
	file_pro_snp_filler_api_proto_depIdxs = nil
}
