// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: internal/pb/order/order.proto

package order

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

type OrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	SpikeId  string `protobuf:"bytes,2,opt,name=SpikeId,proto3" json:"SpikeId,omitempty"`
	Quantity int64  `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *OrderInfo) Reset() {
	*x = OrderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfo) ProtoMessage() {}

func (x *OrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfo.ProtoReflect.Descriptor instead.
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return file_internal_pb_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderInfo) GetSpikeId() string {
	if x != nil {
		return x.SpikeId
	}
	return ""
}

func (x *OrderInfo) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_internal_pb_order_order_proto protoreflect.FileDescriptor

var file_internal_pb_order_order_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x70, 0x69, 0x6b, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x70,
	0x69, 0x6b, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_order_order_proto_rawDescOnce sync.Once
	file_internal_pb_order_order_proto_rawDescData = file_internal_pb_order_order_proto_rawDesc
)

func file_internal_pb_order_order_proto_rawDescGZIP() []byte {
	file_internal_pb_order_order_proto_rawDescOnce.Do(func() {
		file_internal_pb_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_order_order_proto_rawDescData)
	})
	return file_internal_pb_order_order_proto_rawDescData
}

var file_internal_pb_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_pb_order_order_proto_goTypes = []interface{}{
	(*OrderInfo)(nil), // 0: order.OrderInfo
}
var file_internal_pb_order_order_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pb_order_order_proto_init() }
func file_internal_pb_order_order_proto_init() {
	if File_internal_pb_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfo); i {
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
			RawDescriptor: file_internal_pb_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pb_order_order_proto_goTypes,
		DependencyIndexes: file_internal_pb_order_order_proto_depIdxs,
		MessageInfos:      file_internal_pb_order_order_proto_msgTypes,
	}.Build()
	File_internal_pb_order_order_proto = out.File
	file_internal_pb_order_order_proto_rawDesc = nil
	file_internal_pb_order_order_proto_goTypes = nil
	file_internal_pb_order_order_proto_depIdxs = nil
}