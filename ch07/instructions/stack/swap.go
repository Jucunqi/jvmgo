package stack

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

// SWAP the top two operand stack values
type SWAP struct {
	base.NoOperandsInstruction
}

// Execute
/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/
func (S *SWAP) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
