package classfile

type DeprecatedAttributeInfo struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }
type MarkerAttribute struct{}

func (m MarkerAttribute) readInfo(reader *ClassReader) {
	// do nothing
}
