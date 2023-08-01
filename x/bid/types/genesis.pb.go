// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eeta/bid/genesis.proto

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

// GenesisState defines the bid module's genesis state.
type GenesisState struct {
	Params            Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BillboardAuctions []BillboardAuction `protobuf:"bytes,2,rep,name=billboard_auctions,json=billboardAuctions,proto3" json:"billboard_auctions"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_934515db421a9ad3, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBillboardAuctions() []BillboardAuction {
	if m != nil {
		return m.BillboardAuctions
	}
	return nil
}

type BillboardAuction struct {
	BillboardId uint64           `protobuf:"varint,1,opt,name=billboard_id,json=billboardId,proto3" json:"billboard_id,omitempty"`
	Auctions    []AuctionBidding `protobuf:"bytes,2,rep,name=auctions,proto3" json:"auctions"`
}

func (m *BillboardAuction) Reset()         { *m = BillboardAuction{} }
func (m *BillboardAuction) String() string { return proto.CompactTextString(m) }
func (*BillboardAuction) ProtoMessage()    {}
func (*BillboardAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_934515db421a9ad3, []int{1}
}
func (m *BillboardAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BillboardAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BillboardAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BillboardAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillboardAuction.Merge(m, src)
}
func (m *BillboardAuction) XXX_Size() int {
	return m.Size()
}
func (m *BillboardAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_BillboardAuction.DiscardUnknown(m)
}

var xxx_messageInfo_BillboardAuction proto.InternalMessageInfo

func (m *BillboardAuction) GetBillboardId() uint64 {
	if m != nil {
		return m.BillboardId
	}
	return 0
}

func (m *BillboardAuction) GetAuctions() []AuctionBidding {
	if m != nil {
		return m.Auctions
	}
	return nil
}

type AuctionBidding struct {
	Auction *Auction `protobuf:"bytes,1,opt,name=auction,proto3" json:"auction,omitempty"`
	Bids    []Bid    `protobuf:"bytes,2,rep,name=bids,proto3" json:"bids"`
}

func (m *AuctionBidding) Reset()         { *m = AuctionBidding{} }
func (m *AuctionBidding) String() string { return proto.CompactTextString(m) }
func (*AuctionBidding) ProtoMessage()    {}
func (*AuctionBidding) Descriptor() ([]byte, []int) {
	return fileDescriptor_934515db421a9ad3, []int{2}
}
func (m *AuctionBidding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionBidding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionBidding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionBidding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionBidding.Merge(m, src)
}
func (m *AuctionBidding) XXX_Size() int {
	return m.Size()
}
func (m *AuctionBidding) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionBidding.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionBidding proto.InternalMessageInfo

func (m *AuctionBidding) GetAuction() *Auction {
	if m != nil {
		return m.Auction
	}
	return nil
}

func (m *AuctionBidding) GetBids() []Bid {
	if m != nil {
		return m.Bids
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "eeta.bid.GenesisState")
	proto.RegisterType((*BillboardAuction)(nil), "eeta.bid.BillboardAuction")
	proto.RegisterType((*AuctionBidding)(nil), "eeta.bid.AuctionBidding")
}

func init() { proto.RegisterFile("eeta/bid/genesis.proto", fileDescriptor_934515db421a9ad3) }

var fileDescriptor_934515db421a9ad3 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x31, 0x4f, 0xb4, 0x30,
	0x18, 0xc7, 0xe9, 0xfb, 0x92, 0xf3, 0x52, 0x4e, 0xc3, 0x35, 0x6a, 0x08, 0x43, 0x3d, 0x59, 0xbc,
	0x68, 0x02, 0xc9, 0xb9, 0xb9, 0xc9, 0x62, 0x9c, 0x34, 0xb8, 0xb9, 0x98, 0xd6, 0x56, 0xd2, 0xe4,
	0xa4, 0x08, 0x35, 0xd1, 0x4f, 0xa1, 0x1f, 0xeb, 0xc6, 0x1b, 0x9d, 0x8c, 0x81, 0x2f, 0x62, 0xae,
	0x14, 0x50, 0xdc, 0xc8, 0xbf, 0xbf, 0xe7, 0xff, 0x2b, 0x7d, 0xe0, 0x3e, 0xe7, 0x8a, 0x44, 0x54,
	0xb0, 0x28, 0xe5, 0x19, 0x2f, 0x45, 0x19, 0xe6, 0x85, 0x54, 0x12, 0x8d, 0x37, 0x79, 0x48, 0x05,
	0xf3, 0x77, 0x53, 0x99, 0x4a, 0x1d, 0x46, 0x9b, 0xaf, 0xe6, 0xdc, 0xdf, 0xeb, 0xe6, 0x72, 0x52,
	0x90, 0x47, 0x33, 0xe6, 0xa3, 0x2e, 0xa6, 0x82, 0x99, 0xac, 0x57, 0x90, 0xe7, 0x7b, 0x25, 0x64,
	0xd6, 0xe4, 0xc1, 0x1b, 0x80, 0x93, 0x8b, 0x46, 0x7a, 0xa3, 0x88, 0xe2, 0x28, 0x84, 0xa3, 0xa6,
	0xcc, 0x03, 0x33, 0x30, 0x77, 0x16, 0x6e, 0xd8, 0x5e, 0x22, 0xbc, 0xd6, 0x79, 0x6c, 0xaf, 0x3e,
	0x0f, 0xac, 0xc4, 0x50, 0xe8, 0x0a, 0x22, 0x2a, 0x96, 0x4b, 0x2a, 0x49, 0xc1, 0xee, 0x4c, 0x77,
	0xe9, 0xfd, 0x9b, 0xfd, 0x9f, 0x3b, 0x0b, 0xbf, 0x9f, 0x8d, 0x5b, 0xe6, 0xbc, 0x41, 0x4c, 0xcb,
	0x94, 0x0e, 0xf2, 0x32, 0x78, 0x82, 0xee, 0x10, 0x46, 0x87, 0x70, 0xd2, 0x4b, 0x04, 0xd3, 0x57,
	0xb3, 0x13, 0xa7, 0xcb, 0x2e, 0x19, 0x3a, 0x83, 0xe3, 0x81, 0xdd, 0xeb, 0xed, 0xad, 0x54, 0x30,
	0x26, 0xb2, 0xd4, 0xb8, 0x3b, 0x3e, 0x78, 0x80, 0x3b, 0xbf, 0x09, 0x74, 0x02, 0xb7, 0xcc, 0xa9,
	0x79, 0x86, 0xe9, 0x9f, 0xb2, 0xa4, 0x25, 0xd0, 0x11, 0xb4, 0xa9, 0x60, 0xad, 0x76, 0xfb, 0xe7,
	0x4f, 0x33, 0xe3, 0xd2, 0x40, 0x7c, 0xbc, 0xaa, 0x30, 0x58, 0x57, 0x18, 0x7c, 0x55, 0x18, 0xbc,
	0xd7, 0xd8, 0x5a, 0xd7, 0xd8, 0xfa, 0xa8, 0xb1, 0x75, 0xeb, 0xea, 0xf5, 0xbc, 0xe8, 0x05, 0xa9,
	0xd7, 0x9c, 0x97, 0x74, 0xa4, 0xf7, 0x73, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xd7, 0x60,
	0xa0, 0x1c, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BillboardAuctions) > 0 {
		for iNdEx := len(m.BillboardAuctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BillboardAuctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BillboardAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BillboardAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BillboardAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Auctions) > 0 {
		for iNdEx := len(m.Auctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Auctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BillboardId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BillboardId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuctionBidding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionBidding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionBidding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bids) > 0 {
		for iNdEx := len(m.Bids) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bids[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Auction != nil {
		{
			size, err := m.Auction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.BillboardAuctions) > 0 {
		for _, e := range m.BillboardAuctions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *BillboardAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BillboardId != 0 {
		n += 1 + sovGenesis(uint64(m.BillboardId))
	}
	if len(m.Auctions) > 0 {
		for _, e := range m.Auctions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *AuctionBidding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Auction != nil {
		l = m.Auction.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Bids) > 0 {
		for _, e := range m.Bids {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillboardAuctions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BillboardAuctions = append(m.BillboardAuctions, BillboardAuction{})
			if err := m.BillboardAuctions[len(m.BillboardAuctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *BillboardAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: BillboardAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BillboardAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillboardId", wireType)
			}
			m.BillboardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auctions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auctions = append(m.Auctions, AuctionBidding{})
			if err := m.Auctions[len(m.Auctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *AuctionBidding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: AuctionBidding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionBidding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Auction == nil {
				m.Auction = &Auction{}
			}
			if err := m.Auction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bids", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bids = append(m.Bids, Bid{})
			if err := m.Bids[len(m.Bids)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
