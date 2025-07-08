package references

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
)

type PUT_STATIC struct {
	base.Index16Instruction
}

func (p *PUT_STATIC) Execute(frame *rtda.Frame) {

	// 获取当前方法
	currentMethod := frame.Method()

	// 获取当前类
	currentClass := currentMethod.Class()

	// 获取当前类的常量池
	cp := currentClass.ConstantPool()

	// 根据索引获取指定的符号引用
	fieldRef := cp.GetConstant(p.Index).(*heap.FieldRef)

	// 解析属性
	field := fieldRef.ResolvedField()

	// 获取属性所在类
	class := field.Class()

	// 如果类未初始化，执行初始化方法
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 访问校验
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	// 从操作数栈中获取变量引用， 封装为字段的field引用
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		ref := stack.PopRef()
		slots.SetRef(slotId, ref)

	}
}
