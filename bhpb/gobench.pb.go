// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gobench.proto

package bhpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// TODO: single server, single command and single report to make things easier for now
type GoBenchmarkSpec struct {
	Server               *ServerTarget           `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Command              *GoBenchmarkCommandSpec `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Report               *GoBenchmarkReportSpec  `protobuf:"bytes,3,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GoBenchmarkSpec) Reset()         { *m = GoBenchmarkSpec{} }
func (m *GoBenchmarkSpec) String() string { return proto.CompactTextString(m) }
func (*GoBenchmarkSpec) ProtoMessage()    {}
func (*GoBenchmarkSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a56bd951d162f, []int{0}
}
func (m *GoBenchmarkSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoBenchmarkSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoBenchmarkSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoBenchmarkSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoBenchmarkSpec.Merge(m, src)
}
func (m *GoBenchmarkSpec) XXX_Size() int {
	return m.Size()
}
func (m *GoBenchmarkSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GoBenchmarkSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GoBenchmarkSpec proto.InternalMessageInfo

type GoBenchmarkCommandSpec struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Output               string   `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoBenchmarkCommandSpec) Reset()         { *m = GoBenchmarkCommandSpec{} }
func (m *GoBenchmarkCommandSpec) String() string { return proto.CompactTextString(m) }
func (*GoBenchmarkCommandSpec) ProtoMessage()    {}
func (*GoBenchmarkCommandSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a56bd951d162f, []int{1}
}
func (m *GoBenchmarkCommandSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoBenchmarkCommandSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoBenchmarkCommandSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoBenchmarkCommandSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoBenchmarkCommandSpec.Merge(m, src)
}
func (m *GoBenchmarkCommandSpec) XXX_Size() int {
	return m.Size()
}
func (m *GoBenchmarkCommandSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GoBenchmarkCommandSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GoBenchmarkCommandSpec proto.InternalMessageInfo

type GoBenchmarkReportSpec struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoBenchmarkReportSpec) Reset()         { *m = GoBenchmarkReportSpec{} }
func (m *GoBenchmarkReportSpec) String() string { return proto.CompactTextString(m) }
func (*GoBenchmarkReportSpec) ProtoMessage()    {}
func (*GoBenchmarkReportSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a56bd951d162f, []int{2}
}
func (m *GoBenchmarkReportSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoBenchmarkReportSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoBenchmarkReportSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoBenchmarkReportSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoBenchmarkReportSpec.Merge(m, src)
}
func (m *GoBenchmarkReportSpec) XXX_Size() int {
	return m.Size()
}
func (m *GoBenchmarkReportSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GoBenchmarkReportSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GoBenchmarkReportSpec proto.InternalMessageInfo

type GoBenchmarkResult struct {
	PackageId            int64    `protobuf:"varint,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	PackageName          string   `protobuf:"bytes,2,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	CaseId               int64    `protobuf:"varint,3,opt,name=case_id,json=caseId,proto3" json:"case_id,omitempty"`
	CaseName             string   `protobuf:"bytes,4,opt,name=case_name,json=caseName,proto3" json:"case_name,omitempty"`
	Duration             int64    `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	NsPerOp              int64    `protobuf:"varint,6,opt,name=ns_per_op,json=nsPerOp,proto3" json:"ns_per_op,omitempty"`
	AllocPerOp           int64    `protobuf:"varint,7,opt,name=alloc_per_op,json=allocPerOp,proto3" json:"alloc_per_op,omitempty"`
	BytesAllocatedPerOp  int64    `protobuf:"varint,8,opt,name=bytes_allocated_per_op,json=bytesAllocatedPerOp,proto3" json:"bytes_allocated_per_op,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoBenchmarkResult) Reset()         { *m = GoBenchmarkResult{} }
func (m *GoBenchmarkResult) String() string { return proto.CompactTextString(m) }
func (*GoBenchmarkResult) ProtoMessage()    {}
func (*GoBenchmarkResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a56bd951d162f, []int{3}
}
func (m *GoBenchmarkResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoBenchmarkResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoBenchmarkResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoBenchmarkResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoBenchmarkResult.Merge(m, src)
}
func (m *GoBenchmarkResult) XXX_Size() int {
	return m.Size()
}
func (m *GoBenchmarkResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GoBenchmarkResult.DiscardUnknown(m)
}

var xxx_messageInfo_GoBenchmarkResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GoBenchmarkSpec)(nil), "bhpb.GoBenchmarkSpec")
	proto.RegisterType((*GoBenchmarkCommandSpec)(nil), "bhpb.GoBenchmarkCommandSpec")
	proto.RegisterType((*GoBenchmarkReportSpec)(nil), "bhpb.GoBenchmarkReportSpec")
	proto.RegisterType((*GoBenchmarkResult)(nil), "bhpb.GoBenchmarkResult")
}

func init() { proto.RegisterFile("gobench.proto", fileDescriptor_a18a56bd951d162f) }

var fileDescriptor_a18a56bd951d162f = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0xc9, 0x74, 0x26, 0x69, 0xde, 0x14, 0x21, 0x0c, 0x84, 0x28, 0x03, 0xd1, 0x90, 0x15,
	0x42, 0x9a, 0x54, 0xa2, 0x12, 0x7b, 0xca, 0x02, 0x95, 0x05, 0xa0, 0x94, 0x15, 0x9b, 0xc8, 0x49,
	0x4c, 0x1a, 0xb5, 0xb1, 0x2d, 0xc7, 0x41, 0xe2, 0x2a, 0x1c, 0x81, 0x93, 0x74, 0xc9, 0x11, 0xa0,
	0x27, 0x41, 0x7e, 0x4e, 0x28, 0x68, 0xba, 0xf3, 0xff, 0xfe, 0xff, 0xfb, 0xdf, 0x5b, 0x18, 0xee,
	0xd6, 0xa2, 0x60, 0xbc, 0xdc, 0xa4, 0x52, 0x09, 0x2d, 0xc8, 0x79, 0xb1, 0x91, 0x45, 0x74, 0x53,
	0x37, 0x7a, 0xd3, 0x17, 0x69, 0x29, 0xda, 0x79, 0x2d, 0x6a, 0x31, 0x47, 0xb3, 0xe8, 0xbf, 0xa0,
	0x42, 0x81, 0x2f, 0x0b, 0x45, 0xb3, 0x52, 0xb4, 0xad, 0xe0, 0x56, 0x25, 0x3f, 0x1c, 0xb8, 0xf7,
	0x56, 0x2c, 0x4d, 0x69, 0x4b, 0xd5, 0x76, 0x2d, 0x59, 0x49, 0x5e, 0x80, 0xdb, 0x31, 0xf5, 0x95,
	0xa9, 0xd0, 0xb9, 0x76, 0x9e, 0x5f, 0xbe, 0x24, 0xa9, 0xd9, 0x93, 0xae, 0x71, 0xf6, 0x89, 0xaa,
	0x9a, 0xe9, 0x6c, 0x48, 0x90, 0x57, 0xe0, 0x99, 0x3e, 0xca, 0xab, 0xf0, 0x0c, 0xc3, 0x4f, 0x6c,
	0xf8, 0x9f, 0xce, 0x37, 0xd6, 0x37, 0xd5, 0xd9, 0x18, 0x26, 0x0b, 0x70, 0x15, 0x93, 0x42, 0xe9,
	0x70, 0x82, 0xd8, 0xd5, 0x2d, 0x2c, 0x43, 0x1b, 0xa9, 0x21, 0x9a, 0xbc, 0x83, 0xe0, 0x74, 0x2f,
	0x09, 0x8f, 0x67, 0x98, 0x9b, 0xfd, 0xe3, 0xa2, 0x00, 0x5c, 0xd1, 0x6b, 0xd9, 0x6b, 0xbc, 0xcf,
	0xcf, 0x06, 0x95, 0xdc, 0xc0, 0xa3, 0x93, 0xcb, 0xc8, 0x43, 0xb8, 0x68, 0xb8, 0xc9, 0xdb, 0x22,
	0x2b, 0x92, 0xef, 0x67, 0x70, 0xff, 0xbf, 0x7c, 0xd7, 0xef, 0x34, 0x79, 0x0a, 0x20, 0x69, 0xb9,
	0xa5, 0x35, 0xcb, 0x1b, 0xbb, 0x79, 0x92, 0xf9, 0xc3, 0x64, 0x55, 0x91, 0x67, 0x30, 0x1b, 0x6d,
	0x4e, 0x5b, 0x36, 0x5c, 0x70, 0x39, 0xcc, 0xde, 0xd3, 0x96, 0x91, 0xc7, 0xe0, 0x95, 0xb4, 0x43,
	0x7c, 0x82, 0xb8, 0x6b, 0xe4, 0xaa, 0x22, 0x57, 0xe0, 0xa3, 0x81, 0xe0, 0x39, 0x82, 0x53, 0x33,
	0x40, 0x2a, 0x82, 0x69, 0xd5, 0x2b, 0xaa, 0x1b, 0xc1, 0xc3, 0x0b, 0xc4, 0xfe, 0x6a, 0x12, 0x81,
	0xcf, 0xbb, 0x5c, 0x32, 0x95, 0x0b, 0x19, 0xba, 0x68, 0x7a, 0xbc, 0xfb, 0xc8, 0xd4, 0x07, 0x49,
	0xae, 0x61, 0x46, 0x77, 0x3b, 0x51, 0x8e, 0xb6, 0x87, 0x36, 0xe0, 0xcc, 0x26, 0x16, 0x10, 0x14,
	0xdf, 0x34, 0xeb, 0x72, 0x9c, 0x51, 0xcd, 0xaa, 0x31, 0x3b, 0xc5, 0xec, 0x03, 0x74, 0x5f, 0x8f,
	0x26, 0x42, 0xcb, 0x60, 0xff, 0x3b, 0xbe, 0xb3, 0x3f, 0xc4, 0xce, 0xcf, 0x43, 0xec, 0xfc, 0x3a,
	0xc4, 0xce, 0x67, 0xfc, 0x99, 0x85, 0x8b, 0x7f, 0x6c, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf6,
	0x90, 0x03, 0x28, 0xb7, 0x02, 0x00, 0x00,
}

func (m *GoBenchmarkSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoBenchmarkSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoBenchmarkSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Report != nil {
		{
			size, err := m.Report.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGobench(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Command != nil {
		{
			size, err := m.Command.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGobench(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Server != nil {
		{
			size, err := m.Server.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGobench(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GoBenchmarkCommandSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoBenchmarkCommandSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoBenchmarkCommandSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Output) > 0 {
		i -= len(m.Output)
		copy(dAtA[i:], m.Output)
		i = encodeVarintGobench(dAtA, i, uint64(len(m.Output)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Command) > 0 {
		i -= len(m.Command)
		copy(dAtA[i:], m.Command)
		i = encodeVarintGobench(dAtA, i, uint64(len(m.Command)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GoBenchmarkReportSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoBenchmarkReportSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoBenchmarkReportSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Input) > 0 {
		i -= len(m.Input)
		copy(dAtA[i:], m.Input)
		i = encodeVarintGobench(dAtA, i, uint64(len(m.Input)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GoBenchmarkResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoBenchmarkResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoBenchmarkResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.BytesAllocatedPerOp != 0 {
		i = encodeVarintGobench(dAtA, i, uint64(m.BytesAllocatedPerOp))
		i--
		dAtA[i] = 0x40
	}
	if m.AllocPerOp != 0 {
		i = encodeVarintGobench(dAtA, i, uint64(m.AllocPerOp))
		i--
		dAtA[i] = 0x38
	}
	if m.NsPerOp != 0 {
		i = encodeVarintGobench(dAtA, i, uint64(m.NsPerOp))
		i--
		dAtA[i] = 0x30
	}
	if m.Duration != 0 {
		i = encodeVarintGobench(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x28
	}
	if len(m.CaseName) > 0 {
		i -= len(m.CaseName)
		copy(dAtA[i:], m.CaseName)
		i = encodeVarintGobench(dAtA, i, uint64(len(m.CaseName)))
		i--
		dAtA[i] = 0x22
	}
	if m.CaseId != 0 {
		i = encodeVarintGobench(dAtA, i, uint64(m.CaseId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PackageName) > 0 {
		i -= len(m.PackageName)
		copy(dAtA[i:], m.PackageName)
		i = encodeVarintGobench(dAtA, i, uint64(len(m.PackageName)))
		i--
		dAtA[i] = 0x12
	}
	if m.PackageId != 0 {
		i = encodeVarintGobench(dAtA, i, uint64(m.PackageId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGobench(dAtA []byte, offset int, v uint64) int {
	offset -= sovGobench(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GoBenchmarkSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Server != nil {
		l = m.Server.Size()
		n += 1 + l + sovGobench(uint64(l))
	}
	if m.Command != nil {
		l = m.Command.Size()
		n += 1 + l + sovGobench(uint64(l))
	}
	if m.Report != nil {
		l = m.Report.Size()
		n += 1 + l + sovGobench(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoBenchmarkCommandSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovGobench(uint64(l))
	}
	l = len(m.Output)
	if l > 0 {
		n += 1 + l + sovGobench(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoBenchmarkReportSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Input)
	if l > 0 {
		n += 1 + l + sovGobench(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoBenchmarkResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PackageId != 0 {
		n += 1 + sovGobench(uint64(m.PackageId))
	}
	l = len(m.PackageName)
	if l > 0 {
		n += 1 + l + sovGobench(uint64(l))
	}
	if m.CaseId != 0 {
		n += 1 + sovGobench(uint64(m.CaseId))
	}
	l = len(m.CaseName)
	if l > 0 {
		n += 1 + l + sovGobench(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovGobench(uint64(m.Duration))
	}
	if m.NsPerOp != 0 {
		n += 1 + sovGobench(uint64(m.NsPerOp))
	}
	if m.AllocPerOp != 0 {
		n += 1 + sovGobench(uint64(m.AllocPerOp))
	}
	if m.BytesAllocatedPerOp != 0 {
		n += 1 + sovGobench(uint64(m.BytesAllocatedPerOp))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGobench(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGobench(x uint64) (n int) {
	return sovGobench(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoBenchmarkSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGobench
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
			return fmt.Errorf("proto: GoBenchmarkSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoBenchmarkSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
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
				return ErrInvalidLengthGobench
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGobench
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Server == nil {
				m.Server = &ServerTarget{}
			}
			if err := m.Server.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
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
				return ErrInvalidLengthGobench
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGobench
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Command == nil {
				m.Command = &GoBenchmarkCommandSpec{}
			}
			if err := m.Command.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Report", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
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
				return ErrInvalidLengthGobench
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGobench
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Report == nil {
				m.Report = &GoBenchmarkReportSpec{}
			}
			if err := m.Report.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGobench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGobench
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGobench
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
func (m *GoBenchmarkCommandSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGobench
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
			return fmt.Errorf("proto: GoBenchmarkCommandSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoBenchmarkCommandSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGobench
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGobench
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGobench
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGobench
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Output = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGobench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGobench
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGobench
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
func (m *GoBenchmarkReportSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGobench
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
			return fmt.Errorf("proto: GoBenchmarkReportSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoBenchmarkReportSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGobench
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGobench
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Input = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGobench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGobench
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGobench
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
func (m *GoBenchmarkResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGobench
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
			return fmt.Errorf("proto: GoBenchmarkResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoBenchmarkResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackageId", wireType)
			}
			m.PackageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PackageId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGobench
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGobench
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PackageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaseId", wireType)
			}
			m.CaseId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CaseId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaseName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGobench
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGobench
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CaseName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NsPerOp", wireType)
			}
			m.NsPerOp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NsPerOp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocPerOp", wireType)
			}
			m.AllocPerOp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllocPerOp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesAllocatedPerOp", wireType)
			}
			m.BytesAllocatedPerOp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BytesAllocatedPerOp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGobench(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGobench
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGobench
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
func skipGobench(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGobench
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
					return 0, ErrIntOverflowGobench
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGobench
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
				return 0, ErrInvalidLengthGobench
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGobench
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGobench
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGobench        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGobench          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGobench = fmt.Errorf("proto: unexpected end of group")
)
