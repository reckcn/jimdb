package backend

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BlockSyncOffset) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "ulid":
			z.Ulid, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Ulid")
				return
			}
		case "minT":
			z.MinT, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "MinT")
				return
			}
		case "maxT":
			z.MaxT, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "MaxT")
				return
			}
		case "path":
			z.Path, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Path")
				return
			}
		case "Offset":
			z.Offset, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Offset")
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
func (z *BlockSyncOffset) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "ulid"
	err = en.Append(0x85, 0xa4, 0x75, 0x6c, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Ulid)
	if err != nil {
		err = msgp.WrapError(err, "Ulid")
		return
	}
	// write "minT"
	err = en.Append(0xa4, 0x6d, 0x69, 0x6e, 0x54)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.MinT)
	if err != nil {
		err = msgp.WrapError(err, "MinT")
		return
	}
	// write "maxT"
	err = en.Append(0xa4, 0x6d, 0x61, 0x78, 0x54)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.MaxT)
	if err != nil {
		err = msgp.WrapError(err, "MaxT")
		return
	}
	// write "path"
	err = en.Append(0xa4, 0x70, 0x61, 0x74, 0x68)
	if err != nil {
		return
	}
	err = en.WriteString(z.Path)
	if err != nil {
		err = msgp.WrapError(err, "Path")
		return
	}
	// write "Offset"
	err = en.Append(0xa6, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Offset)
	if err != nil {
		err = msgp.WrapError(err, "Offset")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BlockSyncOffset) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "ulid"
	o = append(o, 0x85, 0xa4, 0x75, 0x6c, 0x69, 0x64)
	o = msgp.AppendString(o, z.Ulid)
	// string "minT"
	o = append(o, 0xa4, 0x6d, 0x69, 0x6e, 0x54)
	o = msgp.AppendInt64(o, z.MinT)
	// string "maxT"
	o = append(o, 0xa4, 0x6d, 0x61, 0x78, 0x54)
	o = msgp.AppendInt64(o, z.MaxT)
	// string "path"
	o = append(o, 0xa4, 0x70, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.Path)
	// string "Offset"
	o = append(o, 0xa6, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	o = msgp.AppendInt64(o, z.Offset)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BlockSyncOffset) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "ulid":
			z.Ulid, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ulid")
				return
			}
		case "minT":
			z.MinT, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MinT")
				return
			}
		case "maxT":
			z.MaxT, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MaxT")
				return
			}
		case "path":
			z.Path, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Path")
				return
			}
		case "Offset":
			z.Offset, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Offset")
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
func (z *BlockSyncOffset) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Ulid) + 5 + msgp.Int64Size + 5 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Path) + 7 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *HandshakeStatus) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 byte
		zb0001, err = dc.ReadByte()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = HandshakeStatus(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z HandshakeStatus) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteByte(byte(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z HandshakeStatus) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendByte(o, byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HandshakeStatus) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 byte
		zb0001, bts, err = msgp.ReadByteBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = HandshakeStatus(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z HandshakeStatus) Msgsize() (s int) {
	s = msgp.ByteSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SlaveOfCommand) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "masterAddr":
			z.MasterAddr, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "MasterAddr")
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
func (z SlaveOfCommand) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "masterAddr"
	err = en.Append(0x81, 0xaa, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.MasterAddr)
	if err != nil {
		err = msgp.WrapError(err, "MasterAddr")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SlaveOfCommand) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "masterAddr"
	o = append(o, 0x81, 0xaa, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72)
	o = msgp.AppendString(o, z.MasterAddr)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SlaveOfCommand) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "masterAddr":
			z.MasterAddr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MasterAddr")
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
func (z SlaveOfCommand) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.MasterAddr)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SyncHandshake) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "slaveAddr":
			z.SlaveAddr, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "SlaveAddr")
				return
			}
		case "blocksMinT":
			z.BlocksMinT, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "BlocksMinT")
				return
			}
		case "slaveOfNoOne":
			z.SlaveOfNoOne, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "SlaveOfNoOne")
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
func (z SyncHandshake) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "slaveAddr"
	err = en.Append(0x83, 0xa9, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.SlaveAddr)
	if err != nil {
		err = msgp.WrapError(err, "SlaveAddr")
		return
	}
	// write "blocksMinT"
	err = en.Append(0xaa, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x4d, 0x69, 0x6e, 0x54)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.BlocksMinT)
	if err != nil {
		err = msgp.WrapError(err, "BlocksMinT")
		return
	}
	// write "slaveOfNoOne"
	err = en.Append(0xac, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x4f, 0x66, 0x4e, 0x6f, 0x4f, 0x6e, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.SlaveOfNoOne)
	if err != nil {
		err = msgp.WrapError(err, "SlaveOfNoOne")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SyncHandshake) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "slaveAddr"
	o = append(o, 0x83, 0xa9, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72)
	o = msgp.AppendString(o, z.SlaveAddr)
	// string "blocksMinT"
	o = append(o, 0xaa, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x4d, 0x69, 0x6e, 0x54)
	o = msgp.AppendInt64(o, z.BlocksMinT)
	// string "slaveOfNoOne"
	o = append(o, 0xac, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x4f, 0x66, 0x4e, 0x6f, 0x4f, 0x6e, 0x65)
	o = msgp.AppendBool(o, z.SlaveOfNoOne)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SyncHandshake) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "slaveAddr":
			z.SlaveAddr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SlaveAddr")
				return
			}
		case "blocksMinT":
			z.BlocksMinT, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BlocksMinT")
				return
			}
		case "slaveOfNoOne":
			z.SlaveOfNoOne, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SlaveOfNoOne")
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
func (z SyncHandshake) Msgsize() (s int) {
	s = 1 + 10 + msgp.StringPrefixSize + len(z.SlaveAddr) + 11 + msgp.Int64Size + 13 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SyncHandshakeAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "status":
			{
				var zb0002 byte
				zb0002, err = dc.ReadByte()
				if err != nil {
					err = msgp.WrapError(err, "Status")
					return
				}
				z.Status = HandshakeStatus(zb0002)
			}
		case "relationID":
			z.RelationID, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "RelationID")
				return
			}
		case "message":
			z.Message, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Message")
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
func (z SyncHandshakeAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "status"
	err = en.Append(0x83, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	if err != nil {
		return
	}
	err = en.WriteByte(byte(z.Status))
	if err != nil {
		err = msgp.WrapError(err, "Status")
		return
	}
	// write "relationID"
	err = en.Append(0xaa, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44)
	if err != nil {
		return
	}
	err = en.WriteString(z.RelationID)
	if err != nil {
		err = msgp.WrapError(err, "RelationID")
		return
	}
	// write "message"
	err = en.Append(0xa7, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Message)
	if err != nil {
		err = msgp.WrapError(err, "Message")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SyncHandshakeAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "status"
	o = append(o, 0x83, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendByte(o, byte(z.Status))
	// string "relationID"
	o = append(o, 0xaa, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44)
	o = msgp.AppendString(o, z.RelationID)
	// string "message"
	o = append(o, 0xa7, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Message)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SyncHandshakeAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "status":
			{
				var zb0002 byte
				zb0002, bts, err = msgp.ReadByteBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Status")
					return
				}
				z.Status = HandshakeStatus(zb0002)
			}
		case "relationID":
			z.RelationID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RelationID")
				return
			}
		case "message":
			z.Message, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Message")
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
func (z SyncHandshakeAck) Msgsize() (s int) {
	s = 1 + 7 + msgp.ByteSize + 11 + msgp.StringPrefixSize + len(z.RelationID) + 8 + msgp.StringPrefixSize + len(z.Message)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SyncHeartbeat) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "masterAddr":
			z.MasterAddr, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "MasterAddr")
				return
			}
		case "slaveAddr":
			z.SlaveAddr, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "SlaveAddr")
				return
			}
		case "relationID":
			z.RelationID, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "RelationID")
				return
			}
		case "blkSyncOffset":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "BlkSyncOffset")
					return
				}
				z.BlkSyncOffset = nil
			} else {
				if z.BlkSyncOffset == nil {
					z.BlkSyncOffset = new(BlockSyncOffset)
				}
				err = z.BlkSyncOffset.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "BlkSyncOffset")
					return
				}
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
func (z *SyncHeartbeat) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "masterAddr"
	err = en.Append(0x84, 0xaa, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.MasterAddr)
	if err != nil {
		err = msgp.WrapError(err, "MasterAddr")
		return
	}
	// write "slaveAddr"
	err = en.Append(0xa9, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.SlaveAddr)
	if err != nil {
		err = msgp.WrapError(err, "SlaveAddr")
		return
	}
	// write "relationID"
	err = en.Append(0xaa, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44)
	if err != nil {
		return
	}
	err = en.WriteString(z.RelationID)
	if err != nil {
		err = msgp.WrapError(err, "RelationID")
		return
	}
	// write "blkSyncOffset"
	err = en.Append(0xad, 0x62, 0x6c, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if err != nil {
		return
	}
	if z.BlkSyncOffset == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.BlkSyncOffset.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "BlkSyncOffset")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SyncHeartbeat) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "masterAddr"
	o = append(o, 0x84, 0xaa, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72)
	o = msgp.AppendString(o, z.MasterAddr)
	// string "slaveAddr"
	o = append(o, 0xa9, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72)
	o = msgp.AppendString(o, z.SlaveAddr)
	// string "relationID"
	o = append(o, 0xaa, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44)
	o = msgp.AppendString(o, z.RelationID)
	// string "blkSyncOffset"
	o = append(o, 0xad, 0x62, 0x6c, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if z.BlkSyncOffset == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.BlkSyncOffset.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "BlkSyncOffset")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SyncHeartbeat) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "masterAddr":
			z.MasterAddr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MasterAddr")
				return
			}
		case "slaveAddr":
			z.SlaveAddr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SlaveAddr")
				return
			}
		case "relationID":
			z.RelationID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RelationID")
				return
			}
		case "blkSyncOffset":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.BlkSyncOffset = nil
			} else {
				if z.BlkSyncOffset == nil {
					z.BlkSyncOffset = new(BlockSyncOffset)
				}
				bts, err = z.BlkSyncOffset.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "BlkSyncOffset")
					return
				}
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
func (z *SyncHeartbeat) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.MasterAddr) + 10 + msgp.StringPrefixSize + len(z.SlaveAddr) + 11 + msgp.StringPrefixSize + len(z.RelationID) + 14
	if z.BlkSyncOffset == nil {
		s += msgp.NilSize
	} else {
		s += z.BlkSyncOffset.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SyncHeartbeatAck) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "status":
			err = z.Status.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Status")
				return
			}
		case "message":
			z.Message, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Message")
				return
			}
		case "blkSyncOffset":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "BlkSyncOffset")
					return
				}
				z.BlkSyncOffset = nil
			} else {
				if z.BlkSyncOffset == nil {
					z.BlkSyncOffset = new(BlockSyncOffset)
				}
				err = z.BlkSyncOffset.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "BlkSyncOffset")
					return
				}
			}
		case "data":
			z.Data, err = dc.ReadBytes(z.Data)
			if err != nil {
				err = msgp.WrapError(err, "Data")
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
func (z *SyncHeartbeatAck) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "status"
	err = en.Append(0x84, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	if err != nil {
		return
	}
	err = z.Status.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Status")
		return
	}
	// write "message"
	err = en.Append(0xa7, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Message)
	if err != nil {
		err = msgp.WrapError(err, "Message")
		return
	}
	// write "blkSyncOffset"
	err = en.Append(0xad, 0x62, 0x6c, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if err != nil {
		return
	}
	if z.BlkSyncOffset == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.BlkSyncOffset.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "BlkSyncOffset")
			return
		}
	}
	// write "data"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SyncHeartbeatAck) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "status"
	o = append(o, 0x84, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	o, err = z.Status.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Status")
		return
	}
	// string "message"
	o = append(o, 0xa7, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Message)
	// string "blkSyncOffset"
	o = append(o, 0xad, 0x62, 0x6c, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if z.BlkSyncOffset == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.BlkSyncOffset.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "BlkSyncOffset")
			return
		}
	}
	// string "data"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendBytes(o, z.Data)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SyncHeartbeatAck) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "status":
			bts, err = z.Status.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Status")
				return
			}
		case "message":
			z.Message, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Message")
				return
			}
		case "blkSyncOffset":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.BlkSyncOffset = nil
			} else {
				if z.BlkSyncOffset == nil {
					z.BlkSyncOffset = new(BlockSyncOffset)
				}
				bts, err = z.BlkSyncOffset.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "BlkSyncOffset")
					return
				}
			}
		case "data":
			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
			if err != nil {
				err = msgp.WrapError(err, "Data")
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
func (z *SyncHeartbeatAck) Msgsize() (s int) {
	s = 1 + 7 + z.Status.Msgsize() + 8 + msgp.StringPrefixSize + len(z.Message) + 14
	if z.BlkSyncOffset == nil {
		s += msgp.NilSize
	} else {
		s += z.BlkSyncOffset.Msgsize()
	}
	s += 5 + msgp.BytesPrefixSize + len(z.Data)
	return
}