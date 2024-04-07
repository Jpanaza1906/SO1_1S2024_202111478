// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: cliente.proto

package confproto

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

type RequestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Album string `protobuf:"bytes,2,opt,name=album,proto3" json:"album,omitempty"`
	Year  string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Rank  string `protobuf:"bytes,4,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *RequestId) Reset() {
	*x = RequestId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cliente_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestId) ProtoMessage() {}

func (x *RequestId) ProtoReflect() protoreflect.Message {
	mi := &file_cliente_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestId.ProtoReflect.Descriptor instead.
func (*RequestId) Descriptor() ([]byte, []int) {
	return file_cliente_proto_rawDescGZIP(), []int{0}
}

func (x *RequestId) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestId) GetAlbum() string {
	if x != nil {
		return x.Album
	}
	return ""
}

func (x *RequestId) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *RequestId) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

type ReplyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ReplyInfo) Reset() {
	*x = ReplyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cliente_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyInfo) ProtoMessage() {}

func (x *ReplyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cliente_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyInfo.ProtoReflect.Descriptor instead.
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return file_cliente_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyInfo) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_cliente_proto protoreflect.FileDescriptor

var file_cliente_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x63, 0x6f, 0x6e, 0x66, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x1f, 0x0a, 0x09, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x45, 0x0a, 0x07, 0x67, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cliente_proto_rawDescOnce sync.Once
	file_cliente_proto_rawDescData = file_cliente_proto_rawDesc
)

func file_cliente_proto_rawDescGZIP() []byte {
	file_cliente_proto_rawDescOnce.Do(func() {
		file_cliente_proto_rawDescData = protoimpl.X.CompressGZIP(file_cliente_proto_rawDescData)
	})
	return file_cliente_proto_rawDescData
}

var file_cliente_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cliente_proto_goTypes = []interface{}{
	(*RequestId)(nil), // 0: confproto.requestId
	(*ReplyInfo)(nil), // 1: confproto.replyInfo
}
var file_cliente_proto_depIdxs = []int32{
	0, // 0: confproto.getInfo.returnInfo:input_type -> confproto.requestId
	1, // 1: confproto.getInfo.returnInfo:output_type -> confproto.replyInfo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cliente_proto_init() }
func file_cliente_proto_init() {
	if File_cliente_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cliente_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestId); i {
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
		file_cliente_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyInfo); i {
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
			RawDescriptor: file_cliente_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cliente_proto_goTypes,
		DependencyIndexes: file_cliente_proto_depIdxs,
		MessageInfos:      file_cliente_proto_msgTypes,
	}.Build()
	File_cliente_proto = out.File
	file_cliente_proto_rawDesc = nil
	file_cliente_proto_goTypes = nil
	file_cliente_proto_depIdxs = nil
}
