// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pro/models/putg_data_model.proto

package putg_data_model

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

type PutgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HasJumper    bool   `protobuf:"varint,2,opt,name=hasJumper,proto3" json:"hasJumper,omitempty"`
	HasHole      bool   `protobuf:"varint,3,opt,name=hasHole,proto3" json:"hasHole,omitempty"`
	HasRemovable bool   `protobuf:"varint,4,opt,name=hasRemovable,proto3" json:"hasRemovable,omitempty"`
	HasMounting  bool   `protobuf:"varint,5,opt,name=hasMounting,proto3" json:"hasMounting,omitempty"`
	HasCoating   bool   `protobuf:"varint,6,opt,name=hasCoating,proto3" json:"hasCoating,omitempty"`
}

func (x *PutgData) Reset() {
	*x = PutgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_putg_data_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutgData) ProtoMessage() {}

func (x *PutgData) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_putg_data_model_proto_msgTypes[0]
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
	return file_pro_models_putg_data_model_proto_rawDescGZIP(), []int{0}
}

func (x *PutgData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutgData) GetHasJumper() bool {
	if x != nil {
		return x.HasJumper
	}
	return false
}

func (x *PutgData) GetHasHole() bool {
	if x != nil {
		return x.HasHole
	}
	return false
}

func (x *PutgData) GetHasRemovable() bool {
	if x != nil {
		return x.HasRemovable
	}
	return false
}

func (x *PutgData) GetHasMounting() bool {
	if x != nil {
		return x.HasMounting
	}
	return false
}

func (x *PutgData) GetHasCoating() bool {
	if x != nil {
		return x.HasCoating
	}
	return false
}

var File_pro_models_putg_data_model_proto protoreflect.FileDescriptor

var file_pro_models_putg_data_model_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74,
	0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0xb8, 0x01, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x4a, 0x75, 0x6d, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x4a, 0x75, 0x6d, 0x70, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x68, 0x61, 0x73, 0x48, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x68, 0x61, 0x73, 0x48, 0x6f, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x68, 0x61, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x68, 0x61, 0x73, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e,
	0x0a, 0x0a, 0x68, 0x61, 0x73, 0x43, 0x6f, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x43, 0x6f, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x45,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65,
	0x78, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x32, 0x37, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x75, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_models_putg_data_model_proto_rawDescOnce sync.Once
	file_pro_models_putg_data_model_proto_rawDescData = file_pro_models_putg_data_model_proto_rawDesc
)

func file_pro_models_putg_data_model_proto_rawDescGZIP() []byte {
	file_pro_models_putg_data_model_proto_rawDescOnce.Do(func() {
		file_pro_models_putg_data_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_models_putg_data_model_proto_rawDescData)
	})
	return file_pro_models_putg_data_model_proto_rawDescData
}

var file_pro_models_putg_data_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pro_models_putg_data_model_proto_goTypes = []interface{}{
	(*PutgData)(nil), // 0: putg_data_model.PutgData
}
var file_pro_models_putg_data_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pro_models_putg_data_model_proto_init() }
func file_pro_models_putg_data_model_proto_init() {
	if File_pro_models_putg_data_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_models_putg_data_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pro_models_putg_data_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pro_models_putg_data_model_proto_goTypes,
		DependencyIndexes: file_pro_models_putg_data_model_proto_depIdxs,
		MessageInfos:      file_pro_models_putg_data_model_proto_msgTypes,
	}.Build()
	File_pro_models_putg_data_model_proto = out.File
	file_pro_models_putg_data_model_proto_rawDesc = nil
	file_pro_models_putg_data_model_proto_goTypes = nil
	file_pro_models_putg_data_model_proto_depIdxs = nil
}
