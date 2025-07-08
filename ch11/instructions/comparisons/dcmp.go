package comparisons

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}
type DCMPL struct {
	base.NoOperandsInstruction
}

func (D *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(true)
}

func _dcmp(b bool) {
	stack := rtda.OperandStack{}
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if b {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (D *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(false)
}
