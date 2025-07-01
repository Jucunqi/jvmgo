package control

import (
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}
type ARETURN struct {
	base.NoOperandsInstruction
}
type DRETURN struct {
	base.NoOperandsInstruction
}
type FRETURN struct {
	base.NoOperandsInstruction
}
type IRETURN struct {
	base.NoOperandsInstruction
}
type LRETURN struct {
	base.NoOperandsInstruction
}

func (r *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (i *IRETURN) Execute(frame *rtda.Frame) {

	// 获取当前线程
	thread := frame.Thread()

	// 获取操作数栈中的变量
	currentFrame := thread.PopFrame()
	result := currentFrame.OperandStack().PopInt()

	// 把变量压入调用线程栈帧中的操作数栈
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushInt(result)
}

func (a *ARETURN) Execute(frame *rtda.Frame) {

	// 获取当前线程
	thread := frame.Thread()

	// 获取操作数栈中的变量
	currentFrame := thread.PopFrame()
	result := currentFrame.OperandStack().PopRef()

	// 把变量压入调用线程栈帧中的操作数栈
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushRef(result)
}

func (d *DRETURN) Execute(frame *rtda.Frame) {

	// 获取当前线程
	thread := frame.Thread()

	// 获取操作数栈中的变量
	currentFrame := thread.PopFrame()
	result := currentFrame.OperandStack().PopDouble()

	// 把变量压入调用线程栈帧中的操作数栈
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushDouble(result)
}

func (f *FRETURN) Execute(frame *rtda.Frame) {

	// 获取当前线程
	thread := frame.Thread()

	// 获取操作数栈中的变量
	currentFrame := thread.PopFrame()
	result := currentFrame.OperandStack().PopFloat()

	// 把变量压入调用线程栈帧中的操作数栈
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushFloat(result)
}

func (l *LRETURN) Execute(frame *rtda.Frame) {

	// 获取当前线程
	thread := frame.Thread()

	// 获取操作数栈中的变量
	currentFrame := thread.PopFrame()
	result := currentFrame.OperandStack().PopLong()

	// 把变量压入调用线程栈帧中的操作数栈
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushLong(result)
}
