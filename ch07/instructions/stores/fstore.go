package stores

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

type FSTORE struct {
	base.Index8Instruction
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}
type FSTORE_1 struct {
	base.NoOperandsInstruction
}
type FSTORE_2 struct {
	base.NoOperandsInstruction
}
type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func _FSTORE(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
func (L *FSTORE) Execute(frame *rtda.Frame) {
	_FSTORE(frame, L.Index)
}

func (L *FSTORE_0) Execute(frame *rtda.Frame) {
	_FSTORE(frame, 0)
}

func (L *FSTORE_1) Execute(frame *rtda.Frame) {
	_FSTORE(frame, 1)
}
func (L *FSTORE_2) Execute(frame *rtda.Frame) {
	_FSTORE(frame, 2)
}
func (L *FSTORE_3) Execute(frame *rtda.Frame) {
	_FSTORE(frame, 3)
}
