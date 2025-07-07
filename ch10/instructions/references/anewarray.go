package references

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
	"github.com/Jucunqi/jvmgo/ch10/rtda/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (a *ANEW_ARRAY) Execute(frame *rtda.Frame) {

	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()

	// 获取类符号引用，并解析
	classRef := cp.GetConstant(a.Index).(*heap.ClassRef)
	componentClass := classRef.ResolveClass()

	// 操作数栈中获取数组长度
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
