package control

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (L *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	L.defaultOffset = reader.ReadInt32()
	L.npairs = reader.ReadInt32()
	L.matchOffsets = reader.ReadInt32s(L.npairs * 2)
}

func (L *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	key := stack.PopInt()
	for i := int32(0); i < L.npairs*2; i += 2 {
		value := L.matchOffsets[i]
		if value == key {
			base.Branch(frame, int(L.matchOffsets[i+1]))
			return
		}
	}
	base.Branch(frame, int(L.defaultOffset))
}
