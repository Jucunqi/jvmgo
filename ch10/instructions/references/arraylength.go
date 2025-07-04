package references

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (a *ARRAY_LENGTH) Execute(frame *rtda.Frame) {

	// 从栈顶获取数组引用
	stack := frame.OperandStack()
	arrRef := stack.PopRef()

	if arrRef == nil {
		panic("java.lang.NullPointException")
	}
	// 获取数组长度
	length := arrRef.ArrayLength()

	// 压入操作数栈
	stack.PushInt(length)
}
