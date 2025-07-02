package rtda

import "github.com/Jucunqi/jvmgo/ch08/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack())}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) SetNextPC(pc int) {
	f.nextPC = pc
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) Method() *heap.Method {
	return f.method
}

func (f *Frame) RevertNextPC() {
	f.nextPC = f.thread.pc
}
