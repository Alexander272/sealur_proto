// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/rings_kit_construction_api.proto

package rings_kit_construction_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	rings_kit_construction_model "github.com/Alexander272/sealur_proto/api/pro/models/rings_kit_construction_model"
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

type GetRingsKitConstructions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRingsKitConstructions) Reset() {
	*x = GetRingsKitConstructions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_rings_kit_construction_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRingsKitConstructions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRingsKitConstructions) ProtoMessage() {}

func (x *GetRingsKitConstructions) ProtoReflect() protoreflect.Message {
	mi := &file_pro_rings_kit_construction_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRingsKitConstructions.ProtoReflect.Descriptor instead.
func (*GetRingsKitConstructions) Descriptor() ([]byte, []int) {
	return file_pro_rings_kit_construction_api_proto_rawDescGZIP(), []int{0}
}

type CreateRingsKitConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code          string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Image         string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	SameRings     bool   `protobuf:"varint,4,opt,name=sameRings,proto3" json:"sameRings,omitempty"`
	MaterialTypes string `protobuf:"bytes,5,opt,name=materialTypes,proto3" json:"materialTypes,omitempty"`
	HasThickness  bool   `protobuf:"varint,6,opt,name=hasThickness,proto3" json:"hasThickness,omitempty"`
	DefaultCount  string `protobuf:"bytes,7,opt,name=defaultCount,proto3" json:"defaultCount,omitempty"`
}

func (x *CreateRingsKitConstruction) Reset() {
	*x = CreateRingsKitConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_rings_kit_construction_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRingsKitConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRingsKitConstruction) ProtoMessage() {}

func (x *CreateRingsKitConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_rings_kit_construction_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRingsKitConstruction.ProtoReflect.Descriptor instead.
func (*CreateRingsKitConstruction) Descriptor() ([]byte, []int) {
	return file_pro_rings_kit_construction_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRingsKitConstruction) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateRingsKitConstruction) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRingsKitConstruction) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateRingsKitConstruction) GetSameRings() bool {
	if x != nil {
		return x.SameRings
	}
	return false
}

func (x *CreateRingsKitConstruction) GetMaterialTypes() string {
	if x != nil {
		return x.MaterialTypes
	}
	return ""
}

func (x *CreateRingsKitConstruction) GetHasThickness() bool {
	if x != nil {
		return x.HasThickness
	}
	return false
}

func (x *CreateRingsKitConstruction) GetDefaultCount() string {
	if x != nil {
		return x.DefaultCount
	}
	return ""
}

type UpdateRingsKitConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code          string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Image         string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	SameRings     bool   `protobuf:"varint,5,opt,name=sameRings,proto3" json:"sameRings,omitempty"`
	MaterialTypes string `protobuf:"bytes,6,opt,name=materialTypes,proto3" json:"materialTypes,omitempty"`
	HasThickness  bool   `protobuf:"varint,7,opt,name=hasThickness,proto3" json:"hasThickness,omitempty"`
	DefaultCount  string `protobuf:"bytes,8,opt,name=defaultCount,proto3" json:"defaultCount,omitempty"`
}

func (x *UpdateRingsKitConstruction) Reset() {
	*x = UpdateRingsKitConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_rings_kit_construction_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRingsKitConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRingsKitConstruction) ProtoMessage() {}

func (x *UpdateRingsKitConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_rings_kit_construction_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRingsKitConstruction.ProtoReflect.Descriptor instead.
func (*UpdateRingsKitConstruction) Descriptor() ([]byte, []int) {
	return file_pro_rings_kit_construction_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRingsKitConstruction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRingsKitConstruction) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRingsKitConstruction) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateRingsKitConstruction) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UpdateRingsKitConstruction) GetSameRings() bool {
	if x != nil {
		return x.SameRings
	}
	return false
}

func (x *UpdateRingsKitConstruction) GetMaterialTypes() string {
	if x != nil {
		return x.MaterialTypes
	}
	return ""
}

func (x *UpdateRingsKitConstruction) GetHasThickness() bool {
	if x != nil {
		return x.HasThickness
	}
	return false
}

func (x *UpdateRingsKitConstruction) GetDefaultCount() string {
	if x != nil {
		return x.DefaultCount
	}
	return ""
}

type DeleteRingsKitConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRingsKitConstruction) Reset() {
	*x = DeleteRingsKitConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_rings_kit_construction_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRingsKitConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRingsKitConstruction) ProtoMessage() {}

func (x *DeleteRingsKitConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_rings_kit_construction_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRingsKitConstruction.ProtoReflect.Descriptor instead.
func (*DeleteRingsKitConstruction) Descriptor() ([]byte, []int) {
	return file_pro_rings_kit_construction_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRingsKitConstruction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RingsKitConstructions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RingsKitConstructions *rings_kit_construction_model.RingsKitConstructionMap `protobuf:"bytes,1,opt,name=ringsKitConstructions,proto3" json:"ringsKitConstructions,omitempty"`
}

func (x *RingsKitConstructions) Reset() {
	*x = RingsKitConstructions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_rings_kit_construction_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RingsKitConstructions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RingsKitConstructions) ProtoMessage() {}

func (x *RingsKitConstructions) ProtoReflect() protoreflect.Message {
	mi := &file_pro_rings_kit_construction_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RingsKitConstructions.ProtoReflect.Descriptor instead.
func (*RingsKitConstructions) Descriptor() ([]byte, []int) {
	return file_pro_rings_kit_construction_api_proto_rawDescGZIP(), []int{4}
}

func (x *RingsKitConstructions) GetRingsKitConstructions() *rings_kit_construction_model.RingsKitConstructionMap {
	if x != nil {
		return x.RingsKitConstructions
	}
	return nil
}

var File_pro_rings_kit_construction_api_proto protoreflect.FileDescriptor

var file_pro_rings_kit_construction_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6b, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6b, 0x69,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6b, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe8,
	0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x61, 0x6d, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x73, 0x61, 0x6d, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x54, 0x68, 0x69, 0x63, 0x6b, 0x6e, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x54, 0x68, 0x69, 0x63,
	0x6b, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf8, 0x01, 0x0a, 0x1a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x61, 0x6d, 0x65,
	0x52, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x61, 0x6d,
	0x65, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x68, 0x61, 0x73, 0x54, 0x68, 0x69, 0x63, 0x6b, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x54, 0x68, 0x69, 0x63, 0x6b, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69,
	0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6b, 0x0a, 0x15,
	0x72, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x73, 0x5f, 0x6b, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x73,
	0x4b, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x61, 0x70, 0x52, 0x15, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xa4, 0x03, 0x0a, 0x1b, 0x52, 0x69,
	0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x34, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6b, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x31, 0x2e, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x5f, 0x6b, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5a, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6b,
	0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b,
	0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x36, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6b, 0x69, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x36,
	0x2e, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6b, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x4b, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x24, 0x5a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x5f, 0x6b, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_rings_kit_construction_api_proto_rawDescOnce sync.Once
	file_pro_rings_kit_construction_api_proto_rawDescData = file_pro_rings_kit_construction_api_proto_rawDesc
)

func file_pro_rings_kit_construction_api_proto_rawDescGZIP() []byte {
	file_pro_rings_kit_construction_api_proto_rawDescOnce.Do(func() {
		file_pro_rings_kit_construction_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_rings_kit_construction_api_proto_rawDescData)
	})
	return file_pro_rings_kit_construction_api_proto_rawDescData
}

var file_pro_rings_kit_construction_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_rings_kit_construction_api_proto_goTypes = []interface{}{
	(*GetRingsKitConstructions)(nil),                             // 0: rings_kit_construction_api.GetRingsKitConstructions
	(*CreateRingsKitConstruction)(nil),                           // 1: rings_kit_construction_api.CreateRingsKitConstruction
	(*UpdateRingsKitConstruction)(nil),                           // 2: rings_kit_construction_api.UpdateRingsKitConstruction
	(*DeleteRingsKitConstruction)(nil),                           // 3: rings_kit_construction_api.DeleteRingsKitConstruction
	(*RingsKitConstructions)(nil),                                // 4: rings_kit_construction_api.RingsKitConstructions
	(*rings_kit_construction_model.RingsKitConstructionMap)(nil), // 5: rings_kit_construction_model.RingsKitConstructionMap
	(*response_model.Response)(nil),                              // 6: response_model.Response
}
var file_pro_rings_kit_construction_api_proto_depIdxs = []int32{
	5, // 0: rings_kit_construction_api.RingsKitConstructions.ringsKitConstructions:type_name -> rings_kit_construction_model.RingsKitConstructionMap
	0, // 1: rings_kit_construction_api.RingsKitConstructionService.GetAll:input_type -> rings_kit_construction_api.GetRingsKitConstructions
	1, // 2: rings_kit_construction_api.RingsKitConstructionService.Create:input_type -> rings_kit_construction_api.CreateRingsKitConstruction
	2, // 3: rings_kit_construction_api.RingsKitConstructionService.Update:input_type -> rings_kit_construction_api.UpdateRingsKitConstruction
	3, // 4: rings_kit_construction_api.RingsKitConstructionService.Delete:input_type -> rings_kit_construction_api.DeleteRingsKitConstruction
	4, // 5: rings_kit_construction_api.RingsKitConstructionService.GetAll:output_type -> rings_kit_construction_api.RingsKitConstructions
	6, // 6: rings_kit_construction_api.RingsKitConstructionService.Create:output_type -> response_model.Response
	6, // 7: rings_kit_construction_api.RingsKitConstructionService.Update:output_type -> response_model.Response
	6, // 8: rings_kit_construction_api.RingsKitConstructionService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_rings_kit_construction_api_proto_init() }
func file_pro_rings_kit_construction_api_proto_init() {
	if File_pro_rings_kit_construction_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_rings_kit_construction_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRingsKitConstructions); i {
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
		file_pro_rings_kit_construction_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRingsKitConstruction); i {
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
		file_pro_rings_kit_construction_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRingsKitConstruction); i {
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
		file_pro_rings_kit_construction_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRingsKitConstruction); i {
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
		file_pro_rings_kit_construction_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingsKitConstructions); i {
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
			RawDescriptor: file_pro_rings_kit_construction_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_rings_kit_construction_api_proto_goTypes,
		DependencyIndexes: file_pro_rings_kit_construction_api_proto_depIdxs,
		MessageInfos:      file_pro_rings_kit_construction_api_proto_msgTypes,
	}.Build()
	File_pro_rings_kit_construction_api_proto = out.File
	file_pro_rings_kit_construction_api_proto_rawDesc = nil
	file_pro_rings_kit_construction_api_proto_goTypes = nil
	file_pro_rings_kit_construction_api_proto_depIdxs = nil
}
