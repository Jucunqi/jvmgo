package math

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

type DMUL struct{ base.NoOperandsInstruction }
type FMUL struct{ base.NoOperandsInstruction }
type IMUL struct{ base.NoOperandsInstruction }
type LMUL struct{ base.NoOperandsInstruction }

func (D *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopDouble()
	double1 := stack.PopDouble()
	stack.PushDouble(double1 + double2)
}

func (F *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopFloat()
	float1 := stack.PopFloat()
	stack.PushFloat(float1 + float2)
}

func (F *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopInt()
	float1 := stack.PopInt()
	stack.PushInt(float1 + float2)
}

func (F *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopLong()
	float1 := stack.PopLong()
	stack.PushLong(float1 + float2)
}
