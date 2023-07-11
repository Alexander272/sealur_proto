// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/models/putg_construction_type_model.proto

package putg_construction_type_model

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

type PutgConstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code          string        `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	HasD4         bool          `protobuf:"varint,4,opt,name=hasD4,proto3" json:"hasD4,omitempty"`
	HasD3         bool          `protobuf:"varint,5,opt,name=hasD3,proto3" json:"hasD3,omitempty"`
	HasD2         bool          `protobuf:"varint,6,opt,name=hasD2,proto3" json:"hasD2,omitempty"`
	HasD1         bool          `protobuf:"varint,7,opt,name=hasD1,proto3" json:"hasD1,omitempty"`
	HasRotaryPlug bool          `protobuf:"varint,8,opt,name=hasRotaryPlug,proto3" json:"hasRotaryPlug,omitempty"`
	HasInnerRing  bool          `protobuf:"varint,9,opt,name=hasInnerRing,proto3" json:"hasInnerRing,omitempty"`
	HasOuterRing  bool          `protobuf:"varint,10,opt,name=hasOuterRing,proto3" json:"hasOuterRing,omitempty"`
	BaseId        string        `protobuf:"bytes,11,opt,name=baseId,proto3" json:"baseId,omitempty"`
	Description   string        `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	MinWidth      float64       `protobuf:"fixed64,13,opt,name=minWidth,proto3" json:"minWidth,omitempty"`
	JumperRange   []int64       `protobuf:"varint,14,rep,packed,name=jumperRange,proto3" json:"jumperRange,omitempty"`
	WidthRange    []*WidthRange `protobuf:"bytes,15,rep,name=widthRange,proto3" json:"widthRange,omitempty"`
}

func (x *PutgConstruction) Reset() {
	*x = PutgConstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_putg_construction_type_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutgConstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutgConstruction) ProtoMessage() {}

func (x *PutgConstruction) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_putg_construction_type_model_proto_msgTypes[0]
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
	return file_pro_models_putg_construction_type_model_proto_rawDescGZIP(), []int{0}
}

func (x *PutgConstruction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutgConstruction) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PutgConstruction) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PutgConstruction) GetHasD4() bool {
	if x != nil {
		return x.HasD4
	}
	return false
}

func (x *PutgConstruction) GetHasD3() bool {
	if x != nil {
		return x.HasD3
	}
	return false
}

func (x *PutgConstruction) GetHasD2() bool {
	if x != nil {
		return x.HasD2
	}
	return false
}

func (x *PutgConstruction) GetHasD1() bool {
	if x != nil {
		return x.HasD1
	}
	return false
}

func (x *PutgConstruction) GetHasRotaryPlug() bool {
	if x != nil {
		return x.HasRotaryPlug
	}
	return false
}

func (x *PutgConstruction) GetHasInnerRing() bool {
	if x != nil {
		return x.HasInnerRing
	}
	return false
}

func (x *PutgConstruction) GetHasOuterRing() bool {
	if x != nil {
		return x.HasOuterRing
	}
	return false
}

func (x *PutgConstruction) GetBaseId() string {
	if x != nil {
		return x.BaseId
	}
	return ""
}

func (x *PutgConstruction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PutgConstruction) GetMinWidth() float64 {
	if x != nil {
		return x.MinWidth
	}
	return 0
}

func (x *PutgConstruction) GetJumperRange() []int64 {
	if x != nil {
		return x.JumperRange
	}
	return nil
}

func (x *PutgConstruction) GetWidthRange() []*WidthRange {
	if x != nil {
		return x.WidthRange
	}
	return nil
}

type WidthRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxD3 float64 `protobuf:"fixed64,1,opt,name=maxD3,proto3" json:"maxD3,omitempty"`
	Width float64 `protobuf:"fixed64,2,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *WidthRange) Reset() {
	*x = WidthRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_putg_construction_type_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidthRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidthRange) ProtoMessage() {}

func (x *WidthRange) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_putg_construction_type_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidthRange.ProtoReflect.Descriptor instead.
func (*WidthRange) Descriptor() ([]byte, []int) {
	return file_pro_models_putg_construction_type_model_proto_rawDescGZIP(), []int{1}
}

func (x *WidthRange) GetMaxD3() float64 {
	if x != nil {
		return x.MaxD3
	}
	return 0
}

func (x *WidthRange) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

var File_pro_models_putg_construction_type_model_proto protoreflect.FileDescriptor

var file_pro_models_putg_construction_type_model_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xd4, 0x03,
	0x0a, 0x10, 0x50, 0x75, 0x74, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x68, 0x61, 0x73, 0x44, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x61, 0x73,
	0x44, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x73, 0x44, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x68, 0x61, 0x73, 0x44, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x73, 0x44,
	0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x61, 0x73, 0x44, 0x32, 0x12, 0x14,
	0x0a, 0x05, 0x68, 0x61, 0x73, 0x44, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68,
	0x61, 0x73, 0x44, 0x31, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x72,
	0x79, 0x50, 0x6c, 0x75, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73,
	0x52, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61,
	0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x68, 0x61, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x22,
	0x0a, 0x0c, 0x68, 0x61, 0x73, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x69, 0x6e, 0x67, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x69,
	0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x69, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6d, 0x69, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x75, 0x6d, 0x70,
	0x65, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x6a,
	0x75, 0x6d, 0x70, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x57, 0x69,
	0x64, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0a, 0x77, 0x69, 0x64, 0x74, 0x68, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x22, 0x38, 0x0a, 0x0a, 0x57, 0x69, 0x64, 0x74, 0x68, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x78, 0x44, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x44, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x42, 0x52,
	0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65,
	0x78, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x32, 0x37, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x75, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_models_putg_construction_type_model_proto_rawDescOnce sync.Once
	file_pro_models_putg_construction_type_model_proto_rawDescData = file_pro_models_putg_construction_type_model_proto_rawDesc
)

func file_pro_models_putg_construction_type_model_proto_rawDescGZIP() []byte {
	file_pro_models_putg_construction_type_model_proto_rawDescOnce.Do(func() {
		file_pro_models_putg_construction_type_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_models_putg_construction_type_model_proto_rawDescData)
	})
	return file_pro_models_putg_construction_type_model_proto_rawDescData
}

var file_pro_models_putg_construction_type_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pro_models_putg_construction_type_model_proto_goTypes = []interface{}{
	(*PutgConstruction)(nil), // 0: putg_construction_type_model.PutgConstruction
	(*WidthRange)(nil),       // 1: putg_construction_type_model.WidthRange
}
var file_pro_models_putg_construction_type_model_proto_depIdxs = []int32{
	1, // 0: putg_construction_type_model.PutgConstruction.widthRange:type_name -> putg_construction_type_model.WidthRange
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_models_putg_construction_type_model_proto_init() }
func file_pro_models_putg_construction_type_model_proto_init() {
	if File_pro_models_putg_construction_type_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_models_putg_construction_type_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pro_models_putg_construction_type_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidthRange); i {
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
			RawDescriptor: file_pro_models_putg_construction_type_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pro_models_putg_construction_type_model_proto_goTypes,
		DependencyIndexes: file_pro_models_putg_construction_type_model_proto_depIdxs,
		MessageInfos:      file_pro_models_putg_construction_type_model_proto_msgTypes,
	}.Build()
	File_pro_models_putg_construction_type_model_proto = out.File
	file_pro_models_putg_construction_type_model_proto_rawDesc = nil
	file_pro_models_putg_construction_type_model_proto_goTypes = nil
	file_pro_models_putg_construction_type_model_proto_depIdxs = nil
}
