package heap

import "github.com/Jucunqi/jvmgo/ch07/classfile"

// MemberRef 成员符号引用
type MemberRef struct {
	SymRef            //继承自SymRef
	name       string // 名称
	descriptor string // 描述符
}

func (m *MemberRef) copyMemberRefInfo(info *classfile.ConstantMemberrefInfo) {

	name, _type := info.NameAndDescriptor()
	m.name = name
	m.descriptor = _type
	m.className = info.ClassName()
}
