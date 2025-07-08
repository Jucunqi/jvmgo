package heap

import (
	"github.com/Jucunqi/jvmgo/ch11/classfile"
)

type Method struct {
	ClassMember                                         //继承自ClassMember
	maxStack        uint                                //最大栈深度
	maxLocals       uint                                //最大局部变量表大小
	code            []byte                              //字节码
	argSlotCount    uint                                // 方法参数个数
	exceptionTable  ExceptionTable                      //异常表
	lineNumberTable *classfile.LineNumberTableAttribute // 行号表
}

func (m *Method) copyAttributes(method *classfile.MemberInfo) {
	codeAttribute := method.CodeAttribute()
	if codeAttribute != nil {
		m.maxStack = uint(codeAttribute.MaxStack())
		m.maxLocals = uint(codeAttribute.MaxLocals())
		m.code = codeAttribute.Code()
		m.lineNumberTable = codeAttribute.LineTableAttribute()
		m.exceptionTable = newExceptionTable(codeAttribute.ExceptionTable(), m.class.constantPool)
	}
}

// 创建异常表
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {

	// 创建数组对象
	table := make([]*ExceptionHandler, len(entries))

	// 循环遍历code属性中的异常表
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(int(entry.CatchType()), cp),
		}
	}
	return table
}

func getCatchType(index int, cp *ConstantPool) *ClassRef {

	if index == 0 {
		return nil
	}
	return cp.GetConstant(uint(index)).(*ClassRef)

}

func newMethods(class *Class, methods []*classfile.MemberInfo) []*Method {

	results := make([]*Method, len(methods))
	for i, method := range methods {
		results[i] = newMethod(class, method)
	}
	return results
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyAttributes(cfMethod)
	method.copyMemberInfo(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
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
func (m *Method) calcArgSlotCount(parameterTypes []string) {

	// 遍历参数列表
	for _, paramType := range parameterTypes {
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

// 根据方法返回类型，构建本地方法的code属性，本地方法的第一个指令是0xfe
func (m *Method) injectCodeAttribute(returnType string) {

	// 最大栈深，默认4
	m.maxStack = 4

	// 布局变量表最大 = 参数个数
	m.maxLocals = m.argSlotCount

	// 根据返回类型，构建code字节码
	switch returnType[0] {
	case 'V':
		m.code = []byte{0xfe, 0xb1} // return
	case 'D':
		m.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		m.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		m.code = []byte{0xfe, 0xad} // lreturn
	case 'L', '[':
		m.code = []byte{0xfe, 0xb0} // areturn
	default:
		m.code = []byte{0xfe, 0xac} // ireturn
	}
}

// FindExceptionHandler 查询异常跳转pc
func (m *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := m.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

func (m *Method) GetLineNumber(pc int) int {
	if m.IsNative() {
		return -2
	}
	if m.lineNumberTable == nil {
		return -1
	}
	return m.lineNumberTable.GetLineNumber(pc)
}
