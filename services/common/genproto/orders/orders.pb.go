// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: orders.proto

package orders

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int32                  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	CustomerId    int32                  `protobuf:"varint,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	ProductId     int32                  `protobuf:"varint,3,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity      int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_orders_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Order) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Order) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int32                  `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	ProductId     int32                  `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_orders_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreateOrderRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateOrderRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_orders_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    int32                  `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_orders_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"` // this is a slice of orders of type Order
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_orders_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_orders_proto protoreflect.FileDescriptor

const file_orders_proto_rawDesc = "" +
	"\n" +
	"\forders.proto\"{\n" +
	"\x05Order\x12\x18\n" +
	"\aorderId\x18\x01 \x01(\x05R\aorderId\x12\x1e\n" +
	"\n" +
	"customerId\x18\x02 \x01(\x05R\n" +
	"customerId\x12\x1c\n" +
	"\tproductId\x18\x03 \x01(\x05R\tproductId\x12\x1a\n" +
	"\bquantity\x18\x04 \x01(\x05R\bquantity\"n\n" +
	"\x12CreateOrderRequest\x12\x1e\n" +
	"\n" +
	"customerId\x18\x01 \x01(\x05R\n" +
	"customerId\x12\x1c\n" +
	"\tproductId\x18\x02 \x01(\x05R\tproductId\x12\x1a\n" +
	"\bquantity\x18\x03 \x01(\x05R\bquantity\"-\n" +
	"\x13CreateOrderResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"1\n" +
	"\x0fGetOrderRequest\x12\x1e\n" +
	"\n" +
	"customerId\x18\x01 \x01(\x05R\n" +
	"customerId\"2\n" +
	"\x10GetOrderResponse\x12\x1e\n" +
	"\x06orders\x18\x01 \x03(\v2\x06.OrderR\x06orders2J\n" +
	"\fOrderService\x12:\n" +
	"\vCreateOrder\x12\x13.CreateOrderRequest\x1a\x14.CreateOrderResponse\"\x00B$Z\"github.com/saket6198/common/ordersb\x06proto3"

var (
	file_orders_proto_rawDescOnce sync.Once
	file_orders_proto_rawDescData []byte
)

func file_orders_proto_rawDescGZIP() []byte {
	file_orders_proto_rawDescOnce.Do(func() {
		file_orders_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orders_proto_rawDesc), len(file_orders_proto_rawDesc)))
	})
	return file_orders_proto_rawDescData
}

var file_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orders_proto_goTypes = []any{
	(*Order)(nil),               // 0: Order
	(*CreateOrderRequest)(nil),  // 1: CreateOrderRequest
	(*CreateOrderResponse)(nil), // 2: CreateOrderResponse
	(*GetOrderRequest)(nil),     // 3: GetOrderRequest
	(*GetOrderResponse)(nil),    // 4: GetOrderResponse
}
var file_orders_proto_depIdxs = []int32{
	0, // 0: GetOrderResponse.orders:type_name -> Order
	1, // 1: OrderService.CreateOrder:input_type -> CreateOrderRequest
	2, // 2: OrderService.CreateOrder:output_type -> CreateOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orders_proto_init() }
func file_orders_proto_init() {
	if File_orders_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orders_proto_rawDesc), len(file_orders_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_proto_goTypes,
		DependencyIndexes: file_orders_proto_depIdxs,
		MessageInfos:      file_orders_proto_msgTypes,
	}.Build()
	File_orders_proto = out.File
	file_orders_proto_goTypes = nil
	file_orders_proto_depIdxs = nil
}
