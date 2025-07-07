package loads

import (
	"github.com/Jucunqi/jvmgo/ch10/instructions/base"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
	"github.com/Jucunqi/jvmgo/ch10/rtda/heap"
)

type AALOAD struct {
	base.NoOperandsInstruction
}

type BALOAD struct {
	base.NoOperandsInstruction
}

type CALOAD struct {
	base.NoOperandsInstruction
}

type DALOAD struct {
	base.NoOperandsInstruction
}

type FALOAD struct {
	base.NoOperandsInstruction
}

type IALOAD struct {
	base.NoOperandsInstruction
}

type LALOAD struct {
	base.NoOperandsInstruction
}

type SALOAD struct {
	base.NoOperandsInstruction
}

func (a *AALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 索引越界验证
	refs := arrRef.Refs()
	checkIndex(len(refs), index)

	// 结果压入栈顶
	stack.PushRef(refs[index])
}

func (b *BALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 获取所有数组元素，并索引越界验证
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)

	// 结果压入栈顶
	stack.PushInt(int32(bytes[index]))
}

func (c *CALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 获取所有数组元素，并索引越界验证
	chars := arrRef.Chars()
	checkIndex(len(chars), index)

	// 结果压入栈顶
	stack.PushInt(int32(chars[index]))
}

func (d *DALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 获取所有数组元素，并索引越界验证
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)

	// 结果压入栈顶
	stack.PushDouble(doubles[index])
}

func (f *FALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 获取所有数组元素，并索引越界验证
	floats := arrRef.Floats()
	checkIndex(len(floats), index)

	// 结果压入栈顶
	stack.PushFloat(floats[index])
}

func (i *IALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 获取所有数组元素，并索引越界验证
	ints := arrRef.Ints()
	checkIndex(len(ints), index)

	// 结果压入栈顶
	stack.PushInt(ints[index])
}

func (l *LALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 获取所有数组元素，并索引越界验证
	longs := arrRef.Longs()
	checkIndex(len(longs), index)

	// 结果压入栈顶
	stack.PushLong(longs[index])
}

func (s *SALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 获取所有数组元素，并索引越界验证
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)

	// 结果压入栈顶
	stack.PushInt(int32(shorts[index]))
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
