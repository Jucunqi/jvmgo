package extended

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (I *IFNULL) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		base.Branch(frame, I.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (I *IFNONNULL) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref != nil {
		base.Branch(frame, I.Offset)
	}
}
