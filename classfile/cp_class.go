package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (class *ConstantClassInfo) readInfo(reader *ClassReader) {
	class.nameIndex = reader.readUint16()
}

func (class *ConstantClassInfo) Name() string {
	return class.cp.getUtf8(class.nameIndex)
}
