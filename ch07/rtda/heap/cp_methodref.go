package heap

import "github.com/Jucunqi/jvmgo/ch07/classfile"

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
