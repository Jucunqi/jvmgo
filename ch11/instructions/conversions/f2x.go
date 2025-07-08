package conversions

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

// F2D Convert float to double
type F2D struct{ base.NoOperandsInstruction }

// F2I Convert float to int
type F2I struct{ base.NoOperandsInstruction }

// F2L Convert float to long
type F2L struct{ base.NoOperandsInstruction }

func (f *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float := stack.PopFloat()
	f2 := float64(float)
	stack.PushDouble(f2)
}

func (f *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float := stack.PopFloat()
	i := int32(float)
	stack.PushInt(i)
}

func (f *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float := stack.PopFloat()
	i := int64(float)
	stack.PushLong(i)
}
