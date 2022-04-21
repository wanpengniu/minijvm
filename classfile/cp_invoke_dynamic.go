package classfile

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (method *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	method.referenceKind = reader.readUint8()
	method.referenceIndex = reader.readUint16()
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (method *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	method.descriptorIndex = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (method *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	method.bootstrapMethodAttrIndex = reader.readUint16()
	method.nameAndTypeIndex = reader.readUint16()
}
