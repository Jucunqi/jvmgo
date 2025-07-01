package stack

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

// DUP
/*
bottom -> top
[...][c][b][a]

	\_
	  |
	  V

[...][c][b][a][a]
*/
type DUP struct {
	base.NoOperandsInstruction
}

// DUP_X1 复制栈顶的元素，并将其插入到栈顶以下第二个元素之后
/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
type DUP_X1 struct {
	base.NoOperandsInstruction
}

// DUP_X2 复制栈顶的元素，并将其插入到栈顶以下第三个元素之后
/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/
type DUP_X2 struct {
	base.NoOperandsInstruction
}

// DUP2 复制栈顶两个元素
/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
type DUP2 struct {
	base.NoOperandsInstruction
}

// DUP2_X1 复制栈顶两个槽位然后放在第二个元素
/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
type DUP2_X1 struct {
	base.NoOperandsInstruction
}

// DUP2_X2 复制栈顶两个槽位然后放在第三个元素
/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (D *DUP) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

func (D *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.PopSlot()
	top2 := stack.PopSlot()
	stack.PushSlot(top)
	stack.PushSlot(top2)
	stack.PushSlot(top)
}

func (D *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (D *DUP2) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (D *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (D *DUP2_X2) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

}
