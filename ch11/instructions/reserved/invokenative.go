package reserved

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/native"
	_ "github.com/Jucunqi/jvmgo/ch11/native/java/io"
	_ "github.com/Jucunqi/jvmgo/ch11/native/java/lang"
	_ "github.com/Jucunqi/jvmgo/ch11/native/java/security"
	_ "github.com/Jucunqi/jvmgo/ch11/native/sun/misc"
	_ "github.com/Jucunqi/jvmgo/ch11/native/sun/reflect"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
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
