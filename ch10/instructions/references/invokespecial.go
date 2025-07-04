package references

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
	"github.com/Jucunqi/jvmgo/ch10/rtda/heap"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (i *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {

	// 获取当前类
	currentClass := frame.Method().Class()
	// 获取当前常量池对象
	cp := currentClass.ConstantPool()

	// 获取目标方法的符号引用
	methodRef := cp.GetConstant(i.Index).(*heap.MethodRef)

	// 解析目标类和方法
	resolvedClass := methodRef.ResolveClass()
	resolvedMethod := methodRef.ResolveMethod()

	// 如果方法是构造函数那么符号引用的类和和声明的类必须是同一个类
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	// 不能是静态方法
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 从操作数栈中获取this对象的引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)

	// 空指针
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// 确保protected方法只能被声明该方法的类或子类调用
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass && ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	// 如果调用的是父类中的方法，但不是构造函数，且当前类被ACC_SUPER标识，需要一个额外的过程查找 最终要调用的方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() && resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	// 方法调用
	base.InvokeMethod(frame, methodToBeInvoked)
}
