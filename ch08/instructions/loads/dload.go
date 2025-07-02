package loads

import (
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
)

type DLOAD struct {
	base.Index8Instruction
}

func (D *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, D.Index)
}

func _dload(frame *rtda.Frame, index uint) {

	localVars := frame.LocalVars()
	stack := frame.OperandStack()
	double := localVars.GetDouble(index)
	stack.PushDouble(double)
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (D *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (D *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (D *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (D *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
