// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eeta/bid/auction.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type Auction struct {
	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BillboardId   uint64 `protobuf:"varint,2,opt,name=billboard_id,json=billboardId,proto3" json:"billboard_id,omitempty"`
	Start         uint64 `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	End           uint64 `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
	BidderAddress string `protobuf:"bytes,5,opt,name=bidder_address,json=bidderAddress,proto3" json:"bidder_address,omitempty"`
}

func (m *Auction) Reset()         { *m = Auction{} }
func (m *Auction) String() string { return proto.CompactTextString(m) }
func (*Auction) ProtoMessage()    {}
func (*Auction) Descriptor() ([]byte, []int) {
	return fileDescriptor_81fb1e49ab2fd419, []int{0}
}
func (m *Auction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auction.Merge(m, src)
}
func (m *Auction) XXX_Size() int {
	return m.Size()
}
func (m *Auction) XXX_DiscardUnknown() {
	xxx_messageInfo_Auction.DiscardUnknown(m)
}

var xxx_messageInfo_Auction proto.InternalMessageInfo

func (m *Auction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Auction) GetBillboardId() uint64 {
	if m != nil {
		return m.BillboardId
	}
	return 0
}

func (m *Auction) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Auction) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Auction) GetBidderAddress() string {
	if m != nil {
		return m.BidderAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Auction)(nil), "eeta.bid.Auction")
}

func init() { proto.RegisterFile("eeta/bid/auction.proto", fileDescriptor_81fb1e49ab2fd419) }

var fileDescriptor_81fb1e49ab2fd419 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x9b, 0x69, 0xeb, 0x4f, 0xd4, 0x52, 0x42, 0x95, 0x30, 0x8b, 0x50, 0x05, 0xa1, 0xb8,
	0x68, 0x28, 0x3e, 0x41, 0xdd, 0xb9, 0xed, 0xd2, 0x4d, 0x49, 0xe6, 0x86, 0x72, 0xa1, 0xed, 0x1d,
	0x92, 0x28, 0xfa, 0x0c, 0x6e, 0x7c, 0x2c, 0x97, 0x5d, 0xba, 0x94, 0x99, 0x17, 0x91, 0x49, 0xa4,
	0xbb, 0x93, 0xef, 0x1c, 0x72, 0xf9, 0xf8, 0x8d, 0x73, 0xd1, 0x68, 0x8b, 0xa0, 0xcd, 0x6b, 0x15,
	0x91, 0xf6, 0xf3, 0xda, 0x53, 0x24, 0x71, 0xd6, 0xf1, 0xb9, 0x45, 0x28, 0xaf, 0x8f, 0x8b, 0xda,
	0x78, 0xb3, 0x0b, 0x79, 0x50, 0xaa, 0x8a, 0xc2, 0x8e, 0x82, 0xb6, 0x26, 0x38, 0xfd, 0xb6, 0xb0,
	0x2e, 0x9a, 0x85, 0xae, 0x08, 0xff, 0x3f, 0x28, 0x27, 0x1b, 0xda, 0x50, 0x8a, 0xba, 0x4b, 0x99,
	0xde, 0x7d, 0x32, 0x7e, 0xba, 0xcc, 0x87, 0xc4, 0x88, 0x17, 0x08, 0x92, 0x4d, 0xd9, 0x6c, 0xb0,
	0x2a, 0x10, 0xc4, 0x2d, 0xbf, 0xb4, 0xb8, 0xdd, 0x5a, 0x32, 0x1e, 0xd6, 0x08, 0xb2, 0x48, 0xcd,
	0xc5, 0x91, 0x3d, 0x83, 0x98, 0xf0, 0x61, 0x88, 0xc6, 0x47, 0xd9, 0x4f, 0x5d, 0x7e, 0x88, 0x31,
	0xef, 0xbb, 0x3d, 0xc8, 0x41, 0x62, 0x5d, 0x14, 0xf7, 0x7c, 0x64, 0x11, 0xc0, 0xf9, 0xb5, 0x01,
	0xf0, 0x2e, 0x04, 0x39, 0x9c, 0xb2, 0xd9, 0xf9, 0xea, 0x2a, 0xd3, 0x65, 0x86, 0x4f, 0x0f, 0xdf,
	0x8d, 0x62, 0x87, 0x46, 0xb1, 0xdf, 0x46, 0xb1, 0xaf, 0x56, 0xf5, 0x0e, 0xad, 0xea, 0xfd, 0xb4,
	0xaa, 0xf7, 0x32, 0x4e, 0xd2, 0xef, 0x49, 0x3b, 0x7e, 0xd4, 0x2e, 0xd8, 0x93, 0x24, 0xf0, 0xf8,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x80, 0x85, 0xa2, 0x0c, 0x31, 0x01, 0x00, 0x00,
}

func (m *Auction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BidderAddress) > 0 {
		i -= len(m.BidderAddress)
		copy(dAtA[i:], m.BidderAddress)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.BidderAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.End != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.End))
		i--
		dAtA[i] = 0x20
	}
	if m.Start != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x18
	}
	if m.BillboardId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.BillboardId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Auction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	if m.BillboardId != 0 {
		n += 1 + sovAuction(uint64(m.BillboardId))
	}
	if m.Start != 0 {
		n += 1 + sovAuction(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovAuction(uint64(m.End))
	}
	l = len(m.BidderAddress)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Auction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Auction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillboardId", wireType)
			}
			m.BillboardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BillboardId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
