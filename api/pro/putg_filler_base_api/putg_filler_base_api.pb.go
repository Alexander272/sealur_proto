// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/putg_filler_base_api.proto

package putg_filler_base_api

import (
	putg_filler_model "github.com/Alexander272/sealur_proto/api/pro/models/putg_filler_model"
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

type GetPutgBaseFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPutgBaseFiller) Reset() {
	*x = GetPutgBaseFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_filler_base_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPutgBaseFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPutgBaseFiller) ProtoMessage() {}

func (x *GetPutgBaseFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_filler_base_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPutgBaseFiller.ProtoReflect.Descriptor instead.
func (*GetPutgBaseFiller) Descriptor() ([]byte, []int) {
	return file_pro_putg_filler_base_api_proto_rawDescGZIP(), []int{0}
}

type CreatePutgBaseFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemperatureId string `protobuf:"bytes,1,opt,name=temperatureId,proto3" json:"temperatureId,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code          string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description   string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Designation   string `protobuf:"bytes,5,opt,name=designation,proto3" json:"designation,omitempty"`
}

func (x *CreatePutgBaseFiller) Reset() {
	*x = CreatePutgBaseFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_filler_base_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePutgBaseFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePutgBaseFiller) ProtoMessage() {}

func (x *CreatePutgBaseFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_filler_base_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePutgBaseFiller.ProtoReflect.Descriptor instead.
func (*CreatePutgBaseFiller) Descriptor() ([]byte, []int) {
	return file_pro_putg_filler_base_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePutgBaseFiller) GetTemperatureId() string {
	if x != nil {
		return x.TemperatureId
	}
	return ""
}

func (x *CreatePutgBaseFiller) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePutgBaseFiller) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreatePutgBaseFiller) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePutgBaseFiller) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

type UpdatePutgBaseFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TemperatureId string `protobuf:"bytes,2,opt,name=temperatureId,proto3" json:"temperatureId,omitempty"`
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Code          string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Designation   string `protobuf:"bytes,6,opt,name=designation,proto3" json:"designation,omitempty"`
}

func (x *UpdatePutgBaseFiller) Reset() {
	*x = UpdatePutgBaseFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_filler_base_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePutgBaseFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePutgBaseFiller) ProtoMessage() {}

func (x *UpdatePutgBaseFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_filler_base_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePutgBaseFiller.ProtoReflect.Descriptor instead.
func (*UpdatePutgBaseFiller) Descriptor() ([]byte, []int) {
	return file_pro_putg_filler_base_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePutgBaseFiller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePutgBaseFiller) GetTemperatureId() string {
	if x != nil {
		return x.TemperatureId
	}
	return ""
}

func (x *UpdatePutgBaseFiller) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePutgBaseFiller) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdatePutgBaseFiller) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdatePutgBaseFiller) GetDesignation() string {
	if x != nil {
		return x.Designation
	}
	return ""
}

type DeletePutgBaseFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePutgBaseFiller) Reset() {
	*x = DeletePutgBaseFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_filler_base_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePutgBaseFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePutgBaseFiller) ProtoMessage() {}

func (x *DeletePutgBaseFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_filler_base_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePutgBaseFiller.ProtoReflect.Descriptor instead.
func (*DeletePutgBaseFiller) Descriptor() ([]byte, []int) {
	return file_pro_putg_filler_base_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePutgBaseFiller) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PutgBaseFiller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fillers []*putg_filler_model.PutgFiller `protobuf:"bytes,1,rep,name=fillers,proto3" json:"fillers,omitempty"`
}

func (x *PutgBaseFiller) Reset() {
	*x = PutgBaseFiller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_putg_filler_base_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutgBaseFiller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutgBaseFiller) ProtoMessage() {}

func (x *PutgBaseFiller) ProtoReflect() protoreflect.Message {
	mi := &file_pro_putg_filler_base_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutgBaseFiller.ProtoReflect.Descriptor instead.
func (*PutgBaseFiller) Descriptor() ([]byte, []int) {
	return file_pro_putg_filler_base_api_proto_rawDescGZIP(), []int{4}
}

func (x *PutgBaseFiller) GetFillers() []*putg_filler_model.PutgFiller {
	if x != nil {
		return x.Fillers
	}
	return nil
}

var File_pro_putg_filler_base_api_proto protoreflect.FileDescriptor

var file_pro_putg_filler_base_api_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65,
	0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x22, 0xaa, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x42,
	0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xba, 0x01,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65,
	0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x49, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x75, 0x74, 0x67, 0x46, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x32, 0xdd, 0x02,
	0x0a, 0x15, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x27,
	0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73,
	0x65, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x24, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4e, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x74, 0x67, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a,
	0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x66, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_putg_filler_base_api_proto_rawDescOnce sync.Once
	file_pro_putg_filler_base_api_proto_rawDescData = file_pro_putg_filler_base_api_proto_rawDesc
)

func file_pro_putg_filler_base_api_proto_rawDescGZIP() []byte {
	file_pro_putg_filler_base_api_proto_rawDescOnce.Do(func() {
		file_pro_putg_filler_base_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_putg_filler_base_api_proto_rawDescData)
	})
	return file_pro_putg_filler_base_api_proto_rawDescData
}

var file_pro_putg_filler_base_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_putg_filler_base_api_proto_goTypes = []interface{}{
	(*GetPutgBaseFiller)(nil),            // 0: putg_filler_base_api.GetPutgBaseFiller
	(*CreatePutgBaseFiller)(nil),         // 1: putg_filler_base_api.CreatePutgBaseFiller
	(*UpdatePutgBaseFiller)(nil),         // 2: putg_filler_base_api.UpdatePutgBaseFiller
	(*DeletePutgBaseFiller)(nil),         // 3: putg_filler_base_api.DeletePutgBaseFiller
	(*PutgBaseFiller)(nil),               // 4: putg_filler_base_api.PutgBaseFiller
	(*putg_filler_model.PutgFiller)(nil), // 5: putg_filler_model.PutgFiller
	(*response_model.Response)(nil),      // 6: response_model.Response
}
var file_pro_putg_filler_base_api_proto_depIdxs = []int32{
	5, // 0: putg_filler_base_api.PutgBaseFiller.fillers:type_name -> putg_filler_model.PutgFiller
	0, // 1: putg_filler_base_api.PutgBaseFillerService.Get:input_type -> putg_filler_base_api.GetPutgBaseFiller
	1, // 2: putg_filler_base_api.PutgBaseFillerService.Create:input_type -> putg_filler_base_api.CreatePutgBaseFiller
	2, // 3: putg_filler_base_api.PutgBaseFillerService.Update:input_type -> putg_filler_base_api.UpdatePutgBaseFiller
	3, // 4: putg_filler_base_api.PutgBaseFillerService.Delete:input_type -> putg_filler_base_api.DeletePutgBaseFiller
	4, // 5: putg_filler_base_api.PutgBaseFillerService.Get:output_type -> putg_filler_base_api.PutgBaseFiller
	6, // 6: putg_filler_base_api.PutgBaseFillerService.Create:output_type -> response_model.Response
	6, // 7: putg_filler_base_api.PutgBaseFillerService.Update:output_type -> response_model.Response
	6, // 8: putg_filler_base_api.PutgBaseFillerService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_putg_filler_base_api_proto_init() }
func file_pro_putg_filler_base_api_proto_init() {
	if File_pro_putg_filler_base_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_putg_filler_base_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPutgBaseFiller); i {
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
		file_pro_putg_filler_base_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePutgBaseFiller); i {
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
		file_pro_putg_filler_base_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePutgBaseFiller); i {
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
		file_pro_putg_filler_base_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePutgBaseFiller); i {
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
		file_pro_putg_filler_base_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutgBaseFiller); i {
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
			RawDescriptor: file_pro_putg_filler_base_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_putg_filler_base_api_proto_goTypes,
		DependencyIndexes: file_pro_putg_filler_base_api_proto_depIdxs,
		MessageInfos:      file_pro_putg_filler_base_api_proto_msgTypes,
	}.Build()
	File_pro_putg_filler_base_api_proto = out.File
	file_pro_putg_filler_base_api_proto_rawDesc = nil
	file_pro_putg_filler_base_api_proto_goTypes = nil
	file_pro_putg_filler_base_api_proto_depIdxs = nil
}
