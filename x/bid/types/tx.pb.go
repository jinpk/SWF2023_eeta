// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eeta/bid/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgCreateAuction struct {
	Creator     string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	BillboardId uint64     `protobuf:"varint,2,opt,name=billboardId,proto3" json:"billboardId,omitempty"`
	Start       uint64     `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	End         uint64     `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
	Amount      types.Coin `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount"`
	AdUrl       string     `protobuf:"bytes,6,opt,name=adUrl,proto3" json:"adUrl,omitempty"`
}

func (m *MsgCreateAuction) Reset()         { *m = MsgCreateAuction{} }
func (m *MsgCreateAuction) String() string { return proto.CompactTextString(m) }
func (*MsgCreateAuction) ProtoMessage()    {}
func (*MsgCreateAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2253e84dd08939d, []int{0}
}
func (m *MsgCreateAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAuction.Merge(m, src)
}
func (m *MsgCreateAuction) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAuction.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAuction proto.InternalMessageInfo

func (m *MsgCreateAuction) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateAuction) GetBillboardId() uint64 {
	if m != nil {
		return m.BillboardId
	}
	return 0
}

func (m *MsgCreateAuction) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *MsgCreateAuction) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *MsgCreateAuction) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *MsgCreateAuction) GetAdUrl() string {
	if m != nil {
		return m.AdUrl
	}
	return ""
}

type MsgCreateAuctionResponse struct {
}

func (m *MsgCreateAuctionResponse) Reset()         { *m = MsgCreateAuctionResponse{} }
func (m *MsgCreateAuctionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateAuctionResponse) ProtoMessage()    {}
func (*MsgCreateAuctionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2253e84dd08939d, []int{1}
}
func (m *MsgCreateAuctionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAuctionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAuctionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAuctionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAuctionResponse.Merge(m, src)
}
func (m *MsgCreateAuctionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAuctionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAuctionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAuctionResponse proto.InternalMessageInfo

type MsgBidding struct {
	Creator     string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	AuctionId   uint64     `protobuf:"varint,2,opt,name=auctionId,proto3" json:"auctionId,omitempty"`
	Amount      types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	AdUrl       string     `protobuf:"bytes,4,opt,name=adUrl,proto3" json:"adUrl,omitempty"`
	BillboardId uint64     `protobuf:"varint,5,opt,name=billboardId,proto3" json:"billboardId,omitempty"`
}

func (m *MsgBidding) Reset()         { *m = MsgBidding{} }
func (m *MsgBidding) String() string { return proto.CompactTextString(m) }
func (*MsgBidding) ProtoMessage()    {}
func (*MsgBidding) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2253e84dd08939d, []int{2}
}
func (m *MsgBidding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBidding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBidding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBidding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBidding.Merge(m, src)
}
func (m *MsgBidding) XXX_Size() int {
	return m.Size()
}
func (m *MsgBidding) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBidding.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBidding proto.InternalMessageInfo

func (m *MsgBidding) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgBidding) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func (m *MsgBidding) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *MsgBidding) GetAdUrl() string {
	if m != nil {
		return m.AdUrl
	}
	return ""
}

func (m *MsgBidding) GetBillboardId() uint64 {
	if m != nil {
		return m.BillboardId
	}
	return 0
}

type MsgBiddingResponse struct {
}

func (m *MsgBiddingResponse) Reset()         { *m = MsgBiddingResponse{} }
func (m *MsgBiddingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBiddingResponse) ProtoMessage()    {}
func (*MsgBiddingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2253e84dd08939d, []int{3}
}
func (m *MsgBiddingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBiddingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBiddingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBiddingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBiddingResponse.Merge(m, src)
}
func (m *MsgBiddingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBiddingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBiddingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBiddingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateAuction)(nil), "eeta.bid.MsgCreateAuction")
	proto.RegisterType((*MsgCreateAuctionResponse)(nil), "eeta.bid.MsgCreateAuctionResponse")
	proto.RegisterType((*MsgBidding)(nil), "eeta.bid.MsgBidding")
	proto.RegisterType((*MsgBiddingResponse)(nil), "eeta.bid.MsgBiddingResponse")
}

func init() { proto.RegisterFile("eeta/bid/tx.proto", fileDescriptor_c2253e84dd08939d) }

var fileDescriptor_c2253e84dd08939d = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x4e, 0xfb, 0x30,
	0x10, 0xc0, 0xe3, 0x7f, 0xd2, 0xf6, 0x5f, 0x57, 0x48, 0xc5, 0xca, 0x60, 0xa2, 0x2a, 0x44, 0x99,
	0x2a, 0x06, 0x47, 0x2d, 0x03, 0x13, 0x03, 0xed, 0xc4, 0x50, 0x21, 0x45, 0x62, 0x61, 0x73, 0x62,
	0x2b, 0xb2, 0xd4, 0xc6, 0x55, 0xec, 0xa2, 0xf2, 0x16, 0x2c, 0x3c, 0x0a, 0x8f, 0x80, 0xd4, 0xb1,
	0x23, 0x13, 0x42, 0xed, 0x8b, 0xa0, 0x24, 0x0d, 0xfd, 0xe0, 0x63, 0x60, 0xf3, 0xfd, 0xce, 0x97,
	0xfc, 0xee, 0x7c, 0xf0, 0x98, 0x73, 0x4d, 0x83, 0x48, 0xb0, 0x40, 0xcf, 0xc9, 0x34, 0x93, 0x5a,
	0xa2, 0xff, 0x39, 0x22, 0x91, 0x60, 0x8e, 0x9d, 0xc8, 0x44, 0x16, 0x30, 0xc8, 0x4f, 0x65, 0xde,
	0x71, 0x63, 0xa9, 0x26, 0x52, 0x05, 0x11, 0x55, 0x3c, 0xb8, 0xef, 0x45, 0x5c, 0xd3, 0x5e, 0x10,
	0x4b, 0x91, 0x96, 0x79, 0xff, 0x05, 0xc0, 0xf6, 0x48, 0x25, 0xc3, 0x8c, 0x53, 0xcd, 0xaf, 0x66,
	0xb1, 0x16, 0x32, 0x45, 0x18, 0x36, 0xe2, 0x1c, 0xc8, 0x0c, 0x03, 0x0f, 0x74, 0x9b, 0x61, 0x15,
	0x22, 0x0f, 0xb6, 0x22, 0x31, 0x1e, 0x47, 0x92, 0x66, 0xec, 0x9a, 0xe1, 0x7f, 0x1e, 0xe8, 0x5a,
	0xe1, 0x2e, 0x42, 0x36, 0xac, 0x29, 0x4d, 0x33, 0x8d, 0xcd, 0x22, 0x57, 0x06, 0xa8, 0x0d, 0x4d,
	0x9e, 0x32, 0x6c, 0x15, 0x2c, 0x3f, 0xa2, 0x0b, 0x58, 0xa7, 0x13, 0x39, 0x4b, 0x35, 0xae, 0x79,
	0xa0, 0xdb, 0xea, 0x9f, 0x90, 0xd2, 0x94, 0xe4, 0xa6, 0x64, 0x63, 0x4a, 0x86, 0x52, 0xa4, 0x03,
	0x6b, 0xf1, 0x76, 0x6a, 0x84, 0x9b, 0xeb, 0xf9, 0x0f, 0x28, 0xbb, 0xcd, 0xc6, 0xb8, 0x5e, 0xa8,
	0x95, 0x81, 0xef, 0x40, 0x7c, 0xd8, 0x46, 0xc8, 0xd5, 0x54, 0xa6, 0x8a, 0xfb, 0xcf, 0x00, 0xc2,
	0x91, 0x4a, 0x06, 0x82, 0x31, 0x91, 0x26, 0xbf, 0x74, 0xd7, 0x81, 0x4d, 0x5a, 0xd6, 0x7e, 0xf6,
	0xb6, 0x05, 0x3b, 0xc6, 0xe6, 0x1f, 0x8d, 0xad, 0x1d, 0xe3, 0xc3, 0x51, 0xd6, 0xbe, 0x8c, 0xd2,
	0xb7, 0x21, 0xda, 0x6a, 0x57, 0xdd, 0xf4, 0x9f, 0x00, 0x34, 0x47, 0x2a, 0x41, 0x37, 0xf0, 0x68,
	0xff, 0xd5, 0x1c, 0x52, 0xed, 0x02, 0x39, 0x1c, 0x85, 0xe3, 0xff, 0x9c, 0xab, 0x3e, 0x8c, 0x2e,
	0x61, 0xa3, 0x1a, 0x91, 0xbd, 0x77, 0x7d, 0x43, 0x9d, 0xce, 0x77, 0xb4, 0x2a, 0x1f, 0x9c, 0x2d,
	0x56, 0x2e, 0x58, 0xae, 0x5c, 0xf0, 0xbe, 0x72, 0xc1, 0xe3, 0xda, 0x35, 0x96, 0x6b, 0xd7, 0x78,
	0x5d, 0xbb, 0xc6, 0x5d, 0xbb, 0x58, 0xdb, 0x79, 0xb9, 0xb8, 0x0f, 0x53, 0xae, 0xa2, 0x7a, 0xb1,
	0x7c, 0xe7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0x2a, 0x84, 0x04, 0xd1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateAuction(ctx context.Context, in *MsgCreateAuction, opts ...grpc.CallOption) (*MsgCreateAuctionResponse, error)
	Bidding(ctx context.Context, in *MsgBidding, opts ...grpc.CallOption) (*MsgBiddingResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateAuction(ctx context.Context, in *MsgCreateAuction, opts ...grpc.CallOption) (*MsgCreateAuctionResponse, error) {
	out := new(MsgCreateAuctionResponse)
	err := c.cc.Invoke(ctx, "/eeta.bid.Msg/CreateAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Bidding(ctx context.Context, in *MsgBidding, opts ...grpc.CallOption) (*MsgBiddingResponse, error) {
	out := new(MsgBiddingResponse)
	err := c.cc.Invoke(ctx, "/eeta.bid.Msg/Bidding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateAuction(context.Context, *MsgCreateAuction) (*MsgCreateAuctionResponse, error)
	Bidding(context.Context, *MsgBidding) (*MsgBiddingResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateAuction(ctx context.Context, req *MsgCreateAuction) (*MsgCreateAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuction not implemented")
}
func (*UnimplementedMsgServer) Bidding(ctx context.Context, req *MsgBidding) (*MsgBiddingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bidding not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAuction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eeta.bid.Msg/CreateAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAuction(ctx, req.(*MsgCreateAuction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Bidding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBidding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Bidding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eeta.bid.Msg/Bidding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Bidding(ctx, req.(*MsgBidding))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eeta.bid.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuction",
			Handler:    _Msg_CreateAuction_Handler,
		},
		{
			MethodName: "Bidding",
			Handler:    _Msg_Bidding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eeta/bid/tx.proto",
}

func (m *MsgCreateAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdUrl) > 0 {
		i -= len(m.AdUrl)
		copy(dAtA[i:], m.AdUrl)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AdUrl)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.End != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.End))
		i--
		dAtA[i] = 0x20
	}
	if m.Start != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x18
	}
	if m.BillboardId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BillboardId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateAuctionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAuctionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAuctionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBidding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBidding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBidding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BillboardId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BillboardId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AdUrl) > 0 {
		i -= len(m.AdUrl)
		copy(dAtA[i:], m.AdUrl)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AdUrl)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.AuctionId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBiddingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBiddingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBiddingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BillboardId != 0 {
		n += 1 + sovTx(uint64(m.BillboardId))
	}
	if m.Start != 0 {
		n += 1 + sovTx(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovTx(uint64(m.End))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.AdUrl)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateAuctionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBidding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AuctionId != 0 {
		n += 1 + sovTx(uint64(m.AuctionId))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.AdUrl)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BillboardId != 0 {
		n += 1 + sovTx(uint64(m.BillboardId))
	}
	return n
}

func (m *MsgBiddingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillboardId", wireType)
			}
			m.BillboardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
					return ErrIntOverflowTx
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
					return ErrIntOverflowTx
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
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateAuctionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateAuctionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAuctionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBidding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBidding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBidding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillboardId", wireType)
			}
			m.BillboardId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBiddingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBiddingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBiddingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
