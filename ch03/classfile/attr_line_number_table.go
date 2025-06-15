package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (l *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberLength := reader.readUnit16()
	lineNumberTable := make([]*LineNumberTableEntry, lineNumberLength)
	for i := range lineNumberTable {
		lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUnit16(),
			lineNumber: reader.readUnit16(),
		}
	}
}
