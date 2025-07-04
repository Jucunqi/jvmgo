package stores

import (
	"github.com/Jucunqi/jvmgo/ch09/instructions/base"
	"github.com/Jucunqi/jvmgo/ch09/rtda"
)

type DSTORE struct {
	base.Index8Instruction
}

func (D *DSTORE) Execute(frame *rtda.Frame) {

	_dstore(frame, D.Index)
}

func _dstore(frame *rtda.Frame, index uint) {

	localVars := frame.LocalVars()
	stack := frame.OperandStack()
	double := stack.PopDouble()
	localVars.SetDouble(index, double)
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (D *DSTORE_0) Execute(frame *rtda.Frame) {

	_dstore(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (D *DSTORE_1) Execute(frame *rtda.Frame) {

	_dstore(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (D *DSTORE_2) Execute(frame *rtda.Frame) {

	_dstore(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (D *DSTORE_3) Execute(frame *rtda.Frame) {

	_dstore(frame, 3)
}
