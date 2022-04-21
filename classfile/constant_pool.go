package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { // 注意索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return cp
}

func (pool ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := pool[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (pool ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := pool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := pool.getUtf8(ntInfo.nameIndex)
	_type := pool.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (pool ConstantPool) getClassName(index uint16) string {
	classInfo := pool.getConstantInfo(index).(*ConstantClassInfo)
	return pool.getUtf8(classInfo.nameIndex)
}

func (pool ConstantPool) getUtf8(index uint16) string {
	utf8Info := pool.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
