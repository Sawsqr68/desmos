// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/query_profile.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// QueryProfileRequest is the request type for the Query/Profile RPC method.
type QueryProfileRequest struct {
	// Address or DTag of the user to query the profile for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *QueryProfileRequest) Reset()         { *m = QueryProfileRequest{} }
func (m *QueryProfileRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProfileRequest) ProtoMessage()    {}
func (*QueryProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0e029b5401ec84, []int{0}
}
func (m *QueryProfileRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProfileRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfileRequest.Merge(m, src)
}
func (m *QueryProfileRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfileRequest proto.InternalMessageInfo

// QueryProfileResponse is the response type for the Query/Profile RPC method.
type QueryProfileResponse struct {
	Profile *types.Any `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (m *QueryProfileResponse) Reset()         { *m = QueryProfileResponse{} }
func (m *QueryProfileResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProfileResponse) ProtoMessage()    {}
func (*QueryProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0e029b5401ec84, []int{1}
}
func (m *QueryProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfileResponse.Merge(m, src)
}
func (m *QueryProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfileResponse proto.InternalMessageInfo

func (m *QueryProfileResponse) GetProfile() *types.Any {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryProfileRequest)(nil), "desmos.profiles.v3.QueryProfileRequest")
	proto.RegisterType((*QueryProfileResponse)(nil), "desmos.profiles.v3.QueryProfileResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v3/query_profile.proto", fileDescriptor_7c0e029b5401ec84)
}

var fileDescriptor_7c0e029b5401ec84 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0x49, 0x2d, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0x2d, 0xd6, 0x2f, 0x33, 0xd6, 0x2f,
	0x2c, 0x4d, 0x2d, 0xaa, 0x8c, 0x87, 0x8a, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x41,
	0xd4, 0xe9, 0xc1, 0xd4, 0xe9, 0x95, 0x19, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xa5, 0xf5,
	0x41, 0x2c, 0x88, 0x4a, 0x29, 0xc9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9,
	0x34, 0x4d, 0x3f, 0x31, 0xaf, 0x12, 0x26, 0x95, 0x9c, 0x0f, 0x32, 0x24, 0x1e, 0xa2, 0x07, 0xc2,
	0x81, 0x48, 0x29, 0xf9, 0x72, 0x09, 0x07, 0x82, 0xac, 0x0d, 0x80, 0x98, 0x1f, 0x94, 0x5a, 0x58,
	0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc3, 0xc5, 0x52, 0x5a, 0x9c, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0xe9, 0x24, 0x71, 0x69, 0x8b, 0xae, 0x08, 0x54, 0x9b, 0x63, 0x4a, 0x4a, 0x51, 0x6a, 0x71,
	0x71, 0x70, 0x49, 0x51, 0x66, 0x5e, 0x7a, 0x10, 0x58, 0x95, 0x15, 0x47, 0xc7, 0x02, 0x79, 0x86,
	0x17, 0x0b, 0xe4, 0x19, 0x94, 0x32, 0xb8, 0x44, 0x50, 0x8d, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e,
	0x15, 0x0a, 0xe0, 0x62, 0x87, 0xfa, 0x00, 0x6c, 0x24, 0xb7, 0x91, 0x88, 0x1e, 0xc4, 0xb9, 0x7a,
	0x30, 0xe7, 0xea, 0x39, 0xe6, 0x55, 0x3a, 0x29, 0x9c, 0xda, 0xa2, 0x2b, 0x03, 0xb5, 0x28, 0xb1,
	0xb4, 0x24, 0x43, 0xaf, 0xcc, 0x30, 0x29, 0xb5, 0x24, 0xd1, 0x50, 0xcf, 0x31, 0x39, 0x39, 0xbf,
	0x34, 0xaf, 0xc4, 0x33, 0x08, 0x66, 0x8c, 0x93, 0xf7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e,
	0xcb, 0x31, 0x44, 0x19, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43,
	0x42, 0x4f, 0x37, 0x27, 0x31, 0xa9, 0x18, 0xca, 0xd6, 0x2f, 0x33, 0xd7, 0xaf, 0x40, 0x04, 0x7b,
	0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x15, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc5, 0x99, 0x3e, 0x89, 0x96, 0x01, 0x00, 0x00,
}

func (m *QueryProfileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProfileRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProfileRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryProfile(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Profile != nil {
		{
			size, err := m.Profile.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryProfile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryProfile(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryProfile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryProfileRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryProfile(uint64(l))
	}
	return n
}

func (m *QueryProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Profile != nil {
		l = m.Profile.Size()
		n += 1 + l + sovQueryProfile(uint64(l))
	}
	return n
}

func sovQueryProfile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryProfile(x uint64) (n int) {
	return sovQueryProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProfileRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryProfile
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
			return fmt.Errorf("proto: QueryProfileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProfileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryProfile
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
				return ErrInvalidLengthQueryProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryProfile
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
func (m *QueryProfileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryProfile
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
			return fmt.Errorf("proto: QueryProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryProfile
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
				return ErrInvalidLengthQueryProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Profile == nil {
				m.Profile = &types.Any{}
			}
			if err := m.Profile.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryProfile
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
func skipQueryProfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryProfile
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
					return 0, ErrIntOverflowQueryProfile
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
					return 0, ErrIntOverflowQueryProfile
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
				return 0, ErrInvalidLengthQueryProfile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryProfile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryProfile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryProfile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryProfile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryProfile = fmt.Errorf("proto: unexpected end of group")
)
