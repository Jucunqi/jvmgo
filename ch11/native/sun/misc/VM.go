package misc

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/native"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
