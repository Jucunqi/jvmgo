package lang

import (
	"github.com/Jucunqi/jvmgo/ch11/native"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
)

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

// 不考虑断言，直接返回false
func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}

// private native String getName0();
func getName0(frame *rtda.Frame) {

	// 从局部变量表中获取当前对象
	this := frame.LocalVars().GetThis()

	// 通过类对象的extra属性可以获取类的信息
	class := this.Extra().(*heap.Class)

	// 名称转换java类名，/ -> .
	name := class.JavaName()

	// 转换为String对象
	jString := heap.JString(class.Loader(), name)

	// 压入操作数栈
	frame.OperandStack().PushRef(jString)
}

// 获取基本类型对应的Class对象
// 对应java代码 Class.getPrimitiveClass("int");
func getPrimitiveClass(frame *rtda.Frame) {

	// 从局部变量表中获取类型参数
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}
