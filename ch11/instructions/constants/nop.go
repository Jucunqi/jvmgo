package constants

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

type Nop struct {
	base.NoOperandsInstruction
}

func (n *Nop) Execute(frame *rtda.Frame) {
	// noting to do
}
