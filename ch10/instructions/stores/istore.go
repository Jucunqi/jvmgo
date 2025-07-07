package stores

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
)

type ISTORE struct {
	base.Index8Instruction
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}
type ISTORE_1 struct {
	base.NoOperandsInstruction
}
type ISTORE_2 struct {
	base.NoOperandsInstruction
}
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func _ISTORE(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
func (L *ISTORE) Execute(frame *rtda.Frame) {
	_ISTORE(frame, L.Index)
}

func (L *ISTORE_0) Execute(frame *rtda.Frame) {
	_ISTORE(frame, 0)
}

func (L *ISTORE_1) Execute(frame *rtda.Frame) {
	_ISTORE(frame, 1)
}
func (L *ISTORE_2) Execute(frame *rtda.Frame) {
	_ISTORE(frame, 2)
}
func (L *ISTORE_3) Execute(frame *rtda.Frame) {
	_ISTORE(frame, 3)
}
