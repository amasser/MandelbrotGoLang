// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: mandelbrot.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CalculateRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MagnificationFactor float64 `protobuf:"fixed64,1,opt,name=MagnificationFactor,proto3" json:"MagnificationFactor,omitempty"`
	MaxIterations       float64 `protobuf:"fixed64,2,opt,name=MaxIterations,proto3" json:"MaxIterations,omitempty"`
	PanX                float64 `protobuf:"fixed64,3,opt,name=PanX,proto3" json:"PanX,omitempty"`
	PanY                float64 `protobuf:"fixed64,4,opt,name=PanY,proto3" json:"PanY,omitempty"`
	Index               int32   `protobuf:"varint,5,opt,name=Index,proto3" json:"Index,omitempty"`
	XStart              int32   `protobuf:"varint,6,opt,name=XStart,proto3" json:"XStart,omitempty"`
	XEnd                int32   `protobuf:"varint,7,opt,name=XEnd,proto3" json:"XEnd,omitempty"`
	YStart              int32   `protobuf:"varint,8,opt,name=YStart,proto3" json:"YStart,omitempty"`
	YEnd                int32   `protobuf:"varint,9,opt,name=YEnd,proto3" json:"YEnd,omitempty"`
	Width               int32   `protobuf:"varint,10,opt,name=Width,proto3" json:"Width,omitempty"`
	Height              int32   `protobuf:"varint,11,opt,name=Height,proto3" json:"Height,omitempty"`
}

func (x *CalculateRegionRequest) Reset() {
	*x = CalculateRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mandelbrot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRegionRequest) ProtoMessage() {}

func (x *CalculateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mandelbrot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRegionRequest.ProtoReflect.Descriptor instead.
func (*CalculateRegionRequest) Descriptor() ([]byte, []int) {
	return file_mandelbrot_proto_rawDescGZIP(), []int{0}
}

func (x *CalculateRegionRequest) GetMagnificationFactor() float64 {
	if x != nil {
		return x.MagnificationFactor
	}
	return 0
}

func (x *CalculateRegionRequest) GetMaxIterations() float64 {
	if x != nil {
		return x.MaxIterations
	}
	return 0
}

func (x *CalculateRegionRequest) GetPanX() float64 {
	if x != nil {
		return x.PanX
	}
	return 0
}

func (x *CalculateRegionRequest) GetPanY() float64 {
	if x != nil {
		return x.PanY
	}
	return 0
}

func (x *CalculateRegionRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CalculateRegionRequest) GetXStart() int32 {
	if x != nil {
		return x.XStart
	}
	return 0
}

func (x *CalculateRegionRequest) GetXEnd() int32 {
	if x != nil {
		return x.XEnd
	}
	return 0
}

func (x *CalculateRegionRequest) GetYStart() int32 {
	if x != nil {
		return x.YStart
	}
	return 0
}

func (x *CalculateRegionRequest) GetYEnd() int32 {
	if x != nil {
		return x.YEnd
	}
	return 0
}

func (x *CalculateRegionRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *CalculateRegionRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type CalculateRegionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RGBPixels           []byte  `protobuf:"bytes,1,opt,name=RGBPixels,proto3" json:"RGBPixels,omitempty"`
	ThreadsProcessTimes []int64 `protobuf:"varint,2,rep,packed,name=ThreadsProcessTimes,proto3" json:"ThreadsProcessTimes,omitempty"`
}

func (x *CalculateRegionResponse) Reset() {
	*x = CalculateRegionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mandelbrot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRegionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRegionResponse) ProtoMessage() {}

func (x *CalculateRegionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mandelbrot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRegionResponse.ProtoReflect.Descriptor instead.
func (*CalculateRegionResponse) Descriptor() ([]byte, []int) {
	return file_mandelbrot_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateRegionResponse) GetRGBPixels() []byte {
	if x != nil {
		return x.RGBPixels
	}
	return nil
}

func (x *CalculateRegionResponse) GetThreadsProcessTimes() []int64 {
	if x != nil {
		return x.ThreadsProcessTimes
	}
	return nil
}

var File_mandelbrot_proto protoreflect.FileDescriptor

var file_mandelbrot_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x62, 0x72, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x16, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x4d, 0x61, 0x67, 0x6e, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x13, 0x4d, 0x61, 0x67, 0x6e, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x4d, 0x61, 0x78, 0x49, 0x74, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x4d,
	0x61, 0x78, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x61, 0x6e, 0x58, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x50, 0x61, 0x6e, 0x58,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x6e, 0x59, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x50, 0x61, 0x6e, 0x59, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x58, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x58, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x58, 0x45, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x58, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x59, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x59, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x59, 0x45, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x59, 0x45,
	0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x6d, 0x0a, 0x17, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52,
	0x47, 0x42, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x52, 0x47, 0x42, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x13, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x10, 0x01, 0x52, 0x13, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x32,
	0x69, 0x0a, 0x13, 0x4d, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x62, 0x72, 0x6f, 0x74, 0x53, 0x6c, 0x61,
	0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mandelbrot_proto_rawDescOnce sync.Once
	file_mandelbrot_proto_rawDescData = file_mandelbrot_proto_rawDesc
)

func file_mandelbrot_proto_rawDescGZIP() []byte {
	file_mandelbrot_proto_rawDescOnce.Do(func() {
		file_mandelbrot_proto_rawDescData = protoimpl.X.CompressGZIP(file_mandelbrot_proto_rawDescData)
	})
	return file_mandelbrot_proto_rawDescData
}

var file_mandelbrot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mandelbrot_proto_goTypes = []interface{}{
	(*CalculateRegionRequest)(nil),  // 0: proto.CalculateRegionRequest
	(*CalculateRegionResponse)(nil), // 1: proto.CalculateRegionResponse
}
var file_mandelbrot_proto_depIdxs = []int32{
	0, // 0: proto.MandelbrotSlaveNode.CalculateRegion:input_type -> proto.CalculateRegionRequest
	1, // 1: proto.MandelbrotSlaveNode.CalculateRegion:output_type -> proto.CalculateRegionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mandelbrot_proto_init() }
func file_mandelbrot_proto_init() {
	if File_mandelbrot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mandelbrot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRegionRequest); i {
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
		file_mandelbrot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRegionResponse); i {
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
			RawDescriptor: file_mandelbrot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mandelbrot_proto_goTypes,
		DependencyIndexes: file_mandelbrot_proto_depIdxs,
		MessageInfos:      file_mandelbrot_proto_msgTypes,
	}.Build()
	File_mandelbrot_proto = out.File
	file_mandelbrot_proto_rawDesc = nil
	file_mandelbrot_proto_goTypes = nil
	file_mandelbrot_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MandelbrotSlaveNodeClient is the client API for MandelbrotSlaveNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MandelbrotSlaveNodeClient interface {
	CalculateRegion(ctx context.Context, in *CalculateRegionRequest, opts ...grpc.CallOption) (*CalculateRegionResponse, error)
}

type mandelbrotSlaveNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewMandelbrotSlaveNodeClient(cc grpc.ClientConnInterface) MandelbrotSlaveNodeClient {
	return &mandelbrotSlaveNodeClient{cc}
}

func (c *mandelbrotSlaveNodeClient) CalculateRegion(ctx context.Context, in *CalculateRegionRequest, opts ...grpc.CallOption) (*CalculateRegionResponse, error) {
	out := new(CalculateRegionResponse)
	err := c.cc.Invoke(ctx, "/proto.MandelbrotSlaveNode/CalculateRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MandelbrotSlaveNodeServer is the server API for MandelbrotSlaveNode service.
type MandelbrotSlaveNodeServer interface {
	CalculateRegion(context.Context, *CalculateRegionRequest) (*CalculateRegionResponse, error)
}

// UnimplementedMandelbrotSlaveNodeServer can be embedded to have forward compatible implementations.
type UnimplementedMandelbrotSlaveNodeServer struct {
}

func (*UnimplementedMandelbrotSlaveNodeServer) CalculateRegion(context.Context, *CalculateRegionRequest) (*CalculateRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateRegion not implemented")
}

func RegisterMandelbrotSlaveNodeServer(s *grpc.Server, srv MandelbrotSlaveNodeServer) {
	s.RegisterService(&_MandelbrotSlaveNode_serviceDesc, srv)
}

func _MandelbrotSlaveNode_CalculateRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MandelbrotSlaveNodeServer).CalculateRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MandelbrotSlaveNode/CalculateRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MandelbrotSlaveNodeServer).CalculateRegion(ctx, req.(*CalculateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MandelbrotSlaveNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MandelbrotSlaveNode",
	HandlerType: (*MandelbrotSlaveNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateRegion",
			Handler:    _MandelbrotSlaveNode_CalculateRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mandelbrot.proto",
}
