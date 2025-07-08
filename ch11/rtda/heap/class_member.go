package heap

import "github.com/Jucunqi/jvmgo/ch11/classfile"

type ClassMember struct {
	accessFlags    uint16 // 访问标识符
	name           string // 名称
	descriptor     string // 描述符
	signature      string
	class          *Class // 所属类
	annotationData []byte // RuntimeVisibleAnnotations_attribute
}

func (c *ClassMember) copyMemberInfo(info *classfile.MemberInfo) {
	c.accessFlags = info.AccessFlags()
	c.name = info.Name()
	c.descriptor = info.Descriptor()
}

func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (c *ClassMember) IsPublic() bool {
	return c.accessFlags&ACC_PUBLIC != 0
}
func (c *ClassMember) IsPrivate() bool {
	return c.accessFlags&ACC_PRIVATE != 0
}

func (c *ClassMember) IsProtected() bool {
	return c.accessFlags&ACC_PROTECTED != 0
}

func (c *ClassMember) IsStatic() bool {
	return c.accessFlags&ACC_STATIC != 0
}

func (c *ClassMember) IsFinal() bool {
	return c.accessFlags&ACC_FINAL != 0
}
func (c *ClassMember) IsSynthetic() bool {
	return c.accessFlags&ACC_SYNTHETIC != 0
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// getters
func (self *ClassMember) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassMember) Signature() string {
	return self.signature
}

func (self *ClassMember) AnnotationData() []byte {
	return self.annotationData
}
