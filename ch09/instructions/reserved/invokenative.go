package reserved

import (
	"github.com/Jucunqi/jvmgo/ch09/instructions/base"
	"github.com/Jucunqi/jvmgo/ch09/native"
	_ "github.com/Jucunqi/jvmgo/ch09/native/java/lang"
	"github.com/Jucunqi/jvmgo/ch09/rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (i *INVOKE_NATIVE) Execute(frame *rtda.Frame) {

	// 获取方法名，类名，描述符
	method := frame.Method()
	methodName := method.Name()
	className := method.Class().Name()
	descriptor := method.Descriptor()

	// 根据上述参数，去本地方法注册表中匹配本地方法
	nativeMethod := native.FindNativeMethod(className, methodName, descriptor)

	// 非空校验
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + "." + descriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	// 执行本地方法
	nativeMethod(frame)
}
