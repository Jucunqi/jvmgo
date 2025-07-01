package control

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (G *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, G.Offset)
}
