// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eeta/deposit/tx.proto

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

type MsgMint struct {
	Sender            string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	ReceipientAddress string     `protobuf:"bytes,2,opt,name=receipient_address,json=receipientAddress,proto3" json:"receipient_address,omitempty"`
	Coin              types.Coin `protobuf:"bytes,3,opt,name=coin,proto3" json:"coin"`
}

func (m *MsgMint) Reset()         { *m = MsgMint{} }
func (m *MsgMint) String() string { return proto.CompactTextString(m) }
func (*MsgMint) ProtoMessage()    {}
func (*MsgMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_97bbbb1bffcf622b, []int{0}
}
func (m *MsgMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMint.Merge(m, src)
}
func (m *MsgMint) XXX_Size() int {
	return m.Size()
}
func (m *MsgMint) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMint.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMint proto.InternalMessageInfo

func (m *MsgMint) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgMint) GetReceipientAddress() string {
	if m != nil {
		return m.ReceipientAddress
	}
	return ""
}

func (m *MsgMint) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

type MsgMintResponse struct {
}

func (m *MsgMintResponse) Reset()         { *m = MsgMintResponse{} }
func (m *MsgMintResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintResponse) ProtoMessage()    {}
func (*MsgMintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97bbbb1bffcf622b, []int{1}
}
func (m *MsgMintResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintResponse.Merge(m, src)
}
func (m *MsgMintResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintResponse proto.InternalMessageInfo

type MsgBurn struct {
	Sender        string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	BurnedAddress string     `protobuf:"bytes,2,opt,name=burnedAddress,proto3" json:"burnedAddress,omitempty"`
	Amount        types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgBurn) Reset()         { *m = MsgBurn{} }
func (m *MsgBurn) String() string { return proto.CompactTextString(m) }
func (*MsgBurn) ProtoMessage()    {}
func (*MsgBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_97bbbb1bffcf622b, []int{2}
}
func (m *MsgBurn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurn.Merge(m, src)
}
func (m *MsgBurn) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurn.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurn proto.InternalMessageInfo

func (m *MsgBurn) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgBurn) GetBurnedAddress() string {
	if m != nil {
		return m.BurnedAddress
	}
	return ""
}

func (m *MsgBurn) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgBurnResponse struct {
}

func (m *MsgBurnResponse) Reset()         { *m = MsgBurnResponse{} }
func (m *MsgBurnResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBurnResponse) ProtoMessage()    {}
func (*MsgBurnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97bbbb1bffcf622b, []int{3}
}
func (m *MsgBurnResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnResponse.Merge(m, src)
}
func (m *MsgBurnResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgMint)(nil), "eeta.deposit.MsgMint")
	proto.RegisterType((*MsgMintResponse)(nil), "eeta.deposit.MsgMintResponse")
	proto.RegisterType((*MsgBurn)(nil), "eeta.deposit.MsgBurn")
	proto.RegisterType((*MsgBurnResponse)(nil), "eeta.deposit.MsgBurnResponse")
}

func init() { proto.RegisterFile("eeta/deposit/tx.proto", fileDescriptor_97bbbb1bffcf622b) }

var fileDescriptor_97bbbb1bffcf622b = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4f, 0x3a, 0x31,
	0x14, 0xc7, 0xaf, 0x3f, 0x08, 0xbf, 0x58, 0x35, 0x86, 0x06, 0x0c, 0x92, 0x58, 0x09, 0x71, 0x60,
	0xb1, 0x17, 0x60, 0x70, 0x71, 0x01, 0x67, 0x16, 0x46, 0x17, 0x73, 0xc7, 0xbd, 0x5c, 0x6e, 0xa0,
	0xbd, 0xf4, 0x15, 0x83, 0x9b, 0x0e, 0xee, 0xfe, 0x59, 0x8c, 0x8c, 0x4e, 0xc6, 0xc0, 0x3f, 0x62,
	0xda, 0x2b, 0x01, 0x22, 0x26, 0x6e, 0x77, 0xdf, 0xef, 0xf7, 0xb5, 0x9f, 0xd7, 0xf7, 0x68, 0x1d,
	0xc0, 0x44, 0x61, 0x02, 0xb9, 0xc2, 0xcc, 0x84, 0x66, 0x2e, 0x72, 0xad, 0x8c, 0x62, 0x27, 0x56,
	0x16, 0x5e, 0x6e, 0xd6, 0x52, 0x95, 0x2a, 0x67, 0x84, 0xf6, 0xab, 0xc8, 0x34, 0xf9, 0x44, 0xe1,
	0x54, 0x61, 0x18, 0x47, 0x08, 0xe1, 0x53, 0x37, 0x06, 0x13, 0x75, 0xc3, 0x89, 0xca, 0x64, 0xe1,
	0xb7, 0xdf, 0x08, 0xfd, 0x3f, 0xc2, 0x74, 0x94, 0x49, 0xc3, 0xce, 0x69, 0x05, 0x41, 0x26, 0xa0,
	0x1b, 0xa4, 0x45, 0x3a, 0x47, 0x63, 0xff, 0xc7, 0x6e, 0x28, 0xd3, 0x30, 0x81, 0x2c, 0xcf, 0x40,
	0x9a, 0xc7, 0x28, 0x49, 0x34, 0x20, 0x36, 0xfe, 0xb9, 0x4c, 0x75, 0xeb, 0x0c, 0x0a, 0x83, 0xf5,
	0x69, 0xd9, 0x5e, 0xd0, 0x28, 0xb5, 0x48, 0xe7, 0xb8, 0x77, 0x21, 0x0a, 0x02, 0x61, 0x09, 0x84,
	0x27, 0x10, 0xf7, 0x2a, 0x93, 0xc3, 0xf2, 0xe2, 0xf3, 0x2a, 0x18, 0xbb, 0x70, 0xbb, 0x4a, 0xcf,
	0x3c, 0xc6, 0x18, 0x30, 0x57, 0x12, 0xa1, 0xfd, 0x52, 0xa0, 0x0d, 0x67, 0x5a, 0xfe, 0x8a, 0x76,
	0x4d, 0x4f, 0xe3, 0x99, 0x96, 0x90, 0x0c, 0xf6, 0xa8, 0xf6, 0x45, 0x76, 0x4b, 0x2b, 0xd1, 0x54,
	0xcd, 0xa4, 0xf9, 0x2b, 0x93, 0x8f, 0x7b, 0x2a, 0x4b, 0xb0, 0xa1, 0xea, 0xbd, 0x12, 0x5a, 0x1a,
	0x61, 0xca, 0xee, 0x68, 0x79, 0x6a, 0x1f, 0xad, 0x2e, 0x76, 0xa7, 0x20, 0x7c, 0x13, 0xcd, 0xcb,
	0x83, 0xf2, 0xe6, 0x14, 0x5b, 0xed, 0xfa, 0xfa, 0x59, 0x6d, 0xe5, 0x03, 0xd5, 0xbb, 0x0c, 0x43,
	0xb1, 0x58, 0x71, 0xb2, 0x5c, 0x71, 0xf2, 0xb5, 0xe2, 0xe4, 0x7d, 0xcd, 0x83, 0xe5, 0x9a, 0x07,
	0x1f, 0x6b, 0x1e, 0x3c, 0xd4, 0xdc, 0xa6, 0xcc, 0xb7, 0xbb, 0xf2, 0x9c, 0x03, 0xc6, 0x15, 0x37,
	0xeb, 0xfe, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x66, 0x90, 0x21, 0x48, 0x02, 0x00, 0x00,
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
	Mint(ctx context.Context, in *MsgMint, opts ...grpc.CallOption) (*MsgMintResponse, error)
	Burn(ctx context.Context, in *MsgBurn, opts ...grpc.CallOption) (*MsgBurnResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Mint(ctx context.Context, in *MsgMint, opts ...grpc.CallOption) (*MsgMintResponse, error) {
	out := new(MsgMintResponse)
	err := c.cc.Invoke(ctx, "/eeta.deposit.Msg/mint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Burn(ctx context.Context, in *MsgBurn, opts ...grpc.CallOption) (*MsgBurnResponse, error) {
	out := new(MsgBurnResponse)
	err := c.cc.Invoke(ctx, "/eeta.deposit.Msg/Burn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Mint(context.Context, *MsgMint) (*MsgMintResponse, error)
	Burn(context.Context, *MsgBurn) (*MsgBurnResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Mint(ctx context.Context, req *MsgMint) (*MsgMintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mint not implemented")
}
func (*UnimplementedMsgServer) Burn(ctx context.Context, req *MsgBurn) (*MsgBurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Burn not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eeta.deposit.Msg/Mint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Mint(ctx, req.(*MsgMint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Burn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Burn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eeta.deposit.Msg/Burn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Burn(ctx, req.(*MsgBurn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eeta.deposit.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "mint",
			Handler:    _Msg_Mint_Handler,
		},
		{
			MethodName: "Burn",
			Handler:    _Msg_Burn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eeta/deposit/tx.proto",
}

func (m *MsgMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ReceipientAddress) > 0 {
		i -= len(m.ReceipientAddress)
		copy(dAtA[i:], m.ReceipientAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ReceipientAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBurn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
	if len(m.BurnedAddress) > 0 {
		i -= len(m.BurnedAddress)
		copy(dAtA[i:], m.BurnedAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BurnedAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBurnResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ReceipientAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgMintResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBurn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BurnedAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgBurnResponse) Size() (n int) {
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
func (m *MsgMint) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceipientAddress", wireType)
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
			m.ReceipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
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
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgMintResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMintResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgBurn) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgBurn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnedAddress", wireType)
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
			m.BurnedAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *MsgBurnResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgBurnResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
