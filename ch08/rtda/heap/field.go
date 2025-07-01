package heap

import "github.com/Jucunqi/jvmgo/ch08/classfile"

type Field struct {
	ClassMember          //继承自ClassMember
	constValueIndex uint //常量值索引
	slotId          uint //槽位ID
}

func newFields(class *Class, infos []*classfile.MemberInfo) []*Field {

	fields := make([]*Field, len(infos))
	for i := range fields {
		fields[i] = &Field{}
		fields[i].copyMemberInfo(infos[i])
		fields[i].class = class
		fields[i].copyAttributes(infos[i])
	}
	return fields
}

func (c *Field) IsVolatile() bool {
	return c.accessFlags&ACC_VOLATILE != 0
}

func (c *Field) IsTransient() bool {
	return c.accessFlags&ACC_TRANSIENT != 0
}

func (c *Field) IsEnum() bool {
	return c.accessFlags&ACC_ENUM != 0
}

func (c *Field) isLongOrDouble() bool {
	return c.descriptor == "J" || c.descriptor == "D"
}

func (c *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		c.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (c *Field) ConstValueIndex() uint {
	return c.constValueIndex
}

func (c *Field) SlotId() uint {
	return c.slotId
}

func (c *Field) Descriptor() string {
	return c.descriptor
}

func (c *Field) Class() *Class {
	return c.class
}
