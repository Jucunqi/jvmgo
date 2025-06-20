package conversions

import (
	"github.com/Jucunqi/jvmgo/ch05/instructions/base"
	"github.com/Jucunqi/jvmgo/ch05/rtda"
)

type I2B struct {
	base.NoOperandsInstruction
}
type I2C struct {
	base.NoOperandsInstruction
}

type I2S struct {
	base.NoOperandsInstruction
}

type I2F struct {
	base.NoOperandsInstruction
}

type I2D struct {
	base.NoOperandsInstruction
}

type I2L struct {
	base.NoOperandsInstruction
}

func (i *I2B) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	intVal := stack.PopInt()
	b := int32(int8(intVal))
	stack.PushInt(b)
}

func (i *I2C) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	intVal := stack.PopInt()
	c := int16(intVal)
	stack.PushInt(int32(c))
}

func (i *I2S) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	intVal := stack.PopInt()
	c := int16(intVal)
	stack.PushInt(int32(c))
}

func (i *I2F) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	intVal := stack.PopInt()
	f := float32(intVal)
	stack.PushFloat(f)
}

func (i *I2D) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	intVal := stack.PopInt()
	d := float64(intVal)
	stack.PushDouble(d)
}

func (i *I2L) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	intVal := stack.PopInt()
	l := int64(intVal)
	stack.PushLong(l)
}
