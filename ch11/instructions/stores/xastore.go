package stores

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
)

type AASTORE struct {
	base.NoOperandsInstruction
}

type BASTORE struct {
	base.NoOperandsInstruction
}

type CASTORE struct {
	base.NoOperandsInstruction
}

type DASTORE struct {
	base.NoOperandsInstruction
}

type FASTORE struct {
	base.NoOperandsInstruction
}

type IASTORE struct {
	base.NoOperandsInstruction
}

type LASTORE struct {
	base.NoOperandsInstruction
}

type SASTORE struct {
	base.NoOperandsInstruction
}

func (i *IASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopInt()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	ints := arrRef.Ints()
	checkIndex(len(ints), index)

	// 数组赋值
	ints[index] = int32(val)
}

func (a *AASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopRef()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	refs := arrRef.Refs()
	checkIndex(len(refs), index)

	// 数组赋值
	refs[index] = val
}

func (b *BASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopInt()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)

	// 数组赋值
	bytes[index] = int8(val)
}

func (c *CASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopInt()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	chars := arrRef.Chars()
	checkIndex(len(chars), index)

	// 数组赋值
	chars[index] = uint16(val)
}

func (d *DASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopDouble()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)

	// 数组赋值
	doubles[index] = val
}

func (f *FASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopFloat()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	floats := arrRef.Floats()
	checkIndex(len(floats), index)

	// 数组赋值
	floats[index] = val
}

func (l *LASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopLong()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	longs := arrRef.Longs()
	checkIndex(len(longs), index)

	// 数组赋值
	longs[index] = val
}

func (s *SASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopInt()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)

	// 数组赋值
	shorts[index] = int16(val)
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
