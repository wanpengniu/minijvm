package classfile

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

func (attr *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	attr.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range attr.lineNumberTable {
		attr.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
