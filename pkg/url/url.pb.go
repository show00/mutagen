// Code generated by protoc-gen-go. DO NOT EDIT.
// source: url/url.proto

package url // import "github.com/havoc-io/mutagen/pkg/url"

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

type Protocol int32

const (
	Protocol_Local Protocol = 0
	Protocol_SSH   Protocol = 1
)

var Protocol_name = map[int32]string{
	0: "Local",
	1: "SSH",
}
var Protocol_value = map[string]int32{
	"Local": 0,
	"SSH":   1,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_url_79f64e26259b2ab2, []int{0}
}

type URL struct {
	Protocol             Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=url.Protocol" json:"protocol,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Hostname             string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_url_79f64e26259b2ab2, []int{0}
}
func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (dst *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(dst, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_Local
}

func (m *URL) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *URL) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *URL) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *URL) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*URL)(nil), "url.URL")
	proto.RegisterEnum("url.Protocol", Protocol_name, Protocol_value)
}

func init() { proto.RegisterFile("url/url.proto", fileDescriptor_url_79f64e26259b2ab2) }

var fileDescriptor_url_79f64e26259b2ab2 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x41, 0x4b, 0x87, 0x30,
	0x18, 0xc6, 0x5b, 0xfb, 0x5b, 0xfa, 0x82, 0x21, 0x3b, 0x8d, 0x0e, 0x21, 0x45, 0x60, 0x41, 0x0e,
	0xea, 0x1b, 0x74, 0xea, 0xe0, 0x21, 0x26, 0x5d, 0xba, 0xcd, 0x21, 0x4e, 0x9a, 0xbe, 0x32, 0xb7,
	0x3e, 0x47, 0x1f, 0x39, 0x5c, 0xe9, 0xed, 0x79, 0x7e, 0xbf, 0x07, 0xb6, 0x17, 0xf2, 0xe0, 0xac,
	0x08, 0xce, 0xd6, 0x8b, 0x43, 0x8f, 0x8c, 0x06, 0x67, 0x6f, 0x7f, 0x08, 0xd0, 0x0f, 0xd9, 0xb0,
	0x07, 0x48, 0x23, 0xd5, 0x68, 0x39, 0x29, 0x49, 0x75, 0xf5, 0x9c, 0xd7, 0xdb, 0xf4, 0xfd, 0x1f,
	0xca, 0x43, 0xb3, 0x6b, 0x48, 0xc3, 0xda, 0xbb, 0x59, 0x4d, 0x3d, 0x3f, 0x2f, 0x49, 0x95, 0xc9,
	0xa3, 0x6f, 0xce, 0xe0, 0xea, 0xa3, 0xa3, 0x7f, 0x6e, 0xef, 0x8c, 0xc1, 0x69, 0x41, 0xe7, 0xf9,
	0xa9, 0x24, 0x55, 0x2e, 0x63, 0x8e, 0x4c, 0x79, 0xc3, 0x93, 0xb8, 0x8d, 0xf9, 0xf1, 0x06, 0xd2,
	0xfd, 0x55, 0x96, 0x41, 0xd2, 0xa0, 0x56, 0xb6, 0x38, 0x63, 0x97, 0x40, 0xdb, 0xf6, 0xad, 0x20,
	0xaf, 0xf7, 0x9f, 0x77, 0xc3, 0xe8, 0x4d, 0xe8, 0x6a, 0x8d, 0x93, 0x30, 0xea, 0x1b, 0xf5, 0xd3,
	0x88, 0x62, 0x0a, 0x5e, 0x0d, 0xfd, 0x2c, 0x96, 0xaf, 0x61, 0x3b, 0xb2, 0xbb, 0x88, 0x1f, 0x7e,
	0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xee, 0x81, 0x51, 0xff, 0xf6, 0x00, 0x00, 0x00,
}