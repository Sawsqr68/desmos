// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/msgs_profile.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// MsgSaveProfile represents a message to save a profile.
type MsgSaveProfile struct {
	// DTag of the profile. If it shouldn't be changed, [do-no-modify] can be used
	// instead.
	DTag string `protobuf:"bytes,1,opt,name=dtag,proto3" json:"dtag,omitempty" yaml:"dtag"`
	// Nickname of the profile. If it shouldn't be changed, [do-no-modify] can be
	// used instead.
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty" yaml:"nickname"`
	// Bio of the profile. If it shouldn't be changed, [do-no-modify] can be used
	// instead.
	Bio string `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty" yaml:"bio"`
	// URL to the profile picture. If it shouldn't be changed, [do-no-modify] can
	// be used instead.
	ProfilePicture string `protobuf:"bytes,4,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty" yaml:"profile_picture"`
	// URL to the profile cover. If it shouldn't be changed, [do-no-modify] can be
	// used instead.
	CoverPicture string `protobuf:"bytes,5,opt,name=cover_picture,json=coverPicture,proto3" json:"cover_picture,omitempty" yaml:"cover_picture"`
	// Address of the user associated to the profile
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *MsgSaveProfile) Reset()         { *m = MsgSaveProfile{} }
func (m *MsgSaveProfile) String() string { return proto.CompactTextString(m) }
func (*MsgSaveProfile) ProtoMessage()    {}
func (*MsgSaveProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea75cf4ca5bb3a3, []int{0}
}
func (m *MsgSaveProfile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveProfile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveProfile.Merge(m, src)
}
func (m *MsgSaveProfile) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveProfile.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveProfile proto.InternalMessageInfo

func (m *MsgSaveProfile) GetDTag() string {
	if m != nil {
		return m.DTag
	}
	return ""
}

func (m *MsgSaveProfile) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *MsgSaveProfile) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *MsgSaveProfile) GetProfilePicture() string {
	if m != nil {
		return m.ProfilePicture
	}
	return ""
}

func (m *MsgSaveProfile) GetCoverPicture() string {
	if m != nil {
		return m.CoverPicture
	}
	return ""
}

func (m *MsgSaveProfile) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// MsgSaveProfileResponse defines the Msg/SaveProfile response type.
type MsgSaveProfileResponse struct {
}

func (m *MsgSaveProfileResponse) Reset()         { *m = MsgSaveProfileResponse{} }
func (m *MsgSaveProfileResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSaveProfileResponse) ProtoMessage()    {}
func (*MsgSaveProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea75cf4ca5bb3a3, []int{1}
}
func (m *MsgSaveProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveProfileResponse.Merge(m, src)
}
func (m *MsgSaveProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveProfileResponse proto.InternalMessageInfo

// MsgDeleteProfile represents the message used to delete an existing profile.
type MsgDeleteProfile struct {
	// Address associated to the profile to be deleted
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *MsgDeleteProfile) Reset()         { *m = MsgDeleteProfile{} }
func (m *MsgDeleteProfile) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProfile) ProtoMessage()    {}
func (*MsgDeleteProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea75cf4ca5bb3a3, []int{2}
}
func (m *MsgDeleteProfile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProfile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProfile.Merge(m, src)
}
func (m *MsgDeleteProfile) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProfile.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProfile proto.InternalMessageInfo

func (m *MsgDeleteProfile) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// MsgDeleteProfileResponse defines the Msg/DeleteProfile response type.
type MsgDeleteProfileResponse struct {
}

func (m *MsgDeleteProfileResponse) Reset()         { *m = MsgDeleteProfileResponse{} }
func (m *MsgDeleteProfileResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProfileResponse) ProtoMessage()    {}
func (*MsgDeleteProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea75cf4ca5bb3a3, []int{3}
}
func (m *MsgDeleteProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProfileResponse.Merge(m, src)
}
func (m *MsgDeleteProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProfileResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSaveProfile)(nil), "desmos.profiles.v3.MsgSaveProfile")
	proto.RegisterType((*MsgSaveProfileResponse)(nil), "desmos.profiles.v3.MsgSaveProfileResponse")
	proto.RegisterType((*MsgDeleteProfile)(nil), "desmos.profiles.v3.MsgDeleteProfile")
	proto.RegisterType((*MsgDeleteProfileResponse)(nil), "desmos.profiles.v3.MsgDeleteProfileResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v3/msgs_profile.proto", fileDescriptor_5ea75cf4ca5bb3a3)
}

var fileDescriptor_5ea75cf4ca5bb3a3 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x31, 0x6f, 0x13, 0x31,
	0x18, 0x86, 0x73, 0x24, 0x14, 0x30, 0x90, 0xc2, 0x11, 0x1a, 0x93, 0xe1, 0xae, 0xb2, 0x04, 0xaa,
	0x82, 0x1a, 0xab, 0xca, 0x80, 0x54, 0x89, 0x81, 0x92, 0x0d, 0x55, 0xaa, 0xae, 0x4c, 0x2c, 0x91,
	0xef, 0x62, 0x8c, 0x45, 0xee, 0x7c, 0xb2, 0xdd, 0x13, 0x5d, 0x33, 0x32, 0xf1, 0x53, 0x3a, 0xf0,
	0x23, 0x18, 0x2b, 0x26, 0xa6, 0x13, 0x4a, 0x86, 0xee, 0xf7, 0x0b, 0x50, 0x6c, 0x5f, 0x42, 0x32,
	0x76, 0x49, 0xec, 0xf7, 0x7d, 0xde, 0xef, 0xfb, 0xce, 0xe7, 0x03, 0x2f, 0x27, 0x54, 0xa5, 0x42,
	0xe1, 0x5c, 0x8a, 0xcf, 0x7c, 0x4a, 0x15, 0x2e, 0x86, 0x38, 0x55, 0x4c, 0x8d, 0x9d, 0x30, 0xc8,
	0xa5, 0xd0, 0xc2, 0xf7, 0x2d, 0x36, 0xa8, 0xb1, 0x41, 0x31, 0xec, 0x3d, 0x25, 0x29, 0xcf, 0x04,
	0x36, 0xbf, 0x16, 0xeb, 0x75, 0x98, 0x60, 0xc2, 0x2c, 0xf1, 0x72, 0xe5, 0xd4, 0x6e, 0x22, 0x4c,
	0x8f, 0x54, 0x31, 0x5c, 0x1c, 0x2d, 0xff, 0x9c, 0xf1, 0xc2, 0x1a, 0x63, 0x9b, 0xb0, 0x1b, 0x6b,
	0xa1, 0x59, 0x13, 0xb4, 0x4f, 0x15, 0x3b, 0x27, 0x05, 0x3d, 0xb3, 0x3d, 0xfd, 0xd7, 0xa0, 0x35,
	0xd1, 0x84, 0x41, 0x6f, 0xdf, 0x3b, 0x78, 0x70, 0xd2, 0x9d, 0x97, 0x61, 0x6b, 0xf4, 0x91, 0xb0,
	0xaa, 0x0c, 0x1f, 0x5e, 0x92, 0x74, 0x7a, 0x8c, 0x96, 0x2e, 0x8a, 0x0c, 0xe4, 0x63, 0x70, 0x3f,
	0xe3, 0xc9, 0xd7, 0x8c, 0xa4, 0x14, 0xde, 0x31, 0x81, 0x67, 0x55, 0x19, 0xee, 0x5a, 0xb0, 0x76,
	0x50, 0xb4, 0x82, 0xfc, 0x7d, 0xd0, 0x8c, 0xb9, 0x80, 0x4d, 0xc3, 0xb6, 0xab, 0x32, 0x04, 0x96,
	0x8d, 0xb9, 0x40, 0xd1, 0xd2, 0xf2, 0xdf, 0x83, 0x5d, 0xf7, 0xf8, 0xe3, 0x9c, 0x27, 0xfa, 0x42,
	0x52, 0xd8, 0x32, 0x74, 0xaf, 0x2a, 0xc3, 0x3d, 0x4b, 0x6f, 0x01, 0x28, 0x6a, 0x3b, 0xe5, 0xcc,
	0x0a, 0xfe, 0x5b, 0xf0, 0x38, 0x11, 0x05, 0x95, 0xab, 0x12, 0x77, 0x4d, 0x09, 0x58, 0x95, 0x61,
	0xc7, 0x96, 0xd8, 0xb0, 0x51, 0xf4, 0xc8, 0xec, 0xeb, 0xf8, 0x08, 0xdc, 0x4b, 0x24, 0x25, 0x5a,
	0x48, 0xb8, 0x63, 0x82, 0xfd, 0xaa, 0x0c, 0xdb, 0x2e, 0x68, 0x0d, 0xf4, 0xfb, 0xe7, 0x61, 0xc7,
	0x9d, 0xe5, 0xbb, 0xc9, 0x44, 0x52, 0xa5, 0xce, 0xb5, 0xe4, 0x19, 0x8b, 0xea, 0xe8, 0xf1, 0xab,
	0xd9, 0xcd, 0x55, 0xbf, 0xde, 0x7d, 0xbf, 0xb9, 0xea, 0x3f, 0x77, 0xb7, 0x60, 0xf3, 0xc4, 0x11,
	0x04, 0x7b, 0x9b, 0x4a, 0x44, 0x55, 0x2e, 0x32, 0x45, 0xd1, 0xcc, 0x03, 0x4f, 0x4e, 0x15, 0x1b,
	0xd1, 0x29, 0xd5, 0xab, 0x17, 0xf4, 0xdf, 0x70, 0xde, 0xed, 0x87, 0x3b, 0xd8, 0x1e, 0xae, 0xbb,
	0x1e, 0x6e, 0xa3, 0x1f, 0xea, 0x01, 0xb8, 0xad, 0xd5, 0x03, 0x9e, 0x7c, 0xf8, 0x35, 0x0f, 0xbc,
	0xeb, 0x79, 0xe0, 0xfd, 0x9d, 0x07, 0xde, 0x8f, 0x45, 0xd0, 0xb8, 0x5e, 0x04, 0x8d, 0x3f, 0x8b,
	0xa0, 0xf1, 0xe9, 0x88, 0x71, 0xfd, 0xe5, 0x22, 0x1e, 0x24, 0x22, 0xc5, 0xb6, 0xf2, 0xe1, 0x94,
	0xc4, 0xca, 0xad, 0x71, 0xf1, 0x06, 0x7f, 0x5b, 0x7f, 0x0d, 0xfa, 0x32, 0xa7, 0x2a, 0xde, 0x31,
	0x77, 0x72, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x1c, 0xce, 0xb3, 0x2d, 0x03, 0x00, 0x00,
}

func (m *MsgSaveProfile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveProfile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveProfile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CoverPicture) > 0 {
		i -= len(m.CoverPicture)
		copy(dAtA[i:], m.CoverPicture)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.CoverPicture)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ProfilePicture) > 0 {
		i -= len(m.ProfilePicture)
		copy(dAtA[i:], m.ProfilePicture)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.ProfilePicture)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bio) > 0 {
		i -= len(m.Bio)
		copy(dAtA[i:], m.Bio)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Bio)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DTag) > 0 {
		i -= len(m.DTag)
		copy(dAtA[i:], m.DTag)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.DTag)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSaveProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProfile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProfile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProfile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsProfile(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsProfile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSaveProfile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DTag)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Bio)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.ProfilePicture)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.CoverPicture)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	return n
}

func (m *MsgSaveProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteProfile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	return n
}

func (m *MsgDeleteProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsProfile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsProfile(x uint64) (n int) {
	return sovMsgsProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSaveProfile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgSaveProfile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveProfile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DTag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfilePicture", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfilePicture = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoverPicture", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoverPicture = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgSaveProfileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgSaveProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgDeleteProfile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgDeleteProfile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProfile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgDeleteProfileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgDeleteProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func skipMsgsProfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsProfile
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
					return 0, ErrIntOverflowMsgsProfile
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
					return 0, ErrIntOverflowMsgsProfile
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
				return 0, ErrInvalidLengthMsgsProfile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsProfile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsProfile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsProfile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsProfile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsProfile = fmt.Errorf("proto: unexpected end of group")
)
