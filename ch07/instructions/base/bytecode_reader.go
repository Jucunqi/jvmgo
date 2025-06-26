package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (b *BytecodeReader) Reset(code []byte, pc int) {

	b.code = code
	b.pc = pc
}

func (b *BytecodeReader) ReadUInt8() uint8 {

	i := b.code[b.pc]
	b.pc++
	return i
}

func (b *BytecodeReader) ReadInt8() int8 {

	return int8(b.ReadUInt8())
}

func (b *BytecodeReader) ReadUInt16() uint16 {

	byte1 := uint16(b.ReadUInt8())
	byte2 := uint16(b.ReadUInt8())
	return (byte1 << 8) | byte2
}
func (b *BytecodeReader) ReadInt16() int16 {

	return int16(b.ReadUInt16())
}

func (b *BytecodeReader) ReadInt32() int32 {

	byte1 := int32(b.ReadUInt8())
	byte2 := int32(b.ReadUInt8())
	byte3 := int32(b.ReadUInt8())
	byte4 := int32(b.ReadUInt8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | (byte4)

}

func (b *BytecodeReader) ReadInt32s(count int32) []int32 {
	is := make([]int32, count)
	for i := range is {
		is[i] = b.ReadInt32()
	}
	return is
}

// SkipPadding 跳过1-3位
func (b *BytecodeReader) SkipPadding() {

	i := b.pc % 4
	for i != 0 {
		b.ReadInt8()
	}
}

func (b *BytecodeReader) PC() int {
	return b.pc
}
