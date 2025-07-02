package constants

import (
	"github.com/Jucunqi/jvmgo/ch09/instructions/base"
	"github.com/Jucunqi/jvmgo/ch09/rtda"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}
type DCONST_0 struct {
	base.NoOperandsInstruction
}
type DCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_0 struct {
	base.NoOperandsInstruction
}
type FCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_M1 struct {
	base.NoOperandsInstruction
}
type ICONST_0 struct {
	base.NoOperandsInstruction
}
type ICONST_1 struct {
	base.NoOperandsInstruction
}
type ICONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_3 struct {
	base.NoOperandsInstruction
}
type ICONST_4 struct {
	base.NoOperandsInstruction
}
type ICONST_5 struct {
	base.NoOperandsInstruction
}
type LCONST_0 struct {
	base.NoOperandsInstruction
}
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (A *ACONST_NULL) Execute(frame *rtda.Frame) {

	frame.OperandStack().PushRef(nil)
}

func (D *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (D *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

func (F *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

func (F *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}
func (F *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

func (I *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (I *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

func (I *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}
func (I *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}
func (I *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}
func (I *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}
func (I *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

func (L *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

func (L *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
