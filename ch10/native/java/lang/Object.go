package lang

import (
	"github.com/Jucunqi/jvmgo/ch10/native"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

func getClass(frame *rtda.Frame) {

	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
