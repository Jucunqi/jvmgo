package heap

import "github.com/Jucunqi/jvmgo/ch09/classfile"

// FieldRef 字段符号引用
type FieldRef struct {
	MemberRef        //继承自MemberRef
	field     *Field // 字段
}

func newFieldRef(cp *ConstantPool, info *classfile.ConstantFieldrefInfo) *FieldRef {

	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberrefInfo)
	return ref
}

func (r *FieldRef) ResolvedField() *Field {

	if r.field == nil {

		// 如果属性为空，则解析
		r.resolveFieldRef()
	}
	return r.field
}

func (r *FieldRef) resolveFieldRef() {

	// 解析当前类
	d := r.cp.class

	// 解析属性所在类
	c := r.ResolveClass()

	// 找到属性对象
	field := lookupField(c, r.name, r.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.Lang.IllegalAccessError")
	}
	r.field = field
}

func lookupField(c *Class, name string, descriptor string) *Field {

	// 遍历当前类中所有属性
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	// 遍历接口中的所有属性
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	// 遍历父类中的所有属性
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}
