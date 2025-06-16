package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

func (c *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUnit16()
	c.nameAndTypeIndex = reader.readUnit16()
}

func (c *ConstantMemberrefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}
func (c *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}
