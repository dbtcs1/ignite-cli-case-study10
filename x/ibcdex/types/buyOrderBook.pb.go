// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibcdex/buyOrderBook.proto

package types

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

type BuyOrderBook struct {
	Creator     string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index       string     `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	AmountDenom string     `protobuf:"bytes,3,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	PriceDenom  string     `protobuf:"bytes,4,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	Book        *OrderBook `protobuf:"bytes,5,opt,name=book,proto3" json:"book,omitempty"`
}

func (m *BuyOrderBook) Reset()         { *m = BuyOrderBook{} }
func (m *BuyOrderBook) String() string { return proto.CompactTextString(m) }
func (*BuyOrderBook) ProtoMessage()    {}
func (*BuyOrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5716e4857f17ae7, []int{0}
}
func (m *BuyOrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BuyOrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BuyOrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BuyOrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyOrderBook.Merge(m, src)
}
func (m *BuyOrderBook) XXX_Size() int {
	return m.Size()
}
func (m *BuyOrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyOrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_BuyOrderBook proto.InternalMessageInfo

func (m *BuyOrderBook) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *BuyOrderBook) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *BuyOrderBook) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *BuyOrderBook) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *BuyOrderBook) GetBook() *OrderBook {
	if m != nil {
		return m.Book
	}
	return nil
}

func init() {
	proto.RegisterType((*BuyOrderBook)(nil), "tendermint.interchange.ibcdex.BuyOrderBook")
}

func init() { proto.RegisterFile("ibcdex/buyOrderBook.proto", fileDescriptor_c5716e4857f17ae7) }

var fileDescriptor_c5716e4857f17ae7 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x68, 0x41, 0xb8, 0x4c, 0x56, 0x87, 0x50, 0x09, 0x2b, 0x62, 0xca, 0x64, 0x23,
	0x58, 0x99, 0x22, 0x26, 0x16, 0xa4, 0x8e, 0x6c, 0xb1, 0x73, 0x4a, 0xad, 0x2a, 0xbe, 0xc8, 0x38,
	0x52, 0xfa, 0x16, 0x3c, 0x11, 0x33, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08, 0xaa, 0x5d, 0x20, 0x13,
	0x9b, 0xef, 0xf3, 0xff, 0x7f, 0xb2, 0x8f, 0x5e, 0x19, 0xa5, 0x2b, 0xe8, 0xa5, 0xea, 0x76, 0xcf,
	0xae, 0x02, 0x57, 0x20, 0x6e, 0x45, 0xeb, 0xd0, 0x23, 0xbb, 0xf6, 0x60, 0x2b, 0x70, 0x8d, 0xb1,
	0x5e, 0x18, 0xeb, 0xc1, 0xe9, 0x4d, 0x69, 0x6b, 0x10, 0xb1, 0xb1, 0x5a, 0xd6, 0x58, 0x63, 0x48,
	0xca, 0xc3, 0x29, 0x96, 0x56, 0xec, 0xe8, 0xc3, 0x83, 0x2c, 0xb2, 0x9b, 0x77, 0x42, 0x2f, 0x8b,
	0x89, 0x9f, 0xa5, 0xf4, 0x5c, 0x3b, 0x28, 0x3d, 0xba, 0x94, 0x64, 0x24, 0xbf, 0x58, 0xff, 0x8c,
	0x6c, 0x49, 0xe7, 0xc6, 0x56, 0xd0, 0xa7, 0x27, 0x81, 0xc7, 0x81, 0x65, 0x74, 0x51, 0x36, 0xd8,
	0x59, 0xff, 0x08, 0x16, 0x9b, 0xf4, 0x34, 0xdc, 0x4d, 0x11, 0xe3, 0x94, 0xb6, 0xce, 0x68, 0x88,
	0x81, 0x59, 0x08, 0x4c, 0x08, 0x7b, 0xa0, 0x33, 0x85, 0xb8, 0x4d, 0xe7, 0x19, 0xc9, 0x17, 0x77,
	0xb9, 0xf8, 0xf7, 0x6b, 0xe2, 0xf7, 0xa5, 0xeb, 0xd0, 0x2a, 0x9e, 0x3e, 0x06, 0x4e, 0xf6, 0x03,
	0x27, 0x5f, 0x03, 0x27, 0x6f, 0x23, 0x4f, 0xf6, 0x23, 0x4f, 0x3e, 0x47, 0x9e, 0xbc, 0xdc, 0xd6,
	0xc6, 0x6f, 0x3a, 0x25, 0x34, 0x36, 0xf2, 0xcf, 0x29, 0x27, 0x4e, 0xd9, 0xcb, 0xe3, 0x4a, 0xfc,
	0xae, 0x85, 0x57, 0x75, 0x16, 0x76, 0x72, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x73, 0x1b, 0x7e,
	0x21, 0x79, 0x01, 0x00, 0x00,
}

func (m *BuyOrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BuyOrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuyOrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Book != nil {
		{
			size, err := m.Book.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBuyOrderBook(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBuyOrderBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovBuyOrderBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BuyOrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	if m.Book != nil {
		l = m.Book.Size()
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	return n
}

func sovBuyOrderBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBuyOrderBook(x uint64) (n int) {
	return sovBuyOrderBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BuyOrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuyOrderBook
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
			return fmt.Errorf("proto: BuyOrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuyOrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Book", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Book == nil {
				m.Book = &OrderBook{}
			}
			if err := m.Book.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuyOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBuyOrderBook
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
func skipBuyOrderBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuyOrderBook
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
					return 0, ErrIntOverflowBuyOrderBook
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
					return 0, ErrIntOverflowBuyOrderBook
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
				return 0, ErrInvalidLengthBuyOrderBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBuyOrderBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBuyOrderBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBuyOrderBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuyOrderBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBuyOrderBook = fmt.Errorf("proto: unexpected end of group")
)