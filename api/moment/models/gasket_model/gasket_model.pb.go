// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: moment/models/gasket_model.proto

package gasket_model

import (
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

type Gasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Gasket) Reset() {
	*x = Gasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_gasket_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gasket) ProtoMessage() {}

func (x *Gasket) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_gasket_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gasket.ProtoReflect.Descriptor instead.
func (*Gasket) Descriptor() ([]byte, []int) {
	return file_moment_models_gasket_model_proto_rawDescGZIP(), []int{0}
}

func (x *Gasket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Gasket) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GasketWithThick struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Thickness []float64 `protobuf:"fixed64,3,rep,packed,name=thickness,proto3" json:"thickness,omitempty"`
}

func (x *GasketWithThick) Reset() {
	*x = GasketWithThick{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_gasket_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GasketWithThick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GasketWithThick) ProtoMessage() {}

func (x *GasketWithThick) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_gasket_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GasketWithThick.ProtoReflect.Descriptor instead.
func (*GasketWithThick) Descriptor() ([]byte, []int) {
	return file_moment_models_gasket_model_proto_rawDescGZIP(), []int{1}
}

func (x *GasketWithThick) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GasketWithThick) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GasketWithThick) GetThickness() []float64 {
	if x != nil {
		return x.Thickness
	}
	return nil
}

type GasketType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *GasketType) Reset() {
	*x = GasketType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_gasket_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GasketType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GasketType) ProtoMessage() {}

func (x *GasketType) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_gasket_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GasketType.ProtoReflect.Descriptor instead.
func (*GasketType) Descriptor() ([]byte, []int) {
	return file_moment_models_gasket_model_proto_rawDescGZIP(), []int{2}
}

func (x *GasketType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GasketType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GasketType) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type Env struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Env) Reset() {
	*x = Env{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_gasket_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Env) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Env) ProtoMessage() {}

func (x *Env) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_gasket_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Env.ProtoReflect.Descriptor instead.
func (*Env) Descriptor() ([]byte, []int) {
	return file_moment_models_gasket_model_proto_rawDescGZIP(), []int{3}
}

func (x *Env) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Env) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Full_GasketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GasketId        string  `protobuf:"bytes,2,opt,name=gasketId,proto3" json:"gasketId,omitempty"`
	PermissiblePres float64 `protobuf:"fixed64,3,opt,name=permissiblePres,proto3" json:"permissiblePres,omitempty"`
	Compression     float64 `protobuf:"fixed64,4,opt,name=compression,proto3" json:"compression,omitempty"`
	Epsilon         float64 `protobuf:"fixed64,5,opt,name=epsilon,proto3" json:"epsilon,omitempty"`
	Thickness       float64 `protobuf:"fixed64,6,opt,name=thickness,proto3" json:"thickness,omitempty"`
	TypeId          string  `protobuf:"bytes,7,opt,name=typeId,proto3" json:"typeId,omitempty"`
}

func (x *Full_GasketData) Reset() {
	*x = Full_GasketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_gasket_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Full_GasketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Full_GasketData) ProtoMessage() {}

func (x *Full_GasketData) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_gasket_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Full_GasketData.ProtoReflect.Descriptor instead.
func (*Full_GasketData) Descriptor() ([]byte, []int) {
	return file_moment_models_gasket_model_proto_rawDescGZIP(), []int{4}
}

func (x *Full_GasketData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Full_GasketData) GetGasketId() string {
	if x != nil {
		return x.GasketId
	}
	return ""
}

func (x *Full_GasketData) GetPermissiblePres() float64 {
	if x != nil {
		return x.PermissiblePres
	}
	return 0
}

func (x *Full_GasketData) GetCompression() float64 {
	if x != nil {
		return x.Compression
	}
	return 0
}

func (x *Full_GasketData) GetEpsilon() float64 {
	if x != nil {
		return x.Epsilon
	}
	return 0
}

func (x *Full_GasketData) GetThickness() float64 {
	if x != nil {
		return x.Thickness
	}
	return 0
}

func (x *Full_GasketData) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

type Full_EnvData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EnvId        string  `protobuf:"bytes,2,opt,name=envId,proto3" json:"envId,omitempty"`
	GasketId     string  `protobuf:"bytes,3,opt,name=gasketId,proto3" json:"gasketId,omitempty"`
	M            float64 `protobuf:"fixed64,4,opt,name=m,proto3" json:"m,omitempty"`
	SpecificPres float64 `protobuf:"fixed64,5,opt,name=specificPres,proto3" json:"specificPres,omitempty"`
}

func (x *Full_EnvData) Reset() {
	*x = Full_EnvData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_gasket_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Full_EnvData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Full_EnvData) ProtoMessage() {}

func (x *Full_EnvData) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_gasket_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Full_EnvData.ProtoReflect.Descriptor instead.
func (*Full_EnvData) Descriptor() ([]byte, []int) {
	return file_moment_models_gasket_model_proto_rawDescGZIP(), []int{5}
}

func (x *Full_EnvData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Full_EnvData) GetEnvId() string {
	if x != nil {
		return x.EnvId
	}
	return ""
}

func (x *Full_EnvData) GetGasketId() string {
	if x != nil {
		return x.GasketId
	}
	return ""
}

func (x *Full_EnvData) GetM() float64 {
	if x != nil {
		return x.M
	}
	return 0
}

func (x *Full_EnvData) GetSpecificPres() float64 {
	if x != nil {
		return x.SpecificPres
	}
	return 0
}

var File_moment_models_gasket_model_proto protoreflect.FileDescriptor

var file_moment_models_gasket_model_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x67, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x67, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0x2e, 0x0a, 0x06, 0x47, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x55, 0x0a, 0x0f, 0x47, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x54, 0x68,
	0x69, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x69,
	0x63, 0x6b, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01, 0x52, 0x09, 0x74, 0x68,
	0x69, 0x63, 0x6b, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x48, 0x0a, 0x0a, 0x47, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x22, 0x2b, 0x0a, 0x03, 0x45, 0x6e, 0x76, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xd9,
	0x01, 0x0a, 0x0f, 0x46, 0x75, 0x6c, 0x6c, 0x5f, 0x47, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70,
	0x73, 0x69, 0x6c, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x65, 0x70, 0x73,
	0x69, 0x6c, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x6e, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x46,
	0x75, 0x6c, 0x6c, 0x5f, 0x45, 0x6e, 0x76, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6e, 0x76, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x0c, 0x0a,
	0x01, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x50, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x50, 0x72, 0x65, 0x73, 0x42,
	0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c,
	0x65, 0x78, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x32, 0x37, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x75,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moment_models_gasket_model_proto_rawDescOnce sync.Once
	file_moment_models_gasket_model_proto_rawDescData = file_moment_models_gasket_model_proto_rawDesc
)

func file_moment_models_gasket_model_proto_rawDescGZIP() []byte {
	file_moment_models_gasket_model_proto_rawDescOnce.Do(func() {
		file_moment_models_gasket_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_moment_models_gasket_model_proto_rawDescData)
	})
	return file_moment_models_gasket_model_proto_rawDescData
}

var file_moment_models_gasket_model_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_moment_models_gasket_model_proto_goTypes = []interface{}{
	(*Gasket)(nil),          // 0: gasket_model.Gasket
	(*GasketWithThick)(nil), // 1: gasket_model.GasketWithThick
	(*GasketType)(nil),      // 2: gasket_model.GasketType
	(*Env)(nil),             // 3: gasket_model.Env
	(*Full_GasketData)(nil), // 4: gasket_model.Full_GasketData
	(*Full_EnvData)(nil),    // 5: gasket_model.Full_EnvData
}
var file_moment_models_gasket_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_moment_models_gasket_model_proto_init() }
func file_moment_models_gasket_model_proto_init() {
	if File_moment_models_gasket_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moment_models_gasket_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gasket); i {
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
		file_moment_models_gasket_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GasketWithThick); i {
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
		file_moment_models_gasket_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GasketType); i {
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
		file_moment_models_gasket_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Env); i {
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
		file_moment_models_gasket_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Full_GasketData); i {
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
		file_moment_models_gasket_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Full_EnvData); i {
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
			RawDescriptor: file_moment_models_gasket_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moment_models_gasket_model_proto_goTypes,
		DependencyIndexes: file_moment_models_gasket_model_proto_depIdxs,
		MessageInfos:      file_moment_models_gasket_model_proto_msgTypes,
	}.Build()
	File_moment_models_gasket_model_proto = out.File
	file_moment_models_gasket_model_proto_rawDesc = nil
	file_moment_models_gasket_model_proto_goTypes = nil
	file_moment_models_gasket_model_proto_depIdxs = nil
}
