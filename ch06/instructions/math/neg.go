package math

import (
	"github.com/Jucunqi/jvmgo/ch05/instructions/base"
	"github.com/Jucunqi/jvmgo/ch05/rtda"
)

type DNEG struct{ base.NoOperandsInstruction }
type FNEG struct{ base.NoOperandsInstruction }
type INEG struct{ base.NoOperandsInstruction }
type LNEG struct{ base.NoOperandsInstruction }

func (D *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

func (F *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

func (F *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

func (F *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
