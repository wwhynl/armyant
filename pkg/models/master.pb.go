// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/yiqinguo/armyant/pkg/models/master.proto

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetJobRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=jobId" json:"jobId,omitempty"`
	InstancdId           string   `protobuf:"bytes,2,opt,name=instancdId" json:"instancdId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobRequest) Reset()         { *m = GetJobRequest{} }
func (m *GetJobRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobRequest) ProtoMessage()    {}
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_master_945e62429116b154, []int{0}
}
func (m *GetJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobRequest.Unmarshal(m, b)
}
func (m *GetJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobRequest.Marshal(b, m, deterministic)
}
func (dst *GetJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobRequest.Merge(dst, src)
}
func (m *GetJobRequest) XXX_Size() int {
	return xxx_messageInfo_GetJobRequest.Size(m)
}
func (m *GetJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobRequest proto.InternalMessageInfo

func (m *GetJobRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *GetJobRequest) GetInstancdId() string {
	if m != nil {
		return m.InstancdId
	}
	return ""
}

type ListJobRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobRequest) Reset()         { *m = ListJobRequest{} }
func (m *ListJobRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobRequest) ProtoMessage()    {}
func (*ListJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_master_945e62429116b154, []int{1}
}
func (m *ListJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobRequest.Unmarshal(m, b)
}
func (m *ListJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobRequest.Marshal(b, m, deterministic)
}
func (dst *ListJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobRequest.Merge(dst, src)
}
func (m *ListJobRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobRequest.Size(m)
}
func (m *ListJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetJobRequest)(nil), "models.GetJobRequest")
	proto.RegisterType((*ListJobRequest)(nil), "models.ListJobRequest")
}

func init() {
	proto.RegisterFile("github.com/yiqinguo/armyant/pkg/models/master.proto", fileDescriptor_master_945e62429116b154)
}

var fileDescriptor_master_945e62429116b154 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xaf, 0xcc, 0x2c, 0xcc, 0xcc, 0x4b, 0x2f, 0xcd, 0xd7,
	0x4f, 0x2c, 0xca, 0xad, 0x4c, 0xcc, 0x2b, 0xd1, 0x2f, 0xc8, 0x4e, 0xd7, 0xcf, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0xcf, 0x4d, 0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x83, 0x08, 0x2a, 0xb9, 0x72, 0xf1, 0xba, 0xa7, 0x96, 0x78, 0xe5, 0x27, 0x05, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0x66, 0xe5, 0x27, 0x79, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x72, 0x5c, 0x5c, 0x99, 0x79, 0xc5, 0x25, 0x89, 0x79,
	0xc9, 0x29, 0x9e, 0x29, 0x12, 0x4c, 0x60, 0x29, 0x24, 0x11, 0x25, 0x01, 0x2e, 0x3e, 0x9f, 0xcc,
	0x62, 0x24, 0x73, 0x92, 0xd8, 0xc0, 0xf6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x90,
	0xe9, 0x72, 0x9e, 0x00, 0x00, 0x00,
}
