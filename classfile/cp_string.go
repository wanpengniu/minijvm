package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (str *ConstantStringInfo) readInfo(reader *ClassReader) {
	str.stringIndex = reader.readUint16()
}

func (str *ConstantStringInfo) String() string {
	return str.cp.getUtf8(str.stringIndex)
}
