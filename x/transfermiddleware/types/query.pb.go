// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: banksy/transfermiddleware/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParaTokenInfoRequest is the request type for the Query/Params RPC method.
type QueryParaTokenInfoRequest struct {
	NativeDenom string `protobuf:"bytes,1,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom,omitempty" yaml:"native_denom"`
}

func (m *QueryParaTokenInfoRequest) Reset()         { *m = QueryParaTokenInfoRequest{} }
func (m *QueryParaTokenInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParaTokenInfoRequest) ProtoMessage()    {}
func (*QueryParaTokenInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6411fbf179210b0, []int{0}
}
func (m *QueryParaTokenInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParaTokenInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParaTokenInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParaTokenInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParaTokenInfoRequest.Merge(m, src)
}
func (m *QueryParaTokenInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParaTokenInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParaTokenInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParaTokenInfoRequest proto.InternalMessageInfo

func (m *QueryParaTokenInfoRequest) GetNativeDenom() string {
	if m != nil {
		return m.NativeDenom
	}
	return ""
}

// QueryParaTokenInfoResponse is the response type for the Query/ParaTokenInfo RPC method.
type QueryParaTokenInfoResponse struct {
	IbcDenom    string `protobuf:"bytes,2,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty" yaml:"ibc_denom"`
	ChannelId   string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	NativeDenom string `protobuf:"bytes,7,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom,omitempty" yaml:"native_denom"`
}

func (m *QueryParaTokenInfoResponse) Reset()         { *m = QueryParaTokenInfoResponse{} }
func (m *QueryParaTokenInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParaTokenInfoResponse) ProtoMessage()    {}
func (*QueryParaTokenInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6411fbf179210b0, []int{1}
}
func (m *QueryParaTokenInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParaTokenInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParaTokenInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParaTokenInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParaTokenInfoResponse.Merge(m, src)
}
func (m *QueryParaTokenInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParaTokenInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParaTokenInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParaTokenInfoResponse proto.InternalMessageInfo

func (m *QueryParaTokenInfoResponse) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

func (m *QueryParaTokenInfoResponse) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *QueryParaTokenInfoResponse) GetNativeDenom() string {
	if m != nil {
		return m.NativeDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParaTokenInfoRequest)(nil), "banksy.transfermiddleware.v1beta1.QueryParaTokenInfoRequest")
	proto.RegisterType((*QueryParaTokenInfoResponse)(nil), "banksy.transfermiddleware.v1beta1.QueryParaTokenInfoResponse")
}

func init() {
	proto.RegisterFile("banksy/transfermiddleware/v1beta1/query.proto", fileDescriptor_f6411fbf179210b0)
}

var fileDescriptor_f6411fbf179210b0 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x4a, 0xcc, 0xcb,
	0x2e, 0xae, 0xd4, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4b, 0x2d, 0xca, 0xcd, 0x4c, 0x49, 0xc9,
	0x49, 0x2d, 0x4f, 0x2c, 0x4a, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2c,
	0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x84, 0x28, 0xd7, 0xc3, 0x54,
	0xae, 0x07, 0x55, 0x2e, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xad, 0x0f, 0x62, 0x41, 0x34,
	0x4a, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5,
	0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x64, 0x95, 0xc2, 0xb9, 0x24, 0x03, 0x41,
	0xb6, 0x04, 0x24, 0x16, 0x25, 0x86, 0xe4, 0x67, 0xa7, 0xe6, 0x79, 0xe6, 0xa5, 0xe5, 0x07, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x59, 0x71, 0xf1, 0xe4, 0x25, 0x96, 0x64, 0x96, 0xa5, 0xc6,
	0xa7, 0xa4, 0xe6, 0xe5, 0xe7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0x89, 0x7f, 0xba, 0x27,
	0x2f, 0x5c, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x84, 0x2c, 0xab, 0x14, 0xc4, 0x0d, 0xe1, 0xba, 0x80,
	0x79, 0x07, 0x19, 0xb9, 0xa4, 0xb0, 0x99, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xc8,
	0xc5, 0x99, 0x99, 0x94, 0x0c, 0x35, 0x97, 0x09, 0x6c, 0xae, 0xc8, 0xa7, 0x7b, 0xf2, 0x02, 0x10,
	0x73, 0xe1, 0x52, 0x4a, 0x41, 0x1c, 0x99, 0x49, 0xc9, 0x60, 0x13, 0x85, 0x4c, 0xb8, 0xb8, 0x92,
	0x33, 0x12, 0xf3, 0xf2, 0x52, 0x73, 0xe2, 0x33, 0x53, 0x24, 0x98, 0xc1, 0x7a, 0x44, 0x3f, 0xdd,
	0x93, 0x17, 0x84, 0xe8, 0x41, 0xc8, 0x29, 0x05, 0x71, 0x42, 0x39, 0x9e, 0x29, 0x18, 0x7e, 0x60,
	0x27, 0xde, 0x0f, 0x46, 0x5b, 0x19, 0xb9, 0x58, 0xc1, 0x7e, 0x10, 0x5a, 0xcd, 0xc8, 0xc5, 0x8b,
	0xe2, 0x11, 0x21, 0x1b, 0x3d, 0x82, 0x11, 0xa2, 0x87, 0x33, 0x64, 0xa5, 0x6c, 0xc9, 0xd4, 0x0d,
	0x09, 0x3d, 0x25, 0xd9, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x89, 0x0b, 0x89, 0xea, 0x43, 0x13, 0x51,
	0x41, 0x62, 0x51, 0x62, 0x09, 0x48, 0x59, 0x66, 0x5e, 0x5a, 0xbe, 0x93, 0xc9, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x49, 0x55, 0x60, 0x4b, 0x6f, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0xe0, 0x14, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x14, 0xb2, 0x3a, 0x79,
	0x99, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// ParaTokenInfo queries all token info of a native denom.
	ParaTokenInfo(ctx context.Context, in *QueryParaTokenInfoRequest, opts ...grpc.CallOption) (*QueryParaTokenInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ParaTokenInfo(ctx context.Context, in *QueryParaTokenInfoRequest, opts ...grpc.CallOption) (*QueryParaTokenInfoResponse, error) {
	out := new(QueryParaTokenInfoResponse)
	err := c.cc.Invoke(ctx, "/banksy.transfermiddleware.v1beta1.Query/ParaTokenInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ParaTokenInfo queries all token info of a native denom.
	ParaTokenInfo(context.Context, *QueryParaTokenInfoRequest) (*QueryParaTokenInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ParaTokenInfo(ctx context.Context, req *QueryParaTokenInfoRequest) (*QueryParaTokenInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParaTokenInfo not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ParaTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParaTokenInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ParaTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banksy.transfermiddleware.v1beta1.Query/ParaTokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ParaTokenInfo(ctx, req.(*QueryParaTokenInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "banksy.transfermiddleware.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParaTokenInfo",
			Handler:    _Query_ParaTokenInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "banksy/transfermiddleware/v1beta1/query.proto",
}

func (m *QueryParaTokenInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParaTokenInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParaTokenInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NativeDenom) > 0 {
		i -= len(m.NativeDenom)
		copy(dAtA[i:], m.NativeDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.NativeDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParaTokenInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParaTokenInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParaTokenInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NativeDenom) > 0 {
		i -= len(m.NativeDenom)
		copy(dAtA[i:], m.NativeDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.NativeDenom)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParaTokenInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NativeDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParaTokenInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.NativeDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParaTokenInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParaTokenInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParaTokenInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParaTokenInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParaTokenInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParaTokenInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
