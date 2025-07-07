package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}
type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (l *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUnit16()
	localVariableTable := make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range localVariableTable {
		localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUnit16(),
			length:          reader.readUnit16(),
			nameIndex:       reader.readUnit16(),
			descriptorIndex: reader.readUnit16(),
			index:           reader.readUnit16(),
		}
	}
}
