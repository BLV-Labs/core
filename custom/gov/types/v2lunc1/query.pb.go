// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/gov/v2lunc1/query.proto

package v2lunc1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/types"
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

// QueryProposalRequest is the request type for the Query/Proposal RPC method.
type QueryProposalRequest struct {
	// proposal_id defines the unique id of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *QueryProposalRequest) Reset()         { *m = QueryProposalRequest{} }
func (m *QueryProposalRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProposalRequest) ProtoMessage()    {}
func (*QueryProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_075d72aa1a1cda58, []int{0}
}
func (m *QueryProposalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProposalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProposalRequest.Merge(m, src)
}
func (m *QueryProposalRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProposalRequest proto.InternalMessageInfo

func (m *QueryProposalRequest) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

// QueryVoteRequest is the request type for the Query/Vote RPC method.
type QueryMinimalDepositProposalResponse struct {
	// proposal_id defines the unique id of the proposal.
	MinimalDeposit types.Coin `protobuf:"bytes,1,opt,name=minimal_deposit,json=minimalDeposit,proto3" json:"minimal_deposit"`
}

func (m *QueryMinimalDepositProposalResponse) Reset()         { *m = QueryMinimalDepositProposalResponse{} }
func (m *QueryMinimalDepositProposalResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMinimalDepositProposalResponse) ProtoMessage()    {}
func (*QueryMinimalDepositProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_075d72aa1a1cda58, []int{1}
}
func (m *QueryMinimalDepositProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinimalDepositProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinimalDepositProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinimalDepositProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinimalDepositProposalResponse.Merge(m, src)
}
func (m *QueryMinimalDepositProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinimalDepositProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinimalDepositProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinimalDepositProposalResponse proto.InternalMessageInfo

func (m *QueryMinimalDepositProposalResponse) GetMinimalDeposit() types.Coin {
	if m != nil {
		return m.MinimalDeposit
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*QueryProposalRequest)(nil), "terra.gov.v2lunc1.QueryProposalRequest")
	proto.RegisterType((*QueryMinimalDepositProposalResponse)(nil), "terra.gov.v2lunc1.QueryMinimalDepositProposalResponse")
}

func init() { proto.RegisterFile("terra/gov/v2lunc1/query.proto", fileDescriptor_075d72aa1a1cda58) }

var fileDescriptor_075d72aa1a1cda58 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0xce, 0x88, 0x0a, 0xce, 0x82, 0xb2, 0xc3, 0x1e, 0x4c, 0xd0, 0x89, 0xc4, 0x83, 0x22, 0x6c,
	0x17, 0xc9, 0x82, 0xcb, 0x9e, 0x84, 0xac, 0x17, 0xc1, 0x15, 0x8d, 0xec, 0xc5, 0x4b, 0xe8, 0x99,
	0x69, 0xc7, 0x86, 0xe9, 0xae, 0xd9, 0xa9, 0xee, 0x81, 0x20, 0x5e, 0x7c, 0x02, 0xc1, 0x97, 0xf0,
	0xe8, 0x0b, 0x78, 0x76, 0x8f, 0x0b, 0x5e, 0x3c, 0x89, 0x24, 0x82, 0xaf, 0x21, 0xe9, 0xee, 0x41,
	0x8d, 0x3f, 0x97, 0xa6, 0xbb, 0xbe, 0xaf, 0xaa, 0xbe, 0xfa, 0xba, 0xe2, 0xeb, 0x46, 0x34, 0x0d,
	0x87, 0x12, 0x5b, 0x68, 0x27, 0x95, 0xd5, 0xf9, 0x18, 0x4e, 0xac, 0x68, 0x16, 0xac, 0x6e, 0xd0,
	0x60, 0xb2, 0xed, 0x60, 0x56, 0x62, 0xcb, 0x02, 0x3c, 0x48, 0x73, 0x24, 0x85, 0x04, 0x19, 0x27,
	0x01, 0xed, 0x38, 0x13, 0x86, 0x8f, 0x21, 0x47, 0xa9, 0x7d, 0xca, 0x60, 0xa7, 0xc4, 0x12, 0xdd,
	0x15, 0xd6, 0xb7, 0x10, 0xbd, 0x56, 0x22, 0x96, 0x95, 0x00, 0x5e, 0x4b, 0xe0, 0x5a, 0xa3, 0xe1,
	0x46, 0xa2, 0xa6, 0x80, 0x0e, 0x03, 0xea, 0x5e, 0x99, 0x7d, 0x0e, 0x46, 0x2a, 0x41, 0x86, 0xab,
	0x3a, 0x10, 0xfa, 0x9b, 0x04, 0xae, 0x83, 0xc4, 0x41, 0xba, 0x09, 0x15, 0xb6, 0x71, 0xc5, 0xbb,
	0x54, 0xaf, 0x77, 0xee, 0x25, 0xf9, 0x47, 0x80, 0xb6, 0xb9, 0x92, 0x1a, 0xc1, 0x9d, 0x3e, 0x34,
	0xda, 0x8f, 0x77, 0x9e, 0xac, 0xe7, 0x7f, 0xdc, 0x60, 0x8d, 0xc4, 0xab, 0x99, 0x38, 0xb1, 0x82,
	0x4c, 0x32, 0x8c, 0xb7, 0xea, 0x10, 0x9a, 0xcb, 0xe2, 0x6a, 0x74, 0x23, 0xba, 0x7d, 0x7e, 0x16,
	0x77, 0xa1, 0x07, 0xc5, 0xc8, 0xc4, 0x37, 0x5d, 0xe2, 0x91, 0xd4, 0x52, 0xf1, 0xea, 0xbe, 0xa8,
	0x91, 0xa4, 0xf9, 0x59, 0x86, 0x6a, 0xd4, 0x24, 0x92, 0xa3, 0xf8, 0x8a, 0xf2, 0x8c, 0x79, 0xe1,
	0x29, 0xae, 0xd6, 0xd6, 0xa4, 0xcf, 0x82, 0xb4, 0xb5, 0xaf, 0x2c, 0xf8, 0xca, 0x0e, 0x51, 0xea,
	0xe9, 0xa5, 0xd3, 0x2f, 0xc3, 0xde, 0xbb, 0xef, 0xef, 0xef, 0x44, 0xb3, 0xcb, 0xea, 0xb7, 0xf2,
	0x93, 0x8f, 0x51, 0x7c, 0xc1, 0xb5, 0x4d, 0x3e, 0x44, 0x71, 0xbf, 0xeb, 0x16, 0x34, 0x3c, 0x3c,
	0x7e, 0x74, 0x38, 0x5d, 0x1c, 0x5b, 0x2a, 0x92, 0x5b, 0xec, 0x8f, 0x8f, 0x64, 0x7f, 0x9b, 0x73,
	0x70, 0xf7, 0x5f, 0xc4, 0xff, 0xcf, 0x35, 0xba, 0xf7, 0xfa, 0xd3, 0xb7, 0xb7, 0xe7, 0x0e, 0x92,
	0xfd, 0xe0, 0xb0, 0xdf, 0xa8, 0x31, 0x74, 0x0e, 0x11, 0x04, 0xe5, 0xbb, 0x61, 0x6c, 0x82, 0x97,
	0xbf, 0x18, 0xfa, 0x6a, 0xfa, 0xf4, 0x74, 0x99, 0x46, 0x67, 0xcb, 0x34, 0xfa, 0xba, 0x4c, 0xa3,
	0x37, 0xab, 0xb4, 0x77, 0xb6, 0x4a, 0x7b, 0x9f, 0x57, 0x69, 0xef, 0xd9, 0x41, 0x29, 0xcd, 0x0b,
	0x9b, 0xb1, 0x1c, 0x15, 0xe4, 0x15, 0x27, 0x92, 0xf9, 0xae, 0xdf, 0xda, 0x1c, 0x1b, 0x01, 0xed,
	0x1e, 0xe4, 0x96, 0x0c, 0x2a, 0xd7, 0xd2, 0x2c, 0x6a, 0x41, 0xdd, 0x2a, 0x67, 0x17, 0xdd, 0xa7,
	0xee, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0x87, 0x87, 0x74, 0x01, 0xe6, 0x02, 0x00, 0x00,
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
	// Proposal queries proposal details based on ProposalID.
	ProposalMinimalLUNCByUusd(ctx context.Context, in *QueryProposalRequest, opts ...grpc.CallOption) (*QueryMinimalDepositProposalResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ProposalMinimalLUNCByUusd(ctx context.Context, in *QueryProposalRequest, opts ...grpc.CallOption) (*QueryMinimalDepositProposalResponse, error) {
	out := new(QueryMinimalDepositProposalResponse)
	err := c.cc.Invoke(ctx, "/terra.gov.v2lunc1.Query/ProposalMinimalLUNCByUusd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Proposal queries proposal details based on ProposalID.
	ProposalMinimalLUNCByUusd(context.Context, *QueryProposalRequest) (*QueryMinimalDepositProposalResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ProposalMinimalLUNCByUusd(ctx context.Context, req *QueryProposalRequest) (*QueryMinimalDepositProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposalMinimalLUNCByUusd not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ProposalMinimalLUNCByUusd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProposalMinimalLUNCByUusd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.gov.v2lunc1.Query/ProposalMinimalLUNCByUusd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProposalMinimalLUNCByUusd(ctx, req.(*QueryProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terra.gov.v2lunc1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProposalMinimalLUNCByUusd",
			Handler:    _Query_ProposalMinimalLUNCByUusd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terra/gov/v2lunc1/query.proto",
}

func (m *QueryProposalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProposalRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProposalRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryMinimalDepositProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinimalDepositProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinimalDepositProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinimalDeposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *QueryProposalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	return n
}

func (m *QueryMinimalDepositProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinimalDeposit.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProposalRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryProposalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProposalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryMinimalDepositProposalResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMinimalDepositProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinimalDepositProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimalDeposit", wireType)
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
			if err := m.MinimalDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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