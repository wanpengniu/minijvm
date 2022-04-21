package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

func (member *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	member.classIndex = reader.readUint16()
	member.nameAndTypeIndex = reader.readUint16()
}

func (member *ConstantMemberRefInfo) ClassName() string {
	return member.cp.getClassName(member.classIndex)
}

func (member *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return member.cp.getNameAndType(member.nameAndTypeIndex)
}
