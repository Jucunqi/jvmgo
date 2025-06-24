package references

import (
	"github.com/Jucunqi/jvmgo/ch06/instructions/base"
	"github.com/Jucunqi/jvmgo/ch06/rtda"
	"github.com/Jucunqi/jvmgo/ch06/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (n *NEW) Execute(frame *rtda.Frame) {

	cp := frame.Method().Class().ConstantPool()
	constant := cp.GetConstant(n.Index)
	classRef := constant.(*heap.ClassRef)
	class := classRef.ResolveClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
