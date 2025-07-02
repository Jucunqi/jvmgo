package references

import (
	"github.com/Jucunqi/jvmgo/ch09/instructions/base"
	"github.com/Jucunqi/jvmgo/ch09/rtda"
	"github.com/Jucunqi/jvmgo/ch09/rtda/heap"
)

type MULTI_ANEW_ARRAY struct {
	index      uint16 // 常量池索引，通过索引获取类符号引用
	dimensions uint8  // 表示数组纬度
}

func (m *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	m.index = reader.ReadUInt16()
	m.dimensions = reader.ReadUInt8()
}

func (m *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {

	// 获取常量池
	cp := frame.Method().Class().ConstantPool()

	// 根据常量池索引，获取数组类型符号引用
	classRef := cp.GetConstant(uint(m.index)).(*heap.ClassRef)

	// 解析类
	arrClass := classRef.ResolveClass()

	// 获取操作数栈
	stack := frame.OperandStack()

	// 根据纬度，弹出每个纬度的数组长度
	counts := popAndCheckCounts(stack, int(m.dimensions))

	// 创建多维数组
	arr := newMultiDimensionalArray(counts, arrClass)

	// 压入操作数栈
	stack.PushRef(arr)
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {

	// 创建数组对象
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	// 如果大于一维，遍历数组元素
	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			// 每个元素递归再次创建数组
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}
	return arr
}

func popAndCheckCounts(stack *rtda.OperandStack, dimension int) []int32 {

	result := make([]int32, dimension)

	// 循环获取栈顶元素，为每个纬度的数组长度赋值
	for i := dimension - 1; i >= 0; i-- {
		result[i] = stack.PopInt()
		if result[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return result
}
