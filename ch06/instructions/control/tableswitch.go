package control

import (
	"github.com/Jucunqi/jvmgo/ch06/instructions/base"
	"github.com/Jucunqi/jvmgo/ch06/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (T *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {

	reader.SkipPadding()
	T.defaultOffset = reader.ReadInt32()
	T.low = reader.ReadInt32()
	T.high = reader.ReadInt32()
	count := T.high - T.low + 1
	T.jumpOffsets = reader.ReadInt32s(count)
}

func (T *TABLE_SWITCH) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	i := stack.PopInt()
	if i >= T.low && i <= T.high {
		base.Branch(frame, int(T.jumpOffsets[i-T.low]))
	} else {
		base.Branch(frame, int(T.defaultOffset))
	}
}
