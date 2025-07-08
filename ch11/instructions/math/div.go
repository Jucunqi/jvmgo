package math

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

type DDIV struct{ base.NoOperandsInstruction }
type FDIV struct{ base.NoOperandsInstruction }
type IDIV struct{ base.NoOperandsInstruction }
type LDIV struct{ base.NoOperandsInstruction }

func (D *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopDouble()
	double1 := stack.PopDouble()
	if double2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushDouble(double1 / double2)
}

func (F *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopFloat()
	float1 := stack.PopFloat()
	if float2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushFloat(float1 / float2)
}

func (F *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopInt()
	float1 := stack.PopInt()
	if float2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushInt(float1 / float2)
}

func (F *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopLong()
	float1 := stack.PopLong()
	if float2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushLong(float1 / float2)
}
