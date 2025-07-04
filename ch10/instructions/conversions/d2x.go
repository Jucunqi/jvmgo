package conversions

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
)

type D2F struct {
	base.NoOperandsInstruction
}
type D2I struct {
	base.NoOperandsInstruction
}
type D2L struct {
	base.NoOperandsInstruction
}

func (d *D2I) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	val := stack.PopDouble()
	intVal := int32(val)
	stack.PushInt(intVal)
}

func (d *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double := stack.PopDouble()
	f := float32(double)
	stack.PushFloat(f)
}

func (d *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double := stack.PopDouble()
	l := int64(double)
	stack.PushLong(l)
}
