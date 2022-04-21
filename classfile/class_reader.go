package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (reader *ClassReader) readUint8() uint8 {
	value := reader.data[0]
	reader.data = reader.data[1:]
	return value
}

func (reader *ClassReader) readUint16() uint16 {
	value := binary.BigEndian.Uint16(reader.data)
	reader.data = reader.data[2:]
	return value
}

func (reader *ClassReader) readUint32() uint32 {
	value := binary.BigEndian.Uint32(reader.data)
	reader.data = reader.data[4:]
	return value
}

func (reader *ClassReader) readUint64() uint64 {
	value := binary.BigEndian.Uint64(reader.data)
	reader.data = reader.data[8:]
	return value
}

func (reader *ClassReader) readUint16s() []uint16 {
	n := reader.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = reader.readUint16()
	}
	return s
}

func (reader *ClassReader) readBytes(n uint32) []byte {
	bytes := reader.data[:n]
	reader.data = reader.data[n:]
	return bytes
}
