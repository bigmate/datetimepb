// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: datetime.proto

package datetimepb

import (
	"reflect"
	"sync"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Datetime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	Offset  int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func Now() *Datetime {
	return New(time.Now())
}

func New(t time.Time) *Datetime {
	_, offset := t.Zone()

	return &Datetime{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
		Offset:  int32(offset),
	}
}

func (x *Datetime) AsTime() time.Time {
	loc := time.FixedZone("UNKNOWN", int(x.GetOffset()))

	return time.Unix(x.GetSeconds(), int64(x.GetNanos())).In(loc)
}

func (x *Datetime) Reset() {
	*x = Datetime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datetime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datetime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datetime) ProtoMessage() {}

func (x *Datetime) ProtoReflect() protoreflect.Message {
	mi := &file_datetime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datetime.ProtoReflect.Descriptor instead.
func (*Datetime) Descriptor() ([]byte, []int) {
	return file_datetime_proto_rawDescGZIP(), []int{0}
}

func (x *Datetime) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Datetime) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

func (x *Datetime) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_datetime_proto protoreflect.FileDescriptor

var file_datetime_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x1f,
	0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x67,
	0x6d, 0x61, 0x74, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datetime_proto_rawDescOnce sync.Once
	file_datetime_proto_rawDescData = file_datetime_proto_rawDesc
)

func file_datetime_proto_rawDescGZIP() []byte {
	file_datetime_proto_rawDescOnce.Do(func() {
		file_datetime_proto_rawDescData = protoimpl.X.CompressGZIP(file_datetime_proto_rawDescData)
	})
	return file_datetime_proto_rawDescData
}

var file_datetime_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_datetime_proto_goTypes = []interface{}{
	(*Datetime)(nil), // 0: datetime.Datetime
}
var file_datetime_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datetime_proto_init() }
func file_datetime_proto_init() {
	if File_datetime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datetime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datetime); i {
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
			RawDescriptor: file_datetime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datetime_proto_goTypes,
		DependencyIndexes: file_datetime_proto_depIdxs,
		MessageInfos:      file_datetime_proto_msgTypes,
	}.Build()
	File_datetime_proto = out.File
	file_datetime_proto_rawDesc = nil
	file_datetime_proto_goTypes = nil
	file_datetime_proto_depIdxs = nil
}
