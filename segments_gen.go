package rankdb

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Segments) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			{
				var zb0002 string
				zb0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "ID")
					return
				}
				z.ID = SegmentsID(zb0002)
			}
		case "Segments":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Segments")
				return
			}
			if cap(z.Segments) >= int(zb0003) {
				z.Segments = (z.Segments)[:zb0003]
			} else {
				z.Segments = make([]Segment, zb0003)
			}
			for za0001 := range z.Segments {
				err = z.Segments[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Segments", za0001)
					return
				}
			}
		case "NextID":
			err = z.NextID.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "NextID")
				return
			}
		case "IsIndex":
			z.IsIndex, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "IsIndex")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Segments) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "ID"
	err = en.Append(0x84, 0xa2, 0x49, 0x44)
	if err != nil {
		return
	}
	err = en.WriteString(string(z.ID))
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	// write "Segments"
	err = en.Append(0xa8, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Segments)))
	if err != nil {
		err = msgp.WrapError(err, "Segments")
		return
	}
	for za0001 := range z.Segments {
		err = z.Segments[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Segments", za0001)
			return
		}
	}
	// write "NextID"
	err = en.Append(0xa6, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x44)
	if err != nil {
		return
	}
	err = z.NextID.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "NextID")
		return
	}
	// write "IsIndex"
	err = en.Append(0xa7, 0x49, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsIndex)
	if err != nil {
		err = msgp.WrapError(err, "IsIndex")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Segments) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "ID"
	o = append(o, 0x84, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, string(z.ID))
	// string "Segments"
	o = append(o, 0xa8, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Segments)))
	for za0001 := range z.Segments {
		o, err = z.Segments[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Segments", za0001)
			return
		}
	}
	// string "NextID"
	o = append(o, 0xa6, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x44)
	o, err = z.NextID.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "NextID")
		return
	}
	// string "IsIndex"
	o = append(o, 0xa7, 0x49, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78)
	o = msgp.AppendBool(o, z.IsIndex)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Segments) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			{
				var zb0002 string
				zb0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "ID")
					return
				}
				z.ID = SegmentsID(zb0002)
			}
		case "Segments":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Segments")
				return
			}
			if cap(z.Segments) >= int(zb0003) {
				z.Segments = (z.Segments)[:zb0003]
			} else {
				z.Segments = make([]Segment, zb0003)
			}
			for za0001 := range z.Segments {
				bts, err = z.Segments[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Segments", za0001)
					return
				}
			}
		case "NextID":
			bts, err = z.NextID.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "NextID")
				return
			}
		case "IsIndex":
			z.IsIndex, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IsIndex")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Segments) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(string(z.ID)) + 9 + msgp.ArrayHeaderSize
	for za0001 := range z.Segments {
		s += z.Segments[za0001].Msgsize()
	}
	s += 7 + z.NextID.Msgsize() + 8 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SegmentsID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = SegmentsID(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z SegmentsID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SegmentsID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SegmentsID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = SegmentsID(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SegmentsID) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}
