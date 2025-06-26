package loads

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

type ALOAD struct {
	base.Index8Instruction
}
type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (A *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (A *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (A *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (A *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func (A *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, A.Index)
}

func _aload(frame *rtda.Frame, index uint) {
	localVars := frame.LocalVars()
	stack := frame.OperandStack()
	ref := localVars.GetRef(index)
	stack.PushRef(ref)
}
