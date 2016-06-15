// Code generated by protoc-gen-gogo.
// source: coco_interference_scores.proto
// DO NOT EDIT!

/*
	Package firmament is a generated protocol buffer package.

	It is generated from these files:
		coco_interference_scores.proto
		job_desc.proto
		reference_desc.proto
		resource_desc.proto
		resource_topology_node_desc.proto
		resource_vector.proto
		task_desc.proto
		task_final_report.proto
		whare_map_stats.proto

	It has these top-level messages:
		CoCoInterferenceScores
		JobDescriptor
		ReferenceDescriptor
		ResourceDescriptor
		ResourceTopologyNodeDescriptor
		ResourceVector
		TaskDescriptor
		TaskFinalReport
		WhareMapStats
*/
package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CoCoInterferenceScores struct {
	DevilPenalty  uint32 `protobuf:"varint,1,opt,name=devil_penalty,json=devilPenalty,proto3" json:"devil_penalty,omitempty"`
	RabbitPenalty uint32 `protobuf:"varint,2,opt,name=rabbit_penalty,json=rabbitPenalty,proto3" json:"rabbit_penalty,omitempty"`
	SheepPenalty  uint32 `protobuf:"varint,3,opt,name=sheep_penalty,json=sheepPenalty,proto3" json:"sheep_penalty,omitempty"`
	TurtlePenalty uint32 `protobuf:"varint,4,opt,name=turtle_penalty,json=turtlePenalty,proto3" json:"turtle_penalty,omitempty"`
}

func (m *CoCoInterferenceScores) Reset()         { *m = CoCoInterferenceScores{} }
func (m *CoCoInterferenceScores) String() string { return proto.CompactTextString(m) }
func (*CoCoInterferenceScores) ProtoMessage()    {}
func (*CoCoInterferenceScores) Descriptor() ([]byte, []int) {
	return fileDescriptorCocoInterferenceScores, []int{0}
}

func init() {
	proto.RegisterType((*CoCoInterferenceScores)(nil), "firmament.CoCoInterferenceScores")
}
func (m *CoCoInterferenceScores) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CoCoInterferenceScores) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DevilPenalty != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCocoInterferenceScores(data, i, uint64(m.DevilPenalty))
	}
	if m.RabbitPenalty != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCocoInterferenceScores(data, i, uint64(m.RabbitPenalty))
	}
	if m.SheepPenalty != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintCocoInterferenceScores(data, i, uint64(m.SheepPenalty))
	}
	if m.TurtlePenalty != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintCocoInterferenceScores(data, i, uint64(m.TurtlePenalty))
	}
	return i, nil
}

func encodeFixed64CocoInterferenceScores(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32CocoInterferenceScores(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCocoInterferenceScores(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *CoCoInterferenceScores) Size() (n int) {
	var l int
	_ = l
	if m.DevilPenalty != 0 {
		n += 1 + sovCocoInterferenceScores(uint64(m.DevilPenalty))
	}
	if m.RabbitPenalty != 0 {
		n += 1 + sovCocoInterferenceScores(uint64(m.RabbitPenalty))
	}
	if m.SheepPenalty != 0 {
		n += 1 + sovCocoInterferenceScores(uint64(m.SheepPenalty))
	}
	if m.TurtlePenalty != 0 {
		n += 1 + sovCocoInterferenceScores(uint64(m.TurtlePenalty))
	}
	return n
}

func sovCocoInterferenceScores(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCocoInterferenceScores(x uint64) (n int) {
	return sovCocoInterferenceScores(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CoCoInterferenceScores) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCocoInterferenceScores
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CoCoInterferenceScores: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoCoInterferenceScores: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevilPenalty", wireType)
			}
			m.DevilPenalty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCocoInterferenceScores
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DevilPenalty |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RabbitPenalty", wireType)
			}
			m.RabbitPenalty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCocoInterferenceScores
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RabbitPenalty |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SheepPenalty", wireType)
			}
			m.SheepPenalty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCocoInterferenceScores
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.SheepPenalty |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TurtlePenalty", wireType)
			}
			m.TurtlePenalty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCocoInterferenceScores
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TurtlePenalty |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCocoInterferenceScores(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCocoInterferenceScores
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
func skipCocoInterferenceScores(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCocoInterferenceScores
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowCocoInterferenceScores
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCocoInterferenceScores
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCocoInterferenceScores
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCocoInterferenceScores
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCocoInterferenceScores(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCocoInterferenceScores = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCocoInterferenceScores   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorCocoInterferenceScores = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x4f, 0xce,
	0x8f, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x2d, 0x4a, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e,
	0xce, 0x2f, 0x4a, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xcb, 0x2c, 0xca,
	0x4d, 0xcc, 0x4d, 0xcd, 0x2b, 0x91, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x48, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c,
	0x30, 0x0b, 0xa2, 0x53, 0x69, 0x03, 0x23, 0x97, 0x98, 0x73, 0xbe, 0x73, 0xbe, 0x27, 0x92, 0xd9,
	0xc1, 0x60, 0xa3, 0x85, 0x94, 0xb9, 0x78, 0x53, 0x52, 0xcb, 0x32, 0x73, 0xe2, 0x0b, 0x52, 0xf3,
	0x12, 0x73, 0x4a, 0x2a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x78, 0xc0, 0x82, 0x01, 0x10,
	0x31, 0x21, 0x55, 0x2e, 0xbe, 0xa2, 0xc4, 0xa4, 0xa4, 0xcc, 0x12, 0xb8, 0x2a, 0x26, 0xb0, 0x2a,
	0x5e, 0x88, 0x28, 0x4c, 0x19, 0xd0, 0xac, 0xe2, 0x8c, 0xd4, 0xd4, 0x02, 0xb8, 0x2a, 0x66, 0x88,
	0x59, 0x60, 0x41, 0x24, 0xb3, 0x4a, 0x4a, 0x8b, 0x4a, 0x72, 0x52, 0xe1, 0xaa, 0x58, 0x20, 0x66,
	0x41, 0x44, 0xa1, 0xca, 0x9c, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0x00, 0xc4, 0x0f, 0x80, 0x78,
	0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0x5f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15,
	0xbc, 0x5f, 0x8e, 0x27, 0x01, 0x00, 0x00,
}
