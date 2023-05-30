// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/putg_construction_api.proto

package putg_construction_api

import (
	putg_construction_type_model "github.com/Alexander272/sealur_proto/api/pro/models/putg_construction_type_model"
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

type GetPutgConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO удалить standardId
	// string standardId = 1;
	FlangeTypeId string `protobuf:"bytes,2,opt,name=flangeTypeId,proto3" json:"flangeTypeId,omitempty"`
}

func (x *GetPutgConstruction) Reset() {
	*x = GetPutgConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_construction_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPutgConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPutgConstruction) ProtoMessage() {}

func (x *GetPutgConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_construction_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPutgConstruction.ProtoReflect.Descriptor instead.
func (*GetPutgConstruction) Descriptor() ([]byte, []int) {
	return file_pro_putg_construction_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetPutgConstruction) GetFlangeTypeId() string {
	if x != nil {
		return x.FlangeTypeId
	}
	return ""
}

type GetPutgConstruction_New struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FillerId     string `protobuf:"bytes,1,opt,name=fillerId,proto3" json:"fillerId,omitempty"`
	FlangeTypeId string `protobuf:"bytes,2,opt,name=flangeTypeId,proto3" json:"flangeTypeId,omitempty"`
}

func (x *GetPutgConstruction_New) Reset() {
	*x = GetPutgConstruction_New{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_construction_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPutgConstruction_New) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPutgConstruction_New) ProtoMessage() {}

func (x *GetPutgConstruction_New) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_construction_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPutgConstruction_New.ProtoReflect.Descriptor instead.
func (*GetPutgConstruction_New) Descriptor() ([]byte, []int) {
	return file_pro_putg_construction_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetPutgConstruction_New) GetFillerId() string {
	if x != nil {
		return x.FillerId
	}
	return ""
}

func (x *GetPutgConstruction_New) GetFlangeTypeId() string {
	if x != nil {
		return x.FlangeTypeId
	}
	return ""
}

type CreatePutgConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConstructionId string `protobuf:"bytes,1,opt,name=constructionId,proto3" json:"constructionId,omitempty"`
	StandardId     string `protobuf:"bytes,2,opt,name=standardId,proto3" json:"standardId,omitempty"`
	FlangeTypeId   string `protobuf:"bytes,3,opt,name=flangeTypeId,proto3" json:"flangeTypeId,omitempty"`
}

func (x *CreatePutgConstruction) Reset() {
	*x = CreatePutgConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_construction_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePutgConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePutgConstruction) ProtoMessage() {}

func (x *CreatePutgConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_construction_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePutgConstruction.ProtoReflect.Descriptor instead.
func (*CreatePutgConstruction) Descriptor() ([]byte, []int) {
	return file_pro_putg_construction_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePutgConstruction) GetConstructionId() string {
	if x != nil {
		return x.ConstructionId
	}
	return ""
}

func (x *CreatePutgConstruction) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

func (x *CreatePutgConstruction) GetFlangeTypeId() string {
	if x != nil {
		return x.FlangeTypeId
	}
	return ""
}

type UpdatePutgConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConstructionId string `protobuf:"bytes,2,opt,name=constructionId,proto3" json:"constructionId,omitempty"`
	StandardId     string `protobuf:"bytes,3,opt,name=standardId,proto3" json:"standardId,omitempty"`
	FlangeTypeId   string `protobuf:"bytes,4,opt,name=flangeTypeId,proto3" json:"flangeTypeId,omitempty"`
}

func (x *UpdatePutgConstruction) Reset() {
	*x = UpdatePutgConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_construction_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePutgConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePutgConstruction) ProtoMessage() {}

func (x *UpdatePutgConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_construction_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePutgConstruction.ProtoReflect.Descriptor instead.
func (*UpdatePutgConstruction) Descriptor() ([]byte, []int) {
	return file_pro_putg_construction_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePutgConstruction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePutgConstruction) GetConstructionId() string {
	if x != nil {
		return x.ConstructionId
	}
	return ""
}

func (x *UpdatePutgConstruction) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

func (x *UpdatePutgConstruction) GetFlangeTypeId() string {
	if x != nil {
		return x.FlangeTypeId
	}
	return ""
}

type DeletePutgConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePutgConstruction) Reset() {
	*x = DeletePutgConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_construction_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePutgConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePutgConstruction) ProtoMessage() {}

func (x *DeletePutgConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_construction_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePutgConstruction.ProtoReflect.Descriptor instead.
func (*DeletePutgConstruction) Descriptor() ([]byte, []int) {
	return file_pro_putg_construction_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePutgConstruction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PutgConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constructions []*putg_construction_type_model.PutgConstruction `protobuf:"bytes,1,rep,name=constructions,proto3" json:"constructions,omitempty"`
}

func (x *PutgConstruction) Reset() {
	*x = PutgConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_construction_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutgConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutgConstruction) ProtoMessage() {}

func (x *PutgConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_construction_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutgConstruction.ProtoReflect.Descriptor instead.
func (*PutgConstruction) Descriptor() ([]byte, []int) {
	return file_pro_putg_construction_api_proto_rawDescGZIP(), []int{5}
}

func (x *PutgConstruction) GetConstructions() []*putg_construction_type_model.PutgConstruction {
	if x != nil {
		return x.Constructions
	}
	return nil
}

var File_pro_putg_construction_api_proto protoreflect.FileDescriptor

var file_pro_putg_construction_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70, 0x72, 0x6f, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x4e, 0x65, 0x77, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6c,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x84,
	0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0xd2, 0x03, 0x0a, 0x17, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x27, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x5f,
	0x4e, 0x65, 0x77, 0x12, 0x2e, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x4e, 0x65, 0x77, 0x1a, 0x27, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x74, 0x67,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x51, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x70, 0x75, 0x74, 0x67,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x70,
	0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_putg_construction_api_proto_rawDescOnce sync.Once
	file_pro_putg_construction_api_proto_rawDescData = file_pro_putg_construction_api_proto_rawDesc
)

func file_pro_putg_construction_api_proto_rawDescGZIP() []byte {
	file_pro_putg_construction_api_proto_rawDescOnce.Do(func() {
		file_pro_putg_construction_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_putg_construction_api_proto_rawDescData)
	})
	return file_pro_putg_construction_api_proto_rawDescData
}

var file_pro_putg_construction_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pro_putg_construction_api_proto_goTypes = []interface{}{
	(*GetPutgConstruction)(nil),                           // 0: putg_construction_api.GetPutgConstruction
	(*GetPutgConstruction_New)(nil),                       // 1: putg_construction_api.GetPutgConstruction_New
	(*CreatePutgConstruction)(nil),                        // 2: putg_construction_api.CreatePutgConstruction
	(*UpdatePutgConstruction)(nil),                        // 3: putg_construction_api.UpdatePutgConstruction
	(*DeletePutgConstruction)(nil),                        // 4: putg_construction_api.DeletePutgConstruction
	(*PutgConstruction)(nil),                              // 5: putg_construction_api.PutgConstruction
	(*putg_construction_type_model.PutgConstruction)(nil), // 6: putg_construction_type_model.PutgConstruction
	(*response_model.Response)(nil),                       // 7: response_model.Response
}
var file_pro_putg_construction_api_proto_depIdxs = []int32{
	6, // 0: putg_construction_api.PutgConstruction.constructions:type_name -> putg_construction_type_model.PutgConstruction
	0, // 1: putg_construction_api.PutgConstructionService.Get:input_type -> putg_construction_api.GetPutgConstruction
	1, // 2: putg_construction_api.PutgConstructionService.Get_New:input_type -> putg_construction_api.GetPutgConstruction_New
	2, // 3: putg_construction_api.PutgConstructionService.Create:input_type -> putg_construction_api.CreatePutgConstruction
	3, // 4: putg_construction_api.PutgConstructionService.Update:input_type -> putg_construction_api.UpdatePutgConstruction
	4, // 5: putg_construction_api.PutgConstructionService.Delete:input_type -> putg_construction_api.DeletePutgConstruction
	5, // 6: putg_construction_api.PutgConstructionService.Get:output_type -> putg_construction_api.PutgConstruction
	5, // 7: putg_construction_api.PutgConstructionService.Get_New:output_type -> putg_construction_api.PutgConstruction
	7, // 8: putg_construction_api.PutgConstructionService.Create:output_type -> response_model.Response
	7, // 9: putg_construction_api.PutgConstructionService.Update:output_type -> response_model.Response
	7, // 10: putg_construction_api.PutgConstructionService.Delete:output_type -> response_model.Response
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_putg_construction_api_proto_init() }
func file_pro_putg_construction_api_proto_init() {
	if File_pro_putg_construction_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_putg_construction_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPutgConstruction); i {
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
		file_pro_putg_construction_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPutgConstruction_New); i {
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
		file_pro_putg_construction_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePutgConstruction); i {
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
		file_pro_putg_construction_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePutgConstruction); i {
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
		file_pro_putg_construction_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePutgConstruction); i {
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
		file_pro_putg_construction_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutgConstruction); i {
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
			RawDescriptor: file_pro_putg_construction_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_putg_construction_api_proto_goTypes,
		DependencyIndexes: file_pro_putg_construction_api_proto_depIdxs,
		MessageInfos:      file_pro_putg_construction_api_proto_msgTypes,
	}.Build()
	File_pro_putg_construction_api_proto = out.File
	file_pro_putg_construction_api_proto_rawDesc = nil
	file_pro_putg_construction_api_proto_goTypes = nil
	file_pro_putg_construction_api_proto_depIdxs = nil
}
