package classfile

import "fmt"

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (class *ClassFile) read(reader *ClassReader) {
	class.readAndCheckMagic(reader)
	class.readAndCheckVersion(reader)
	class.constantPool = readConstantPool(reader)
	class.accessFlags = reader.readUint16()
	class.thisClass = reader.readUint16()
	class.superClass = reader.readUint16()
	class.interfaces = reader.readUint16s()
	class.fields = readMembers(reader, class.constantPool)
	class.methods = readMembers(reader, class.constantPool)
	class.attributes = readAttributes(reader, class.constantPool)
}

func (class *ClassFile) ClassName() string {
	return class.constantPool.getClassName(class.thisClass)
}

func (class *ClassFile) SuperClassName() string {
	if class.superClass > 0 {
		return class.constantPool.getClassName(class.superClass)
	}

	// only java.lang.Object has no supper class
	return ""
}

func (class *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(class.interfaces))
	for i, cpIndex := range class.interfaces {
		interfaceNames[i] = class.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (class *ClassFile) readAndCheckMagic(reader *ClassReader) {
	class.magic = reader.readUint32()

	// Java虚拟机规范规定，如果加载的class文件不符合要求的格式，Java虚拟机实现就抛出java.lang.ClassFormatError异常
	if class.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (class *ClassFile) readAndCheckVersion(reader *ClassReader) {
	class.minorVersion = reader.readUint16()
	class.majorVersion = reader.readUint16()

	// 特定的Java虚拟机实现只能支持版本号在某个范围内的class文件。Oracle的实现是完全向后兼容的，比如Java SE 8支持版本号为45.0~52.0的class文件
	// 如果版本号不在支持的范围内，Java虚拟机实现就抛出java.lang.UnsupportedClassVersionError异常
	switch class.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if class.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (class *ClassFile) MajorVersion() uint16 {
	return class.majorVersion
}

func (class *ClassFile) MinorVersion() uint16 {
	return class.minorVersion
}

func (class *ClassFile) ConstantPool() ConstantPool {
	return class.constantPool
}

func (class *ClassFile) AccessFlags() uint16 {
	return class.accessFlags
}

func (class *ClassFile) Fields() []*MemberInfo {
	return class.fields
}

func (class *ClassFile) Methods() []*MemberInfo {
	return class.methods
}
