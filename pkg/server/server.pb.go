// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/yiqinguo/armyant/pkg/server/server.proto

package server

/*
import "k8s.io/apimachinery/pkg/api/resource/generated.proto";
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import models "github.com/yiqinguo/armyant/pkg/models"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The response message containing the greetings
type Response struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_2599d8476497a96c, []int{0}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "server.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiserverClient is the client API for Apiserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiserverClient interface {
	// rpc CreateJob (models.Job) returns (Response) {}
	ReportResult(ctx context.Context, in *models.Stats, opts ...grpc.CallOption) (*Response, error)
	ReportStatus(ctx context.Context, in *models.Status, opts ...grpc.CallOption) (*Response, error)
	CreateJob(ctx context.Context, in *models.Job, opts ...grpc.CallOption) (*Response, error)
	ListJobs(ctx context.Context, in *models.ListJobRequest, opts ...grpc.CallOption) (*Response, error)
	GetJob(ctx context.Context, in *models.GetJobRequest, opts ...grpc.CallOption) (*models.JobResultResponse, error)
	GetJobResult(ctx context.Context, in *models.GetJobRequest, opts ...grpc.CallOption) (*models.JobResultResponse, error)
	UpdateJob(ctx context.Context, in *models.Job, opts ...grpc.CallOption) (*Response, error)
}

type apiserverClient struct {
	cc *grpc.ClientConn
}

func NewApiserverClient(cc *grpc.ClientConn) ApiserverClient {
	return &apiserverClient{cc}
}

func (c *apiserverClient) ReportResult(ctx context.Context, in *models.Stats, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/server.Apiserver/ReportResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiserverClient) ReportStatus(ctx context.Context, in *models.Status, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/server.Apiserver/ReportStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiserverClient) CreateJob(ctx context.Context, in *models.Job, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/server.Apiserver/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiserverClient) ListJobs(ctx context.Context, in *models.ListJobRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/server.Apiserver/ListJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiserverClient) GetJob(ctx context.Context, in *models.GetJobRequest, opts ...grpc.CallOption) (*models.JobResultResponse, error) {
	out := new(models.JobResultResponse)
	err := c.cc.Invoke(ctx, "/server.Apiserver/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiserverClient) GetJobResult(ctx context.Context, in *models.GetJobRequest, opts ...grpc.CallOption) (*models.JobResultResponse, error) {
	out := new(models.JobResultResponse)
	err := c.cc.Invoke(ctx, "/server.Apiserver/GetJobResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiserverClient) UpdateJob(ctx context.Context, in *models.Job, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/server.Apiserver/UpdateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiserverServer is the server API for Apiserver service.
type ApiserverServer interface {
	// rpc CreateJob (models.Job) returns (Response) {}
	ReportResult(context.Context, *models.Stats) (*Response, error)
	ReportStatus(context.Context, *models.Status) (*Response, error)
	CreateJob(context.Context, *models.Job) (*Response, error)
	ListJobs(context.Context, *models.ListJobRequest) (*Response, error)
	GetJob(context.Context, *models.GetJobRequest) (*models.JobResultResponse, error)
	GetJobResult(context.Context, *models.GetJobRequest) (*models.JobResultResponse, error)
	UpdateJob(context.Context, *models.Job) (*Response, error)
}

func RegisterApiserverServer(s *grpc.Server, srv ApiserverServer) {
	s.RegisterService(&_Apiserver_serviceDesc, srv)
}

func _Apiserver_ReportResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.Stats)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiserverServer).ReportResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Apiserver/ReportResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiserverServer).ReportResult(ctx, req.(*models.Stats))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apiserver_ReportStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiserverServer).ReportStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Apiserver/ReportStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiserverServer).ReportStatus(ctx, req.(*models.Status))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apiserver_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiserverServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Apiserver/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiserverServer).CreateJob(ctx, req.(*models.Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apiserver_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiserverServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Apiserver/ListJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiserverServer).ListJobs(ctx, req.(*models.ListJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apiserver_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiserverServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Apiserver/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiserverServer).GetJob(ctx, req.(*models.GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apiserver_GetJobResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiserverServer).GetJobResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Apiserver/GetJobResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiserverServer).GetJobResult(ctx, req.(*models.GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apiserver_UpdateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiserverServer).UpdateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Apiserver/UpdateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiserverServer).UpdateJob(ctx, req.(*models.Job))
	}
	return interceptor(ctx, in, info, handler)
}

var _Apiserver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Apiserver",
	HandlerType: (*ApiserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportResult",
			Handler:    _Apiserver_ReportResult_Handler,
		},
		{
			MethodName: "ReportStatus",
			Handler:    _Apiserver_ReportStatus_Handler,
		},
		{
			MethodName: "CreateJob",
			Handler:    _Apiserver_CreateJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _Apiserver_ListJobs_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _Apiserver_GetJob_Handler,
		},
		{
			MethodName: "GetJobResult",
			Handler:    _Apiserver_GetJobResult_Handler,
		},
		{
			MethodName: "UpdateJob",
			Handler:    _Apiserver_UpdateJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/yiqinguo/armyant/pkg/server/server.proto",
}

func init() {
	proto.RegisterFile("github.com/yiqinguo/armyant/pkg/server/server.proto", fileDescriptor_server_2599d8476497a96c)
}

var fileDescriptor_server_2599d8476497a96c = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x9d, 0x93, 0xba, 0xc6, 0x29, 0x12, 0x50, 0x6a, 0x4f, 0x63, 0xa7, 0x1d, 0x46, 0x8b,
	0x2d, 0x88, 0x07, 0x2f, 0xea, 0x41, 0x28, 0x9e, 0x22, 0x7e, 0x80, 0x74, 0xfd, 0x53, 0x8b, 0x6d,
	0x93, 0xe5, 0x9f, 0x0c, 0xf6, 0x0d, 0xfc, 0xd8, 0xb2, 0xa6, 0x81, 0x1e, 0x0a, 0x0e, 0x4f, 0xc9,
	0x0b, 0xef, 0x97, 0x97, 0x97, 0x3f, 0x49, 0xcb, 0x4a, 0x7f, 0x99, 0x3c, 0xda, 0x88, 0x26, 0xde,
	0x57, 0xdb, 0xaa, 0x2d, 0x8d, 0x88, 0xb9, 0x6a, 0xf6, 0xbc, 0xd5, 0xb1, 0xfc, 0x2e, 0x63, 0x04,
	0xb5, 0x03, 0xd5, 0x2f, 0x91, 0x54, 0x42, 0x0b, 0xea, 0x59, 0x15, 0x26, 0x7f, 0xc1, 0x8d, 0x28,
	0xa0, 0xc6, 0x18, 0x6b, 0xbe, 0x03, 0xcb, 0x86, 0xe9, 0x91, 0x4c, 0xc3, 0x51, 0xbb, 0xc0, 0xe5,
	0x23, 0x99, 0x31, 0x40, 0x29, 0x5a, 0x04, 0x4a, 0xc9, 0xd9, 0x46, 0x14, 0x10, 0x4c, 0x16, 0x93,
	0xd5, 0x94, 0x75, 0x7b, 0x1a, 0x90, 0xf3, 0x06, 0x10, 0x79, 0x09, 0xc1, 0xe9, 0x62, 0xb2, 0xf2,
	0x99, 0x93, 0xc9, 0xcf, 0x94, 0xf8, 0xcf, 0xb2, 0xb2, 0x0f, 0xa6, 0xf7, 0x64, 0xce, 0x40, 0x0a,
	0xa5, 0x19, 0xa0, 0xa9, 0x35, 0xbd, 0x8c, 0x6c, 0x5a, 0xf4, 0xa1, 0xb9, 0xc6, 0xf0, 0x3a, 0xea,
	0x6b, 0xba, 0xb0, 0xe5, 0x09, 0x4d, 0x1c, 0x72, 0xb0, 0x18, 0xa4, 0x57, 0x43, 0xc4, 0x8c, 0x33,
	0x6b, 0xe2, 0xbf, 0x2a, 0xe0, 0x1a, 0x32, 0x91, 0xd3, 0x0b, 0x07, 0x64, 0x22, 0x1f, 0x75, 0x3f,
	0x90, 0xd9, 0x7b, 0x85, 0x3a, 0x13, 0x39, 0xd2, 0x5b, 0x67, 0xee, 0x4f, 0x18, 0x6c, 0x0d, 0xa0,
	0x1e, 0xe5, 0x9e, 0x88, 0xf7, 0x06, 0x07, 0x13, 0xbd, 0x71, 0x94, 0xd5, 0x0e, 0xba, 0x1b, 0x24,
	0xdb, 0xc2, 0x03, 0xfa, 0x85, 0xcc, 0x9d, 0xbb, 0xfb, 0x8a, 0xff, 0xdc, 0xb1, 0x26, 0xfe, 0xa7,
	0x2c, 0x8e, 0xec, 0x99, 0x7b, 0xdd, 0x2c, 0xd3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xd0,
	0x32, 0xba, 0x73, 0x02, 0x00, 0x00,
}
