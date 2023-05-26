// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/material_api.proto

package material_api

import (
	material_model "github.com/Alexander272/sealur_proto/api/pro/models/material_model"
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

type GetAllMaterials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllMaterials) Reset() {
	*x = GetAllMaterials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_material_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMaterials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMaterials) ProtoMessage() {}

func (x *GetAllMaterials) ProtoReflect() protoreflect.Message {
	mi := &file_pro_material_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMaterials.ProtoReflect.Descriptor instead.
func (*GetAllMaterials) Descriptor() ([]byte, []int) {
	return file_pro_material_api_proto_rawDescGZIP(), []int{0}
}

type CreateMaterial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code     string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	ShortEn  string `protobuf:"bytes,3,opt,name=shortEn,proto3" json:"shortEn,omitempty"`
	ShortRus string `protobuf:"bytes,4,opt,name=shortRus,proto3" json:"shortRus,omitempty"`
}

func (x *CreateMaterial) Reset() {
	*x = CreateMaterial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_material_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMaterial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMaterial) ProtoMessage() {}

func (x *CreateMaterial) ProtoReflect() protoreflect.Message {
	mi := &file_pro_material_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMaterial.ProtoReflect.Descriptor instead.
func (*CreateMaterial) Descriptor() ([]byte, []int) {
	return file_pro_material_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMaterial) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateMaterial) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateMaterial) GetShortEn() string {
	if x != nil {
		return x.ShortEn
	}
	return ""
}

func (x *CreateMaterial) GetShortRus() string {
	if x != nil {
		return x.ShortRus
	}
	return ""
}

type CreateSeveralMaterial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Materials []*CreateMaterial `protobuf:"bytes,1,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *CreateSeveralMaterial) Reset() {
	*x = CreateSeveralMaterial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_material_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeveralMaterial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeveralMaterial) ProtoMessage() {}

func (x *CreateSeveralMaterial) ProtoReflect() protoreflect.Message {
	mi := &file_pro_material_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeveralMaterial.ProtoReflect.Descriptor instead.
func (*CreateSeveralMaterial) Descriptor() ([]byte, []int) {
	return file_pro_material_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSeveralMaterial) GetMaterials() []*CreateMaterial {
	if x != nil {
		return x.Materials
	}
	return nil
}

type UpdateMaterial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	ShortEn  string `protobuf:"bytes,4,opt,name=shortEn,proto3" json:"shortEn,omitempty"`
	ShortRus string `protobuf:"bytes,5,opt,name=shortRus,proto3" json:"shortRus,omitempty"`
}

func (x *UpdateMaterial) Reset() {
	*x = UpdateMaterial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_material_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMaterial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMaterial) ProtoMessage() {}

func (x *UpdateMaterial) ProtoReflect() protoreflect.Message {
	mi := &file_pro_material_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMaterial.ProtoReflect.Descriptor instead.
func (*UpdateMaterial) Descriptor() ([]byte, []int) {
	return file_pro_material_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMaterial) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMaterial) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMaterial) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateMaterial) GetShortEn() string {
	if x != nil {
		return x.ShortEn
	}
	return ""
}

func (x *UpdateMaterial) GetShortRus() string {
	if x != nil {
		return x.ShortRus
	}
	return ""
}

type DeleteMaterial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMaterial) Reset() {
	*x = DeleteMaterial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_material_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMaterial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMaterial) ProtoMessage() {}

func (x *DeleteMaterial) ProtoReflect() protoreflect.Message {
	mi := &file_pro_material_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMaterial.ProtoReflect.Descriptor instead.
func (*DeleteMaterial) Descriptor() ([]byte, []int) {
	return file_pro_material_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMaterial) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Materials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Materials []*material_model.Material `protobuf:"bytes,1,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *Materials) Reset() {
	*x = Materials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_material_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Materials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Materials) ProtoMessage() {}

func (x *Materials) ProtoReflect() protoreflect.Message {
	mi := &file_pro_material_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Materials.ProtoReflect.Descriptor instead.
func (*Materials) Descriptor() ([]byte, []int) {
	return file_pro_material_api_proto_rawDescGZIP(), []int{5}
}

func (x *Materials) GetMaterials() []*material_model.Material {
	if x != nil {
		return x.Materials
	}
	return nil
}

var File_pro_material_api_proto protoreflect.FileDescriptor

var file_pro_material_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x22,
	0x70, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x45, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x75,
	0x73, 0x22, 0x53, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x09, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x09, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x09, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x32, 0xe9, 0x02, 0x0a, 0x0f, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x23, 0x2e,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x16, 0x5a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_material_api_proto_rawDescOnce sync.Once
	file_pro_material_api_proto_rawDescData = file_pro_material_api_proto_rawDesc
)

func file_pro_material_api_proto_rawDescGZIP() []byte {
	file_pro_material_api_proto_rawDescOnce.Do(func() {
		file_pro_material_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_material_api_proto_rawDescData)
	})
	return file_pro_material_api_proto_rawDescData
}

var file_pro_material_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pro_material_api_proto_goTypes = []interface{}{
	(*GetAllMaterials)(nil),         // 0: material_api.GetAllMaterials
	(*CreateMaterial)(nil),          // 1: material_api.CreateMaterial
	(*CreateSeveralMaterial)(nil),   // 2: material_api.CreateSeveralMaterial
	(*UpdateMaterial)(nil),          // 3: material_api.UpdateMaterial
	(*DeleteMaterial)(nil),          // 4: material_api.DeleteMaterial
	(*Materials)(nil),               // 5: material_api.Materials
	(*material_model.Material)(nil), // 6: pro_material_model.Material
	(*response_model.Response)(nil), // 7: response_model.Response
}
var file_pro_material_api_proto_depIdxs = []int32{
	1, // 0: material_api.CreateSeveralMaterial.materials:type_name -> material_api.CreateMaterial
	6, // 1: material_api.Materials.materials:type_name -> pro_material_model.Material
	0, // 2: material_api.MaterialService.GetAll:input_type -> material_api.GetAllMaterials
	1, // 3: material_api.MaterialService.Create:input_type -> material_api.CreateMaterial
	2, // 4: material_api.MaterialService.CreateSeveral:input_type -> material_api.CreateSeveralMaterial
	3, // 5: material_api.MaterialService.Update:input_type -> material_api.UpdateMaterial
	4, // 6: material_api.MaterialService.Delete:input_type -> material_api.DeleteMaterial
	5, // 7: material_api.MaterialService.GetAll:output_type -> material_api.Materials
	7, // 8: material_api.MaterialService.Create:output_type -> response_model.Response
	7, // 9: material_api.MaterialService.CreateSeveral:output_type -> response_model.Response
	7, // 10: material_api.MaterialService.Update:output_type -> response_model.Response
	7, // 11: material_api.MaterialService.Delete:output_type -> response_model.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pro_material_api_proto_init() }
func file_pro_material_api_proto_init() {
	if File_pro_material_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_material_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMaterials); i {
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
		file_pro_material_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMaterial); i {
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
		file_pro_material_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeveralMaterial); i {
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
		file_pro_material_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMaterial); i {
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
		file_pro_material_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMaterial); i {
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
		file_pro_material_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Materials); i {
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
			RawDescriptor: file_pro_material_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_material_api_proto_goTypes,
		DependencyIndexes: file_pro_material_api_proto_depIdxs,
		MessageInfos:      file_pro_material_api_proto_msgTypes,
	}.Build()
	File_pro_material_api_proto = out.File
	file_pro_material_api_proto_rawDesc = nil
	file_pro_material_api_proto_goTypes = nil
	file_pro_material_api_proto_depIdxs = nil
}
