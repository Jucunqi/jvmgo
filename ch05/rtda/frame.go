package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, maxLocals uint, maxStack uint) *Frame {
	return &Frame{thread: thread, localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack)}
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
