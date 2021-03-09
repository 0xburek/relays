// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: x/relay/types/query_services.proto

package types

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EmptyParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyParams) Reset() {
	*x = EmptyParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_relay_types_query_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyParams) ProtoMessage() {}

func (x *EmptyParams) ProtoReflect() protoreflect.Message {
	mi := &file_x_relay_types_query_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyParams.ProtoReflect.Descriptor instead.
func (*EmptyParams) Descriptor() ([]byte, []int) {
	return file_x_relay_types_query_services_proto_rawDescGZIP(), []int{0}
}

var File_x_relay_types_query_services_proto protoreflect.FileDescriptor

var file_x_relay_types_query_services_proto_rawDesc = []byte{
	0x0a, 0x22, 0x78, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x1a, 0x1b, 0x78, 0x2f, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0xba, 0x06, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x47, 0x0a, 0x0a, 0x49, 0x73, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12,
	0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x49, 0x73, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x1a, 0x19, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x49, 0x73,
	0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x12, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x6f, 0x72, 0x67, 0x4c, 0x43, 0x41, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x6f, 0x72, 0x67, 0x4c, 0x43, 0x41, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x42, 0x65, 0x73, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x12, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x47, 0x65, 0x74, 0x42, 0x65, 0x73, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x22, 0x00,
	0x12, 0x65, 0x0a, 0x14, 0x48, 0x65, 0x61, 0x76, 0x69, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x65, 0x61, 0x76,
	0x69, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x1a, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x48, 0x65, 0x61, 0x76, 0x69, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x6e, 0x63,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x1a, 0x49, 0x73, 0x4d, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x6e, 0x63,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x12, 0x2c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x49, 0x73, 0x4d, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x6e, 0x63, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x1a, 0x29, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x49, 0x73, 0x4d, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x1c, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_x_relay_types_query_services_proto_rawDescOnce sync.Once
	file_x_relay_types_query_services_proto_rawDescData = file_x_relay_types_query_services_proto_rawDesc
)

func file_x_relay_types_query_services_proto_rawDescGZIP() []byte {
	file_x_relay_types_query_services_proto_rawDescOnce.Do(func() {
		file_x_relay_types_query_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_x_relay_types_query_services_proto_rawDescData)
	})
	return file_x_relay_types_query_services_proto_rawDescData
}

var file_x_relay_types_query_services_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_x_relay_types_query_services_proto_goTypes = []interface{}{
	(*EmptyParams)(nil),                           // 0: relay.EmptyParams
	(*QueryParamsIsAncestor)(nil),                 // 1: relay.QueryParamsIsAncestor
	(*QueryParamsFindAncestor)(nil),               // 2: relay.QueryParamsFindAncestor
	(*QueryParamsHeaviestFromAncestor)(nil),       // 3: relay.QueryParamsHeaviestFromAncestor
	(*QueryParamsIsMostRecentCommonAncestor)(nil), // 4: relay.QueryParamsIsMostRecentCommonAncestor
	(*QueryParamsGetRequest)(nil),                 // 5: relay.QueryParamsGetRequest
	(*QueryParamsCheckRequests)(nil),              // 6: relay.QueryParamsCheckRequests
	(*QueryParamsCheckProof)(nil),                 // 7: relay.QueryParamsCheckProof
	(*QueryResIsAncestor)(nil),                    // 8: relay.QueryResIsAncestor
	(*QueryResGetRelayGenesis)(nil),               // 9: relay.QueryResGetRelayGenesis
	(*QueryResGetLastReorgLCA)(nil),               // 10: relay.QueryResGetLastReorgLCA
	(*QueryResGetBestDigest)(nil),                 // 11: relay.QueryResGetBestDigest
	(*QueryResFindAncestor)(nil),                  // 12: relay.QueryResFindAncestor
	(*QueryResHeaviestFromAncestor)(nil),          // 13: relay.QueryResHeaviestFromAncestor
	(*QueryResIsMostRecentCommonAncestor)(nil),    // 14: relay.QueryResIsMostRecentCommonAncestor
	(*QueryResGetRequest)(nil),                    // 15: relay.QueryResGetRequest
	(*QueryResCheckRequests)(nil),                 // 16: relay.QueryResCheckRequests
	(*QueryResCheckProof)(nil),                    // 17: relay.QueryResCheckProof
}
var file_x_relay_types_query_services_proto_depIdxs = []int32{
	1,  // 0: relay.Query.IsAncestor:input_type -> relay.QueryParamsIsAncestor
	0,  // 1: relay.Query.GetRelayGenesis:input_type -> relay.EmptyParams
	0,  // 2: relay.Query.GetLastReorgLCA:input_type -> relay.EmptyParams
	0,  // 3: relay.Query.GetBestDigest:input_type -> relay.EmptyParams
	2,  // 4: relay.Query.FindAncestor:input_type -> relay.QueryParamsFindAncestor
	3,  // 5: relay.Query.HeaviestFromAncestor:input_type -> relay.QueryParamsHeaviestFromAncestor
	4,  // 6: relay.Query.IsMostRecentCommonAncestor:input_type -> relay.QueryParamsIsMostRecentCommonAncestor
	5,  // 7: relay.Query.GetRequest:input_type -> relay.QueryParamsGetRequest
	6,  // 8: relay.Query.CheckRequests:input_type -> relay.QueryParamsCheckRequests
	7,  // 9: relay.Query.CheckProof:input_type -> relay.QueryParamsCheckProof
	8,  // 10: relay.Query.IsAncestor:output_type -> relay.QueryResIsAncestor
	9,  // 11: relay.Query.GetRelayGenesis:output_type -> relay.QueryResGetRelayGenesis
	10, // 12: relay.Query.GetLastReorgLCA:output_type -> relay.QueryResGetLastReorgLCA
	11, // 13: relay.Query.GetBestDigest:output_type -> relay.QueryResGetBestDigest
	12, // 14: relay.Query.FindAncestor:output_type -> relay.QueryResFindAncestor
	13, // 15: relay.Query.HeaviestFromAncestor:output_type -> relay.QueryResHeaviestFromAncestor
	14, // 16: relay.Query.IsMostRecentCommonAncestor:output_type -> relay.QueryResIsMostRecentCommonAncestor
	15, // 17: relay.Query.GetRequest:output_type -> relay.QueryResGetRequest
	16, // 18: relay.Query.CheckRequests:output_type -> relay.QueryResCheckRequests
	17, // 19: relay.Query.CheckProof:output_type -> relay.QueryResCheckProof
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_x_relay_types_query_services_proto_init() }
func file_x_relay_types_query_services_proto_init() {
	if File_x_relay_types_query_services_proto != nil {
		return
	}
	file_x_relay_types_querier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_x_relay_types_query_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyParams); i {
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
			RawDescriptor: file_x_relay_types_query_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_x_relay_types_query_services_proto_goTypes,
		DependencyIndexes: file_x_relay_types_query_services_proto_depIdxs,
		MessageInfos:      file_x_relay_types_query_services_proto_msgTypes,
	}.Build()
	File_x_relay_types_query_services_proto = out.File
	file_x_relay_types_query_services_proto_rawDesc = nil
	file_x_relay_types_query_services_proto_goTypes = nil
	file_x_relay_types_query_services_proto_depIdxs = nil
}