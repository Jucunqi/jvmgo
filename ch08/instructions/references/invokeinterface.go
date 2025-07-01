package references

import (
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
	"github.com/Jucunqi/jvmgo/ch08/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index uint
}

func (i *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	i.index = uint(reader.ReadInt16())
	reader.ReadInt8() // count
	reader.ReadInt8() // must be 0  给其他oracle虚拟机用的
}

func (i *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {

	// 获取常量池
	cp := frame.Method().Class().ConstantPool()

	// 从常量池中获取符号引用
	methodRef := cp.GetConstant(i.index).(*heap.InterfaceMethodRef)

	// 解析方法
	resolvedMethod := methodRef.ResolvedInterfaceMethod()

	// 访问标识拦截
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// 如果引用对象所指对象没有实现解析出来的接口，则报错
	if !ref.Class().IsImplements(methodRef.ResolveClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 根据操作数栈中弹出的对象引用，找到多态方法真实调用的对象，从中找到执行方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())

	// 校验
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	// 方法调用
	base.InvokeMethod(frame, methodToBeInvoked)
}
