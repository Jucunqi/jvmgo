package constants

import (
	"github.com/Jucunqi/jvmgo/ch05/instructions/base"
	"github.com/Jucunqi/jvmgo/ch05/rtda"
)

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (B *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	B.val = reader.ReadInt8()
}

func (B *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(B.val))
}

func (S *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	S.val = reader.ReadInt16()
}

func (S *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(S.val))
}
