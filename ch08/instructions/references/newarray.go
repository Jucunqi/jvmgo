package references

import (
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
	"github.com/Jucunqi/jvmgo/ch08/rtda/heap"
)

const (
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

type NEW_ARRAY struct {
	atype uint8
}

func (n *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	n.atype = reader.ReadUInt8()
}

func (n *NEW_ARRAY) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出元素，代表数组长度
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	// 获取类加载器对象
	classLoder := frame.Method().Class().Loader()

	// 解析数组类
	arrClass := getPrimitiveArrayClass(classLoder, n.atype)

	// 创建数组对象
	arr := arrClass.NewArray(uint(count))

	// 压入操作数栈
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loder *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loder.LoadClass("[Z")
	case AT_BYTE:
		return loder.LoadClass("[B")
	case AT_CHAR:
		return loder.LoadClass("[C")
	case AT_SHORT:
		return loder.LoadClass("[S")
	case AT_INT:
		return loder.LoadClass("[I")
	case AT_LONG:
		return loder.LoadClass("[J")
	case AT_FLOAT:
		return loder.LoadClass("[F")
	case AT_DOUBLE:
		return loder.LoadClass("[D")
	default:
		panic("Invalid atype!")

	}
}
