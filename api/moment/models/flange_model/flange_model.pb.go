// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: moment/models/flange_model.proto

package flange_model

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

type Bolt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Diameter float64 `protobuf:"fixed64,3,opt,name=diameter,proto3" json:"diameter,omitempty"`
	Area     float64 `protobuf:"fixed64,4,opt,name=area,proto3" json:"area,omitempty"`
	IsInch   bool    `protobuf:"varint,5,opt,name=isInch,proto3" json:"isInch,omitempty"`
}

func (x *Bolt) Reset() {
	*x = Bolt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_flange_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bolt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bolt) ProtoMessage() {}

func (x *Bolt) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_flange_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bolt.ProtoReflect.Descriptor instead.
func (*Bolt) Descriptor() ([]byte, []int) {
	return file_moment_models_flange_model_proto_rawDescGZIP(), []int{0}
}

func (x *Bolt) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bolt) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Bolt) GetDiameter() float64 {
	if x != nil {
		return x.Diameter
	}
	return 0
}

func (x *Bolt) GetArea() float64 {
	if x != nil {
		return x.Area
	}
	return 0
}

func (x *Bolt) GetIsInch() bool {
	if x != nil {
		return x.IsInch
	}
	return false
}

type TypeFlange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *TypeFlange) Reset() {
	*x = TypeFlange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_flange_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeFlange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeFlange) ProtoMessage() {}

func (x *TypeFlange) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_flange_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeFlange.ProtoReflect.Descriptor instead.
func (*TypeFlange) Descriptor() ([]byte, []int) {
	return file_moment_models_flange_model_proto_rawDescGZIP(), []int{1}
}

func (x *TypeFlange) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TypeFlange) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TypeFlange) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type StandartWithSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title          string                   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	TypeId         string                   `protobuf:"bytes,3,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Sizes          *BasisFlangeSizeResponse `protobuf:"bytes,4,opt,name=sizes,proto3" json:"sizes,omitempty"`
	TitleDn        string                   `protobuf:"bytes,5,opt,name=titleDn,proto3" json:"titleDn,omitempty"`
	TitlePn        string                   `protobuf:"bytes,6,opt,name=titlePn,proto3" json:"titlePn,omitempty"`
	IsNeedRow      bool                     `protobuf:"varint,7,opt,name=isNeedRow,proto3" json:"isNeedRow,omitempty"`
	Rows           []string                 `protobuf:"bytes,8,rep,name=rows,proto3" json:"rows,omitempty"`
	IsInch         bool                     `protobuf:"varint,9,opt,name=isInch,proto3" json:"isInch,omitempty"`
	HasDesignation bool                     `protobuf:"varint,10,opt,name=hasDesignation,proto3" json:"hasDesignation,omitempty"`
}

func (x *StandartWithSize) Reset() {
	*x = StandartWithSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_flange_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandartWithSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandartWithSize) ProtoMessage() {}

func (x *StandartWithSize) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_flange_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandartWithSize.ProtoReflect.Descriptor instead.
func (*StandartWithSize) Descriptor() ([]byte, []int) {
	return file_moment_models_flange_model_proto_rawDescGZIP(), []int{2}
}

func (x *StandartWithSize) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StandartWithSize) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *StandartWithSize) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *StandartWithSize) GetSizes() *BasisFlangeSizeResponse {
	if x != nil {
		return x.Sizes
	}
	return nil
}

func (x *StandartWithSize) GetTitleDn() string {
	if x != nil {
		return x.TitleDn
	}
	return ""
}

func (x *StandartWithSize) GetTitlePn() string {
	if x != nil {
		return x.TitlePn
	}
	return ""
}

func (x *StandartWithSize) GetIsNeedRow() bool {
	if x != nil {
		return x.IsNeedRow
	}
	return false
}

func (x *StandartWithSize) GetRows() []string {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *StandartWithSize) GetIsInch() bool {
	if x != nil {
		return x.IsInch
	}
	return false
}

func (x *StandartWithSize) GetHasDesignation() bool {
	if x != nil {
		return x.HasDesignation
	}
	return false
}

type Standart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title          string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	TypeId         string   `protobuf:"bytes,3,opt,name=typeId,proto3" json:"typeId,omitempty"`
	TitleDn        string   `protobuf:"bytes,4,opt,name=titleDn,proto3" json:"titleDn,omitempty"`
	TitlePn        string   `protobuf:"bytes,5,opt,name=titlePn,proto3" json:"titlePn,omitempty"`
	IsNeedRow      bool     `protobuf:"varint,6,opt,name=isNeedRow,proto3" json:"isNeedRow,omitempty"`
	Rows           []string `protobuf:"bytes,7,rep,name=rows,proto3" json:"rows,omitempty"`
	IsInch         bool     `protobuf:"varint,8,opt,name=isInch,proto3" json:"isInch,omitempty"`
	HasDesignation bool     `protobuf:"varint,9,opt,name=hasDesignation,proto3" json:"hasDesignation,omitempty"`
}

func (x *Standart) Reset() {
	*x = Standart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_flange_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Standart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Standart) ProtoMessage() {}

func (x *Standart) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_flange_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Standart.ProtoReflect.Descriptor instead.
func (*Standart) Descriptor() ([]byte, []int) {
	return file_moment_models_flange_model_proto_rawDescGZIP(), []int{3}
}

func (x *Standart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Standart) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Standart) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *Standart) GetTitleDn() string {
	if x != nil {
		return x.TitleDn
	}
	return ""
}

func (x *Standart) GetTitlePn() string {
	if x != nil {
		return x.TitlePn
	}
	return ""
}

func (x *Standart) GetIsNeedRow() bool {
	if x != nil {
		return x.IsNeedRow
	}
	return false
}

func (x *Standart) GetRows() []string {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *Standart) GetIsInch() bool {
	if x != nil {
		return x.IsInch
	}
	return false
}

func (x *Standart) GetHasDesignation() bool {
	if x != nil {
		return x.HasDesignation
	}
	return false
}

type BasisFlangeSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dn string                `protobuf:"bytes,1,opt,name=dn,proto3" json:"dn,omitempty"`
	Pn []*BasisFlangeSize_Pn `protobuf:"bytes,2,rep,name=pn,proto3" json:"pn,omitempty"`
}

func (x *BasisFlangeSize) Reset() {
	*x = BasisFlangeSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_flange_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasisFlangeSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasisFlangeSize) ProtoMessage() {}

func (x *BasisFlangeSize) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_flange_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasisFlangeSize.ProtoReflect.Descriptor instead.
func (*BasisFlangeSize) Descriptor() ([]byte, []int) {
	return file_moment_models_flange_model_proto_rawDescGZIP(), []int{4}
}

func (x *BasisFlangeSize) GetDn() string {
	if x != nil {
		return x.Dn
	}
	return ""
}

func (x *BasisFlangeSize) GetPn() []*BasisFlangeSize_Pn {
	if x != nil {
		return x.Pn
	}
	return nil
}

type BasisFlangeSizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SizeRow1 []*BasisFlangeSize `protobuf:"bytes,1,rep,name=sizeRow1,proto3" json:"sizeRow1,omitempty"`
	SizeRow2 []*BasisFlangeSize `protobuf:"bytes,2,rep,name=sizeRow2,proto3" json:"sizeRow2,omitempty"`
}

func (x *BasisFlangeSizeResponse) Reset() {
	*x = BasisFlangeSizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_flange_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasisFlangeSizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasisFlangeSizeResponse) ProtoMessage() {}

func (x *BasisFlangeSizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_flange_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasisFlangeSizeResponse.ProtoReflect.Descriptor instead.
func (*BasisFlangeSizeResponse) Descriptor() ([]byte, []int) {
	return file_moment_models_flange_model_proto_rawDescGZIP(), []int{5}
}

func (x *BasisFlangeSizeResponse) GetSizeRow1() []*BasisFlangeSize {
	if x != nil {
		return x.SizeRow1
	}
	return nil
}

func (x *BasisFlangeSizeResponse) GetSizeRow2() []*BasisFlangeSize {
	if x != nil {
		return x.SizeRow2
	}
	return nil
}

type FullFlangeSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StandId string  `protobuf:"bytes,2,opt,name=standId,proto3" json:"standId,omitempty"`
	Pn      float64 `protobuf:"fixed64,3,opt,name=pn,proto3" json:"pn,omitempty"`
	D       float64 `protobuf:"fixed64,4,opt,name=d,proto3" json:"d,omitempty"`
	D6      float64 `protobuf:"fixed64,5,opt,name=d6,proto3" json:"d6,omitempty"`
	DOut    float64 `protobuf:"fixed64,6,opt,name=dOut,proto3" json:"dOut,omitempty"`
	H       float64 `protobuf:"fixed64,7,opt,name=h,proto3" json:"h,omitempty"`
	S0      float64 `protobuf:"fixed64,8,opt,name=s0,proto3" json:"s0,omitempty"`
	S1      float64 `protobuf:"fixed64,9,opt,name=s1,proto3" json:"s1,omitempty"`
	Length  float64 `protobuf:"fixed64,10,opt,name=length,proto3" json:"length,omitempty"`
	Count   int32   `protobuf:"varint,11,opt,name=count,proto3" json:"count,omitempty"`
	BoltId  string  `protobuf:"bytes,12,opt,name=boltId,proto3" json:"boltId,omitempty"`
	Dn      string  `protobuf:"bytes,13,opt,name=dn,proto3" json:"dn,omitempty"`
	Dmm     float64 `protobuf:"fixed64,14,opt,name=dmm,proto3" json:"dmm,omitempty"`
	X       float64 `protobuf:"fixed64,15,opt,name=x,proto3" json:"x,omitempty"`
	A       float64 `protobuf:"fixed64,16,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *FullFlangeSize) Reset() {
	*x = FullFlangeSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_flange_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullFlangeSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullFlangeSize) ProtoMessage() {}

func (x *FullFlangeSize) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_flange_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullFlangeSize.ProtoReflect.Descriptor instead.
func (*FullFlangeSize) Descriptor() ([]byte, []int) {
	return file_moment_models_flange_model_proto_rawDescGZIP(), []int{6}
}

func (x *FullFlangeSize) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FullFlangeSize) GetStandId() string {
	if x != nil {
		return x.StandId
	}
	return ""
}

func (x *FullFlangeSize) GetPn() float64 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *FullFlangeSize) GetD() float64 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *FullFlangeSize) GetD6() float64 {
	if x != nil {
		return x.D6
	}
	return 0
}

func (x *FullFlangeSize) GetDOut() float64 {
	if x != nil {
		return x.DOut
	}
	return 0
}

func (x *FullFlangeSize) GetH() float64 {
	if x != nil {
		return x.H
	}
	return 0
}

func (x *FullFlangeSize) GetS0() float64 {
	if x != nil {
		return x.S0
	}
	return 0
}

func (x *FullFlangeSize) GetS1() float64 {
	if x != nil {
		return x.S1
	}
	return 0
}

func (x *FullFlangeSize) GetLength() float64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *FullFlangeSize) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FullFlangeSize) GetBoltId() string {
	if x != nil {
		return x.BoltId
	}
	return ""
}

func (x *FullFlangeSize) GetDn() string {
	if x != nil {
		return x.Dn
	}
	return ""
}

func (x *FullFlangeSize) GetDmm() float64 {
	if x != nil {
		return x.Dmm
	}
	return 0
}

func (x *FullFlangeSize) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *FullFlangeSize) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

type BasisFlangeSize_Pn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pn       float64 `protobuf:"fixed64,1,opt,name=pn,proto3" json:"pn,omitempty"`
	IsEmptyD bool    `protobuf:"varint,2,opt,name=isEmptyD,proto3" json:"isEmptyD,omitempty"`
}

func (x *BasisFlangeSize_Pn) Reset() {
	*x = BasisFlangeSize_Pn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_flange_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasisFlangeSize_Pn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasisFlangeSize_Pn) ProtoMessage() {}

func (x *BasisFlangeSize_Pn) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_flange_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasisFlangeSize_Pn.ProtoReflect.Descriptor instead.
func (*BasisFlangeSize_Pn) Descriptor() ([]byte, []int) {
	return file_moment_models_flange_model_proto_rawDescGZIP(), []int{4, 0}
}

func (x *BasisFlangeSize_Pn) GetPn() float64 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *BasisFlangeSize_Pn) GetIsEmptyD() bool {
	if x != nil {
		return x.IsEmptyD
	}
	return false
}

var File_moment_models_flange_model_proto protoreflect.FileDescriptor

var file_moment_models_flange_model_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0x74, 0x0a, 0x04, 0x42, 0x6f, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x64, 0x69, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x49, 0x6e, 0x63, 0x68, 0x22, 0x48, 0x0a, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x46, 0x6c,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x22, 0xb3, 0x02, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x73, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x44, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x44, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x50, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x50, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x52, 0x6f,
	0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x52,
	0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x63, 0x68,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x63, 0x68, 0x12, 0x26,
	0x0a, 0x0e, 0x68, 0x61, 0x73, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x44, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xee, 0x01, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x44, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x44, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x50, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x50, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x52,
	0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64,
	0x52, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x63,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x63, 0x68, 0x12,
	0x26, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x69,
	0x73, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x64,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x6e, 0x12, 0x30, 0x0a, 0x02, 0x70,
	0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x73, 0x46, 0x6c, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x2e, 0x50, 0x6e, 0x52, 0x02, 0x70, 0x6e, 0x1a, 0x30, 0x0a,
	0x02, 0x50, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x02, 0x70, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x22,
	0x8f, 0x01, 0x0a, 0x17, 0x42, 0x61, 0x73, 0x69, 0x73, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x73,
	0x69, 0x7a, 0x65, 0x52, 0x6f, 0x77, 0x31, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73,
	0x69, 0x73, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x73, 0x69,
	0x7a, 0x65, 0x52, 0x6f, 0x77, 0x31, 0x12, 0x39, 0x0a, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x6f,
	0x77, 0x32, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6c, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x73, 0x46, 0x6c, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x6f, 0x77,
	0x32, 0x22, 0xae, 0x02, 0x0a, 0x0e, 0x46, 0x75, 0x6c, 0x6c, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x70, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x70, 0x6e, 0x12, 0x0c,
	0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x64, 0x36, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x64, 0x36, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x4f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x64, 0x4f, 0x75, 0x74,
	0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x68, 0x12, 0x0e,
	0x0a, 0x02, 0x73, 0x30, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x73, 0x30, 0x12, 0x0e,
	0x0a, 0x02, 0x73, 0x31, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x73, 0x31, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x6f, 0x6c, 0x74, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f,
	0x6c, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x64, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6d, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x64, 0x6d, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x01, 0x61, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x32, 0x37, 0x32, 0x2f, 0x73, 0x65,
	0x61, 0x6c, 0x75, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x66, 0x6c, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_moment_models_flange_model_proto_rawDescOnce sync.Once
	file_moment_models_flange_model_proto_rawDescData = file_moment_models_flange_model_proto_rawDesc
)

func file_moment_models_flange_model_proto_rawDescGZIP() []byte {
	file_moment_models_flange_model_proto_rawDescOnce.Do(func() {
		file_moment_models_flange_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_moment_models_flange_model_proto_rawDescData)
	})
	return file_moment_models_flange_model_proto_rawDescData
}

var file_moment_models_flange_model_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_moment_models_flange_model_proto_goTypes = []interface{}{
	(*Bolt)(nil),                    // 0: flange_model.Bolt
	(*TypeFlange)(nil),              // 1: flange_model.TypeFlange
	(*StandartWithSize)(nil),        // 2: flange_model.StandartWithSize
	(*Standart)(nil),                // 3: flange_model.Standart
	(*BasisFlangeSize)(nil),         // 4: flange_model.BasisFlangeSize
	(*BasisFlangeSizeResponse)(nil), // 5: flange_model.BasisFlangeSizeResponse
	(*FullFlangeSize)(nil),          // 6: flange_model.FullFlangeSize
	(*BasisFlangeSize_Pn)(nil),      // 7: flange_model.BasisFlangeSize.Pn
}
var file_moment_models_flange_model_proto_depIdxs = []int32{
	5, // 0: flange_model.StandartWithSize.sizes:type_name -> flange_model.BasisFlangeSizeResponse
	7, // 1: flange_model.BasisFlangeSize.pn:type_name -> flange_model.BasisFlangeSize.Pn
	4, // 2: flange_model.BasisFlangeSizeResponse.sizeRow1:type_name -> flange_model.BasisFlangeSize
	4, // 3: flange_model.BasisFlangeSizeResponse.sizeRow2:type_name -> flange_model.BasisFlangeSize
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_moment_models_flange_model_proto_init() }
func file_moment_models_flange_model_proto_init() {
	if File_moment_models_flange_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moment_models_flange_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bolt); i {
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
		file_moment_models_flange_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeFlange); i {
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
		file_moment_models_flange_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandartWithSize); i {
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
		file_moment_models_flange_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Standart); i {
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
		file_moment_models_flange_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasisFlangeSize); i {
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
		file_moment_models_flange_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasisFlangeSizeResponse); i {
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
		file_moment_models_flange_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullFlangeSize); i {
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
		file_moment_models_flange_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasisFlangeSize_Pn); i {
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
			RawDescriptor: file_moment_models_flange_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moment_models_flange_model_proto_goTypes,
		DependencyIndexes: file_moment_models_flange_model_proto_depIdxs,
		MessageInfos:      file_moment_models_flange_model_proto_msgTypes,
	}.Build()
	File_moment_models_flange_model_proto = out.File
	file_moment_models_flange_model_proto_rawDesc = nil
	file_moment_models_flange_model_proto_goTypes = nil
	file_moment_models_flange_model_proto_depIdxs = nil
}
