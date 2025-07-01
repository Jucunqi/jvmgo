package heap

import "github.com/Jucunqi/jvmgo/ch08/classfile"

// 方法符号引用
type MethodRef struct {
	MemberRef         //继承自MemberRef
	method    *Method // 方法
}

func (r *MethodRef) Name() string {
	return r.name
}

func (r *MethodRef) Descriptor() string {
	return r.descriptor
}

func newMethodRef(cp *ConstantPool, info *classfile.ConstantMethodrefInfo) *MethodRef {

	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberrefInfo)
	return ref
}

func (r *MethodRef) ResolveMethod() *Method {

	if r.method == nil {
		r.resolveMethodRef()
	}
	return r.method
}

func (r *MethodRef) resolveMethodRef() {

	// 获取当前类
	currentClass := r.cp.class

	// 获取方法所在类
	methodClass := r.ResolveClass()

	// 如果方法所在类是个接口，抛出异常
	if methodClass.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 遍历类中的方法，获取当前方法
	method := lookupMethod(methodClass, r.name, r.descriptor)

	// 如果方法为空，抛出异常
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	// 验证访问权限
	if !method.isAccessibleTo(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	r.method = method
}

func lookupMethod(class *Class, name string, descriptor string) *Method {

	// 封装通过方法，从类中根据名称和描述符匹配方法
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterface(class.interfaces, name, descriptor)
	}
	return method
}

func lookupMethodInInterface(ifaces []*Class, name string, descriptor string) *Method {

	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterface(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
