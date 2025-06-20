package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (u *UnparsedAttribute) readInfo(reader *ClassReader) {

	u.info = reader.readBytes(u.length)
}
