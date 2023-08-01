// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eeta/billboard/tx.proto

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

type MsgCreateBillboard struct {
	Creator     string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name        string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Url         string     `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	BoardType   string     `protobuf:"bytes,5,opt,name=boardType,proto3" json:"boardType,omitempty"`
	MinPrice    types.Coin `protobuf:"bytes,6,opt,name=minPrice,proto3" json:"minPrice"`
}

func (m *MsgCreateBillboard) Reset()         { *m = MsgCreateBillboard{} }
func (m *MsgCreateBillboard) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBillboard) ProtoMessage()    {}
func (*MsgCreateBillboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_d22ad6fda9c78063, []int{0}
}
func (m *MsgCreateBillboard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBillboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBillboard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBillboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBillboard.Merge(m, src)
}
func (m *MsgCreateBillboard) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBillboard) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBillboard.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBillboard proto.InternalMessageInfo

func (m *MsgCreateBillboard) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateBillboard) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreateBillboard) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgCreateBillboard) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MsgCreateBillboard) GetBoardType() string {
	if m != nil {
		return m.BoardType
	}
	return ""
}

func (m *MsgCreateBillboard) GetMinPrice() types.Coin {
	if m != nil {
		return m.MinPrice
	}
	return types.Coin{}
}

type MsgCreateBillboardResponse struct {
}

func (m *MsgCreateBillboardResponse) Reset()         { *m = MsgCreateBillboardResponse{} }
func (m *MsgCreateBillboardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBillboardResponse) ProtoMessage()    {}
func (*MsgCreateBillboardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d22ad6fda9c78063, []int{1}
}
func (m *MsgCreateBillboardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBillboardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBillboardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBillboardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBillboardResponse.Merge(m, src)
}
func (m *MsgCreateBillboardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBillboardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBillboardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBillboardResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateBillboard)(nil), "eeta.billboard.MsgCreateBillboard")
	proto.RegisterType((*MsgCreateBillboardResponse)(nil), "eeta.billboard.MsgCreateBillboardResponse")
}

func init() { proto.RegisterFile("eeta/billboard/tx.proto", fileDescriptor_d22ad6fda9c78063) }

var fileDescriptor_d22ad6fda9c78063 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0xa7, 0x1f, 0x7c, 0x28, 0x25, 0x51, 0xd3, 0x18, 0xad, 0x13, 0x52, 0x09, 0x2b, 0xe2,
	0xa2, 0x23, 0xb8, 0x74, 0x07, 0x6b, 0x12, 0x43, 0x5c, 0xb9, 0xeb, 0x0c, 0x37, 0x63, 0x13, 0x98,
	0x3b, 0x69, 0xab, 0x81, 0xb7, 0xf0, 0xb1, 0x58, 0xe2, 0xce, 0x95, 0x31, 0xf0, 0x22, 0x66, 0xca,
	0x1f, 0x11, 0x16, 0xee, 0xce, 0xfc, 0xce, 0x9d, 0x7b, 0x73, 0x7a, 0xe8, 0x25, 0x80, 0x53, 0x51,
	0xac, 0x47, 0xa3, 0x18, 0x95, 0x19, 0x46, 0x6e, 0x22, 0x73, 0x83, 0x0e, 0xd9, 0x49, 0x61, 0xc8,
	0xad, 0x11, 0x8a, 0xbd, 0xc1, 0xad, 0x5a, 0xcd, 0x87, 0x22, 0x41, 0x3b, 0x46, 0x1b, 0xc5, 0xca,
	0x42, 0xf4, 0xda, 0x8e, 0xc1, 0xa9, 0x76, 0x94, 0xa0, 0xce, 0xd6, 0xfe, 0x79, 0x8a, 0x29, 0x7a,
	0x19, 0x15, 0x6a, 0x45, 0x9b, 0xef, 0x84, 0xb2, 0xbe, 0x4d, 0x7b, 0x06, 0x94, 0x83, 0xee, 0x66,
	0x25, 0xe3, 0xf4, 0x28, 0x29, 0x10, 0x1a, 0x4e, 0x1a, 0xa4, 0x55, 0x1d, 0x6c, 0x3e, 0x19, 0xa3,
	0xe5, 0x4c, 0x8d, 0x81, 0xff, 0xf3, 0xd8, 0x6b, 0xd6, 0xa0, 0xb5, 0x21, 0xd8, 0xc4, 0xe8, 0xdc,
	0x69, 0xcc, 0x78, 0xc9, 0x5b, 0xbb, 0x88, 0x9d, 0xd1, 0xd2, 0x8b, 0x19, 0xf1, 0xb2, 0x77, 0x0a,
	0xc9, 0xea, 0xb4, 0xea, 0x4f, 0x3d, 0x4e, 0x73, 0xe0, 0xff, 0x3d, 0xff, 0x01, 0xec, 0x9e, 0x1e,
	0x8f, 0x75, 0xf6, 0x60, 0x74, 0x02, 0xbc, 0xd2, 0x20, 0xad, 0x5a, 0xe7, 0x4a, 0xae, 0xf2, 0xc9,
	0x22, 0x9f, 0x5c, 0xe7, 0x93, 0x3d, 0xd4, 0x59, 0xb7, 0x3c, 0xfb, 0xbc, 0x0e, 0x06, 0xdb, 0x1f,
	0x9a, 0x75, 0x1a, 0x1e, 0x46, 0x1a, 0x80, 0xcd, 0x31, 0xb3, 0xd0, 0x79, 0xa6, 0xa5, 0xbe, 0x4d,
	0x99, 0xa2, 0xa7, 0xfb, 0xa1, 0x9b, 0xf2, 0xf7, 0x93, 0xcb, 0xc3, 0x2d, 0xe1, 0xcd, 0xdf, 0x33,
	0x9b, 0x4b, 0xdd, 0xdb, 0xd9, 0x42, 0x90, 0xf9, 0x42, 0x90, 0xaf, 0x85, 0x20, 0x6f, 0x4b, 0x11,
	0xcc, 0x97, 0x22, 0xf8, 0x58, 0x8a, 0xe0, 0xe9, 0xc2, 0x77, 0x39, 0xd9, 0xad, 0x7d, 0x9a, 0x83,
	0x8d, 0x2b, 0xbe, 0x94, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x0b, 0x8d, 0x80, 0x15,
	0x02, 0x00, 0x00,
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
	CreateBillboard(ctx context.Context, in *MsgCreateBillboard, opts ...grpc.CallOption) (*MsgCreateBillboardResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBillboard(ctx context.Context, in *MsgCreateBillboard, opts ...grpc.CallOption) (*MsgCreateBillboardResponse, error) {
	out := new(MsgCreateBillboardResponse)
	err := c.cc.Invoke(ctx, "/eeta.billboard.Msg/CreateBillboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateBillboard(context.Context, *MsgCreateBillboard) (*MsgCreateBillboardResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBillboard(ctx context.Context, req *MsgCreateBillboard) (*MsgCreateBillboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBillboard not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBillboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBillboard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBillboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eeta.billboard.Msg/CreateBillboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBillboard(ctx, req.(*MsgCreateBillboard))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eeta.billboard.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBillboard",
			Handler:    _Msg_CreateBillboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eeta/billboard/tx.proto",
}

func (m *MsgCreateBillboard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBillboard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBillboard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.BoardType) > 0 {
		i -= len(m.BoardType)
		copy(dAtA[i:], m.BoardType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BoardType)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgCreateBillboardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBillboardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBillboardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateBillboard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BoardType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MinPrice.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCreateBillboardResponse) Size() (n int) {
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
func (m *MsgCreateBillboard) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateBillboard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBillboard: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
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
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoardType", wireType)
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
			m.BoardType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPrice", wireType)
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
			if err := m.MinPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreateBillboardResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateBillboardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBillboardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
