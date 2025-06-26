package conversions

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

type L2I struct {
	base.NoOperandsInstruction
}

type L2F struct {
	base.NoOperandsInstruction
}

type L2D struct {
	base.NoOperandsInstruction
}

func (l *L2I) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	long := stack.PopLong()
	i := int32(long)
	stack.PushInt(i)
}

func (l *L2F) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	long := stack.PopLong()
	f := float32(long)
	stack.PushFloat(f)
}

func (l *L2D) Execute(frame *rtda.Frame) {
	stack := rtda.OperandStack{}
	long := stack.PopLong()
	d := float64(long)
	stack.PushDouble(d)
}
