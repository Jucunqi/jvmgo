package heap

import "github.com/Jucunqi/jvmgo/ch07/classfile"

// 字段符号引用
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
		r.resolveFieldRef()
	}
	return r.field
}

func (r *FieldRef) resolveFieldRef() {

	d := r.cp.class
	c := r.ResolveClass()
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

	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}
