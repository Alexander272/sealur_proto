// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/putg_data_api.proto

package putg_data_api

import (
	putg_data_model "github.com/Alexander272/sealur_proto/api/pro/models/putg_data_model"
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

type GetPutgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FillerId       string `protobuf:"bytes,1,opt,name=fillerId,proto3" json:"fillerId,omitempty"`
	ConstructionId string `protobuf:"bytes,2,opt,name=constructionId,proto3" json:"constructionId,omitempty"`
}

func (x *GetPutgData) Reset() {
	*x = GetPutgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_data_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPutgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPutgData) ProtoMessage() {}

func (x *GetPutgData) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_data_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPutgData.ProtoReflect.Descriptor instead.
func (*GetPutgData) Descriptor() ([]byte, []int) {
	return file_pro_putg_data_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetPutgData) GetFillerId() string {
	if x != nil {
		return x.FillerId
	}
	return ""
}

func (x *GetPutgData) GetConstructionId() string {
	if x != nil {
		return x.ConstructionId
	}
	return ""
}

type CreatePutgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FillerId     string `protobuf:"bytes,1,opt,name=fillerId,proto3" json:"fillerId,omitempty"`
	HasJumper    bool   `protobuf:"varint,2,opt,name=hasJumper,proto3" json:"hasJumper,omitempty"`
	HasHole      bool   `protobuf:"varint,3,opt,name=hasHole,proto3" json:"hasHole,omitempty"`
	HasRemovable bool   `protobuf:"varint,4,opt,name=hasRemovable,proto3" json:"hasRemovable,omitempty"`
	HasMounting  bool   `protobuf:"varint,5,opt,name=hasMounting,proto3" json:"hasMounting,omitempty"`
	HasCoating   bool   `protobuf:"varint,6,opt,name=hasCoating,proto3" json:"hasCoating,omitempty"`
}

func (x *CreatePutgData) Reset() {
	*x = CreatePutgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_data_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePutgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePutgData) ProtoMessage() {}

func (x *CreatePutgData) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_data_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePutgData.ProtoReflect.Descriptor instead.
func (*CreatePutgData) Descriptor() ([]byte, []int) {
	return file_pro_putg_data_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePutgData) GetFillerId() string {
	if x != nil {
		return x.FillerId
	}
	return ""
}

func (x *CreatePutgData) GetHasJumper() bool {
	if x != nil {
		return x.HasJumper
	}
	return false
}

func (x *CreatePutgData) GetHasHole() bool {
	if x != nil {
		return x.HasHole
	}
	return false
}

func (x *CreatePutgData) GetHasRemovable() bool {
	if x != nil {
		return x.HasRemovable
	}
	return false
}

func (x *CreatePutgData) GetHasMounting() bool {
	if x != nil {
		return x.HasMounting
	}
	return false
}

func (x *CreatePutgData) GetHasCoating() bool {
	if x != nil {
		return x.HasCoating
	}
	return false
}

type UpdatePutgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HasJumper    bool   `protobuf:"varint,2,opt,name=hasJumper,proto3" json:"hasJumper,omitempty"`
	HasHole      bool   `protobuf:"varint,3,opt,name=hasHole,proto3" json:"hasHole,omitempty"`
	HasRemovable bool   `protobuf:"varint,4,opt,name=hasRemovable,proto3" json:"hasRemovable,omitempty"`
	HasMounting  bool   `protobuf:"varint,5,opt,name=hasMounting,proto3" json:"hasMounting,omitempty"`
	HasCoating   bool   `protobuf:"varint,6,opt,name=hasCoating,proto3" json:"hasCoating,omitempty"`
	FillerId     string `protobuf:"bytes,7,opt,name=fillerId,proto3" json:"fillerId,omitempty"`
}

func (x *UpdatePutgData) Reset() {
	*x = UpdatePutgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_data_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePutgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePutgData) ProtoMessage() {}

func (x *UpdatePutgData) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_data_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePutgData.ProtoReflect.Descriptor instead.
func (*UpdatePutgData) Descriptor() ([]byte, []int) {
	return file_pro_putg_data_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePutgData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePutgData) GetHasJumper() bool {
	if x != nil {
		return x.HasJumper
	}
	return false
}

func (x *UpdatePutgData) GetHasHole() bool {
	if x != nil {
		return x.HasHole
	}
	return false
}

func (x *UpdatePutgData) GetHasRemovable() bool {
	if x != nil {
		return x.HasRemovable
	}
	return false
}

func (x *UpdatePutgData) GetHasMounting() bool {
	if x != nil {
		return x.HasMounting
	}
	return false
}

func (x *UpdatePutgData) GetHasCoating() bool {
	if x != nil {
		return x.HasCoating
	}
	return false
}

func (x *UpdatePutgData) GetFillerId() string {
	if x != nil {
		return x.FillerId
	}
	return ""
}

type DeletePutgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePutgData) Reset() {
	*x = DeletePutgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_data_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePutgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePutgData) ProtoMessage() {}

func (x *DeletePutgData) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_data_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePutgData.ProtoReflect.Descriptor instead.
func (*DeletePutgData) Descriptor() ([]byte, []int) {
	return file_pro_putg_data_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePutgData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PutgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*putg_data_model.PutgData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PutgData) Reset() {
	*x = PutgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_data_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutgData) ProtoMessage() {}

func (x *PutgData) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_data_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutgData.ProtoReflect.Descriptor instead.
func (*PutgData) Descriptor() ([]byte, []int) {
	return file_pro_putg_data_api_proto_rawDescGZIP(), []int{4}
}

func (x *PutgData) GetData() []*putg_data_model.PutgData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pro_putg_data_api_proto protoreflect.FileDescriptor

var file_pro_putg_data_api_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x75, 0x74, 0x67, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xca,
	0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x68, 0x61, 0x73, 0x4a, 0x75, 0x6d, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x68, 0x61, 0x73, 0x4a, 0x75, 0x6d, 0x70, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x61, 0x73, 0x48, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61,
	0x73, 0x48, 0x6f, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x61, 0x73,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x68, 0x61, 0x73, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x68,
	0x61, 0x73, 0x43, 0x6f, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x68, 0x61, 0x73, 0x43, 0x6f, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xda, 0x01, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x68, 0x61, 0x73, 0x4a, 0x75, 0x6d, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x4a, 0x75, 0x6d, 0x70, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x68, 0x61, 0x73, 0x48, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68,
	0x61, 0x73, 0x48, 0x6f, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61,
	0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x61,
	0x73, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x61, 0x73, 0x43, 0x6f, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x68, 0x61, 0x73, 0x43, 0x6f, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x08, 0x50, 0x75,
	0x74, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x96, 0x02, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1a, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x70,
	0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x74,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17,
	0x5a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_putg_data_api_proto_rawDescOnce sync.Once
	file_pro_putg_data_api_proto_rawDescData = file_pro_putg_data_api_proto_rawDesc
)

func file_pro_putg_data_api_proto_rawDescGZIP() []byte {
	file_pro_putg_data_api_proto_rawDescOnce.Do(func() {
		file_pro_putg_data_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_putg_data_api_proto_rawDescData)
	})
	return file_pro_putg_data_api_proto_rawDescData
}

var file_pro_putg_data_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_putg_data_api_proto_goTypes = []interface{}{
	(*GetPutgData)(nil),              // 0: putg_data_api.GetPutgData
	(*CreatePutgData)(nil),           // 1: putg_data_api.CreatePutgData
	(*UpdatePutgData)(nil),           // 2: putg_data_api.UpdatePutgData
	(*DeletePutgData)(nil),           // 3: putg_data_api.DeletePutgData
	(*PutgData)(nil),                 // 4: putg_data_api.PutgData
	(*putg_data_model.PutgData)(nil), // 5: putg_data_model.PutgData
	(*response_model.Response)(nil),  // 6: response_model.Response
}
var file_pro_putg_data_api_proto_depIdxs = []int32{
	5, // 0: putg_data_api.PutgData.data:type_name -> putg_data_model.PutgData
	0, // 1: putg_data_api.PutgDataService.Get:input_type -> putg_data_api.GetPutgData
	1, // 2: putg_data_api.PutgDataService.Create:input_type -> putg_data_api.CreatePutgData
	2, // 3: putg_data_api.PutgDataService.Update:input_type -> putg_data_api.UpdatePutgData
	3, // 4: putg_data_api.PutgDataService.Delete:input_type -> putg_data_api.DeletePutgData
	4, // 5: putg_data_api.PutgDataService.Get:output_type -> putg_data_api.PutgData
	6, // 6: putg_data_api.PutgDataService.Create:output_type -> response_model.Response
	6, // 7: putg_data_api.PutgDataService.Update:output_type -> response_model.Response
	6, // 8: putg_data_api.PutgDataService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_putg_data_api_proto_init() }
func file_pro_putg_data_api_proto_init() {
	if File_pro_putg_data_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_putg_data_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPutgData); i {
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
		file_pro_putg_data_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePutgData); i {
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
		file_pro_putg_data_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePutgData); i {
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
		file_pro_putg_data_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePutgData); i {
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
		file_pro_putg_data_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutgData); i {
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
			RawDescriptor: file_pro_putg_data_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_putg_data_api_proto_goTypes,
		DependencyIndexes: file_pro_putg_data_api_proto_depIdxs,
		MessageInfos:      file_pro_putg_data_api_proto_msgTypes,
	}.Build()
	File_pro_putg_data_api_proto = out.File
	file_pro_putg_data_api_proto_rawDesc = nil
	file_pro_putg_data_api_proto_goTypes = nil
	file_pro_putg_data_api_proto_depIdxs = nil
}
