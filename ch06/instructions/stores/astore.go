package stores

import (
	"github.com/Jucunqi/jvmgo/ch06/instructions/base"
	"github.com/Jucunqi/jvmgo/ch06/rtda"
)

type ASTORE struct {
	base.Index8Instruction
}

func (A *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, A.Index)
}

func _astore(frame *rtda.Frame, index uint) {

	localVars := frame.LocalVars()
	stack := frame.OperandStack()
	ref := stack.PopRef()
	localVars.SetRef(index, ref)
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (A *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (A *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (A *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (A *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
