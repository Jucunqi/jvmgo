package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (c *ConstantStringInfo) readInfo(reader *ClassReader) {
	c.stringIndex = reader.readUnit16()
}

func (c *ConstantStringInfo) String() string {
	return c.cp.getUtf8(c.stringIndex)
}
