package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUnit16()
	c.descriptorIndex = reader.readUnit16()
}
