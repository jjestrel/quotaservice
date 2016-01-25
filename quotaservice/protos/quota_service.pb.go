// Code generated by protoc-gen-go.
// source: quotaservice/protos/quota_service.proto
// DO NOT EDIT!

/*
Package quotaservice is a generated protocol buffer package.

It is generated from these files:
	quotaservice/protos/quota_service.proto

It has these top-level messages:
	AllowRequest
	AllowResponse
*/
package quotaservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AllowRequest struct {
	FromService string `protobuf:"bytes,1,opt,name=from_service" json:"from_service,omitempty"`
	ToService   string `protobuf:"bytes,2,opt,name=to_service" json:"to_service,omitempty"`
}

func (m *AllowRequest) Reset()                    { *m = AllowRequest{} }
func (m *AllowRequest) String() string            { return proto.CompactTextString(m) }
func (*AllowRequest) ProtoMessage()               {}
func (*AllowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllowResponse struct {
	Granted bool `protobuf:"varint,1,opt,name=granted" json:"granted,omitempty"`
}

func (m *AllowResponse) Reset()                    { *m = AllowResponse{} }
func (m *AllowResponse) String() string            { return proto.CompactTextString(m) }
func (*AllowResponse) ProtoMessage()               {}
func (*AllowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*AllowRequest)(nil), "quotaservice.AllowRequest")
	proto.RegisterType((*AllowResponse)(nil), "quotaservice.AllowResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for QuotaService service

type QuotaServiceClient interface {
	Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*AllowResponse, error)
}

type quotaServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuotaServiceClient(cc *grpc.ClientConn) QuotaServiceClient {
	return &quotaServiceClient{cc}
}

func (c *quotaServiceClient) Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*AllowResponse, error) {
	out := new(AllowResponse)
	err := grpc.Invoke(ctx, "/quotaservice.QuotaService/Allow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QuotaService service

type QuotaServiceServer interface {
	Allow(context.Context, *AllowRequest) (*AllowResponse, error)
}

func RegisterQuotaServiceServer(s *grpc.Server, srv QuotaServiceServer) {
	s.RegisterService(&_QuotaService_serviceDesc, srv)
}

func _QuotaService_Allow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AllowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(QuotaServiceServer).Allow(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _QuotaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quotaservice.QuotaService",
	HandlerType: (*QuotaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allow",
			Handler:    _QuotaService_Allow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x2c, 0xcd, 0x2f,
	0x49, 0x2c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6,
	0x07, 0x8b, 0xc5, 0x43, 0x05, 0xf5, 0xc0, 0x82, 0x42, 0x3c, 0xc8, 0x0a, 0x95, 0x2c, 0xb8, 0x78,
	0x1c, 0x73, 0x72, 0xf2, 0xcb, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x78,
	0xd2, 0x8a, 0xf2, 0x73, 0x61, 0x7a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x85, 0x84, 0xb8, 0xb8,
	0x4a, 0xf2, 0xe1, 0x62, 0x4c, 0x20, 0x31, 0x25, 0x05, 0x2e, 0x5e, 0xa8, 0xce, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0x21, 0x7e, 0x2e, 0xf6, 0xf4, 0xa2, 0xc4, 0xbc, 0x92, 0xd4, 0x14, 0xb0, 0x2e,
	0x0e, 0xa3, 0x20, 0x2e, 0x9e, 0x40, 0x90, 0x5d, 0xc1, 0x10, 0x7d, 0x42, 0x4e, 0x5c, 0xac, 0x60,
	0x1d, 0x42, 0x52, 0x7a, 0xc8, 0x6e, 0xd0, 0x43, 0x76, 0x80, 0x94, 0x34, 0x56, 0x39, 0x88, 0x15,
	0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x4f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xb8, 0x8b,
	0x3f, 0xef, 0x00, 0x00, 0x00,
}
