// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/tokenconverter/v1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	NonTransferableDenom string   `protobuf:"bytes,1,opt,name=non_transferable_denom,json=nonTransferableDenom,proto3" json:"non_transferable_denom,omitempty"`
	TransferableDenom    string   `protobuf:"bytes,2,opt,name=transferable_denom,json=transferableDenom,proto3" json:"transferable_denom,omitempty"`
	AllowedAddresses     []string `protobuf:"bytes,3,rep,name=allowed_addresses,json=allowedAddresses,proto3" json:"allowed_addresses,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e32774d966b58991, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetNonTransferableDenom() string {
	if m != nil {
		return m.NonTransferableDenom
	}
	return ""
}

func (m *Params) GetTransferableDenom() string {
	if m != nil {
		return m.TransferableDenom
	}
	return ""
}

func (m *Params) GetAllowedAddresses() []string {
	if m != nil {
		return m.AllowedAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "sunrise.tokenconverter.v1.Params")
}

func init() {
	proto.RegisterFile("sunrise/tokenconverter/v1/params.proto", fileDescriptor_e32774d966b58991)
}

var fileDescriptor_e32774d966b58991 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x2e, 0xcd, 0x2b,
	0xca, 0x2c, 0x4e, 0xd5, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0x4b, 0xce, 0xcf, 0x2b, 0x4b, 0x2d, 0x2a,
	0x49, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x84, 0xaa, 0xd3, 0x43, 0x55, 0xa7, 0x57, 0x66, 0x28, 0x25, 0x92, 0x9e,
	0x9f, 0x9e, 0x0f, 0x56, 0xa5, 0x0f, 0x62, 0x41, 0x34, 0x28, 0x2d, 0x60, 0xe4, 0x62, 0x0b, 0x00,
	0x9b, 0x20, 0x64, 0xc2, 0x25, 0x96, 0x97, 0x9f, 0x17, 0x5f, 0x52, 0x94, 0x98, 0x57, 0x9c, 0x96,
	0x5a, 0x94, 0x98, 0x94, 0x93, 0x1a, 0x9f, 0x92, 0x9a, 0x97, 0x9f, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0x24, 0x92, 0x97, 0x9f, 0x17, 0x82, 0x24, 0xe9, 0x02, 0x92, 0x13, 0xd2, 0xe5, 0x12,
	0xc2, 0xa2, 0x83, 0x09, 0xac, 0x43, 0xb0, 0x04, 0x43, 0xb9, 0x36, 0x97, 0x60, 0x62, 0x4e, 0x4e,
	0x7e, 0x79, 0x6a, 0x4a, 0x7c, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0x71, 0x6a, 0xb1, 0x04, 0xb3,
	0x02, 0xb3, 0x06, 0x67, 0x90, 0x00, 0x54, 0xc2, 0x11, 0x26, 0x6e, 0xc5, 0xf2, 0x62, 0x81, 0x3c,
	0xa3, 0x53, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0xa5, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x3d, 0x9e, 0x93, 0x58, 0x99, 0x5a,
	0x04, 0xe3, 0xe8, 0x57, 0xa0, 0x87, 0x57, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xef,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xed, 0x66, 0x16, 0x56, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.NonTransferableDenom != that1.NonTransferableDenom {
		return false
	}
	if this.TransferableDenom != that1.TransferableDenom {
		return false
	}
	if len(this.AllowedAddresses) != len(that1.AllowedAddresses) {
		return false
	}
	for i := range this.AllowedAddresses {
		if this.AllowedAddresses[i] != that1.AllowedAddresses[i] {
			return false
		}
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedAddresses) > 0 {
		for iNdEx := len(m.AllowedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedAddresses[iNdEx])
			copy(dAtA[i:], m.AllowedAddresses[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AllowedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TransferableDenom) > 0 {
		i -= len(m.TransferableDenom)
		copy(dAtA[i:], m.TransferableDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TransferableDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NonTransferableDenom) > 0 {
		i -= len(m.NonTransferableDenom)
		copy(dAtA[i:], m.NonTransferableDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.NonTransferableDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NonTransferableDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TransferableDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.AllowedAddresses) > 0 {
		for _, s := range m.AllowedAddresses {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NonTransferableDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NonTransferableDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferableDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferableDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedAddresses = append(m.AllowedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
