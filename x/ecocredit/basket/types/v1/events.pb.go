// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/basket/v1/events.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// EventCreate is an event emitted when a basket is created.
type EventCreate struct {
	// basket_denom is the basket bank denom.
	BasketDenom string `protobuf:"bytes,1,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// curator is the address of the basket curator who is able to change certain
	// basket settings.
	//
	// Deprecated (Since Revision 1): This field is still populated and will be
	// removed in the next version.
	Curator string `protobuf:"bytes,2,opt,name=curator,proto3" json:"curator,omitempty"` // Deprecated: Do not use.
}

func (m *EventCreate) Reset()         { *m = EventCreate{} }
func (m *EventCreate) String() string { return proto.CompactTextString(m) }
func (*EventCreate) ProtoMessage()    {}
func (*EventCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc7fc2fbcbd93cbc, []int{0}
}
func (m *EventCreate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreate.Merge(m, src)
}
func (m *EventCreate) XXX_Size() int {
	return m.Size()
}
func (m *EventCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreate.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreate proto.InternalMessageInfo

func (m *EventCreate) GetBasketDenom() string {
	if m != nil {
		return m.BasketDenom
	}
	return ""
}

// Deprecated: Do not use.
func (m *EventCreate) GetCurator() string {
	if m != nil {
		return m.Curator
	}
	return ""
}

// EventPut is an event emitted when credits are put into a basket in return for
// basket tokens.
type EventPut struct {
	// owner is the owner of the credits put into the basket.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// basket_denom is the basket bank denom that the credits were added to.
	BasketDenom string `protobuf:"bytes,2,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// credits are the credits that were added to the basket.
	//
	// Deprecated (Since Revision 1): This field is still populated and will be
	// removed in the next version.
	Credits []*BasketCredit `protobuf:"bytes,3,rep,name=credits,proto3" json:"credits,omitempty"` // Deprecated: Do not use.
	// amount is the integer number of basket tokens converted from credits.
	//
	// Deprecated (Since Revision 1): This field is still populated and will be
	// removed in the next version.
	Amount string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"` // Deprecated: Do not use.
}

func (m *EventPut) Reset()         { *m = EventPut{} }
func (m *EventPut) String() string { return proto.CompactTextString(m) }
func (*EventPut) ProtoMessage()    {}
func (*EventPut) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc7fc2fbcbd93cbc, []int{1}
}
func (m *EventPut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPut.Merge(m, src)
}
func (m *EventPut) XXX_Size() int {
	return m.Size()
}
func (m *EventPut) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPut.DiscardUnknown(m)
}

var xxx_messageInfo_EventPut proto.InternalMessageInfo

func (m *EventPut) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *EventPut) GetBasketDenom() string {
	if m != nil {
		return m.BasketDenom
	}
	return ""
}

// Deprecated: Do not use.
func (m *EventPut) GetCredits() []*BasketCredit {
	if m != nil {
		return m.Credits
	}
	return nil
}

// Deprecated: Do not use.
func (m *EventPut) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

// EventTake is an event emitted when credits are taken from a basket starting
// from the oldest credits first.
type EventTake struct {
	// owner is the owner of the credits taken from the basket.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// basket_denom is the basket bank denom that credits were taken from.
	BasketDenom string `protobuf:"bytes,2,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// credits are the credits that were taken from the basket.
	//
	// Deprecated (Since Revision 1): This field is still populated and will be
	// removed in the next version.
	Credits []*BasketCredit `protobuf:"bytes,3,rep,name=credits,proto3" json:"credits,omitempty"` // Deprecated: Do not use.
	// amount is the integer number of basket tokens converted to credits.
	//
	// Deprecated (Since Revision 1): This field is still populated and will be
	// removed in the next version.
	Amount string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"` // Deprecated: Do not use.
}

func (m *EventTake) Reset()         { *m = EventTake{} }
func (m *EventTake) String() string { return proto.CompactTextString(m) }
func (*EventTake) ProtoMessage()    {}
func (*EventTake) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc7fc2fbcbd93cbc, []int{2}
}
func (m *EventTake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTake.Merge(m, src)
}
func (m *EventTake) XXX_Size() int {
	return m.Size()
}
func (m *EventTake) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTake.DiscardUnknown(m)
}

var xxx_messageInfo_EventTake proto.InternalMessageInfo

func (m *EventTake) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *EventTake) GetBasketDenom() string {
	if m != nil {
		return m.BasketDenom
	}
	return ""
}

// Deprecated: Do not use.
func (m *EventTake) GetCredits() []*BasketCredit {
	if m != nil {
		return m.Credits
	}
	return nil
}

// Deprecated: Do not use.
func (m *EventTake) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCreate)(nil), "regen.ecocredit.basket.v1.EventCreate")
	proto.RegisterType((*EventPut)(nil), "regen.ecocredit.basket.v1.EventPut")
	proto.RegisterType((*EventTake)(nil), "regen.ecocredit.basket.v1.EventTake")
}

func init() {
	proto.RegisterFile("regen/ecocredit/basket/v1/events.proto", fileDescriptor_bc7fc2fbcbd93cbc)
}

var fileDescriptor_bc7fc2fbcbd93cbc = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xeb, 0xf6, 0xfb, 0x5a, 0xea, 0x32, 0x45, 0x0c, 0xa1, 0x42, 0x56, 0x89, 0x04, 0x74,
	0xc1, 0x56, 0xe0, 0x09, 0x48, 0xe9, 0x8a, 0x50, 0xc5, 0x04, 0x03, 0xca, 0x9f, 0xab, 0x50, 0x85,
	0xd8, 0x95, 0xe3, 0xa4, 0xf0, 0x16, 0x3c, 0x05, 0xbc, 0x0a, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08,
	0x8a, 0x1d, 0xba, 0x44, 0xdd, 0xd9, 0x7c, 0x8f, 0x8f, 0xcf, 0xfd, 0x59, 0x3a, 0xf8, 0x54, 0x42,
	0x0c, 0x9c, 0x41, 0x28, 0x42, 0x09, 0xd1, 0x52, 0xb1, 0xc0, 0xcf, 0x12, 0x50, 0xac, 0x70, 0x19,
	0x14, 0xc0, 0x55, 0x46, 0x57, 0x52, 0x28, 0x61, 0x1d, 0x6a, 0x1f, 0xdd, 0xfa, 0xa8, 0xf1, 0xd1,
	0xc2, 0x1d, 0x9f, 0xec, 0x8e, 0x50, 0xaf, 0x2b, 0x68, 0x12, 0x9c, 0x1b, 0x3c, 0x9a, 0xd7, 0x89,
	0x33, 0x09, 0xbe, 0x02, 0xeb, 0x18, 0xef, 0x1b, 0xdf, 0x63, 0x04, 0x5c, 0xa4, 0x36, 0x9a, 0xa0,
	0xe9, 0x70, 0x31, 0x32, 0xda, 0x75, 0x2d, 0x59, 0x47, 0x78, 0x10, 0xe6, 0xd2, 0x57, 0x42, 0xda,
	0xdd, 0xfa, 0xd6, 0xeb, 0xda, 0x68, 0xf1, 0x2b, 0x39, 0xef, 0x08, 0xef, 0xe9, 0xc0, 0xdb, 0x5c,
	0x59, 0x07, 0xf8, 0xbf, 0x58, 0x73, 0x90, 0x4d, 0x8c, 0x19, 0x5a, 0x3b, 0xba, 0xed, 0x1d, 0x73,
	0x3c, 0x30, 0xd4, 0x99, 0xdd, 0x9b, 0xf4, 0xa6, 0xa3, 0x8b, 0x33, 0xba, 0xf3, 0xa7, 0xd4, 0xd3,
	0xa7, 0x99, 0x96, 0x1b, 0x18, 0xf3, 0xd6, 0x1a, 0xe3, 0xbe, 0x9f, 0x8a, 0x9c, 0x2b, 0xfb, 0xdf,
	0x96, 0xb4, 0x51, 0x9c, 0x0f, 0x84, 0x87, 0x1a, 0xf4, 0xce, 0x4f, 0xe0, 0x2f, 0x93, 0x7a, 0x0f,
	0x9f, 0x25, 0x41, 0x9b, 0x92, 0xa0, 0xef, 0x92, 0xa0, 0xb7, 0x8a, 0x74, 0x36, 0x15, 0xe9, 0x7c,
	0x55, 0xa4, 0x73, 0x7f, 0x15, 0x2f, 0xd5, 0x53, 0x1e, 0xd0, 0x50, 0xa4, 0x4c, 0x6f, 0x3d, 0xe7,
	0xa0, 0xd6, 0x42, 0x26, 0xcd, 0xf4, 0x0c, 0x51, 0x0c, 0x92, 0xbd, 0xb4, 0x5b, 0xa0, 0x2b, 0xc0,
	0x0a, 0x37, 0xe8, 0xeb, 0x1a, 0x5c, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x98, 0x18, 0x55,
	0x72, 0x02, 0x00, 0x00,
}

func (m *EventCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Curator) > 0 {
		i -= len(m.Curator)
		copy(dAtA[i:], m.Curator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Curator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BasketDenom) > 0 {
		i -= len(m.BasketDenom)
		copy(dAtA[i:], m.BasketDenom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BasketDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventPut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Credits) > 0 {
		for iNdEx := len(m.Credits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Credits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BasketDenom) > 0 {
		i -= len(m.BasketDenom)
		copy(dAtA[i:], m.BasketDenom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BasketDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventTake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Credits) > 0 {
		for iNdEx := len(m.Credits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Credits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BasketDenom) > 0 {
		i -= len(m.BasketDenom)
		copy(dAtA[i:], m.BasketDenom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BasketDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BasketDenom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Curator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventPut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.BasketDenom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Credits) > 0 {
		for _, e := range m.Credits {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventTake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.BasketDenom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Credits) > 0 {
		for _, e := range m.Credits {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasketDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Curator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Curator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventPut) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventPut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasketDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credits = append(m.Credits, &BasketCredit{})
			if err := m.Credits[len(m.Credits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventTake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventTake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasketDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credits = append(m.Credits, &BasketCredit{})
			if err := m.Credits[len(m.Credits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
