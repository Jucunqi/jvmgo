package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (c *CodeAttribute) readInfo(reader *ClassReader) {
	c.maxStack = reader.readUnit16()
	c.maxLocals = reader.readUnit16()
	codeLen := reader.readUnit32()
	c.code = reader.readBytes(codeLen)
	c.exceptionTable = readExceptionTable(reader)
	c.attributes = readAttributes(reader, c.cp)
}

func (c *CodeAttribute) Code() []byte {

	return c.code
}

func (c *CodeAttribute) MaxLocals() uint16 {
	return c.maxLocals
}

func (c *CodeAttribute) MaxStack() uint16 {
	return c.maxStack
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUnit16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUnit16(),
			endPc:     reader.readUnit16(),
			handlerPc: reader.readUnit16(),
			catchType: reader.readUnit16(),
		}
	}
	return exceptionTable
}
