package heap

import (
	"github.com/Jucunqi/jvmgo/ch09/classfile"
)

type Method struct {
	ClassMember         //继承自ClassMember
	maxStack     uint   //最大栈深度
	maxLocals    uint   //最大局部变量表大小
	code         []byte //字节码
	argSlotCount uint   // 方法参数个数
}

func (m *Method) copyAttributes(method *classfile.MemberInfo) {
	codeAttribute := method.CodeAttribute()
	if codeAttribute != nil {
		m.maxStack = uint(codeAttribute.MaxStack())
		m.maxLocals = uint(codeAttribute.MaxLocals())
		m.code = codeAttribute.Code()
	}
}

func newMethod(class *Class, methods []*classfile.MemberInfo) []*Method {

	results := make([]*Method, len(methods))
	for i, method := range methods {
		results[i] = &Method{}
		results[i].class = class
		results[i].copyAttributes(method)
		results[i].copyMemberInfo(method)
		results[i].calcArgSlotCount()
	}
	return results
}

func (m *Method) IsSynchronized() bool {
	return 0 != m.accessFlags&ACC_SYNCHRONIZED
}
func (m *Method) IsBridge() bool {
	return 0 != m.accessFlags&ACC_BRIDGE
}
func (m *Method) IsVarargs() bool {
	return 0 != m.accessFlags&ACC_VARARGS
}
func (m *Method) IsNative() bool {
	return 0 != m.accessFlags&ACC_NATIVE
}
func (m *Method) IsAbstract() bool {
	return 0 != m.accessFlags&ACC_ABSTRACT
}
func (m *Method) IsStrict() bool {
	return 0 != m.accessFlags&ACC_STRICT
}

// getters
func (m *Method) MaxStack() uint {
	return m.maxStack
}
func (m *Method) MaxLocals() uint {
	return m.maxLocals
}
func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) Class() *Class {
	return m.class
}

func (m *Method) Name() string {
	return m.name
}

func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}

// 计算方法参数数量
func (m *Method) calcArgSlotCount() {

	// 解析方法描述符(参数类型数组和返回类型)
	parsedDescriptor := parseMethodDescriptor(m.descriptor)

	// 遍历参数列表
	for _, paramType := range parsedDescriptor.parameterTypes {
		m.argSlotCount++

		// Long 和 Double 占两个槽位
		if paramType == "J" || paramType == "D" {
			m.argSlotCount++
		}
	}

	// 如果不是静态方法，再加一个this
	if !m.IsStatic() {
		m.argSlotCount++
	}
}

func (m *Method) Descriptor() string {
	return m.descriptor
}
