// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pooltoy/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type QueryListUsersRequest struct {
}

func (m *QueryListUsersRequest) Reset()         { *m = QueryListUsersRequest{} }
func (m *QueryListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryListUsersRequest) ProtoMessage()    {}
func (*QueryListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc5dcdc78c1b6fd, []int{0}
}
func (m *QueryListUsersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListUsersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListUsersRequest.Merge(m, src)
}
func (m *QueryListUsersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListUsersRequest proto.InternalMessageInfo

type QueryListUsersResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (m *QueryListUsersResponse) Reset()         { *m = QueryListUsersResponse{} }
func (m *QueryListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryListUsersResponse) ProtoMessage()    {}
func (*QueryListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc5dcdc78c1b6fd, []int{1}
}
func (m *QueryListUsersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListUsersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListUsersResponse.Merge(m, src)
}
func (m *QueryListUsersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListUsersResponse proto.InternalMessageInfo

func (m *QueryListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryListUsersRequest)(nil), "pooltoy.v1beta1.QueryListUsersRequest")
	proto.RegisterType((*QueryListUsersResponse)(nil), "pooltoy.v1beta1.QueryListUsersResponse")
}

func init() { proto.RegisterFile("pooltoy/v1beta1/query.proto", fileDescriptor_4dc5dcdc78c1b6fd) }

var fileDescriptor_4dc5dcdc78c1b6fd = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x9b, 0xff, 0xa7, 0x2e, 0x46, 0xb4, 0x10, 0xac, 0xda, 0x28, 0x83, 0x64, 0xa1, 0x82,
	0x90, 0xa1, 0xed, 0x1b, 0x08, 0xee, 0x74, 0x61, 0xc1, 0x8d, 0x1b, 0x99, 0x84, 0x31, 0x1d, 0x48,
	0xe7, 0xa6, 0x73, 0x27, 0xd2, 0x2c, 0xf5, 0x09, 0x04, 0x5f, 0xca, 0x65, 0xc1, 0x8d, 0x4b, 0x49,
	0x7c, 0x10, 0x49, 0x26, 0x2d, 0x18, 0x05, 0x77, 0xf7, 0xe6, 0x3b, 0xe7, 0xe4, 0xdc, 0x21, 0x07,
	0x29, 0x40, 0x62, 0x20, 0x67, 0x0f, 0xc3, 0x50, 0x18, 0x3e, 0x64, 0xf3, 0x4c, 0xe8, 0x3c, 0x48,
	0x35, 0x18, 0x70, 0x7b, 0x0d, 0x0c, 0x1a, 0xe8, 0x0d, 0x22, 0xc0, 0x19, 0xe0, 0x5d, 0x8d, 0x99,
	0x5d, 0xac, 0xd6, 0x3b, 0x8c, 0x01, 0xe2, 0x44, 0x30, 0x9e, 0x4a, 0xc6, 0x95, 0x02, 0xc3, 0x8d,
	0x04, 0xb5, 0xa2, 0x3b, 0x31, 0xc4, 0x60, 0x5d, 0xd5, 0xd4, 0x7c, 0x1d, 0x34, 0x9e, 0x7a, 0x0b,
	0xb3, 0x7b, 0xc6, 0x55, 0xf3, 0x6b, 0xcf, 0x6b, 0xf7, 0xca, 0x50, 0x68, 0xcb, 0xfc, 0x3d, 0xd2,
	0xbf, 0xae, 0x5a, 0x5e, 0x4a, 0x34, 0x37, 0x28, 0x34, 0x4e, 0xc4, 0x3c, 0x13, 0x68, 0xfc, 0x0b,
	0xb2, 0xdb, 0x06, 0x98, 0x82, 0x42, 0xe1, 0x9e, 0x91, 0x6e, 0x15, 0x80, 0xfb, 0xce, 0xd1, 0xff,
	0xd3, 0xcd, 0x51, 0x3f, 0x68, 0x5d, 0x16, 0x54, 0xf2, 0x89, 0xd5, 0x8c, 0x1e, 0x1d, 0xd2, 0xad,
	0x73, 0xdc, 0x05, 0xd9, 0xfe, 0x1e, 0xe8, 0x1e, 0xff, 0x70, 0xfe, 0x5a, 0xc5, 0x3b, 0xf9, 0x53,
	0x67, 0x9b, 0xf9, 0xfd, 0xa7, 0xb7, 0xcf, 0x97, 0x7f, 0x3d, 0x77, 0x8b, 0xad, 0x2e, 0xae, 0x4a,
	0x9c, 0x5f, 0xbd, 0x16, 0xd4, 0x59, 0x16, 0xd4, 0xf9, 0x28, 0xa8, 0xf3, 0x5c, 0xd2, 0xce, 0xb2,
	0xa4, 0x9d, 0xf7, 0x92, 0x76, 0x6e, 0xc7, 0xb1, 0x34, 0xd3, 0x2c, 0x0c, 0x22, 0x98, 0x31, 0xa9,
	0x8c, 0xd0, 0xd1, 0x94, 0x4b, 0x15, 0x0a, 0x9d, 0x48, 0xb5, 0xce, 0x58, 0xac, 0x27, 0x93, 0xa7,
	0x02, 0xc3, 0x8d, 0xfa, 0xe5, 0xc6, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x1a, 0x58, 0xf6,
	0xef, 0x01, 0x00, 0x00,
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
	QueryListUsers(ctx context.Context, in *QueryListUsersRequest, opts ...grpc.CallOption) (*QueryListUsersResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) QueryListUsers(ctx context.Context, in *QueryListUsersRequest, opts ...grpc.CallOption) (*QueryListUsersResponse, error) {
	out := new(QueryListUsersResponse)
	err := c.cc.Invoke(ctx, "/pooltoy.v1beta1.Query/QueryListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	QueryListUsers(context.Context, *QueryListUsersRequest) (*QueryListUsersResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) QueryListUsers(ctx context.Context, req *QueryListUsersRequest) (*QueryListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryListUsers not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_QueryListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pooltoy.v1beta1.Query/QueryListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryListUsers(ctx, req.(*QueryListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pooltoy.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryListUsers",
			Handler:    _Query_QueryListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pooltoy/v1beta1/query.proto",
}

func (m *QueryListUsersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListUsersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListUsersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryListUsersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListUsersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListUsersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Users) > 0 {
		for iNdEx := len(m.Users) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Users[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryListUsersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryListUsersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryListUsersRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryListUsersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListUsersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryListUsersResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryListUsersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListUsersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &User{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
