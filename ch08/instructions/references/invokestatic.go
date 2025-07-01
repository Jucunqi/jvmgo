package references

import (
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
	"github.com/Jucunqi/jvmgo/ch08/rtda/heap"
)

// INVOKE_STATIC 调用静态方法
type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (i *INVOKE_STATIC) Execute(frame *rtda.Frame) {

	// 获取当前方法所在类的常量池
	cp := frame.Method().Class().ConstantPool()

	// 常量池中获取方法的符号引用
	methodRef := cp.GetConstant(i.Index).(*heap.MethodRef)

	// 解析方法符号引用，获取方法对象
	method := methodRef.ResolveMethod()

	// 获取方法所在类
	class := method.Class()

	// 若类未初始化，执行类的初始化方法
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 校验方法必须为静态
	if !method.IsStatic() {
		panic("java.lang.IncompatibleClasChangeError")
	}

	// 方法调用
	base.InvokeMethod(frame, method)
}
