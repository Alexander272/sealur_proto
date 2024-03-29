// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/putg_flange_type_api.proto

package putg_flange_type_api

import (
	putg_flange_type_model "github.com/Alexander272/sealur_proto/api/pro/models/putg_flange_type_model"
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

type GetPutgFlangeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StandardId string `protobuf:"bytes,1,opt,name=standardId,proto3" json:"standardId,omitempty"`
}

func (x *GetPutgFlangeType) Reset() {
	*x = GetPutgFlangeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_flange_type_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPutgFlangeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPutgFlangeType) ProtoMessage() {}

func (x *GetPutgFlangeType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_flange_type_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPutgFlangeType.ProtoReflect.Descriptor instead.
func (*GetPutgFlangeType) Descriptor() ([]byte, []int) {
	return file_pro_putg_flange_type_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetPutgFlangeType) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

type CreatePutgFlangeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StandardId string `protobuf:"bytes,1,opt,name=standardId,proto3" json:"standardId,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CreatePutgFlangeType) Reset() {
	*x = CreatePutgFlangeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_flange_type_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePutgFlangeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePutgFlangeType) ProtoMessage() {}

func (x *CreatePutgFlangeType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_flange_type_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePutgFlangeType.ProtoReflect.Descriptor instead.
func (*CreatePutgFlangeType) Descriptor() ([]byte, []int) {
	return file_pro_putg_flange_type_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePutgFlangeType) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

func (x *CreatePutgFlangeType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePutgFlangeType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type UpdatePutgFlangeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	StandardId string `protobuf:"bytes,4,opt,name=standardId,proto3" json:"standardId,omitempty"`
}

func (x *UpdatePutgFlangeType) Reset() {
	*x = UpdatePutgFlangeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_flange_type_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePutgFlangeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePutgFlangeType) ProtoMessage() {}

func (x *UpdatePutgFlangeType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_flange_type_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePutgFlangeType.ProtoReflect.Descriptor instead.
func (*UpdatePutgFlangeType) Descriptor() ([]byte, []int) {
	return file_pro_putg_flange_type_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePutgFlangeType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePutgFlangeType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePutgFlangeType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdatePutgFlangeType) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

type DeletePutgFlangeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePutgFlangeType) Reset() {
	*x = DeletePutgFlangeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_flange_type_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePutgFlangeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePutgFlangeType) ProtoMessage() {}

func (x *DeletePutgFlangeType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_flange_type_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePutgFlangeType.ProtoReflect.Descriptor instead.
func (*DeletePutgFlangeType) Descriptor() ([]byte, []int) {
	return file_pro_putg_flange_type_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePutgFlangeType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PutgFlangeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlangeTypes []*putg_flange_type_model.PutgFlangeType `protobuf:"bytes,1,rep,name=flangeTypes,proto3" json:"flangeTypes,omitempty"`
}

func (x *PutgFlangeType) Reset() {
	*x = PutgFlangeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_flange_type_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutgFlangeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutgFlangeType) ProtoMessage() {}

func (x *PutgFlangeType) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_flange_type_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutgFlangeType.ProtoReflect.Descriptor instead.
func (*PutgFlangeType) Descriptor() ([]byte, []int) {
	return file_pro_putg_flange_type_api_proto_rawDescGZIP(), []int{4}
}

func (x *PutgFlangeType) GetFlangeTypes() []*putg_flange_type_model.PutgFlangeType {
	if x != nil {
		return x.FlangeTypes
	}
	return nil
}

var File_pro_putg_flange_type_api_proto protoreflect.FileDescriptor

var file_pro_putg_flange_type_api_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x6c, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x33, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x70, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f,
	0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x32, 0xdd, 0x02,
	0x0a, 0x15, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x27,
	0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x24, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66,
	0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66,
	0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66,
	0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66,
	0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a,
	0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x6c,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_putg_flange_type_api_proto_rawDescOnce sync.Once
	file_pro_putg_flange_type_api_proto_rawDescData = file_pro_putg_flange_type_api_proto_rawDesc
)

func file_pro_putg_flange_type_api_proto_rawDescGZIP() []byte {
	file_pro_putg_flange_type_api_proto_rawDescOnce.Do(func() {
		file_pro_putg_flange_type_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_putg_flange_type_api_proto_rawDescData)
	})
	return file_pro_putg_flange_type_api_proto_rawDescData
}

var file_pro_putg_flange_type_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_putg_flange_type_api_proto_goTypes = []interface{}{
	(*GetPutgFlangeType)(nil),                     // 0: putg_flange_type_api.GetPutgFlangeType
	(*CreatePutgFlangeType)(nil),                  // 1: putg_flange_type_api.CreatePutgFlangeType
	(*UpdatePutgFlangeType)(nil),                  // 2: putg_flange_type_api.UpdatePutgFlangeType
	(*DeletePutgFlangeType)(nil),                  // 3: putg_flange_type_api.DeletePutgFlangeType
	(*PutgFlangeType)(nil),                        // 4: putg_flange_type_api.PutgFlangeType
	(*putg_flange_type_model.PutgFlangeType)(nil), // 5: putg_flange_type_model.PutgFlangeType
	(*response_model.Response)(nil),               // 6: response_model.Response
}
var file_pro_putg_flange_type_api_proto_depIdxs = []int32{
	5, // 0: putg_flange_type_api.PutgFlangeType.flangeTypes:type_name -> putg_flange_type_model.PutgFlangeType
	0, // 1: putg_flange_type_api.PutgFlangeTypeService.Get:input_type -> putg_flange_type_api.GetPutgFlangeType
	1, // 2: putg_flange_type_api.PutgFlangeTypeService.Create:input_type -> putg_flange_type_api.CreatePutgFlangeType
	2, // 3: putg_flange_type_api.PutgFlangeTypeService.Update:input_type -> putg_flange_type_api.UpdatePutgFlangeType
	3, // 4: putg_flange_type_api.PutgFlangeTypeService.Delete:input_type -> putg_flange_type_api.DeletePutgFlangeType
	4, // 5: putg_flange_type_api.PutgFlangeTypeService.Get:output_type -> putg_flange_type_api.PutgFlangeType
	6, // 6: putg_flange_type_api.PutgFlangeTypeService.Create:output_type -> response_model.Response
	6, // 7: putg_flange_type_api.PutgFlangeTypeService.Update:output_type -> response_model.Response
	6, // 8: putg_flange_type_api.PutgFlangeTypeService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_putg_flange_type_api_proto_init() }
func file_pro_putg_flange_type_api_proto_init() {
	if File_pro_putg_flange_type_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_putg_flange_type_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPutgFlangeType); i {
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
		file_pro_putg_flange_type_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePutgFlangeType); i {
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
		file_pro_putg_flange_type_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePutgFlangeType); i {
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
		file_pro_putg_flange_type_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePutgFlangeType); i {
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
		file_pro_putg_flange_type_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutgFlangeType); i {
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
			RawDescriptor: file_pro_putg_flange_type_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_putg_flange_type_api_proto_goTypes,
		DependencyIndexes: file_pro_putg_flange_type_api_proto_depIdxs,
		MessageInfos:      file_pro_putg_flange_type_api_proto_msgTypes,
	}.Build()
	File_pro_putg_flange_type_api_proto = out.File
	file_pro_putg_flange_type_api_proto_rawDesc = nil
	file_pro_putg_flange_type_api_proto_goTypes = nil
	file_pro_putg_flange_type_api_proto_depIdxs = nil
}
