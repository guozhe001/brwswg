// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: protobufIntro.proto

package __

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

// Enumerations provide the ordering of numbers for a given set of elements. The default order of values is from zero to n.
type Schedule_Days int32

const (
	Schedule_SUNDAY    Schedule_Days = 0
	Schedule_MONDAY    Schedule_Days = 1
	Schedule_TUESDAY   Schedule_Days = 2
	Schedule_WEDNESDAY Schedule_Days = 3
	Schedule_THURSDAY  Schedule_Days = 4
	Schedule_FRIDAY    Schedule_Days = 5
	Schedule_SATURDAY  Schedule_Days = 6
)

// Enum value maps for Schedule_Days.
var (
	Schedule_Days_name = map[int32]string{
		0: "SUNDAY",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
	}
	Schedule_Days_value = map[string]int32{
		"SUNDAY":    0,
		"MONDAY":    1,
		"TUESDAY":   2,
		"WEDNESDAY": 3,
		"THURSDAY":  4,
		"FRIDAY":    5,
		"SATURDAY":  6,
	}
)

func (x Schedule_Days) Enum() *Schedule_Days {
	p := new(Schedule_Days)
	*p = x
	return p
}

func (x Schedule_Days) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schedule_Days) Descriptor() protoreflect.EnumDescriptor {
	return file_protobufIntro_proto_enumTypes[0].Descriptor()
}

func (Schedule_Days) Type() protoreflect.EnumType {
	return &file_protobufIntro_proto_enumTypes[0]
}

func (x Schedule_Days) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schedule_Days.Descriptor instead.
func (Schedule_Days) EnumDescriptor() ([]byte, []int) {
	return file_protobufIntro_proto_rawDescGZIP(), []int{0, 0}
}

// Protobuf3 allows an option called allow aliases to assign two different members the same value.
// 如果需要定义全部不相同的tag，则需要把option allow_alias = true;删除，否则会报编译错误
type Schedule_EnumAllowingAlias int32

const (
	Schedule_UNKNOWN Schedule_EnumAllowingAlias = 0
	Schedule_STARTED Schedule_EnumAllowingAlias = 1
	Schedule_RUNNING Schedule_EnumAllowingAlias = 1
)

// Enum value maps for Schedule_EnumAllowingAlias.
var (
	Schedule_EnumAllowingAlias_name = map[int32]string{
		0: "UNKNOWN",
		1: "STARTED",
		// Duplicate value: 1: "RUNNING",
	}
	Schedule_EnumAllowingAlias_value = map[string]int32{
		"UNKNOWN": 0,
		"STARTED": 1,
		"RUNNING": 1,
	}
)

func (x Schedule_EnumAllowingAlias) Enum() *Schedule_EnumAllowingAlias {
	p := new(Schedule_EnumAllowingAlias)
	*p = x
	return p
}

func (x Schedule_EnumAllowingAlias) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schedule_EnumAllowingAlias) Descriptor() protoreflect.EnumDescriptor {
	return file_protobufIntro_proto_enumTypes[1].Descriptor()
}

func (Schedule_EnumAllowingAlias) Type() protoreflect.EnumType {
	return &file_protobufIntro_proto_enumTypes[1]
}

func (x Schedule_EnumAllowingAlias) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schedule_EnumAllowingAlias.Descriptor instead.
func (Schedule_EnumAllowingAlias) EnumDescriptor() ([]byte, []int) {
	return file_protobufIntro_proto_rawDescGZIP(), []int{0, 1}
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufIntro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_protobufIntro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_protobufIntro_proto_rawDescGZIP(), []int{0}
}

// 嵌套消息
type Site struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Latency int64    `protobuf:"varint,2,opt,name=latency,proto3" json:"latency,omitempty"`
	Proxies []*Proxy `protobuf:"bytes,3,rep,name=proxies,proto3" json:"proxies,omitempty"`
}

func (x *Site) Reset() {
	*x = Site{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufIntro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Site) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Site) ProtoMessage() {}

func (x *Site) ProtoReflect() protoreflect.Message {
	mi := &file_protobufIntro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Site.ProtoReflect.Descriptor instead.
func (*Site) Descriptor() ([]byte, []int) {
	return file_protobufIntro_proto_rawDescGZIP(), []int{1}
}

func (x *Site) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Site) GetLatency() int64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

func (x *Site) GetProxies() []*Proxy {
	if x != nil {
		return x.Proxies
	}
	return nil
}

type Proxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Utl     string `protobuf:"bytes,1,opt,name=utl,proto3" json:"utl,omitempty"`
	Latency int64  `protobuf:"varint,2,opt,name=latency,proto3" json:"latency,omitempty"`
}

func (x *Proxy) Reset() {
	*x = Proxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufIntro_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proxy) ProtoMessage() {}

func (x *Proxy) ProtoReflect() protoreflect.Message {
	mi := &file_protobufIntro_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proxy.ProtoReflect.Descriptor instead.
func (*Proxy) Descriptor() ([]byte, []int) {
	return file_protobufIntro_proto_rawDescGZIP(), []int{2}
}

func (x *Proxy) GetUtl() string {
	if x != nil {
		return x.Utl
	}
	return ""
}

func (x *Proxy) GetLatency() int64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

var File_protobufIntro_proto protoreflect.FileDescriptor

var file_protobufIntro_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x22, 0xae, 0x01, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x62,
	0x0a, 0x04, 0x44, 0x61, 0x79, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x54, 0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57,
	0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x48,
	0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x49, 0x44,
	0x41, 0x59, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59,
	0x10, 0x06, 0x22, 0x3e, 0x0a, 0x11, 0x45, 0x6e, 0x75, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x22, 0x5f, 0x0a, 0x04, 0x53, 0x69, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78,
	0x69, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x74, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x74, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobufIntro_proto_rawDescOnce sync.Once
	file_protobufIntro_proto_rawDescData = file_protobufIntro_proto_rawDesc
)

func file_protobufIntro_proto_rawDescGZIP() []byte {
	file_protobufIntro_proto_rawDescOnce.Do(func() {
		file_protobufIntro_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufIntro_proto_rawDescData)
	})
	return file_protobufIntro_proto_rawDescData
}

var file_protobufIntro_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protobufIntro_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobufIntro_proto_goTypes = []interface{}{
	(Schedule_Days)(0),              // 0: protofiles.Schedule.Days
	(Schedule_EnumAllowingAlias)(0), // 1: protofiles.Schedule.EnumAllowingAlias
	(*Schedule)(nil),                // 2: protofiles.Schedule
	(*Site)(nil),                    // 3: protofiles.Site
	(*Proxy)(nil),                   // 4: protofiles.Proxy
}
var file_protobufIntro_proto_depIdxs = []int32{
	4, // 0: protofiles.Site.proxies:type_name -> protofiles.Proxy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobufIntro_proto_init() }
func file_protobufIntro_proto_init() {
	if File_protobufIntro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobufIntro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_protobufIntro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Site); i {
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
		file_protobufIntro_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proxy); i {
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
			RawDescriptor: file_protobufIntro_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobufIntro_proto_goTypes,
		DependencyIndexes: file_protobufIntro_proto_depIdxs,
		EnumInfos:         file_protobufIntro_proto_enumTypes,
		MessageInfos:      file_protobufIntro_proto_msgTypes,
	}.Build()
	File_protobufIntro_proto = out.File
	file_protobufIntro_proto_rawDesc = nil
	file_protobufIntro_proto_goTypes = nil
	file_protobufIntro_proto_depIdxs = nil
}