// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opencensus/proto/agent/trace/v1/trace_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-gogo/agent/common/v1"
	v12 "github.com/census-instrumentation/opencensus-proto/gen-gogo/resource/v1"
	v11 "github.com/census-instrumentation/opencensus-proto/gen-gogo/trace/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CurrentLibraryConfig struct {
	// This is required only in the first message on the stream or if the
	// previous sent CurrentLibraryConfig message has a different Node (e.g.
	// when the same RPC is used to configure multiple Applications).
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Current configuration.
	Config               *v11.TraceConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CurrentLibraryConfig) Reset()         { *m = CurrentLibraryConfig{} }
func (m *CurrentLibraryConfig) String() string { return proto.CompactTextString(m) }
func (*CurrentLibraryConfig) ProtoMessage()    {}
func (*CurrentLibraryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7027f99caf7ac6a5, []int{0}
}
func (m *CurrentLibraryConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrentLibraryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrentLibraryConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrentLibraryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentLibraryConfig.Merge(m, src)
}
func (m *CurrentLibraryConfig) XXX_Size() int {
	return m.Size()
}
func (m *CurrentLibraryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentLibraryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentLibraryConfig proto.InternalMessageInfo

func (m *CurrentLibraryConfig) GetNode() *v1.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *CurrentLibraryConfig) GetConfig() *v11.TraceConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type UpdatedLibraryConfig struct {
	// This field is ignored when the RPC is used to configure only one Application.
	// This is required only in the first message on the stream or if the
	// previous sent UpdatedLibraryConfig message has a different Node.
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Requested updated configuration.
	Config               *v11.TraceConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdatedLibraryConfig) Reset()         { *m = UpdatedLibraryConfig{} }
func (m *UpdatedLibraryConfig) String() string { return proto.CompactTextString(m) }
func (*UpdatedLibraryConfig) ProtoMessage()    {}
func (*UpdatedLibraryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7027f99caf7ac6a5, []int{1}
}
func (m *UpdatedLibraryConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatedLibraryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatedLibraryConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatedLibraryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatedLibraryConfig.Merge(m, src)
}
func (m *UpdatedLibraryConfig) XXX_Size() int {
	return m.Size()
}
func (m *UpdatedLibraryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatedLibraryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatedLibraryConfig proto.InternalMessageInfo

func (m *UpdatedLibraryConfig) GetNode() *v1.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *UpdatedLibraryConfig) GetConfig() *v11.TraceConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type ExportTraceServiceRequest struct {
	// This is required only in the first message on the stream or if the
	// previous sent ExportTraceServiceRequest message has a different Node (e.g.
	// when the same RPC is used to send Spans from multiple Applications).
	Node *v1.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// A list of Spans that belong to the last received Node.
	Spans []*v11.Span `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	// The resource for the spans in this message that do not have an explicit
	// resource set.
	// If unset, the most recently set resource in the RPC stream applies. It is
	// valid to never be set within a stream, e.g. when no resource info is known.
	Resource             *v12.Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ExportTraceServiceRequest) Reset()         { *m = ExportTraceServiceRequest{} }
func (m *ExportTraceServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportTraceServiceRequest) ProtoMessage()    {}
func (*ExportTraceServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7027f99caf7ac6a5, []int{2}
}
func (m *ExportTraceServiceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportTraceServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportTraceServiceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportTraceServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportTraceServiceRequest.Merge(m, src)
}
func (m *ExportTraceServiceRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExportTraceServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportTraceServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportTraceServiceRequest proto.InternalMessageInfo

func (m *ExportTraceServiceRequest) GetNode() *v1.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *ExportTraceServiceRequest) GetSpans() []*v11.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *ExportTraceServiceRequest) GetResource() *v12.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type ExportTraceServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportTraceServiceResponse) Reset()         { *m = ExportTraceServiceResponse{} }
func (m *ExportTraceServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportTraceServiceResponse) ProtoMessage()    {}
func (*ExportTraceServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7027f99caf7ac6a5, []int{3}
}
func (m *ExportTraceServiceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportTraceServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportTraceServiceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportTraceServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportTraceServiceResponse.Merge(m, src)
}
func (m *ExportTraceServiceResponse) XXX_Size() int {
	return m.Size()
}
func (m *ExportTraceServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportTraceServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportTraceServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CurrentLibraryConfig)(nil), "opencensus.proto.agent.trace.v1.CurrentLibraryConfig")
	proto.RegisterType((*UpdatedLibraryConfig)(nil), "opencensus.proto.agent.trace.v1.UpdatedLibraryConfig")
	proto.RegisterType((*ExportTraceServiceRequest)(nil), "opencensus.proto.agent.trace.v1.ExportTraceServiceRequest")
	proto.RegisterType((*ExportTraceServiceResponse)(nil), "opencensus.proto.agent.trace.v1.ExportTraceServiceResponse")
}

func init() {
	proto.RegisterFile("opencensus/proto/agent/trace/v1/trace_service.proto", fileDescriptor_7027f99caf7ac6a5)
}

var fileDescriptor_7027f99caf7ac6a5 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0xdd, 0x9e, 0xd5, 0x41, 0x7a, 0xbd, 0x18, 0x3c, 0xc4, 0x20, 0x99, 0x25, 0xa0, 0x0c, 0x68,
	0x3a, 0x66, 0x96, 0xbd, 0xac, 0x20, 0x38, 0x83, 0x17, 0x11, 0x5d, 0xb2, 0xea, 0xc1, 0xcb, 0x92,
	0x49, 0xca, 0x98, 0xc3, 0x74, 0xb7, 0xdd, 0x9d, 0xa0, 0xe0, 0xdd, 0xbb, 0x07, 0xf1, 0x4f, 0xfc,
	0x05, 0x8f, 0xe2, 0x17, 0xc8, 0x78, 0xf3, 0x2b, 0x24, 0x5d, 0x99, 0xdd, 0xd5, 0x99, 0x18, 0xd0,
	0xcb, 0xde, 0x8a, 0xf4, 0x7b, 0xaf, 0x5e, 0x75, 0xbd, 0x0e, 0xdd, 0x13, 0x12, 0x78, 0x06, 0x5c,
	0x57, 0x3a, 0x92, 0x4a, 0x18, 0x11, 0xa5, 0x05, 0x70, 0x13, 0x19, 0x95, 0x66, 0x10, 0xd5, 0x31,
	0x16, 0xc7, 0x1a, 0x54, 0x5d, 0x66, 0xc0, 0x2c, 0xc4, 0x19, 0x9d, 0x92, 0xf0, 0x0b, 0xb3, 0x24,
	0x66, 0xb1, 0xac, 0x8e, 0xbd, 0xb0, 0x43, 0x35, 0x13, 0x8b, 0x85, 0xe0, 0x8d, 0x2c, 0x56, 0xc8,
	0xf6, 0x6e, 0xad, 0xc1, 0x15, 0x68, 0x51, 0x29, 0x74, 0xb0, 0xaa, 0x5b, 0xf0, 0x8d, 0x35, 0xf0,
	0xef, 0x5e, 0x5b, 0xd8, 0xed, 0x1e, 0xd8, 0x71, 0x26, 0xf8, 0xcb, 0xb2, 0x40, 0x74, 0xf0, 0x81,
	0xd0, 0xab, 0xb3, 0x4a, 0x29, 0xe0, 0xe6, 0x51, 0x39, 0x57, 0xa9, 0x7a, 0x3b, 0xb3, 0xc7, 0xce,
	0x01, 0xbd, 0xc0, 0x45, 0x0e, 0x2e, 0xd9, 0x25, 0xe3, 0x9d, 0xc9, 0x4d, 0xd6, 0x31, 0x79, 0x3b,
	0x4e, 0x1d, 0xb3, 0xc7, 0x22, 0x87, 0xc4, 0x72, 0x9c, 0x7b, 0x74, 0x88, 0x4d, 0xdc, 0x41, 0x17,
	0x7b, 0x75, 0x63, 0xec, 0x69, 0x53, 0x60, 0xcf, 0xa4, 0x65, 0x59, 0x53, 0xcf, 0x64, 0x9e, 0x1a,
	0xc8, 0xcf, 0x8f, 0xa9, 0x6f, 0x84, 0x5e, 0x7b, 0xf0, 0x46, 0x0a, 0x65, 0xec, 0xe9, 0x11, 0x06,
	0x23, 0x81, 0xd7, 0x15, 0x68, 0xf3, 0x5f, 0xce, 0xf6, 0xe9, 0x45, 0x2d, 0x53, 0xae, 0xdd, 0xc1,
	0xee, 0xf6, 0x78, 0x67, 0x32, 0xfa, 0x8b, 0xb1, 0x23, 0x99, 0xf2, 0x04, 0xd1, 0xce, 0x94, 0x5e,
	0x5a, 0x25, 0xc4, 0xdd, 0xee, 0x6a, 0x7b, 0x92, 0xa1, 0x3a, 0x66, 0x49, 0x5b, 0x27, 0x27, 0xbc,
	0xe0, 0x3a, 0xf5, 0x36, 0xcd, 0xa4, 0xa5, 0xe0, 0x1a, 0x26, 0x1f, 0x07, 0xf4, 0xf2, 0xd9, 0x03,
	0xe7, 0x1d, 0x1d, 0xb6, 0x9b, 0xd8, 0x67, 0x3d, 0x4f, 0x81, 0x6d, 0x4a, 0x95, 0xd7, 0x4f, 0xdb,
	0xb4, 0xf7, 0x60, 0x6b, 0x4c, 0xee, 0x10, 0xe7, 0x3d, 0xa1, 0x43, 0x74, 0xeb, 0x1c, 0xf4, 0xea,
	0x74, 0xae, 0xca, 0xbb, 0xfb, 0x4f, 0x5c, 0xbc, 0x12, 0x74, 0x32, 0xfd, 0x4c, 0xbe, 0x2c, 0x7d,
	0xf2, 0x75, 0xe9, 0x93, 0xef, 0x4b, 0x9f, 0x7c, 0xfa, 0xe1, 0x6f, 0xd1, 0xa0, 0x14, 0x7d, 0xba,
	0xd3, 0x2b, 0x67, 0x25, 0x0f, 0x1b, 0xc4, 0x21, 0x79, 0xf1, 0xb0, 0x28, 0xcd, 0xab, 0x6a, 0xde,
	0x44, 0x23, 0x42, 0x72, 0x58, 0x72, 0x6d, 0x54, 0xb5, 0x00, 0x6e, 0x52, 0x53, 0x0a, 0x1e, 0x9d,
	0xea, 0x86, 0xf8, 0xa2, 0x0b, 0xe0, 0x61, 0x21, 0x8a, 0x3f, 0xff, 0x59, 0x3f, 0x07, 0xa3, 0x27,
	0x12, 0xf8, 0x0c, 0x2d, 0xd8, 0x06, 0xec, 0xbe, 0xb5, 0x60, 0x1b, 0xb3, 0xe7, 0xf1, 0x7c, 0x68,
	0x05, 0xf6, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x21, 0x0f, 0x5b, 0xa3, 0xff, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	// After initialization, this RPC must be kept alive for the entire life of
	// the application. The agent pushes configs down to applications via a
	// stream.
	Config(ctx context.Context, opts ...grpc.CallOption) (TraceService_ConfigClient, error)
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, opts ...grpc.CallOption) (TraceService_ExportClient, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) Config(ctx context.Context, opts ...grpc.CallOption) (TraceService_ConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[0], "/opencensus.proto.agent.trace.v1.TraceService/Config", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceConfigClient{stream}
	return x, nil
}

type TraceService_ConfigClient interface {
	Send(*CurrentLibraryConfig) error
	Recv() (*UpdatedLibraryConfig, error)
	grpc.ClientStream
}

type traceServiceConfigClient struct {
	grpc.ClientStream
}

func (x *traceServiceConfigClient) Send(m *CurrentLibraryConfig) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceConfigClient) Recv() (*UpdatedLibraryConfig, error) {
	m := new(UpdatedLibraryConfig)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *traceServiceClient) Export(ctx context.Context, opts ...grpc.CallOption) (TraceService_ExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[1], "/opencensus.proto.agent.trace.v1.TraceService/Export", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceExportClient{stream}
	return x, nil
}

type TraceService_ExportClient interface {
	Send(*ExportTraceServiceRequest) error
	Recv() (*ExportTraceServiceResponse, error)
	grpc.ClientStream
}

type traceServiceExportClient struct {
	grpc.ClientStream
}

func (x *traceServiceExportClient) Send(m *ExportTraceServiceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceExportClient) Recv() (*ExportTraceServiceResponse, error) {
	m := new(ExportTraceServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	// After initialization, this RPC must be kept alive for the entire life of
	// the application. The agent pushes configs down to applications via a
	// stream.
	Config(TraceService_ConfigServer) error
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(TraceService_ExportServer) error
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_Config_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).Config(&traceServiceConfigServer{stream})
}

type TraceService_ConfigServer interface {
	Send(*UpdatedLibraryConfig) error
	Recv() (*CurrentLibraryConfig, error)
	grpc.ServerStream
}

type traceServiceConfigServer struct {
	grpc.ServerStream
}

func (x *traceServiceConfigServer) Send(m *UpdatedLibraryConfig) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceConfigServer) Recv() (*CurrentLibraryConfig, error) {
	m := new(CurrentLibraryConfig)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TraceService_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).Export(&traceServiceExportServer{stream})
}

type TraceService_ExportServer interface {
	Send(*ExportTraceServiceResponse) error
	Recv() (*ExportTraceServiceRequest, error)
	grpc.ServerStream
}

type traceServiceExportServer struct {
	grpc.ServerStream
}

func (x *traceServiceExportServer) Send(m *ExportTraceServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceExportServer) Recv() (*ExportTraceServiceRequest, error) {
	m := new(ExportTraceServiceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opencensus.proto.agent.trace.v1.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Config",
			Handler:       _TraceService_Config_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Export",
			Handler:       _TraceService_Export_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "opencensus/proto/agent/trace/v1/trace_service.proto",
}

func (m *CurrentLibraryConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrentLibraryConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTraceService(dAtA, i, uint64(m.Node.Size()))
		n1, err1 := m.Node.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if m.Config != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTraceService(dAtA, i, uint64(m.Config.Size()))
		n2, err2 := m.Config.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UpdatedLibraryConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatedLibraryConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTraceService(dAtA, i, uint64(m.Node.Size()))
		n3, err3 := m.Node.MarshalTo(dAtA[i:])
		if err3 != nil {
			return 0, err3
		}
		i += n3
	}
	if m.Config != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTraceService(dAtA, i, uint64(m.Config.Size()))
		n4, err4 := m.Config.MarshalTo(dAtA[i:])
		if err4 != nil {
			return 0, err4
		}
		i += n4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ExportTraceServiceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportTraceServiceRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTraceService(dAtA, i, uint64(m.Node.Size()))
		n5, err5 := m.Node.MarshalTo(dAtA[i:])
		if err5 != nil {
			return 0, err5
		}
		i += n5
	}
	if len(m.Spans) > 0 {
		for _, msg := range m.Spans {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTraceService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Resource != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTraceService(dAtA, i, uint64(m.Resource.Size()))
		n6, err6 := m.Resource.MarshalTo(dAtA[i:])
		if err6 != nil {
			return 0, err6
		}
		i += n6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ExportTraceServiceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportTraceServiceResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTraceService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CurrentLibraryConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovTraceService(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovTraceService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdatedLibraryConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovTraceService(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovTraceService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExportTraceServiceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovTraceService(uint64(l))
	}
	if len(m.Spans) > 0 {
		for _, e := range m.Spans {
			l = e.Size()
			n += 1 + l + sovTraceService(uint64(l))
		}
	}
	if m.Resource != nil {
		l = m.Resource.Size()
		n += 1 + l + sovTraceService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExportTraceServiceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTraceService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTraceService(x uint64) (n int) {
	return sovTraceService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CurrentLibraryConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CurrentLibraryConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrentLibraryConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &v1.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &v11.TraceConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdatedLibraryConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdatedLibraryConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatedLibraryConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &v1.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &v11.TraceConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportTraceServiceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportTraceServiceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportTraceServiceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &v1.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spans = append(m.Spans, &v11.Span{})
			if err := m.Spans[len(m.Spans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &v12.Resource{}
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportTraceServiceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportTraceServiceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportTraceServiceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTraceService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTraceService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraceService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTraceService
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTraceService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTraceService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTraceService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTraceService
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTraceService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraceService   = fmt.Errorf("proto: integer overflow")
)
