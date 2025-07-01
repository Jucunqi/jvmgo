package math

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

// Add double
type DADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }
type IADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }

func (D *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopDouble()
	double1 := stack.PopDouble()
	stack.PushDouble(double1 + double2)
}

func (F *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopFloat()
	float1 := stack.PopFloat()
	stack.PushFloat(float1 + float2)
}

func (F *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopInt()
	float1 := stack.PopInt()
	stack.PushInt(float1 + float2)
}

func (F *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopLong()
	float1 := stack.PopLong()
	stack.PushLong(float1 + float2)
}
