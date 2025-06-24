package heap

import "github.com/Jucunqi/jvmgo/ch06/classfile"

// 接口方法符号引用
type InterfaceMethodref struct {
	MemberRef         //继承自MemberRef
	method    *Method // 方法
}

func newInterfaceMethodRef(cp *ConstantPool, classInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodref {
	ref := &InterfaceMethodref{}
	ref.cp = cp
	ref.copyMemberRefInfo(&classInfo.ConstantMemberrefInfo)
	return ref
}
