// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/explorer_auto_optimizer_setting.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Settings for the
// <a href="https://support.google.com/google-ads/answer/190596">
// Display Campaign Optimizer</a>, initially termed "Explorer".
type ExplorerAutoOptimizerSetting struct {
	// Indicates whether the optimizer is turned on.
	OptIn                *wrappers.BoolValue `protobuf:"bytes,1,opt,name=opt_in,json=optIn,proto3" json:"opt_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExplorerAutoOptimizerSetting) Reset()         { *m = ExplorerAutoOptimizerSetting{} }
func (m *ExplorerAutoOptimizerSetting) String() string { return proto.CompactTextString(m) }
func (*ExplorerAutoOptimizerSetting) ProtoMessage()    {}
func (*ExplorerAutoOptimizerSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee4d755cec81c290, []int{0}
}

func (m *ExplorerAutoOptimizerSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExplorerAutoOptimizerSetting.Unmarshal(m, b)
}
func (m *ExplorerAutoOptimizerSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExplorerAutoOptimizerSetting.Marshal(b, m, deterministic)
}
func (m *ExplorerAutoOptimizerSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExplorerAutoOptimizerSetting.Merge(m, src)
}
func (m *ExplorerAutoOptimizerSetting) XXX_Size() int {
	return xxx_messageInfo_ExplorerAutoOptimizerSetting.Size(m)
}
func (m *ExplorerAutoOptimizerSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ExplorerAutoOptimizerSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ExplorerAutoOptimizerSetting proto.InternalMessageInfo

func (m *ExplorerAutoOptimizerSetting) GetOptIn() *wrappers.BoolValue {
	if m != nil {
		return m.OptIn
	}
	return nil
}

func init() {
	proto.RegisterType((*ExplorerAutoOptimizerSetting)(nil), "google.ads.googleads.v1.common.ExplorerAutoOptimizerSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/explorer_auto_optimizer_setting.proto", fileDescriptor_ee4d755cec81c290)
}

var fileDescriptor_ee4d755cec81c290 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x95, 0x22, 0x3a, 0x84, 0xad, 0x13, 0xaa, 0xaa, 0x0a, 0x3a, 0x31, 0x3d, 0xcb, 0xb0,
	0x99, 0xc9, 0x05, 0x54, 0x31, 0x51, 0x40, 0xca, 0x80, 0x22, 0x45, 0x6e, 0x63, 0x2c, 0x4b, 0x89,
	0x9f, 0x65, 0x3b, 0x05, 0xf1, 0x39, 0x8c, 0x7c, 0x0a, 0x9f, 0xc2, 0x37, 0x30, 0xa0, 0xc6, 0x49,
	0x36, 0x3a, 0xf9, 0xda, 0xbe, 0xef, 0xbe, 0xa3, 0x9b, 0xde, 0x2a, 0x44, 0x55, 0x49, 0x22, 0x4a,
	0x4f, 0xa2, 0xdc, 0xab, 0x1d, 0x25, 0x5b, 0xac, 0x6b, 0x34, 0x44, 0xbe, 0xdb, 0x0a, 0x9d, 0x74,
	0x85, 0x68, 0x02, 0x16, 0x68, 0x83, 0xae, 0xf5, 0x87, 0x74, 0x85, 0x97, 0x21, 0x68, 0xa3, 0xc0,
	0x3a, 0x0c, 0x38, 0x99, 0xc7, 0x51, 0x10, 0xa5, 0x87, 0x21, 0x05, 0x76, 0x14, 0x62, 0xca, 0x74,
	0xd6, 0x6f, 0xb1, 0x9a, 0x08, 0x63, 0x30, 0x88, 0xa0, 0xd1, 0xf8, 0x38, 0x3d, 0xed, 0xa6, 0x49,
	0x7b, 0xdb, 0x34, 0xaf, 0xe4, 0xcd, 0x09, 0x6b, 0xa5, 0xeb, 0xfe, 0x17, 0x8f, 0xe9, 0xec, 0xae,
	0xc3, 0xe0, 0x4d, 0xc0, 0x87, 0x1e, 0xe2, 0x39, 0x32, 0x4c, 0x68, 0x3a, 0x46, 0x1b, 0x0a, 0x6d,
	0x4e, 0x93, 0xb3, 0xe4, 0xe2, 0xe4, 0x72, 0xda, 0x31, 0x40, 0x1f, 0x08, 0x4b, 0xc4, 0x2a, 0x13,
	0x55, 0x23, 0x9f, 0x8e, 0xd1, 0x86, 0x7b, 0xb3, 0xfc, 0x4d, 0xd2, 0xc5, 0x16, 0x6b, 0x38, 0xcc,
	0xbd, 0x3c, 0x3f, 0xb4, 0x77, 0xbd, 0x4f, 0x5f, 0x27, 0x2f, 0x5d, 0x85, 0xa0, 0xb0, 0x12, 0x46,
	0x01, 0x3a, 0x45, 0x94, 0x34, 0xed, 0xee, 0xbe, 0x52, 0xab, 0xfd, 0x7f, 0x0d, 0x5f, 0xc7, 0xe3,
	0x73, 0x74, 0xb4, 0xe2, 0xfc, 0x6b, 0x34, 0x5f, 0xc5, 0x30, 0x5e, 0x7a, 0x88, 0x72, 0xaf, 0x32,
	0x0a, 0x37, 0xad, 0xed, 0xbb, 0x37, 0xe4, 0xbc, 0xf4, 0xf9, 0x60, 0xc8, 0x33, 0x9a, 0x47, 0xc3,
	0xcf, 0x68, 0x11, 0x5f, 0x19, 0xe3, 0xa5, 0x67, 0x6c, 0xb0, 0x30, 0x96, 0x51, 0xc6, 0xa2, 0x69,
	0x33, 0x6e, 0xe9, 0xae, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x71, 0x0c, 0x35, 0xfe, 0x01,
	0x00, 0x00,
}
