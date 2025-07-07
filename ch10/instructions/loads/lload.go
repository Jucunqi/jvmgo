package loads

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
)

type LLOAD struct {
	base.Index8Instruction
}

func (L *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, L.Index)
}

func _lload(frame *rtda.Frame, index uint) {
	vars := frame.LocalVars()
	stack := frame.OperandStack()
	val := vars.GetLong(index)
	stack.PushLong(val)
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (L *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (L *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (L *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (L *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
