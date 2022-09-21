// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: getambassador.io/v2/Host_nojson.proto

package getambassador_io_v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

type Host struct {
	*v1.TypeMeta `protobuf:"bytes,1,opt,name=type_meta,json=typeMeta,proto3,embedded=type_meta" json:",omitempty"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	*v1.ObjectMeta `protobuf:"bytes,2,opt,name=metadata,proto3,embedded=metadata" json:"metadata,omitempty"`
	// Specification of the desired behavior of the Host.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *HostSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	// Most recently observed status of the Host.
	// This data may not be up to date.
	// Populated by AES.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status               *HostStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_c48ae35a58af4346, []int{0}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Host.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return m.Size()
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetSpec() *HostSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Host) GetStatus() *HostStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Host)(nil), "getambassador.io.v2.Host")
}

func init() {
	proto.RegisterFile("getambassador.io/v2/Host_nojson.proto", fileDescriptor_c48ae35a58af4346)
}

var fileDescriptor_c48ae35a58af4346 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0x3b, 0x31,
	0x10, 0xc6, 0xd9, 0xfe, 0x97, 0xd2, 0x7f, 0xf4, 0x14, 0x2f, 0x4b, 0xc1, 0xad, 0x08, 0x82, 0x07,
	0x9d, 0xd8, 0x2a, 0xe8, 0xb9, 0x27, 0x2f, 0x22, 0xac, 0xde, 0x4b, 0x76, 0x3b, 0xa6, 0xdb, 0x92,
	0x4d, 0xd8, 0x4c, 0x17, 0xf6, 0x89, 0x7c, 0x15, 0x8f, 0x7d, 0x82, 0x22, 0x3d, 0xfa, 0x14, 0x92,
	0xb4, 0x88, 0x88, 0x05, 0x6f, 0x99, 0x99, 0xdf, 0xf7, 0xcd, 0x97, 0x61, 0x67, 0x0a, 0x49, 0xea,
	0x5c, 0x3a, 0x27, 0xa7, 0xa6, 0x86, 0xd2, 0x88, 0x66, 0x24, 0xee, 0x8d, 0xa3, 0x49, 0x65, 0xe6,
	0xce, 0x54, 0x60, 0x6b, 0x43, 0x86, 0x1f, 0xfd, 0xc4, 0xa0, 0x19, 0xf5, 0x6f, 0x16, 0x77, 0xce,
	0x2b, 0xa4, 0x2d, 0xb5, 0x2c, 0x66, 0x65, 0x85, 0x75, 0x2b, 0xec, 0x42, 0xf9, 0x86, 0x13, 0x1a,
	0x49, 0x8a, 0x66, 0x28, 0x14, 0x56, 0x58, 0x4b, 0xc2, 0xe9, 0xd6, 0xaa, 0x7f, 0xa9, 0x4a, 0x9a,
	0x2d, 0x73, 0x28, 0x8c, 0x16, 0xca, 0x28, 0x23, 0x42, 0x3b, 0x5f, 0xbe, 0x84, 0x2a, 0x14, 0xe1,
	0xb5, 0xc3, 0xd3, 0x7d, 0x01, 0xb7, 0xf3, 0xd3, 0xd7, 0x0e, 0x8b, 0x7d, 0xc9, 0x27, 0xec, 0x3f,
	0xb5, 0x16, 0x27, 0x7e, 0x6f, 0x12, 0x9d, 0x44, 0xe7, 0x07, 0x23, 0x80, 0x6d, 0x42, 0xf8, 0x9e,
	0x10, 0xec, 0x42, 0xf9, 0x86, 0x03, 0x4f, 0x42, 0x33, 0x84, 0xe7, 0xd6, 0xe2, 0x03, 0x92, 0x1c,
	0xf3, 0xd5, 0x7a, 0x10, 0x7d, 0xac, 0x07, 0xec, 0xc2, 0xe8, 0x92, 0x50, 0x5b, 0x6a, 0xb3, 0x1e,
	0xed, 0xa6, 0x3c, 0x63, 0x3d, 0xaf, 0x98, 0x4a, 0x92, 0x49, 0x27, 0xf8, 0x5f, 0xfd, 0xcd, 0xff,
	0x31, 0x9f, 0x63, 0x41, 0x61, 0x43, 0xec, 0x37, 0x64, 0x5f, 0x3e, 0x7c, 0xc8, 0x62, 0x67, 0xb1,
	0x48, 0xfe, 0x05, 0xbf, 0x63, 0xf8, 0xe5, 0xcc, 0xe0, 0x7f, 0xf7, 0x64, 0xb1, 0xc8, 0x02, 0xca,
	0x6f, 0x59, 0xd7, 0x91, 0xa4, 0xa5, 0x4b, 0xe2, 0x20, 0x1a, 0xec, 0x17, 0x05, 0x2c, 0xdb, 0xe1,
	0xe3, 0xc3, 0xb7, 0x4d, 0x1a, 0xad, 0x36, 0x69, 0xf4, 0xbe, 0x49, 0xa3, 0xbc, 0x1b, 0xce, 0x77,
	0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0xef, 0x2d, 0xdb, 0xa2, 0x01, 0x02, 0x00, 0x00,
}

func (m *Host) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Host) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Host) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != nil {
		{
			size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostNojson(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Spec != nil {
		{
			size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostNojson(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ObjectMeta != nil {
		{
			size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostNojson(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.TypeMeta != nil {
		{
			size, err := m.TypeMeta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostNojson(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHostNojson(dAtA []byte, offset int, v uint64) int {
	offset -= sovHostNojson(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Host) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TypeMeta != nil {
		l = m.TypeMeta.Size()
		n += 1 + l + sovHostNojson(uint64(l))
	}
	if m.ObjectMeta != nil {
		l = m.ObjectMeta.Size()
		n += 1 + l + sovHostNojson(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovHostNojson(uint64(l))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovHostNojson(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHostNojson(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHostNojson(x uint64) (n int) {
	return sovHostNojson(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Host) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHostNojson
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
			return fmt.Errorf("proto: Host: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Host: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostNojson
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
				return ErrInvalidLengthHostNojson
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostNojson
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TypeMeta == nil {
				m.TypeMeta = &v1.TypeMeta{}
			}
			if err := m.TypeMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostNojson
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
				return ErrInvalidLengthHostNojson
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostNojson
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ObjectMeta == nil {
				m.ObjectMeta = &v1.ObjectMeta{}
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostNojson
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
				return ErrInvalidLengthHostNojson
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostNojson
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &HostSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostNojson
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
				return ErrInvalidLengthHostNojson
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostNojson
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &HostStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHostNojson(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHostNojson
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHostNojson
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
func skipHostNojson(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHostNojson
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
					return 0, ErrIntOverflowHostNojson
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
					return 0, ErrIntOverflowHostNojson
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
				return 0, ErrInvalidLengthHostNojson
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHostNojson
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHostNojson
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
				next, err := skipHostNojson(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHostNojson
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
	ErrInvalidLengthHostNojson = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHostNojson   = fmt.Errorf("proto: integer overflow")
)