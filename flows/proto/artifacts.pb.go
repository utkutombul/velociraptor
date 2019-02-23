// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifacts.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Artifacts struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifacts) Reset()         { *m = Artifacts{} }
func (m *Artifacts) String() string { return proto.CompactTextString(m) }
func (*Artifacts) ProtoMessage()    {}
func (*Artifacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_f6c9f82cfb728c46, []int{0}
}
func (m *Artifacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifacts.Unmarshal(m, b)
}
func (m *Artifacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifacts.Marshal(b, m, deterministic)
}
func (dst *Artifacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifacts.Merge(dst, src)
}
func (m *Artifacts) XXX_Size() int {
	return xxx_messageInfo_Artifacts.Size(m)
}
func (m *Artifacts) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifacts.DiscardUnknown(m)
}

var xxx_messageInfo_Artifacts proto.InternalMessageInfo

func (m *Artifacts) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type ArtifactParameters struct {
	Env                  []*proto1.VQLEnv `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	Files                []*proto1.VQLEnv `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArtifactParameters) Reset()         { *m = ArtifactParameters{} }
func (m *ArtifactParameters) String() string { return proto.CompactTextString(m) }
func (*ArtifactParameters) ProtoMessage()    {}
func (*ArtifactParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_f6c9f82cfb728c46, []int{1}
}
func (m *ArtifactParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactParameters.Unmarshal(m, b)
}
func (m *ArtifactParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactParameters.Marshal(b, m, deterministic)
}
func (dst *ArtifactParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactParameters.Merge(dst, src)
}
func (m *ArtifactParameters) XXX_Size() int {
	return xxx_messageInfo_ArtifactParameters.Size(m)
}
func (m *ArtifactParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactParameters proto.InternalMessageInfo

func (m *ArtifactParameters) GetEnv() []*proto1.VQLEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *ArtifactParameters) GetFiles() []*proto1.VQLEnv {
	if m != nil {
		return m.Files
	}
	return nil
}

type ArtifactCollectorArgs struct {
	Artifacts            *Artifacts          `protobuf:"bytes,1,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	Parameters           *ArtifactParameters `protobuf:"bytes,5,opt,name=parameters,proto3" json:"parameters,omitempty"`
	OpsPerSecond         float32             `protobuf:"fixed32,6,opt,name=ops_per_second,json=opsPerSecond,proto3" json:"ops_per_second,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ArtifactCollectorArgs) Reset()         { *m = ArtifactCollectorArgs{} }
func (m *ArtifactCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorArgs) ProtoMessage()    {}
func (*ArtifactCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_f6c9f82cfb728c46, []int{2}
}
func (m *ArtifactCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorArgs.Unmarshal(m, b)
}
func (m *ArtifactCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorArgs.Marshal(b, m, deterministic)
}
func (dst *ArtifactCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorArgs.Merge(dst, src)
}
func (m *ArtifactCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorArgs.Size(m)
}
func (m *ArtifactCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorArgs proto.InternalMessageInfo

func (m *ArtifactCollectorArgs) GetArtifacts() *Artifacts {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetParameters() *ArtifactParameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetOpsPerSecond() float32 {
	if m != nil {
		return m.OpsPerSecond
	}
	return 0
}

func init() {
	proto.RegisterType((*Artifacts)(nil), "proto.Artifacts")
	proto.RegisterType((*ArtifactParameters)(nil), "proto.ArtifactParameters")
	proto.RegisterType((*ArtifactCollectorArgs)(nil), "proto.ArtifactCollectorArgs")
}

func init() { proto.RegisterFile("artifacts.proto", fileDescriptor_artifacts_f6c9f82cfb728c46) }

var fileDescriptor_artifacts_f6c9f82cfb728c46 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0xe5, 0x86, 0x04, 0x75, 0x0b, 0x05, 0xad, 0x84, 0x64, 0x22, 0x0e, 0x23, 0x73, 0x31,
	0x1c, 0x9c, 0x92, 0x4a, 0x1c, 0x90, 0x8a, 0x94, 0x84, 0xc0, 0xa5, 0x52, 0x83, 0xa9, 0x10, 0xb7,
	0x6a, 0x63, 0x4f, 0x92, 0x85, 0xf5, 0xee, 0x76, 0x76, 0xeb, 0x94, 0xf7, 0xe3, 0xc4, 0x63, 0xc0,
	0x85, 0x87, 0xe0, 0x80, 0xb2, 0xce, 0x9f, 0x82, 0x40, 0xc2, 0x17, 0x5b, 0xb3, 0xb3, 0xdf, 0xf7,
	0x9b, 0x91, 0xd9, 0x3d, 0x41, 0x5e, 0xce, 0x44, 0xe1, 0x5d, 0x66, 0xc9, 0x78, 0xc3, 0xdb, 0xe1,
	0xd5, 0x7d, 0xb1, 0x5c, 0x2e, 0xb3, 0x1a, 0x95, 0x29, 0x64, 0x89, 0xd7, 0x59, 0x61, 0xaa, 0xde,
	0xdc, 0x28, 0xa1, 0xe7, 0xbd, 0xa6, 0x48, 0xc2, 0x7a, 0x43, 0xbd, 0xd0, 0xdc, 0x73, 0x58, 0x09,
	0xed, 0x65, 0xd1, 0x20, 0xba, 0x27, 0xff, 0x77, 0x57, 0x14, 0x5e, 0x1a, 0xed, 0xd6, 0x8c, 0xfa,
	0x52, 0x35, 0xd7, 0x93, 0x11, 0xdb, 0x1f, 0x6c, 0x42, 0xf1, 0xe7, 0xac, 0xad, 0x45, 0x85, 0x2e,
	0x8e, 0xa0, 0x95, 0xee, 0x0f, 0xe1, 0xdb, 0xcf, 0xef, 0x5f, 0xa2, 0x2e, 0x8f, 0xcf, 0x17, 0x08,
	0xdb, 0xe8, 0xe0, 0x0d, 0x28, 0x71, 0xa5, 0x8b, 0x45, 0x96, 0x37, 0xed, 0xc9, 0x8f, 0x88, 0xf1,
	0x0d, 0x65, 0x22, 0x48, 0x54, 0xe8, 0x91, 0x1c, 0x2f, 0x59, 0x0b, 0x75, 0x1d, 0xb7, 0xa0, 0x95,
	0x1e, 0xf4, 0xef, 0x36, 0xc2, 0xec, 0xfd, 0xdb, 0xd3, 0xb1, 0xae, 0x87, 0xa3, 0xc0, 0x3e, 0xe1,
	0xc7, 0x63, 0x5d, 0x4b, 0x32, 0xba, 0x42, 0xed, 0xa1, 0x16, 0x24, 0xc5, 0x54, 0x61, 0x70, 0x4c,
	0x11, 0x2c, 0x99, 0x5a, 0x96, 0x58, 0xc2, 0xcc, 0x10, 0xf8, 0x05, 0xc2, 0xe5, 0x15, 0xd2, 0xe7,
	0x2c, 0xe9, 0x04, 0x89, 0xcb, 0x57, 0x78, 0xae, 0x58, 0x7b, 0x26, 0x15, 0xba, 0xf8, 0xd6, 0xdf,
	0x3c, 0x6f, 0x82, 0x67, 0xc0, 0x7b, 0xbb, 0x5c, 0x6b, 0xf8, 0x4c, 0x2a, 0x85, 0x25, 0x48, 0x0d,
	0x33, 0x32, 0x15, 0xe0, 0xb5, 0x35, 0xe4, 0x57, 0xae, 0x15, 0x2c, 0x4b, 0x0e, 0xc7, 0x9b, 0xc2,
	0xeb, 0x55, 0x21, 0x6f, 0x24, 0xc9, 0xd7, 0x3d, 0xf6, 0x60, 0x33, 0xea, 0xc8, 0x28, 0x85, 0x85,
	0x37, 0x34, 0xa0, 0xb9, 0xe3, 0x1f, 0xd8, 0xfe, 0x76, 0x47, 0x71, 0x04, 0x51, 0x7a, 0xd0, 0xbf,
	0xbf, 0xce, 0xb2, 0xdd, 0xf0, 0x30, 0x0d, 0x71, 0x92, 0x7f, 0xaf, 0x34, 0xe9, 0x9c, 0x86, 0x8f,
	0x7c, 0x07, 0xe3, 0x9f, 0x18, 0xb3, 0xdb, 0xf4, 0x71, 0x3b, 0xa0, 0x1f, 0xfe, 0x81, 0xde, 0x8d,
	0x37, 0x3c, 0x0a, 0x8e, 0xa7, 0xfc, 0xd1, 0xef, 0x23, 0xfb, 0x9b, 0xc6, 0x2c, 0x61, 0xbb, 0xd3,
	0xfc, 0x06, 0x9e, 0x7f, 0x64, 0x87, 0xc6, 0xba, 0x0b, 0x8b, 0x74, 0xe1, 0xb0, 0x30, 0xba, 0x8c,
	0x3b, 0x10, 0xa5, 0x7b, 0xc3, 0x57, 0x81, 0xfa, 0x92, 0x3f, 0x3e, 0xb3, 0x48, 0x22, 0xfc, 0x4c,
	0x60, 0x91, 0xa0, 0x69, 0x82, 0xf4, 0x7c, 0x41, 0xc6, 0x7b, 0x25, 0xf5, 0xfc, 0x49, 0x96, 0x1c,
	0x9e, 0x59, 0x07, 0x13, 0x24, 0x78, 0x17, 0x4e, 0xfb, 0xb7, 0x9f, 0x1d, 0x85, 0x27, 0xbf, 0x63,
	0xac, 0x9b, 0x20, 0x35, 0xe5, 0x69, 0x27, 0xcc, 0x70, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xe6, 0x12, 0xc6, 0x18, 0x03, 0x00, 0x00,
}
