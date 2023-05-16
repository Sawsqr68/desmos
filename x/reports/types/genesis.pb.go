// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/reports/v1/genesis.proto

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

// GenesisState defines the reports module's genesis state.
type GenesisState struct {
	SubspacesData []SubspaceDataEntry `protobuf:"bytes,1,rep,name=subspaces_data,json=subspacesData,proto3" json:"subspaces_data"`
	Reasons       []Reason            `protobuf:"bytes,2,rep,name=reasons,proto3" json:"reasons"`
	Reports       []Report            `protobuf:"bytes,3,rep,name=reports,proto3" json:"reports"`
	Params        Params              `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c495c10dc36e45, []int{0}
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

// SubspaceDataEntry contains the data related to a single subspace
type SubspaceDataEntry struct {
	// Id of the subspace to which the data relates
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	// Id of the next reason inside the subspace
	ReasonID uint32 `protobuf:"varint,2,opt,name=reason_id,json=reasonId,proto3" json:"reason_id,omitempty"`
	// Id of the next report inside the subspace
	ReportID uint64 `protobuf:"varint,3,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
}

func (m *SubspaceDataEntry) Reset()         { *m = SubspaceDataEntry{} }
func (m *SubspaceDataEntry) String() string { return proto.CompactTextString(m) }
func (*SubspaceDataEntry) ProtoMessage()    {}
func (*SubspaceDataEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c495c10dc36e45, []int{1}
}
func (m *SubspaceDataEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubspaceDataEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubspaceDataEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubspaceDataEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubspaceDataEntry.Merge(m, src)
}
func (m *SubspaceDataEntry) XXX_Size() int {
	return m.Size()
}
func (m *SubspaceDataEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SubspaceDataEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SubspaceDataEntry proto.InternalMessageInfo

func (m *SubspaceDataEntry) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *SubspaceDataEntry) GetReasonID() uint32 {
	if m != nil {
		return m.ReasonID
	}
	return 0
}

func (m *SubspaceDataEntry) GetReportID() uint64 {
	if m != nil {
		return m.ReportID
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.reports.v1.GenesisState")
	proto.RegisterType((*SubspaceDataEntry)(nil), "desmos.reports.v1.SubspaceDataEntry")
}

func init() { proto.RegisterFile("desmos/reports/v1/genesis.proto", fileDescriptor_e4c495c10dc36e45) }

var fileDescriptor_e4c495c10dc36e45 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x4f, 0xf2, 0x40,
	0x1c, 0xc7, 0x7b, 0xd0, 0xf0, 0xf0, 0x1c, 0x7f, 0x12, 0x9a, 0x67, 0xe8, 0xc3, 0xd0, 0x36, 0xc4,
	0x01, 0x07, 0x7b, 0xa2, 0x31, 0x46, 0xc7, 0x06, 0x63, 0x70, 0xd2, 0xb2, 0xb9, 0x90, 0x2b, 0xbd,
	0x54, 0x12, 0xca, 0x35, 0xbd, 0x83, 0xc8, 0x3b, 0x70, 0x74, 0x72, 0x70, 0xf2, 0xe5, 0x30, 0x32,
	0x3a, 0x11, 0x53, 0xde, 0x88, 0xe9, 0xdd, 0x15, 0x07, 0xd4, 0xed, 0x97, 0xfe, 0x3e, 0xdf, 0x4f,
	0xef, 0x9b, 0xfc, 0xa0, 0x1d, 0x12, 0x16, 0x53, 0x86, 0x52, 0x92, 0xd0, 0x94, 0x33, 0xb4, 0xe8,
	0xa1, 0x88, 0xcc, 0x08, 0x9b, 0x30, 0x37, 0x49, 0x29, 0xa7, 0x46, 0x4b, 0x02, 0xae, 0x02, 0xdc,
	0x45, 0xaf, 0xfd, 0x2f, 0xa2, 0x11, 0x15, 0x5b, 0x94, 0x4f, 0x12, 0x6c, 0x5b, 0xfb, 0xa6, 0x98,
	0x86, 0x64, 0xaa, 0x44, 0x9d, 0x97, 0x12, 0xac, 0x5f, 0x4b, 0xf5, 0x90, 0x63, 0x4e, 0x8c, 0x3b,
	0xd8, 0x64, 0xf3, 0x80, 0x25, 0x78, 0x4c, 0xd8, 0x28, 0xc4, 0x1c, 0x9b, 0xc0, 0x29, 0x77, 0x6b,
	0x27, 0x07, 0xee, 0xde, 0x2f, 0xdd, 0xa1, 0x02, 0xfb, 0x98, 0xe3, 0xab, 0x19, 0x4f, 0x97, 0x9e,
	0xbe, 0xda, 0xd8, 0x9a, 0xdf, 0xd8, 0x19, 0xf2, 0x8d, 0x71, 0x01, 0xff, 0xa4, 0x04, 0x33, 0x3a,
	0x63, 0x66, 0x49, 0xb8, 0xfe, 0x7f, 0xe3, 0xf2, 0x05, 0xa1, 0x04, 0x05, 0x2f, 0xa3, 0x82, 0x31,
	0xcb, 0xbf, 0x44, 0xf3, 0xf1, 0x2b, 0x2a, 0x16, 0xc6, 0x39, 0xac, 0x24, 0x38, 0xc5, 0x31, 0x33,
	0x75, 0x07, 0xfc, 0x90, 0xbc, 0x15, 0x80, 0x4a, 0x2a, 0xfc, 0x52, 0x7f, 0x7a, 0xb3, 0xb5, 0xce,
	0x2b, 0x80, 0xad, 0xbd, 0x7e, 0x06, 0x82, 0xb5, 0xa2, 0xdb, 0x68, 0x12, 0x9a, 0xc0, 0x01, 0x5d,
	0xdd, 0x6b, 0x66, 0x1b, 0x1b, 0x16, 0xec, 0xa0, 0xef, 0xc3, 0x02, 0x19, 0x84, 0xc6, 0x21, 0xfc,
	0x2b, 0xbb, 0xe4, 0x78, 0xc9, 0x01, 0xdd, 0x86, 0x57, 0xcf, 0x36, 0x76, 0x55, 0xd6, 0x1d, 0xf4,
	0xfd, 0xaa, 0x5c, 0x17, 0x68, 0xfe, 0xb4, 0x1c, 0x2d, 0x0b, 0xb3, 0x42, 0xf3, 0x8f, 0x12, 0x15,
	0x53, 0xe8, 0xdd, 0xac, 0x32, 0x0b, 0xac, 0x33, 0x0b, 0x7c, 0x64, 0x16, 0x78, 0xde, 0x5a, 0xda,
	0x7a, 0x6b, 0x69, 0xef, 0x5b, 0x4b, 0xbb, 0x3f, 0x8e, 0x26, 0xfc, 0x61, 0x1e, 0xb8, 0x63, 0x1a,
	0x23, 0xd9, 0xf7, 0x68, 0x8a, 0x03, 0xa6, 0x66, 0xb4, 0x38, 0x43, 0x8f, 0xbb, 0x5b, 0xe0, 0xcb,
	0x84, 0xb0, 0xa0, 0x22, 0x0e, 0xe1, 0xf4, 0x33, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x91, 0x25, 0x83,
	0x74, 0x02, 0x00, 0x00,
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Reports) > 0 {
		for iNdEx := len(m.Reports) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reports[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Reasons) > 0 {
		for iNdEx := len(m.Reasons) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reasons[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.SubspacesData) > 0 {
		for iNdEx := len(m.SubspacesData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubspacesData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SubspaceDataEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubspaceDataEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubspaceDataEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReportID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ReportID))
		i--
		dAtA[i] = 0x18
	}
	if m.ReasonID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ReasonID))
		i--
		dAtA[i] = 0x10
	}
	if m.SubspaceID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
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
	if len(m.SubspacesData) > 0 {
		for _, e := range m.SubspacesData {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Reasons) > 0 {
		for _, e := range m.Reasons {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Reports) > 0 {
		for _, e := range m.Reports {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *SubspaceDataEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.ReasonID != 0 {
		n += 1 + sovGenesis(uint64(m.ReasonID))
	}
	if m.ReportID != 0 {
		n += 1 + sovGenesis(uint64(m.ReportID))
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
				return fmt.Errorf("proto: wrong wireType = %d for field SubspacesData", wireType)
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
			m.SubspacesData = append(m.SubspacesData, SubspaceDataEntry{})
			if err := m.SubspacesData[len(m.SubspacesData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reasons", wireType)
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
			m.Reasons = append(m.Reasons, Reason{})
			if err := m.Reasons[len(m.Reasons)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reports", wireType)
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
			m.Reports = append(m.Reports, Report{})
			if err := m.Reports[len(m.Reports)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
func (m *SubspaceDataEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SubspaceDataEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubspaceDataEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReasonID", wireType)
			}
			m.ReasonID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReasonID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportID", wireType)
			}
			m.ReportID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReportID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
