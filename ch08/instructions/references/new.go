package references

import (
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
	"github.com/Jucunqi/jvmgo/ch08/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (n *NEW) Execute(frame *rtda.Frame) {

	// 获取当前类的常量池对象
	cp := frame.Method().Class().ConstantPool()

	// 根据索引获取待创建对象的常量池
	constant := cp.GetConstant(n.Index)

	// 转换为类符号引用
	classRef := constant.(*heap.ClassRef)

	// 解析类
	class := classRef.ResolveClass()

	// 如果类未初始化，则需要进行类的初始化 -> 执行<clinit>
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 类访问标识校验
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	// 创建对象并将对象的引用压入操作数栈
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
