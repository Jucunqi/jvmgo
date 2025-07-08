package native

import (
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

type NativeMethod func(frame *rtda.Frame)

// 存储本地方法实现
var registry = map[string]NativeMethod{}

// Register 本地方法注册逻辑
func Register(className string, methodName string, methodDescriptor string, method NativeMethod) {

	// 定义缓存key
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// FindNativeMethod 查找本地方法
func FindNativeMethod(className string, methodName string, methodDescriptor string) NativeMethod {

	// 根据key查询本地方法
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}

	// 如果方法名为‘registerNatives’，不做任何操作，这个方法是java注册本地方法用的，这里我们自己实现注册逻辑
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}
