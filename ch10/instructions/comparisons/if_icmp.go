package comparisons

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}
type IF_ICMPNE struct {
	base.BranchInstruction
}
type IF_ICMPLT struct {
	base.BranchInstruction
}
type IF_ICMPLE struct {
	base.BranchInstruction
}
type IF_ICMPGT struct {
	base.BranchInstruction
}
type IF_ICMPGE struct {
	base.BranchInstruction
}

func (I *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 == v2 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 < v2 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 <= v2 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		base.Branch(frame, I.Offset)
	}
}
func (I *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 >= v2 {
		base.Branch(frame, I.Offset)
	}
}
