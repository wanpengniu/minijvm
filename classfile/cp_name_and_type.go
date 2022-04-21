package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (nameAndType *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	nameAndType.nameIndex = reader.readUint16()
	nameAndType.descriptorIndex = reader.readUint16()
}
