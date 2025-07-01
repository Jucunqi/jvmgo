package references

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
	"github.com/Jucunqi/jvmgo/ch08/rtda/heap"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (i *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {

	// 获取当前类的常量池
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()

	// 获取方法的符号引用
	methodRef := cp.GetConstant(i.Index).(*heap.MethodRef)

	// 方法解析
	resolvedMethod := methodRef.ResolveMethod()

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 从操作数栈中获取this对象引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)

	if ref == nil {
		// hack！
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPinterException")
	}
	// 校验protect权限
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() && ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtda.OperandStack, descriptor string) {

	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(B)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(I)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
