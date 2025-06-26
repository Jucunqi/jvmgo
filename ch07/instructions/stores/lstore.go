package stores

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

type LSTORE struct {
	base.Index8Instruction
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}
type LSTORE_1 struct {
	base.NoOperandsInstruction
}
type LSTORE_2 struct {
	base.NoOperandsInstruction
}
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
func (L *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, L.Index)
}

func (L *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (L *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}
func (L *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}
func (L *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
