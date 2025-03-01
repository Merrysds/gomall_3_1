// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: order_page.proto

package order

import (
	_ "github.com/cloudwego/gomall07/app/frontend/hertz_gen/api"
	common "github.com/cloudwego/gomall07/app/frontend/hertz_gen/frontend/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_order_page_proto protoreflect.FileDescriptor

var file_order_page_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x1a, 0x0d, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x69, 0x64, 0x6c, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x57, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0a, 0xca, 0xc1, 0x18, 0x06, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x65, 0x67, 0x6f, 0x2f, 0x67, 0x6f,
	0x6d, 0x61, 0x6c, 0x6c, 0x30, 0x37, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_order_page_proto_goTypes = []interface{}{
	(*common.Empty)(nil), // 0: frontend.common.Empty
}
var file_order_page_proto_depIdxs = []int32{
	0, // 0: frontend.order.OrderService.OrderList:input_type -> frontend.common.Empty
	0, // 1: frontend.order.OrderService.OrderList:output_type -> frontend.common.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_page_proto_init() }
func file_order_page_proto_init() {
	if File_order_page_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_page_proto_goTypes,
		DependencyIndexes: file_order_page_proto_depIdxs,
	}.Build()
	File_order_page_proto = out.File
	file_order_page_proto_rawDesc = nil
	file_order_page_proto_goTypes = nil
	file_order_page_proto_depIdxs = nil
}
