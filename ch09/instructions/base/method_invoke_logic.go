package base

import (
	"github.com/Jucunqi/jvmgo/ch09/rtda"
	"github.com/Jucunqi/jvmgo/ch09/rtda/heap"
)

// InvokeMethod 方法调用逻辑实现
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {

	// 获取当前栈帧所在的线程
	thread := invokerFrame.Thread()

	// 因为方法调用需要向当前线程的栈中压入一个栈帧，所以创建一个方法的栈帧
	newFrame := thread.NewFrame(method)

	// 压入虚拟机栈
	thread.PushFrame(newFrame)

	// 参数传递：1. 获取参数数量
	argSlotCount := int(method.ArgSlotCount())

	// 参数传递：2. 从调用方的操作数栈中参数变量，放入方法栈帧的局部变量表中
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

}
