package base

import "github.com/Jucunqi/jvmgo/ch09/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction 用于封装无操作数指令
type NoOperandsInstruction struct {
}

// BranchInstruction 用于封装跳转指令
type BranchInstruction struct {
	Offset int
}

// Index8Instruction 封装从局部变量表从获取或存储的指令
type Index8Instruction struct {
	Index uint
}

// Index16Instruction 封装从局部变量表从获取或存储的指令 对于某些占用两个位置的变量 比如 double long
type Index16Instruction struct {
	Index uint
}

func (n *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

	// noting to do
}

func (b *BranchInstruction) FetchOperands(reader *BytecodeReader) {

	b.Offset = int(reader.ReadInt16())

}

func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {

	i.Index = uint(int(reader.ReadInt8()))

}

func (i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(int(reader.ReadInt16()))
}
