package math

import (
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
	"math"
)

// DREM 求余数double类型
type DREM struct {
	base.NoOperandsInstruction
}

type FREM struct {
	base.NoOperandsInstruction
}

type IREM struct {
	base.NoOperandsInstruction
}

type LREM struct {
	base.NoOperandsInstruction
}

func (D *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopDouble()
	double1 := stack.PopDouble()

	i := math.Mod(double1, double2)
	stack.PushDouble(i)
}

func (D *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopFloat()
	float1 := stack.PopFloat()

	i := float32(math.Mod(float64(float1), float64(float2)))
	stack.PushFloat(i)
}

func (I *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	int2 := stack.PopInt()
	int1 := stack.PopInt()
	if int2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushInt(int1 % int2)
}

func (L *LREM) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	long2 := stack.PopLong()
	long1 := stack.PopLong()
	if long2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushLong(long1 % long2)
}
