package classfile

type MarkerAttribute struct {
}

type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttribute struct {
	MarkerAttribute
}

func (attr *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
