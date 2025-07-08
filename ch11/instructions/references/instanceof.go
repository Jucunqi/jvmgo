package references

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
)

// INSTANCE_OF 判断某个对象是否是类的实例，或者实现了某些接口
type INSTANCE_OF struct {
	base.Index16Instruction
}

func (i *INSTANCE_OF) Execute(frame *rtda.Frame) {

	// 从操作数栈中弹出对象引用
	stack := frame.OperandStack()
	ref := stack.PopRef()

	if ref == nil {
		stack.PushInt(0)
		return
	}

	// 在常量池中根据操作数获取类或者接口
	cp := frame.Method().Class().ConstantPool()

	classRef := cp.GetConstant(i.Index).(*heap.ClassRef)
	class := classRef.ResolveClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
