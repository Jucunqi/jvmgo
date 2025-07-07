package heap

import "github.com/Jucunqi/jvmgo/ch10/classfile"

// InterfaceMethodRef 接口方法符号引用
type InterfaceMethodRef struct {
	MemberRef         //继承自MemberRef
	method    *Method // 方法
}

func newInterfaceMethodRef(cp *ConstantPool, classInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&classInfo.ConstantMemberrefInfo)
	return ref
}

// ResolvedInterfaceMethod 解析接口方法符号引用
func (i *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {

	if i.method == nil {
		i.resolveInterfaceMethodRef()
	}
	return i.method
}

// 解析逻辑
func (i *InterfaceMethodRef) resolveInterfaceMethodRef() {

	// 当前类
	currentClass := i.cp.class

	// 方法所在类
	methodClass := i.ResolveClass()

	// 如果方法所在类不是接口，则抛出异常
	if !methodClass.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 递归遍历接口及接口的接口
	method := lookupInterfaceMethod(methodClass, i.name, i.descriptor)

	// 若未找到，抛出异常
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	// 校验权限
	if !method.isAccessibleTo(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	// 方法赋值
	i.method = method
}

func (i *InterfaceMethodRef) Name() string {
	return i.name
}

func (i *InterfaceMethodRef) Descriptor() string {
	return i.descriptor
}

func lookupInterfaceMethod(iface *Class, name string, descriptor string) *Method {

	// 先检查当前接口中的方法
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	// 接着递归检查接口及接口的接口中方法
	return lookupMethodInInterface(iface.interfaces, name, descriptor)
}
