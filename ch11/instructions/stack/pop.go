package stack

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (P *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

func (P *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
