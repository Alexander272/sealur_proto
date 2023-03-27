// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: moment/models/device_model.proto

package device_model

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

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Pressure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Pressure) Reset() {
	*x = Pressure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pressure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pressure) ProtoMessage() {}

func (x *Pressure) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pressure.ProtoReflect.Descriptor instead.
func (*Pressure) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{1}
}

func (x *Pressure) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pressure) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TubeCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value int32  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TubeCount) Reset() {
	*x = TubeCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TubeCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TubeCount) ProtoMessage() {}

func (x *TubeCount) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TubeCount.ProtoReflect.Descriptor instead.
func (*TubeCount) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{2}
}

func (x *TubeCount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TubeCount) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type FinningFactor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DevId string `protobuf:"bytes,2,opt,name=devId,proto3" json:"devId,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FinningFactor) Reset() {
	*x = FinningFactor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinningFactor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinningFactor) ProtoMessage() {}

func (x *FinningFactor) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinningFactor.ProtoReflect.Descriptor instead.
func (*FinningFactor) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{3}
}

func (x *FinningFactor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FinningFactor) GetDevId() string {
	if x != nil {
		return x.DevId
	}
	return ""
}

func (x *FinningFactor) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SectionExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DevId string `protobuf:"bytes,2,opt,name=devId,proto3" json:"devId,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SectionExecution) Reset() {
	*x = SectionExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionExecution) ProtoMessage() {}

func (x *SectionExecution) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionExecution.ProtoReflect.Descriptor instead.
func (*SectionExecution) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{4}
}

func (x *SectionExecution) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SectionExecution) GetDevId() string {
	if x != nil {
		return x.DevId
	}
	return ""
}

func (x *SectionExecution) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TubeLength struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DevId string `protobuf:"bytes,2,opt,name=devId,proto3" json:"devId,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TubeLength) Reset() {
	*x = TubeLength{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TubeLength) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TubeLength) ProtoMessage() {}

func (x *TubeLength) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TubeLength.ProtoReflect.Descriptor instead.
func (*TubeLength) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{5}
}

func (x *TubeLength) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TubeLength) GetDevId() string {
	if x != nil {
		return x.DevId
	}
	return ""
}

func (x *TubeLength) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type NumberOfMoves struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DevId   string `protobuf:"bytes,2,opt,name=devId,proto3" json:"devId,omitempty"`
	CountId string `protobuf:"bytes,3,opt,name=countId,proto3" json:"countId,omitempty"`
	Value   string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NumberOfMoves) Reset() {
	*x = NumberOfMoves{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberOfMoves) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberOfMoves) ProtoMessage() {}

func (x *NumberOfMoves) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberOfMoves.ProtoReflect.Descriptor instead.
func (*NumberOfMoves) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{6}
}

func (x *NumberOfMoves) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NumberOfMoves) GetDevId() string {
	if x != nil {
		return x.DevId
	}
	return ""
}

func (x *NumberOfMoves) GetCountId() string {
	if x != nil {
		return x.CountId
	}
	return ""
}

func (x *NumberOfMoves) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type NameGasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FinId  string `protobuf:"bytes,2,opt,name=finId,proto3" json:"finId,omitempty"`
	NumId  string `protobuf:"bytes,3,opt,name=numId,proto3" json:"numId,omitempty"`
	PresId string `protobuf:"bytes,4,opt,name=presId,proto3" json:"presId,omitempty"`
	Title  string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *NameGasket) Reset() {
	*x = NameGasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameGasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameGasket) ProtoMessage() {}

func (x *NameGasket) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameGasket.ProtoReflect.Descriptor instead.
func (*NameGasket) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{7}
}

func (x *NameGasket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NameGasket) GetFinId() string {
	if x != nil {
		return x.FinId
	}
	return ""
}

func (x *NameGasket) GetNumId() string {
	if x != nil {
		return x.NumId
	}
	return ""
}

func (x *NameGasket) GetPresId() string {
	if x != nil {
		return x.PresId
	}
	return ""
}

func (x *NameGasket) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type NameGasketSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SizeLong  float64 `protobuf:"fixed64,2,opt,name=sizeLong,proto3" json:"sizeLong,omitempty"`
	SizeTrans float64 `protobuf:"fixed64,3,opt,name=sizeTrans,proto3" json:"sizeTrans,omitempty"`
	Width     float64 `protobuf:"fixed64,4,opt,name=width,proto3" json:"width,omitempty"`
	Thick1    float64 `protobuf:"fixed64,5,opt,name=thick1,proto3" json:"thick1,omitempty"`
	Thick2    float64 `protobuf:"fixed64,6,opt,name=thick2,proto3" json:"thick2,omitempty"`
	Thick3    float64 `protobuf:"fixed64,7,opt,name=thick3,proto3" json:"thick3,omitempty"`
	Thick4    float64 `protobuf:"fixed64,8,opt,name=thick4,proto3" json:"thick4,omitempty"`
}

func (x *NameGasketSize) Reset() {
	*x = NameGasketSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameGasketSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameGasketSize) ProtoMessage() {}

func (x *NameGasketSize) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameGasketSize.ProtoReflect.Descriptor instead.
func (*NameGasketSize) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{8}
}

func (x *NameGasketSize) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NameGasketSize) GetSizeLong() float64 {
	if x != nil {
		return x.SizeLong
	}
	return 0
}

func (x *NameGasketSize) GetSizeTrans() float64 {
	if x != nil {
		return x.SizeTrans
	}
	return 0
}

func (x *NameGasketSize) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *NameGasketSize) GetThick1() float64 {
	if x != nil {
		return x.Thick1
	}
	return 0
}

func (x *NameGasketSize) GetThick2() float64 {
	if x != nil {
		return x.Thick2
	}
	return 0
}

func (x *NameGasketSize) GetThick3() float64 {
	if x != nil {
		return x.Thick3
	}
	return 0
}

func (x *NameGasketSize) GetThick4() float64 {
	if x != nil {
		return x.Thick4
	}
	return 0
}

type FullNameGasket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FinId     string  `protobuf:"bytes,2,opt,name=finId,proto3" json:"finId,omitempty"`
	NumId     string  `protobuf:"bytes,3,opt,name=numId,proto3" json:"numId,omitempty"`
	PresId    string  `protobuf:"bytes,4,opt,name=presId,proto3" json:"presId,omitempty"`
	Title     string  `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	SizeLong  float64 `protobuf:"fixed64,6,opt,name=sizeLong,proto3" json:"sizeLong,omitempty"`
	SizeTrans float64 `protobuf:"fixed64,7,opt,name=sizeTrans,proto3" json:"sizeTrans,omitempty"`
	Width     float64 `protobuf:"fixed64,8,opt,name=width,proto3" json:"width,omitempty"`
	Thick1    float64 `protobuf:"fixed64,9,opt,name=thick1,proto3" json:"thick1,omitempty"`
	Thick2    float64 `protobuf:"fixed64,10,opt,name=thick2,proto3" json:"thick2,omitempty"`
	Thick3    float64 `protobuf:"fixed64,11,opt,name=thick3,proto3" json:"thick3,omitempty"`
	Thick4    float64 `protobuf:"fixed64,12,opt,name=thick4,proto3" json:"thick4,omitempty"`
}

func (x *FullNameGasket) Reset() {
	*x = FullNameGasket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_device_model_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullNameGasket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullNameGasket) ProtoMessage() {}

func (x *FullNameGasket) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_device_model_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullNameGasket.ProtoReflect.Descriptor instead.
func (*FullNameGasket) Descriptor() ([]byte, []int) {
	return file_moment_models_device_model_proto_rawDescGZIP(), []int{9}
}

func (x *FullNameGasket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FullNameGasket) GetFinId() string {
	if x != nil {
		return x.FinId
	}
	return ""
}

func (x *FullNameGasket) GetNumId() string {
	if x != nil {
		return x.NumId
	}
	return ""
}

func (x *FullNameGasket) GetPresId() string {
	if x != nil {
		return x.PresId
	}
	return ""
}

func (x *FullNameGasket) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FullNameGasket) GetSizeLong() float64 {
	if x != nil {
		return x.SizeLong
	}
	return 0
}

func (x *FullNameGasket) GetSizeTrans() float64 {
	if x != nil {
		return x.SizeTrans
	}
	return 0
}

func (x *FullNameGasket) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *FullNameGasket) GetThick1() float64 {
	if x != nil {
		return x.Thick1
	}
	return 0
}

func (x *FullNameGasket) GetThick2() float64 {
	if x != nil {
		return x.Thick2
	}
	return 0
}

func (x *FullNameGasket) GetThick3() float64 {
	if x != nil {
		return x.Thick3
	}
	return 0
}

func (x *FullNameGasket) GetThick4() float64 {
	if x != nil {
		return x.Thick4
	}
	return 0
}

var File_moment_models_device_model_proto protoreflect.FileDescriptor

var file_moment_models_device_model_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x30, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x54, 0x75, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4b, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x76, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x76, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x76, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x76, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x48, 0x0a, 0x0a, 0x54, 0x75, 0x62, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x76, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x64, 0x65, 0x76, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x65, 0x0a, 0x0d,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4d, 0x6f, 0x76, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x76, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65,
	0x76, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x76, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x47, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x73, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x0e,
	0x4e, 0x61, 0x6d, 0x65, 0x47, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x7a, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73,
	0x69, 0x7a, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x74, 0x68, 0x69, 0x63, 0x6b, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x32,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x32, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x74, 0x68, 0x69, 0x63, 0x6b, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x34,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x34, 0x22, 0xaa,
	0x02, 0x0a, 0x0e, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x47, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x73, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x69, 0x7a, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73,
	0x69, 0x7a, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x68, 0x69, 0x63, 0x6b, 0x31, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x68, 0x69,
	0x63, 0x6b, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x32, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x68, 0x69, 0x63, 0x6b, 0x33, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x68, 0x69,
	0x63, 0x6b, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x34, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x68, 0x69, 0x63, 0x6b, 0x34, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x6e,
	0x64, 0x65, 0x72, 0x32, 0x37, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x75, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moment_models_device_model_proto_rawDescOnce sync.Once
	file_moment_models_device_model_proto_rawDescData = file_moment_models_device_model_proto_rawDesc
)

func file_moment_models_device_model_proto_rawDescGZIP() []byte {
	file_moment_models_device_model_proto_rawDescOnce.Do(func() {
		file_moment_models_device_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_moment_models_device_model_proto_rawDescData)
	})
	return file_moment_models_device_model_proto_rawDescData
}

var file_moment_models_device_model_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_moment_models_device_model_proto_goTypes = []interface{}{
	(*Device)(nil),           // 0: device_model.Device
	(*Pressure)(nil),         // 1: device_model.Pressure
	(*TubeCount)(nil),        // 2: device_model.TubeCount
	(*FinningFactor)(nil),    // 3: device_model.FinningFactor
	(*SectionExecution)(nil), // 4: device_model.SectionExecution
	(*TubeLength)(nil),       // 5: device_model.TubeLength
	(*NumberOfMoves)(nil),    // 6: device_model.NumberOfMoves
	(*NameGasket)(nil),       // 7: device_model.NameGasket
	(*NameGasketSize)(nil),   // 8: device_model.NameGasketSize
	(*FullNameGasket)(nil),   // 9: device_model.FullNameGasket
}
var file_moment_models_device_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_moment_models_device_model_proto_init() }
func file_moment_models_device_model_proto_init() {
	if File_moment_models_device_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moment_models_device_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_moment_models_device_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pressure); i {
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
		file_moment_models_device_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TubeCount); i {
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
		file_moment_models_device_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinningFactor); i {
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
		file_moment_models_device_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionExecution); i {
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
		file_moment_models_device_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TubeLength); i {
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
		file_moment_models_device_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberOfMoves); i {
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
		file_moment_models_device_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameGasket); i {
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
		file_moment_models_device_model_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameGasketSize); i {
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
		file_moment_models_device_model_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullNameGasket); i {
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
			RawDescriptor: file_moment_models_device_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moment_models_device_model_proto_goTypes,
		DependencyIndexes: file_moment_models_device_model_proto_depIdxs,
		MessageInfos:      file_moment_models_device_model_proto_msgTypes,
	}.Build()
	File_moment_models_device_model_proto = out.File
	file_moment_models_device_model_proto_rawDesc = nil
	file_moment_models_device_model_proto_goTypes = nil
	file_moment_models_device_model_proto_depIdxs = nil
}
