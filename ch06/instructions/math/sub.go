package math

import (
	"github.com/Jucunqi/jvmgo/ch06/instructions/base"
	"github.com/Jucunqi/jvmgo/ch06/rtda"
)

type DSUB struct {
	base.NoOperandsInstruction
}
type FSUB struct {
	base.NoOperandsInstruction
}
type ISUB struct {
	base.NoOperandsInstruction
}
type LSUB struct {
	base.NoOperandsInstruction
}

func (D *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopDouble()
	double1 := stack.PopDouble()
	stack.PushDouble(double1 - double2)
}
func (D *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopFloat()
	double1 := stack.PopFloat()
	stack.PushFloat(double1 - double2)
}
func (D *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopInt()
	double1 := stack.PopInt()
	stack.PushInt(double1 - double2)
}
func (D *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopLong()
	double1 := stack.PopLong()
	stack.PushLong(double1 - double2)
}
