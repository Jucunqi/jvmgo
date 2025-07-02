package comparisons

import (
	"github.com/Jucunqi/jvmgo/ch09/instructions/base"
	"github.com/Jucunqi/jvmgo/ch09/rtda"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (L *LCMP) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := 0
	if v1 > v2 {
		result = 1
	} else if v1 < v2 {
		result = -1
	}

	stack.PushInt(int32(result))
}
