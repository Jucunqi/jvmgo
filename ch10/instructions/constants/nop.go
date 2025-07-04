package constants

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
)

type Nop struct {
	base.NoOperandsInstruction
}

func (n *Nop) Execute(frame *rtda.Frame) {
	// noting to do
}
