package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (integer *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	integer.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (float *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	float.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (long *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	long.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (double *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	double.val = math.Float64frombits(bytes)
}
