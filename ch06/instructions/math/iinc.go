package math

import (
	"github.com/Jucunqi/jvmgo/ch06/instructions/base"
	"github.com/Jucunqi/jvmgo/ch06/rtda"
)

// IINC 从局部变量表中的int变量增加常量值，局部变量表索引和常量值都由指令的操作数提供
type IINC struct {
	Index uint
	Const int32
}

func (I *IINC) FetchOperands(reader *base.BytecodeReader) {

	I.Index = uint(reader.ReadUInt8())
	I.Const = int32(reader.ReadInt8())
}

func (I *IINC) Execute(frame *rtda.Frame) {
	vars := frame.LocalVars()
	val := vars.GetInt(I.Index)
	val += I.Const
	vars.SetInt(I.Index, val)
}
