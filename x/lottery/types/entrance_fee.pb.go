// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lottery/entrance_fee.proto

package types

import (
	fmt "fmt"
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

type EntranceFee struct {
}

func (m *EntranceFee) Reset()         { *m = EntranceFee{} }
func (m *EntranceFee) String() string { return proto.CompactTextString(m) }
func (*EntranceFee) ProtoMessage()    {}
func (*EntranceFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_db5c062f03d5f14d, []int{0}
}
func (m *EntranceFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntranceFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntranceFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntranceFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntranceFee.Merge(m, src)
}
func (m *EntranceFee) XXX_Size() int {
	return m.Size()
}
func (m *EntranceFee) XXX_DiscardUnknown() {
	xxx_messageInfo_EntranceFee.DiscardUnknown(m)
}

var xxx_messageInfo_EntranceFee proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EntranceFee)(nil), "lotterychainnel.lottery.EntranceFee")
}

func init() { proto.RegisterFile("lottery/entrance_fee.proto", fileDescriptor_db5c062f03d5f14d) }

var fileDescriptor_db5c062f03d5f14d = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x4f, 0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0x4b, 0x4e, 0x8d, 0x4f, 0x4b, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0xca, 0x25, 0x67, 0x24, 0x66, 0xe6, 0xe5,
	0xa5, 0xe6, 0xe8, 0x41, 0xf9, 0x4a, 0xbc, 0x5c, 0xdc, 0xae, 0x50, 0xe5, 0x6e, 0xa9, 0xa9, 0x4e,
	0xd6, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x08, 0xd5, 0xa1, 0x0b,
	0x36, 0x42, 0x37, 0x2f, 0x35, 0x47, 0xbf, 0x42, 0x1f, 0x66, 0x63, 0x49, 0x65, 0x41, 0x6a, 0x71,
	0x12, 0x1b, 0xd8, 0x2e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x1f, 0xa3, 0x74, 0x89,
	0x00, 0x00, 0x00,
}

func (m *EntranceFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntranceFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntranceFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintEntranceFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovEntranceFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EntranceFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovEntranceFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntranceFee(x uint64) (n int) {
	return sovEntranceFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EntranceFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntranceFee
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
			return fmt.Errorf("proto: EntranceFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntranceFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEntranceFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntranceFee
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEntranceFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntranceFee
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
					return 0, ErrIntOverflowEntranceFee
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
					return 0, ErrIntOverflowEntranceFee
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
				return 0, ErrInvalidLengthEntranceFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEntranceFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEntranceFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEntranceFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntranceFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntranceFee = fmt.Errorf("proto: unexpected end of group")
)