// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: cart/service/v1/cart_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type CartServiceErrorReason int32

const (
	CartServiceErrorReason_UNKNOWN_ERROR CartServiceErrorReason = 0
)

// Enum value maps for CartServiceErrorReason.
var (
	CartServiceErrorReason_name = map[int32]string{
		0: "UNKNOWN_ERROR",
	}
	CartServiceErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR": 0,
	}
)

func (x CartServiceErrorReason) Enum() *CartServiceErrorReason {
	p := new(CartServiceErrorReason)
	*p = x
	return p
}

func (x CartServiceErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CartServiceErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_cart_service_v1_cart_error_proto_enumTypes[0].Descriptor()
}

func (CartServiceErrorReason) Type() protoreflect.EnumType {
	return &file_cart_service_v1_cart_error_proto_enumTypes[0]
}

func (x CartServiceErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CartServiceErrorReason.Descriptor instead.
func (CartServiceErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_cart_service_v1_cart_error_proto_rawDescGZIP(), []int{0}
}

var File_cart_service_v1_cart_error_proto protoreflect.FileDescriptor

var file_cart_service_v1_cart_error_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x31, 0x0a, 0x16, 0x43, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x00, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x39, 0x50, 0x01, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_service_v1_cart_error_proto_rawDescOnce sync.Once
	file_cart_service_v1_cart_error_proto_rawDescData = file_cart_service_v1_cart_error_proto_rawDesc
)

func file_cart_service_v1_cart_error_proto_rawDescGZIP() []byte {
	file_cart_service_v1_cart_error_proto_rawDescOnce.Do(func() {
		file_cart_service_v1_cart_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_service_v1_cart_error_proto_rawDescData)
	})
	return file_cart_service_v1_cart_error_proto_rawDescData
}

var file_cart_service_v1_cart_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cart_service_v1_cart_error_proto_goTypes = []interface{}{
	(CartServiceErrorReason)(0), // 0: cart.service.v1.CartServiceErrorReason
}
var file_cart_service_v1_cart_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cart_service_v1_cart_error_proto_init() }
func file_cart_service_v1_cart_error_proto_init() {
	if File_cart_service_v1_cart_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cart_service_v1_cart_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cart_service_v1_cart_error_proto_goTypes,
		DependencyIndexes: file_cart_service_v1_cart_error_proto_depIdxs,
		EnumInfos:         file_cart_service_v1_cart_error_proto_enumTypes,
	}.Build()
	File_cart_service_v1_cart_error_proto = out.File
	file_cart_service_v1_cart_error_proto_rawDesc = nil
	file_cart_service_v1_cart_error_proto_goTypes = nil
	file_cart_service_v1_cart_error_proto_depIdxs = nil
}