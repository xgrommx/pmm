// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: managementpb/dbaas/templates.proto

package dbaasv1beta1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Template CR name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Template CR kind.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_templates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_templates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_templates_proto_rawDescGZIP(), []int{0}
}

func (x *Template) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Template) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type ListTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes cluster name.
	KubernetesClusterName string `protobuf:"bytes,1,opt,name=kubernetes_cluster_name,json=kubernetesClusterName,proto3" json:"kubernetes_cluster_name,omitempty"`
	// DB cluster type.
	ClusterType DBClusterType `protobuf:"varint,2,opt,name=cluster_type,json=clusterType,proto3,enum=dbaas.v1beta1.DBClusterType" json:"cluster_type,omitempty"`
}

func (x *ListTemplatesRequest) Reset() {
	*x = ListTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_templates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplatesRequest) ProtoMessage() {}

func (x *ListTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_templates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplatesRequest.ProtoReflect.Descriptor instead.
func (*ListTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_templates_proto_rawDescGZIP(), []int{1}
}

func (x *ListTemplatesRequest) GetKubernetesClusterName() string {
	if x != nil {
		return x.KubernetesClusterName
	}
	return ""
}

func (x *ListTemplatesRequest) GetClusterType() DBClusterType {
	if x != nil {
		return x.ClusterType
	}
	return DBClusterType_DB_CLUSTER_TYPE_INVALID
}

type ListTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Templates []*Template `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
}

func (x *ListTemplatesResponse) Reset() {
	*x = ListTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_templates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplatesResponse) ProtoMessage() {}

func (x *ListTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_templates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplatesResponse.ProtoReflect.Descriptor instead.
func (*ListTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_templates_proto_rawDescGZIP(), []int{2}
}

func (x *ListTemplatesResponse) GetTemplates() []*Template {
	if x != nil {
		return x.Templates
	}
	return nil
}

var File_managementpb_dbaas_templates_proto protoreflect.FileDescriptor

var file_managementpb_dbaas_templates_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64,
	0x62, 0x61, 0x61, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2f, 0x64, 0x62,
	0x61, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x9f, 0x01,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x17, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52,
	0x15, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x64,
	0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x42, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x20, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x4e, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x62,
	0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x32,
	0x98, 0x01, 0x0a, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x8a, 0x01,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x23, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x44, 0x42, 0x61, 0x61, 0x53, 0x2f, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0xb4, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64, 0x62, 0x61, 0x61,
	0x73, 0x3b, 0x64, 0x62, 0x61, 0x61, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x44, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x44, 0x62, 0x61, 0x61, 0x73, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x44, 0x62, 0x61, 0x61, 0x73, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0e, 0x44, 0x62, 0x61, 0x61, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_dbaas_templates_proto_rawDescOnce sync.Once
	file_managementpb_dbaas_templates_proto_rawDescData = file_managementpb_dbaas_templates_proto_rawDesc
)

func file_managementpb_dbaas_templates_proto_rawDescGZIP() []byte {
	file_managementpb_dbaas_templates_proto_rawDescOnce.Do(func() {
		file_managementpb_dbaas_templates_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_dbaas_templates_proto_rawDescData)
	})
	return file_managementpb_dbaas_templates_proto_rawDescData
}

var (
	file_managementpb_dbaas_templates_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_managementpb_dbaas_templates_proto_goTypes  = []interface{}{
		(*Template)(nil),              // 0: dbaas.v1beta1.Template
		(*ListTemplatesRequest)(nil),  // 1: dbaas.v1beta1.ListTemplatesRequest
		(*ListTemplatesResponse)(nil), // 2: dbaas.v1beta1.ListTemplatesResponse
		(DBClusterType)(0),            // 3: dbaas.v1beta1.DBClusterType
	}
)

var file_managementpb_dbaas_templates_proto_depIdxs = []int32{
	3, // 0: dbaas.v1beta1.ListTemplatesRequest.cluster_type:type_name -> dbaas.v1beta1.DBClusterType
	0, // 1: dbaas.v1beta1.ListTemplatesResponse.templates:type_name -> dbaas.v1beta1.Template
	1, // 2: dbaas.v1beta1.Templates.ListTemplates:input_type -> dbaas.v1beta1.ListTemplatesRequest
	2, // 3: dbaas.v1beta1.Templates.ListTemplates:output_type -> dbaas.v1beta1.ListTemplatesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_managementpb_dbaas_templates_proto_init() }
func file_managementpb_dbaas_templates_proto_init() {
	if File_managementpb_dbaas_templates_proto != nil {
		return
	}
	file_managementpb_dbaas_dbaas_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_dbaas_templates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_managementpb_dbaas_templates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplatesRequest); i {
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
		file_managementpb_dbaas_templates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplatesResponse); i {
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
			RawDescriptor: file_managementpb_dbaas_templates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_dbaas_templates_proto_goTypes,
		DependencyIndexes: file_managementpb_dbaas_templates_proto_depIdxs,
		MessageInfos:      file_managementpb_dbaas_templates_proto_msgTypes,
	}.Build()
	File_managementpb_dbaas_templates_proto = out.File
	file_managementpb_dbaas_templates_proto_rawDesc = nil
	file_managementpb_dbaas_templates_proto_goTypes = nil
	file_managementpb_dbaas_templates_proto_depIdxs = nil
}
