// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pro/models/putg_standard_model.proto

package putg_standard_model

import (
	flange_standard_model "github.com/Alexander272/sealur_proto/api/pro/models/flange_standard_model"
	standard_model "github.com/Alexander272/sealur_proto/api/pro/models/standard_model"
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

type PutgStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// TODO title & code тут похоже не нужны
	Title          string                                `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code           string                                `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	DnTitle        string                                `protobuf:"bytes,4,opt,name=dnTitle,proto3" json:"dnTitle,omitempty"`
	PnTitle        string                                `protobuf:"bytes,5,opt,name=pnTitle,proto3" json:"pnTitle,omitempty"`
	Standard       *standard_model.Standard              `protobuf:"bytes,6,opt,name=standard,proto3" json:"standard,omitempty"`
	FlangeStandard *flange_standard_model.FlangeStandard `protobuf:"bytes,7,opt,name=flangeStandard,proto3" json:"flangeStandard,omitempty"`
}

func (x *PutgStandard) Reset() {
	*x = PutgStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_models_putg_standard_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutgStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutgStandard) ProtoMessage() {}

func (x *PutgStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_models_putg_standard_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutgStandard.ProtoReflect.Descriptor instead.
func (*PutgStandard) Descriptor() ([]byte, []int) {
	return file_pro_models_putg_standard_model_proto_rawDescGZIP(), []int{0}
}

func (x *PutgStandard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PutgStandard) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PutgStandard) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PutgStandard) GetDnTitle() string {
	if x != nil {
		return x.DnTitle
	}
	return ""
}

func (x *PutgStandard) GetPnTitle() string {
	if x != nil {
		return x.PnTitle
	}
	return ""
}

func (x *PutgStandard) GetStandard() *standard_model.Standard {
	if x != nil {
		return x.Standard
	}
	return nil
}

func (x *PutgStandard) GetFlangeStandard() *flange_standard_model.FlangeStandard {
	if x != nil {
		return x.FlangeStandard
	}
	return nil
}

var File_pro_models_putg_standard_model_proto protoreflect.FileDescriptor

var file_pro_models_putg_standard_model_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x75, 0x74,
	0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x75, 0x74, 0x67, 0x5f, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x26, 0x70, 0x72, 0x6f,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x0c, 0x50, 0x75, 0x74, 0x67, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6e, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6e, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52,
	0x08, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x4d, 0x0a, 0x0e, 0x66, 0x6c, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x6c, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x0e, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x32, 0x37, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x75, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x70, 0x75, 0x74, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_models_putg_standard_model_proto_rawDescOnce sync.Once
	file_pro_models_putg_standard_model_proto_rawDescData = file_pro_models_putg_standard_model_proto_rawDesc
)

func file_pro_models_putg_standard_model_proto_rawDescGZIP() []byte {
	file_pro_models_putg_standard_model_proto_rawDescOnce.Do(func() {
		file_pro_models_putg_standard_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_models_putg_standard_model_proto_rawDescData)
	})
	return file_pro_models_putg_standard_model_proto_rawDescData
}

var file_pro_models_putg_standard_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pro_models_putg_standard_model_proto_goTypes = []interface{}{
	(*PutgStandard)(nil),                         // 0: putg_standard_model.PutgStandard
	(*standard_model.Standard)(nil),              // 1: standard_model.Standard
	(*flange_standard_model.FlangeStandard)(nil), // 2: flange_standard_model.FlangeStandard
}
var file_pro_models_putg_standard_model_proto_depIdxs = []int32{
	1, // 0: putg_standard_model.PutgStandard.standard:type_name -> standard_model.Standard
	2, // 1: putg_standard_model.PutgStandard.flangeStandard:type_name -> flange_standard_model.FlangeStandard
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pro_models_putg_standard_model_proto_init() }
func file_pro_models_putg_standard_model_proto_init() {
	if File_pro_models_putg_standard_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_models_putg_standard_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutgStandard); i {
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
			RawDescriptor: file_pro_models_putg_standard_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pro_models_putg_standard_model_proto_goTypes,
		DependencyIndexes: file_pro_models_putg_standard_model_proto_depIdxs,
		MessageInfos:      file_pro_models_putg_standard_model_proto_msgTypes,
	}.Build()
	File_pro_models_putg_standard_model_proto = out.File
	file_pro_models_putg_standard_model_proto_rawDesc = nil
	file_pro_models_putg_standard_model_proto_goTypes = nil
	file_pro_models_putg_standard_model_proto_depIdxs = nil
}