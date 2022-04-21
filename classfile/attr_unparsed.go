package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (attr *UnparsedAttribute) readInfo(reader *ClassReader) {
	attr.info = reader.readBytes(attr.length)
}
