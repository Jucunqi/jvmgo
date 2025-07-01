package classfile

type ConstantUtf8Info struct {
	str string
}

func (c *ConstantUtf8Info) readInfo(reader *ClassReader) {

	length := uint32(reader.readUnit16())
	bytes := reader.readBytes(length)
	c.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
