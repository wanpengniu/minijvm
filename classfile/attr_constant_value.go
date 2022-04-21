package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (attr *ConstantValueAttribute) readInfo(reader *ClassReader) {
	attr.constantValueIndex = reader.readUint16()
}

func (attr *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return attr.constantValueIndex
}
