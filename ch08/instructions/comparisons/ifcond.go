package comparisons

import (
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
)

type IFEQ struct {
	base.BranchInstruction
}
type IFNE struct {
	base.BranchInstruction
}
type IFLT struct {
	base.BranchInstruction
}
type IFLE struct {
	base.BranchInstruction
}
type IFGT struct {
	base.BranchInstruction
}
type IFGE struct {
	base.BranchInstruction
}

func (I *IFEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	if i == 0 {
		base.Branch(frame, I.Offset)
	}
}

func (I *IFNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	if i != 0 {
		base.Branch(frame, I.Offset)
	}
}

func (I *IFLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	if i < 0 {
		base.Branch(frame, I.Offset)
	}
}

func (I *IFLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	if i <= 0 {
		base.Branch(frame, I.Offset)
	}
}

func (I *IFGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	if i > 0 {
		base.Branch(frame, I.Offset)
	}
}

func (I *IFGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	if i >= 0 {
		base.Branch(frame, I.Offset)
	}
}
