package lang

import (
	"github.com/Jucunqi/jvmgo/ch11/native"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
)

// 对应Java本地方法   public static native void arraycopy(Object src,  int  srcPos, Object dest, int destPos, int length);
func arraycopy(frame *rtda.Frame) {

	// 从局部变量表中拿到5个参数
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	// 非空校验
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	// 源数组和目标数组必须兼容
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}

	// 检查索引位置
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	// 数组拷贝
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src *heap.Object, dest *heap.Object) bool {

	srcClass := src.Class()
	descClass := dest.Class()

	// 必须都是数组
	if !srcClass.IsArray() || !descClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() ||
		descClass.ComponentClass().IsPrimitive() {
		return srcClass == descClass
	}
	return true
}

func identityHashCode(frame *rtda.Frame) {
	vars := frame.LocalVars()
	ref := vars.GetRef(0)
	hash := ref.HashCode()
	frame.OperandStack().PushInt(hash)
}

func init() {
	native.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register("java/lang/System", "identityHashCode", "(Ljava/lang/Object;)I", identityHashCode)
}
