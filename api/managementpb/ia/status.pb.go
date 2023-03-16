// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: managementpb/ia/status.proto

package iav1beta1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status represents Alert Rule's and Alert's combined status.
type Status int32

const (
	Status_STATUS_INVALID Status = 0
	// No alert.
	Status_CLEAR Status = 1
	// Pending, but not triggering alert.
	Status_PENDING Status = 2
	// Triggering (firing) alert.
	Status_TRIGGERING Status = 3
	// Silenced alert.
	Status_SILENCED Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_INVALID",
		1: "CLEAR",
		2: "PENDING",
		3: "TRIGGERING",
		4: "SILENCED",
	}
	Status_value = map[string]int32{
		"STATUS_INVALID": 0,
		"CLEAR":          1,
		"PENDING":        2,
		"TRIGGERING":     3,
		"SILENCED":       4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_ia_status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_managementpb_ia_status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_ia_status_proto_rawDescGZIP(), []int{0}
}

var File_managementpb_ia_status_proto protoreflect.FileDescriptor

var file_managementpb_ia_status_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x69,
	0x61, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2a, 0x52, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x45, 0x41,
	0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x49, 0x4c, 0x45, 0x4e, 0x43, 0x45, 0x44, 0x10, 0x04, 0x42, 0x9c,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72,
	0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x69, 0x61, 0x3b, 0x69, 0x61, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x49,
	0x61, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x0a, 0x49, 0x61, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x16, 0x49, 0x61, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0b, 0x49, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_ia_status_proto_rawDescOnce sync.Once
	file_managementpb_ia_status_proto_rawDescData = file_managementpb_ia_status_proto_rawDesc
)

func file_managementpb_ia_status_proto_rawDescGZIP() []byte {
	file_managementpb_ia_status_proto_rawDescOnce.Do(func() {
		file_managementpb_ia_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_ia_status_proto_rawDescData)
	})
	return file_managementpb_ia_status_proto_rawDescData
}

var (
	file_managementpb_ia_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_managementpb_ia_status_proto_goTypes   = []interface{}{
		(Status)(0), // 0: ia.v1beta1.Status
	}
)

var file_managementpb_ia_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managementpb_ia_status_proto_init() }
func file_managementpb_ia_status_proto_init() {
	if File_managementpb_ia_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_ia_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managementpb_ia_status_proto_goTypes,
		DependencyIndexes: file_managementpb_ia_status_proto_depIdxs,
		EnumInfos:         file_managementpb_ia_status_proto_enumTypes,
	}.Build()
	File_managementpb_ia_status_proto = out.File
	file_managementpb_ia_status_proto_rawDesc = nil
	file_managementpb_ia_status_proto_goTypes = nil
	file_managementpb_ia_status_proto_depIdxs = nil
}
