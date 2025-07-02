package loads

import (
	"github.com/Jucunqi/jvmgo/ch09/instructions/base"
	"github.com/Jucunqi/jvmgo/ch09/rtda"
)

type FLOAD struct {
	base.Index8Instruction
}

func (F *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, F.Index)
}

func _fload(frame *rtda.Frame, index uint) {
	localVars := frame.LocalVars()
	stack := frame.OperandStack()
	float := localVars.GetFloat(index)
	stack.PushFloat(float)
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (F *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (F *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (F *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (F *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
