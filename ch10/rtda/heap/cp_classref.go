package heap

import "github.com/Jucunqi/jvmgo/ch10/classfile"

// 类符号引用
type ClassRef struct {
	SymRef //继承自SymRef
}

func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {

	classRef := &ClassRef{}
	classRef.cp = cp
	classRef.className = classInfo.Name()
	return classRef
}
