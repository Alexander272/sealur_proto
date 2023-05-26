// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/snp_size_api.proto

package snp_size_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	snp_size_model "github.com/Alexander272/sealur_proto/api/pro/models/snp_size_model"
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

type GetSnpSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string standardId = 1;
	// string flangeId = 2;
	TypeId string `protobuf:"bytes,1,opt,name=typeId,proto3" json:"typeId,omitempty"`
	HasD2  bool   `protobuf:"varint,2,opt,name=hasD2,proto3" json:"hasD2,omitempty"`
}

func (x *GetSnpSize) Reset() {
	*x = GetSnpSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_size_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnpSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnpSize) ProtoMessage() {}

func (x *GetSnpSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_size_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnpSize.ProtoReflect.Descriptor instead.
func (*GetSnpSize) Descriptor() ([]byte, []int) {
	return file_pro_snp_size_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetSnpSize) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *GetSnpSize) GetHasD2() bool {
	if x != nil {
		return x.HasD2
	}
	return false
}

type CreateSnpSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnpTypeId string               `protobuf:"bytes,1,opt,name=snpTypeId,proto3" json:"snpTypeId,omitempty"`
	Count     int64                `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Dn        string               `protobuf:"bytes,3,opt,name=dn,proto3" json:"dn,omitempty"`
	DnMm      string               `protobuf:"bytes,4,opt,name=dnMm,proto3" json:"dnMm,omitempty"`
	Pn        []*snp_size_model.Pn `protobuf:"bytes,5,rep,name=pn,proto3" json:"pn,omitempty"`
	D4        string               `protobuf:"bytes,6,opt,name=d4,proto3" json:"d4,omitempty"`
	D3        string               `protobuf:"bytes,7,opt,name=d3,proto3" json:"d3,omitempty"`
	D2        string               `protobuf:"bytes,8,opt,name=d2,proto3" json:"d2,omitempty"`
	D1        string               `protobuf:"bytes,9,opt,name=d1,proto3" json:"d1,omitempty"`
	H         []string             `protobuf:"bytes,10,rep,name=h,proto3" json:"h,omitempty"`
	S2        []string             `protobuf:"bytes,11,rep,name=s2,proto3" json:"s2,omitempty"`
	S3        []string             `protobuf:"bytes,12,rep,name=s3,proto3" json:"s3,omitempty"`
}

func (x *CreateSnpSize) Reset() {
	*x = CreateSnpSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_size_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnpSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnpSize) ProtoMessage() {}

func (x *CreateSnpSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_size_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnpSize.ProtoReflect.Descriptor instead.
func (*CreateSnpSize) Descriptor() ([]byte, []int) {
	return file_pro_snp_size_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSnpSize) GetSnpTypeId() string {
	if x != nil {
		return x.SnpTypeId
	}
	return ""
}

func (x *CreateSnpSize) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CreateSnpSize) GetDn() string {
	if x != nil {
		return x.Dn
	}
	return ""
}

func (x *CreateSnpSize) GetDnMm() string {
	if x != nil {
		return x.DnMm
	}
	return ""
}

func (x *CreateSnpSize) GetPn() []*snp_size_model.Pn {
	if x != nil {
		return x.Pn
	}
	return nil
}

func (x *CreateSnpSize) GetD4() string {
	if x != nil {
		return x.D4
	}
	return ""
}

func (x *CreateSnpSize) GetD3() string {
	if x != nil {
		return x.D3
	}
	return ""
}

func (x *CreateSnpSize) GetD2() string {
	if x != nil {
		return x.D2
	}
	return ""
}

func (x *CreateSnpSize) GetD1() string {
	if x != nil {
		return x.D1
	}
	return ""
}

func (x *CreateSnpSize) GetH() []string {
	if x != nil {
		return x.H
	}
	return nil
}

func (x *CreateSnpSize) GetS2() []string {
	if x != nil {
		return x.S2
	}
	return nil
}

func (x *CreateSnpSize) GetS3() []string {
	if x != nil {
		return x.S3
	}
	return nil
}

type CreateSeveralSnpSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sizes []*CreateSnpSize `protobuf:"bytes,1,rep,name=sizes,proto3" json:"sizes,omitempty"`
}

func (x *CreateSeveralSnpSize) Reset() {
	*x = CreateSeveralSnpSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_size_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeveralSnpSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeveralSnpSize) ProtoMessage() {}

func (x *CreateSeveralSnpSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_size_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeveralSnpSize.ProtoReflect.Descriptor instead.
func (*CreateSeveralSnpSize) Descriptor() ([]byte, []int) {
	return file_pro_snp_size_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSeveralSnpSize) GetSizes() []*CreateSnpSize {
	if x != nil {
		return x.Sizes
	}
	return nil
}

type UpdateSnpSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SnpTypeId string               `protobuf:"bytes,2,opt,name=snpTypeId,proto3" json:"snpTypeId,omitempty"`
	Count     int64                `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Dn        string               `protobuf:"bytes,4,opt,name=dn,proto3" json:"dn,omitempty"`
	DnMm      string               `protobuf:"bytes,5,opt,name=dnMm,proto3" json:"dnMm,omitempty"`
	Pn        []*snp_size_model.Pn `protobuf:"bytes,6,rep,name=pn,proto3" json:"pn,omitempty"`
	D4        string               `protobuf:"bytes,7,opt,name=d4,proto3" json:"d4,omitempty"`
	D3        string               `protobuf:"bytes,8,opt,name=d3,proto3" json:"d3,omitempty"`
	D2        string               `protobuf:"bytes,9,opt,name=d2,proto3" json:"d2,omitempty"`
	D1        string               `protobuf:"bytes,10,opt,name=d1,proto3" json:"d1,omitempty"`
	H         []string             `protobuf:"bytes,11,rep,name=h,proto3" json:"h,omitempty"`
	S2        []string             `protobuf:"bytes,12,rep,name=s2,proto3" json:"s2,omitempty"`
	S3        []string             `protobuf:"bytes,13,rep,name=s3,proto3" json:"s3,omitempty"`
}

func (x *UpdateSnpSize) Reset() {
	*x = UpdateSnpSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_size_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSnpSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnpSize) ProtoMessage() {}

func (x *UpdateSnpSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_size_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnpSize.ProtoReflect.Descriptor instead.
func (*UpdateSnpSize) Descriptor() ([]byte, []int) {
	return file_pro_snp_size_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSnpSize) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSnpSize) GetSnpTypeId() string {
	if x != nil {
		return x.SnpTypeId
	}
	return ""
}

func (x *UpdateSnpSize) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *UpdateSnpSize) GetDn() string {
	if x != nil {
		return x.Dn
	}
	return ""
}

func (x *UpdateSnpSize) GetDnMm() string {
	if x != nil {
		return x.DnMm
	}
	return ""
}

func (x *UpdateSnpSize) GetPn() []*snp_size_model.Pn {
	if x != nil {
		return x.Pn
	}
	return nil
}

func (x *UpdateSnpSize) GetD4() string {
	if x != nil {
		return x.D4
	}
	return ""
}

func (x *UpdateSnpSize) GetD3() string {
	if x != nil {
		return x.D3
	}
	return ""
}

func (x *UpdateSnpSize) GetD2() string {
	if x != nil {
		return x.D2
	}
	return ""
}

func (x *UpdateSnpSize) GetD1() string {
	if x != nil {
		return x.D1
	}
	return ""
}

func (x *UpdateSnpSize) GetH() []string {
	if x != nil {
		return x.H
	}
	return nil
}

func (x *UpdateSnpSize) GetS2() []string {
	if x != nil {
		return x.S2
	}
	return nil
}

func (x *UpdateSnpSize) GetS3() []string {
	if x != nil {
		return x.S3
	}
	return nil
}

type DeleteSnpSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSnpSize) Reset() {
	*x = DeleteSnpSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_size_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnpSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnpSize) ProtoMessage() {}

func (x *DeleteSnpSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_size_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnpSize.ProtoReflect.Descriptor instead.
func (*DeleteSnpSize) Descriptor() ([]byte, []int) {
	return file_pro_snp_size_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSnpSize) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SnpSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sizes []*snp_size_model.SnpSize `protobuf:"bytes,1,rep,name=sizes,proto3" json:"sizes,omitempty"`
}

func (x *SnpSize) Reset() {
	*x = SnpSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_size_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnpSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnpSize) ProtoMessage() {}

func (x *SnpSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_size_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnpSize.ProtoReflect.Descriptor instead.
func (*SnpSize) Descriptor() ([]byte, []int) {
	return file_pro_snp_size_api_proto_rawDescGZIP(), []int{5}
}

func (x *SnpSize) GetSizes() []*snp_size_model.SnpSize {
	if x != nil {
		return x.Sizes
	}
	return nil
}

var File_pro_snp_size_api_proto protoreflect.FileDescriptor

var file_pro_snp_size_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53,
	0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x68, 0x61, 0x73, 0x44, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68,
	0x61, 0x73, 0x44, 0x32, 0x22, 0xf9, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6e, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6e, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6e,
	0x4d, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x6e, 0x4d, 0x6d, 0x12, 0x22,
	0x0a, 0x02, 0x70, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6e, 0x70,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6e, 0x52, 0x02,
	0x70, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x64, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x64, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x64, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x31, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x64, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x01, 0x68,
	0x12, 0x0e, 0x0a, 0x02, 0x73, 0x32, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x73, 0x32,
	0x12, 0x0e, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x73, 0x33,
	0x22, 0x49, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61,
	0x6c, 0x53, 0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x69, 0x7a, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70,
	0x53, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6e, 0x4d, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x6e, 0x4d, 0x6d, 0x12, 0x22, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x50, 0x6e, 0x52, 0x02, 0x70, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x34, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x33, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x32, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x31, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x01, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x32, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x02, 0x73, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x02, 0x73, 0x33, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x07, 0x53, 0x6e, 0x70, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x53, 0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x73, 0x69, 0x7a,
	0x65, 0x73, 0x32, 0xda, 0x02, 0x0a, 0x0e, 0x53, 0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73,
	0x6e, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6e, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3f, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70,
	0x53, 0x69, 0x7a, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x12,
	0x22, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x6e, 0x70, 0x53,
	0x69, 0x7a, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x70,
	0x53, 0x69, 0x7a, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e,
	0x70, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x16, 0x5a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_snp_size_api_proto_rawDescOnce sync.Once
	file_pro_snp_size_api_proto_rawDescData = file_pro_snp_size_api_proto_rawDesc
)

func file_pro_snp_size_api_proto_rawDescGZIP() []byte {
	file_pro_snp_size_api_proto_rawDescOnce.Do(func() {
		file_pro_snp_size_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_snp_size_api_proto_rawDescData)
	})
	return file_pro_snp_size_api_proto_rawDescData
}

var file_pro_snp_size_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pro_snp_size_api_proto_goTypes = []interface{}{
	(*GetSnpSize)(nil),              // 0: snp_size_api.GetSnpSize
	(*CreateSnpSize)(nil),           // 1: snp_size_api.CreateSnpSize
	(*CreateSeveralSnpSize)(nil),    // 2: snp_size_api.CreateSeveralSnpSize
	(*UpdateSnpSize)(nil),           // 3: snp_size_api.UpdateSnpSize
	(*DeleteSnpSize)(nil),           // 4: snp_size_api.DeleteSnpSize
	(*SnpSize)(nil),                 // 5: snp_size_api.SnpSize
	(*snp_size_model.Pn)(nil),       // 6: snp_size_model.Pn
	(*snp_size_model.SnpSize)(nil),  // 7: snp_size_model.SnpSize
	(*response_model.Response)(nil), // 8: response_model.Response
}
var file_pro_snp_size_api_proto_depIdxs = []int32{
	6, // 0: snp_size_api.CreateSnpSize.pn:type_name -> snp_size_model.Pn
	1, // 1: snp_size_api.CreateSeveralSnpSize.sizes:type_name -> snp_size_api.CreateSnpSize
	6, // 2: snp_size_api.UpdateSnpSize.pn:type_name -> snp_size_model.Pn
	7, // 3: snp_size_api.SnpSize.sizes:type_name -> snp_size_model.SnpSize
	0, // 4: snp_size_api.SnpSizeService.Get:input_type -> snp_size_api.GetSnpSize
	1, // 5: snp_size_api.SnpSizeService.Create:input_type -> snp_size_api.CreateSnpSize
	2, // 6: snp_size_api.SnpSizeService.CreateSeveral:input_type -> snp_size_api.CreateSeveralSnpSize
	3, // 7: snp_size_api.SnpSizeService.Update:input_type -> snp_size_api.UpdateSnpSize
	4, // 8: snp_size_api.SnpSizeService.Delete:input_type -> snp_size_api.DeleteSnpSize
	5, // 9: snp_size_api.SnpSizeService.Get:output_type -> snp_size_api.SnpSize
	8, // 10: snp_size_api.SnpSizeService.Create:output_type -> response_model.Response
	8, // 11: snp_size_api.SnpSizeService.CreateSeveral:output_type -> response_model.Response
	8, // 12: snp_size_api.SnpSizeService.Update:output_type -> response_model.Response
	8, // 13: snp_size_api.SnpSizeService.Delete:output_type -> response_model.Response
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pro_snp_size_api_proto_init() }
func file_pro_snp_size_api_proto_init() {
	if File_pro_snp_size_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_snp_size_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnpSize); i {
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
		file_pro_snp_size_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnpSize); i {
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
		file_pro_snp_size_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeveralSnpSize); i {
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
		file_pro_snp_size_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSnpSize); i {
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
		file_pro_snp_size_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnpSize); i {
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
		file_pro_snp_size_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnpSize); i {
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
			RawDescriptor: file_pro_snp_size_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_snp_size_api_proto_goTypes,
		DependencyIndexes: file_pro_snp_size_api_proto_depIdxs,
		MessageInfos:      file_pro_snp_size_api_proto_msgTypes,
	}.Build()
	File_pro_snp_size_api_proto = out.File
	file_pro_snp_size_api_proto_rawDesc = nil
	file_pro_snp_size_api_proto_goTypes = nil
	file_pro_snp_size_api_proto_depIdxs = nil
}
