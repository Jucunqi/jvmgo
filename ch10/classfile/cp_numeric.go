package classfile

import (
	"math"
)

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	c.val = int32(reader.readUnit32())
}

func (c *ConstantIntegerInfo) Value() int32 {
	return c.val
}

func (c *ConstantFloatInfo) readInfo(reader *ClassReader) {
	data := reader.readUnit32()
	c.val = math.Float32frombits(data)
}

func (c *ConstantFloatInfo) Value() float32 {
	return c.val
}

func (c *ConstantLongInfo) readInfo(reader *ClassReader) {

	bytes := reader.readUnit64()
	c.val = int64(bytes)
}

func (c *ConstantLongInfo) Value() int64 {
	return c.val
}

func (c *ConstantDoubleInfo) readInfo(reader *ClassReader) {

	bytes := reader.readUnit64()
	c.val = math.Float64frombits(bytes)
}

func (c *ConstantDoubleInfo) Value() float64 {
	return c.val
}
