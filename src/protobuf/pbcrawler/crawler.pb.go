// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/pbcrawler/crawler.proto

package pbcrawler

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// WebCrawler simple message
type WebCrawlerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebCrawlerRequest) Reset()         { *m = WebCrawlerRequest{} }
func (m *WebCrawlerRequest) String() string { return proto.CompactTextString(m) }
func (*WebCrawlerRequest) ProtoMessage()    {}
func (*WebCrawlerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_937b0d7bee9b0fb9, []int{0}
}

func (m *WebCrawlerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebCrawlerRequest.Unmarshal(m, b)
}
func (m *WebCrawlerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebCrawlerRequest.Marshal(b, m, deterministic)
}
func (m *WebCrawlerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebCrawlerRequest.Merge(m, src)
}
func (m *WebCrawlerRequest) XXX_Size() int {
	return xxx_messageInfo_WebCrawlerRequest.Size(m)
}
func (m *WebCrawlerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebCrawlerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebCrawlerRequest proto.InternalMessageInfo

func (m *WebCrawlerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WebCrawlerRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type WebCrawlerResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebCrawlerResponse) Reset()         { *m = WebCrawlerResponse{} }
func (m *WebCrawlerResponse) String() string { return proto.CompactTextString(m) }
func (*WebCrawlerResponse) ProtoMessage()    {}
func (*WebCrawlerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_937b0d7bee9b0fb9, []int{1}
}

func (m *WebCrawlerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebCrawlerResponse.Unmarshal(m, b)
}
func (m *WebCrawlerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebCrawlerResponse.Marshal(b, m, deterministic)
}
func (m *WebCrawlerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebCrawlerResponse.Merge(m, src)
}
func (m *WebCrawlerResponse) XXX_Size() int {
	return xxx_messageInfo_WebCrawlerResponse.Size(m)
}
func (m *WebCrawlerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WebCrawlerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WebCrawlerResponse proto.InternalMessageInfo

func (m *WebCrawlerResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *WebCrawlerResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WebCrawlerResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Web Crawler batch messages
type WebCrawlerBatchRequest struct {
	Requests             []*WebCrawlerBatchRequest_Request `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *WebCrawlerBatchRequest) Reset()         { *m = WebCrawlerBatchRequest{} }
func (m *WebCrawlerBatchRequest) String() string { return proto.CompactTextString(m) }
func (*WebCrawlerBatchRequest) ProtoMessage()    {}
func (*WebCrawlerBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_937b0d7bee9b0fb9, []int{2}
}

func (m *WebCrawlerBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebCrawlerBatchRequest.Unmarshal(m, b)
}
func (m *WebCrawlerBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebCrawlerBatchRequest.Marshal(b, m, deterministic)
}
func (m *WebCrawlerBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebCrawlerBatchRequest.Merge(m, src)
}
func (m *WebCrawlerBatchRequest) XXX_Size() int {
	return xxx_messageInfo_WebCrawlerBatchRequest.Size(m)
}
func (m *WebCrawlerBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebCrawlerBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebCrawlerBatchRequest proto.InternalMessageInfo

func (m *WebCrawlerBatchRequest) GetRequests() []*WebCrawlerBatchRequest_Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

type WebCrawlerBatchRequest_Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebCrawlerBatchRequest_Request) Reset()         { *m = WebCrawlerBatchRequest_Request{} }
func (m *WebCrawlerBatchRequest_Request) String() string { return proto.CompactTextString(m) }
func (*WebCrawlerBatchRequest_Request) ProtoMessage()    {}
func (*WebCrawlerBatchRequest_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_937b0d7bee9b0fb9, []int{2, 0}
}

func (m *WebCrawlerBatchRequest_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebCrawlerBatchRequest_Request.Unmarshal(m, b)
}
func (m *WebCrawlerBatchRequest_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebCrawlerBatchRequest_Request.Marshal(b, m, deterministic)
}
func (m *WebCrawlerBatchRequest_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebCrawlerBatchRequest_Request.Merge(m, src)
}
func (m *WebCrawlerBatchRequest_Request) XXX_Size() int {
	return xxx_messageInfo_WebCrawlerBatchRequest_Request.Size(m)
}
func (m *WebCrawlerBatchRequest_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_WebCrawlerBatchRequest_Request.DiscardUnknown(m)
}

var xxx_messageInfo_WebCrawlerBatchRequest_Request proto.InternalMessageInfo

func (m *WebCrawlerBatchRequest_Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WebCrawlerBatchRequest_Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type WebCrawlerBatchResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebCrawlerBatchResponse) Reset()         { *m = WebCrawlerBatchResponse{} }
func (m *WebCrawlerBatchResponse) String() string { return proto.CompactTextString(m) }
func (*WebCrawlerBatchResponse) ProtoMessage()    {}
func (*WebCrawlerBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_937b0d7bee9b0fb9, []int{3}
}

func (m *WebCrawlerBatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebCrawlerBatchResponse.Unmarshal(m, b)
}
func (m *WebCrawlerBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebCrawlerBatchResponse.Marshal(b, m, deterministic)
}
func (m *WebCrawlerBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebCrawlerBatchResponse.Merge(m, src)
}
func (m *WebCrawlerBatchResponse) XXX_Size() int {
	return xxx_messageInfo_WebCrawlerBatchResponse.Size(m)
}
func (m *WebCrawlerBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WebCrawlerBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WebCrawlerBatchResponse proto.InternalMessageInfo

func (m *WebCrawlerBatchResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *WebCrawlerBatchResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WebCrawlerBatchResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*WebCrawlerRequest)(nil), "pbcrawler.WebCrawlerRequest")
	proto.RegisterType((*WebCrawlerResponse)(nil), "pbcrawler.WebCrawlerResponse")
	proto.RegisterType((*WebCrawlerBatchRequest)(nil), "pbcrawler.WebCrawlerBatchRequest")
	proto.RegisterType((*WebCrawlerBatchRequest_Request)(nil), "pbcrawler.WebCrawlerBatchRequest.Request")
	proto.RegisterType((*WebCrawlerBatchResponse)(nil), "pbcrawler.WebCrawlerBatchResponse")
}

func init() { proto.RegisterFile("protobuf/pbcrawler/crawler.proto", fileDescriptor_937b0d7bee9b0fb9) }

var fileDescriptor_937b0d7bee9b0fb9 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x48, 0x4a, 0x2e, 0x4a, 0x2c, 0xcf, 0x49, 0x2d, 0xd2, 0x87,
	0xd2, 0x7a, 0x60, 0x29, 0x21, 0x4e, 0xb8, 0x84, 0x92, 0x25, 0x97, 0x60, 0x78, 0x6a, 0x92, 0x33,
	0x84, 0x17, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b,
	0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x09, 0x70, 0x31, 0x97, 0x16, 0xe5,
	0x48, 0x30, 0x81, 0x85, 0x40, 0x4c, 0xa5, 0x28, 0x2e, 0x21, 0x64, 0xad, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x60, 0xdd, 0x1c, 0x41,
	0x50, 0x1e, 0xdc, 0x4c, 0x26, 0x24, 0x33, 0x25, 0xb8, 0xd8, 0x0b, 0x12, 0x2b, 0x73, 0xf2, 0x13,
	0x53, 0x24, 0x98, 0x15, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0xa5, 0x09, 0x8c, 0x5c, 0x62, 0x08,
	0xc3, 0x9d, 0x12, 0x4b, 0x92, 0x33, 0x60, 0x8e, 0x73, 0xe5, 0xe2, 0x28, 0x82, 0x30, 0x41, 0x56,
	0x30, 0x6b, 0x70, 0x1b, 0x69, 0xea, 0xc1, 0xfd, 0xa3, 0x87, 0x5d, 0x93, 0x1e, 0x94, 0x0e, 0x82,
	0x6b, 0x95, 0xd2, 0xe7, 0x62, 0x27, 0xcd, 0xbb, 0xf1, 0x5c, 0xe2, 0x18, 0x86, 0x53, 0xd3, 0xcf,
	0x46, 0x07, 0x18, 0x91, 0xe3, 0x22, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x9b, 0x8b,
	0x0b, 0x21, 0x28, 0x24, 0x83, 0xd5, 0xab, 0x50, 0x8f, 0x48, 0xc9, 0xe2, 0x90, 0x85, 0x38, 0x53,
	0x89, 0x41, 0x28, 0x8e, 0x8b, 0x1f, 0xcd, 0x0f, 0x42, 0x8a, 0x04, 0x03, 0x4f, 0x4a, 0x09, 0x9f,
	0x12, 0x98, 0xd9, 0x1a, 0x8c, 0x06, 0x8c, 0x4e, 0xdc, 0x51, 0x88, 0xa4, 0x95, 0xc4, 0x06, 0x4e,
	0x6c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x1d, 0xf9, 0xb3, 0x90, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WebCrawlerServiceClient is the client API for WebCrawlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebCrawlerServiceClient interface {
	// Unary
	WebCrawler(ctx context.Context, in *WebCrawlerRequest, opts ...grpc.CallOption) (*WebCrawlerResponse, error)
	// BiDi
	WebCrawlerBatch(ctx context.Context, opts ...grpc.CallOption) (WebCrawlerService_WebCrawlerBatchClient, error)
}

type webCrawlerServiceClient struct {
	cc *grpc.ClientConn
}

func NewWebCrawlerServiceClient(cc *grpc.ClientConn) WebCrawlerServiceClient {
	return &webCrawlerServiceClient{cc}
}

func (c *webCrawlerServiceClient) WebCrawler(ctx context.Context, in *WebCrawlerRequest, opts ...grpc.CallOption) (*WebCrawlerResponse, error) {
	out := new(WebCrawlerResponse)
	err := c.cc.Invoke(ctx, "/pbcrawler.WebCrawlerService/WebCrawler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webCrawlerServiceClient) WebCrawlerBatch(ctx context.Context, opts ...grpc.CallOption) (WebCrawlerService_WebCrawlerBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WebCrawlerService_serviceDesc.Streams[0], "/pbcrawler.WebCrawlerService/WebCrawlerBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &webCrawlerServiceWebCrawlerBatchClient{stream}
	return x, nil
}

type WebCrawlerService_WebCrawlerBatchClient interface {
	Send(*WebCrawlerBatchRequest) error
	Recv() (*WebCrawlerBatchResponse, error)
	grpc.ClientStream
}

type webCrawlerServiceWebCrawlerBatchClient struct {
	grpc.ClientStream
}

func (x *webCrawlerServiceWebCrawlerBatchClient) Send(m *WebCrawlerBatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *webCrawlerServiceWebCrawlerBatchClient) Recv() (*WebCrawlerBatchResponse, error) {
	m := new(WebCrawlerBatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WebCrawlerServiceServer is the server API for WebCrawlerService service.
type WebCrawlerServiceServer interface {
	// Unary
	WebCrawler(context.Context, *WebCrawlerRequest) (*WebCrawlerResponse, error)
	// BiDi
	WebCrawlerBatch(WebCrawlerService_WebCrawlerBatchServer) error
}

// UnimplementedWebCrawlerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWebCrawlerServiceServer struct {
}

func (*UnimplementedWebCrawlerServiceServer) WebCrawler(ctx context.Context, req *WebCrawlerRequest) (*WebCrawlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebCrawler not implemented")
}
func (*UnimplementedWebCrawlerServiceServer) WebCrawlerBatch(srv WebCrawlerService_WebCrawlerBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method WebCrawlerBatch not implemented")
}

func RegisterWebCrawlerServiceServer(s *grpc.Server, srv WebCrawlerServiceServer) {
	s.RegisterService(&_WebCrawlerService_serviceDesc, srv)
}

func _WebCrawlerService_WebCrawler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebCrawlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebCrawlerServiceServer).WebCrawler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbcrawler.WebCrawlerService/WebCrawler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebCrawlerServiceServer).WebCrawler(ctx, req.(*WebCrawlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebCrawlerService_WebCrawlerBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WebCrawlerServiceServer).WebCrawlerBatch(&webCrawlerServiceWebCrawlerBatchServer{stream})
}

type WebCrawlerService_WebCrawlerBatchServer interface {
	Send(*WebCrawlerBatchResponse) error
	Recv() (*WebCrawlerBatchRequest, error)
	grpc.ServerStream
}

type webCrawlerServiceWebCrawlerBatchServer struct {
	grpc.ServerStream
}

func (x *webCrawlerServiceWebCrawlerBatchServer) Send(m *WebCrawlerBatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *webCrawlerServiceWebCrawlerBatchServer) Recv() (*WebCrawlerBatchRequest, error) {
	m := new(WebCrawlerBatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _WebCrawlerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbcrawler.WebCrawlerService",
	HandlerType: (*WebCrawlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WebCrawler",
			Handler:    _WebCrawlerService_WebCrawler_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WebCrawlerBatch",
			Handler:       _WebCrawlerService_WebCrawlerBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/pbcrawler/crawler.proto",
}