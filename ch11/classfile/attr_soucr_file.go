package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (s *SourceFileAttribute) readInfo(reader *ClassReader) {
	s.sourceFileIndex = reader.readUnit16()
}

func (s *SourceFileAttribute) FileName() string {
	return s.cp.getUtf8(s.sourceFileIndex)
}
