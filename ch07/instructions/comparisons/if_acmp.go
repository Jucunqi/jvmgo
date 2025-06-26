package comparisons

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}
type IF_ACMPNE struct {
	base.BranchInstruction
}

func (I *IF_ACMPEQ) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v1 == v2 {
		base.Branch(frame, I.Offset)
	}
}

func (I *IF_ACMPNE) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v1 != v2 {
		base.Branch(frame, I.Offset)
	}
}
