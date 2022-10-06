// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: moment/models/material_model.proto

package material_model

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

type Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Material) Reset() {
	*x = Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_material_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Material) ProtoMessage() {}

func (x *Material) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_material_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Material.ProtoReflect.Descriptor instead.
func (*Material) Descriptor() ([]byte, []int) {
	return file_moment_models_material_model_proto_rawDescGZIP(), []int{0}
}

func (x *Material) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Material) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type MaterialWithIsEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title             string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	IsEmptyAlpha      bool   `protobuf:"varint,3,opt,name=IsEmptyAlpha,json=isEmptyAlpha,proto3" json:"IsEmptyAlpha,omitempty"`
	IsEmptyElasticity bool   `protobuf:"varint,4,opt,name=IsEmptyElasticity,json=isEmptyElasticity,proto3" json:"IsEmptyElasticity,omitempty"`
	IsEmptyVoltage    bool   `protobuf:"varint,5,opt,name=IsEmptyVoltage,json=isEmptyVoltage,proto3" json:"IsEmptyVoltage,omitempty"`
}

func (x *MaterialWithIsEmpty) Reset() {
	*x = MaterialWithIsEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_material_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialWithIsEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialWithIsEmpty) ProtoMessage() {}

func (x *MaterialWithIsEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_material_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialWithIsEmpty.ProtoReflect.Descriptor instead.
func (*MaterialWithIsEmpty) Descriptor() ([]byte, []int) {
	return file_moment_models_material_model_proto_rawDescGZIP(), []int{1}
}

func (x *MaterialWithIsEmpty) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialWithIsEmpty) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MaterialWithIsEmpty) GetIsEmptyAlpha() bool {
	if x != nil {
		return x.IsEmptyAlpha
	}
	return false
}

func (x *MaterialWithIsEmpty) GetIsEmptyElasticity() bool {
	if x != nil {
		return x.IsEmptyElasticity
	}
	return false
}

func (x *MaterialWithIsEmpty) GetIsEmptyVoltage() bool {
	if x != nil {
		return x.IsEmptyVoltage
	}
	return false
}

type Voltage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float64 `protobuf:"fixed64,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Voltage     float64 `protobuf:"fixed64,2,opt,name=voltage,proto3" json:"voltage,omitempty"`
}

func (x *Voltage) Reset() {
	*x = Voltage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_material_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voltage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voltage) ProtoMessage() {}

func (x *Voltage) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_material_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voltage.ProtoReflect.Descriptor instead.
func (*Voltage) Descriptor() ([]byte, []int) {
	return file_moment_models_material_model_proto_rawDescGZIP(), []int{2}
}

func (x *Voltage) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Voltage) GetVoltage() float64 {
	if x != nil {
		return x.Voltage
	}
	return 0
}

type Elasticity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float64 `protobuf:"fixed64,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Elasticity  float64 `protobuf:"fixed64,2,opt,name=elasticity,proto3" json:"elasticity,omitempty"`
}

func (x *Elasticity) Reset() {
	*x = Elasticity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_material_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Elasticity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Elasticity) ProtoMessage() {}

func (x *Elasticity) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_material_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Elasticity.ProtoReflect.Descriptor instead.
func (*Elasticity) Descriptor() ([]byte, []int) {
	return file_moment_models_material_model_proto_rawDescGZIP(), []int{3}
}

func (x *Elasticity) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Elasticity) GetElasticity() float64 {
	if x != nil {
		return x.Elasticity
	}
	return 0
}

type Alpha struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float64 `protobuf:"fixed64,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Alpha       float64 `protobuf:"fixed64,2,opt,name=alpha,proto3" json:"alpha,omitempty"`
}

func (x *Alpha) Reset() {
	*x = Alpha{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_models_material_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alpha) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alpha) ProtoMessage() {}

func (x *Alpha) ProtoReflect() protoreflect.Message {
	mi := &file_moment_models_material_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alpha.ProtoReflect.Descriptor instead.
func (*Alpha) Descriptor() ([]byte, []int) {
	return file_moment_models_material_model_proto_rawDescGZIP(), []int{4}
}

func (x *Alpha) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Alpha) GetAlpha() float64 {
	if x != nil {
		return x.Alpha
	}
	return 0
}

var File_moment_models_material_model_proto protoreflect.FileDescriptor

var file_moment_models_material_model_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x30, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x13, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x73, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x45, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x22, 0x45,
	0x0a, 0x07, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x76, 0x6f,
	0x6c, 0x74, 0x61, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x0a, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x69, 0x74, 0x79, 0x22, 0x3f, 0x0a, 0x05, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x32, 0x37,
	0x32, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x75, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moment_models_material_model_proto_rawDescOnce sync.Once
	file_moment_models_material_model_proto_rawDescData = file_moment_models_material_model_proto_rawDesc
)

func file_moment_models_material_model_proto_rawDescGZIP() []byte {
	file_moment_models_material_model_proto_rawDescOnce.Do(func() {
		file_moment_models_material_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_moment_models_material_model_proto_rawDescData)
	})
	return file_moment_models_material_model_proto_rawDescData
}

var file_moment_models_material_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_moment_models_material_model_proto_goTypes = []interface{}{
	(*Material)(nil),            // 0: material_model.Material
	(*MaterialWithIsEmpty)(nil), // 1: material_model.MaterialWithIsEmpty
	(*Voltage)(nil),             // 2: material_model.Voltage
	(*Elasticity)(nil),          // 3: material_model.Elasticity
	(*Alpha)(nil),               // 4: material_model.Alpha
}
var file_moment_models_material_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_moment_models_material_model_proto_init() }
func file_moment_models_material_model_proto_init() {
	if File_moment_models_material_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moment_models_material_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Material); i {
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
		file_moment_models_material_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialWithIsEmpty); i {
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
		file_moment_models_material_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voltage); i {
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
		file_moment_models_material_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Elasticity); i {
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
		file_moment_models_material_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alpha); i {
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
			RawDescriptor: file_moment_models_material_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_moment_models_material_model_proto_goTypes,
		DependencyIndexes: file_moment_models_material_model_proto_depIdxs,
		MessageInfos:      file_moment_models_material_model_proto_msgTypes,
	}.Build()
	File_moment_models_material_model_proto = out.File
	file_moment_models_material_model_proto_rawDesc = nil
	file_moment_models_material_model_proto_goTypes = nil
	file_moment_models_material_model_proto_depIdxs = nil
}