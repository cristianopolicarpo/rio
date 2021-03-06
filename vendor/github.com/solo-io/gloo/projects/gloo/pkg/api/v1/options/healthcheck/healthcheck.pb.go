// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/healthcheck/healthcheck.proto

package healthcheck

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Add this config to a Listener/Gateway to Enable Envoy Health Checks on that port
type HealthCheck struct {
	// match health check requests using this exact path
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6406d046cf2deb6, []int{0}
}
func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "healthcheck.options.gloo.solo.io.HealthCheck")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/healthcheck/healthcheck.proto", fileDescriptor_d6406d046cf2deb6)
}

var fileDescriptor_d6406d046cf2deb6 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0x31, 0x4e, 0xc3, 0x40,
	0x10, 0x94, 0xa5, 0x08, 0x09, 0xd3, 0x59, 0x14, 0x90, 0x22, 0x0a, 0x54, 0x34, 0xdc, 0x0a, 0xf1,
	0x03, 0xa0, 0xa0, 0xa2, 0x48, 0x49, 0x67, 0x5b, 0x9b, 0xbb, 0xc3, 0x86, 0x59, 0xdd, 0x6d, 0x20,
	0x4f, 0xe2, 0x09, 0xbc, 0x87, 0x3f, 0xd0, 0xa3, 0xbb, 0x33, 0x92, 0x2b, 0x94, 0x6e, 0x76, 0x77,
	0x66, 0x67, 0x34, 0xf5, 0xc6, 0x7a, 0x75, 0xbb, 0xce, 0xf4, 0x78, 0xa5, 0x88, 0x11, 0xd7, 0x1e,
	0x64, 0x47, 0x80, 0x24, 0xe0, 0x85, 0x7b, 0x8d, 0x65, 0x6a, 0xc5, 0xd3, 0xfb, 0x0d, 0x41, 0xd4,
	0xe3, 0x2d, 0x92, 0xe3, 0x76, 0x54, 0xd7, 0x3b, 0xee, 0x87, 0x39, 0x36, 0x12, 0xa0, 0x68, 0xd6,
	0xf3, 0xd5, 0x24, 0x31, 0xe9, 0x8d, 0x49, 0x0e, 0xc6, 0x63, 0x79, 0x6a, 0x61, 0x91, 0xc9, 0x94,
	0x50, 0xd1, 0x2d, 0xcf, 0x73, 0x80, 0xc1, 0xeb, 0x9f, 0x5d, 0xe0, 0xed, 0x74, 0x6a, 0x78, 0xaf,
	0x85, 0xcf, 0x7b, 0x9d, 0x76, 0x2b, 0x0b, 0xd8, 0x91, 0x29, 0x4f, 0xdd, 0x6e, 0x4b, 0x1f, 0xa1,
	0x15, 0xe1, 0x10, 0xcb, 0xfd, 0xf2, 0xa2, 0x3e, 0x79, 0xcc, 0x41, 0xee, 0x53, 0x90, 0xa6, 0xa9,
	0x17, 0xd2, 0xaa, 0x3b, 0xab, 0xd6, 0xd5, 0xd5, 0xf1, 0x26, 0xe3, 0xbb, 0xa7, 0xaf, 0x9f, 0x45,
	0xf5, 0xf9, 0xbd, 0xaa, 0x9e, 0x1f, 0x0e, 0xeb, 0x41, 0x06, 0xfb, 0x4f, 0x17, 0xdd, 0x51, 0x76,
	0xbe, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xc1, 0x10, 0xc6, 0x56, 0x01, 0x00, 0x00,
}

func (this *HealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
