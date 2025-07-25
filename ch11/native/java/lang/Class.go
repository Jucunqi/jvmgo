package lang

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/native"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
	"strings"
)

const jlClass = "java/lang/Class"

func init() {
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	native.Register(jlClass, "isInterface", "()Z", isInterface)
	native.Register(jlClass, "isPrimitive", "()Z", isPrimitive)
	native.Register(jlClass, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
	native.Register(jlClass, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	native.Register(jlClass, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructors0)
	native.Register(jlClass, "getModifiers", "()I", getModifiers)
	native.Register(jlClass, "getSuperclass", "()Ljava/lang/Class;", getSuperclass)
	native.Register(jlClass, "getInterfaces0", "()[Ljava/lang/Class;", getInterfaces0)
	native.Register(jlClass, "isArray", "()Z", isArray)
	native.Register(jlClass, "getDeclaredMethods0", "(Z)[Ljava/lang/reflect/Method;", getDeclaredMethods0)
	native.Register(jlClass, "getComponentType", "()Ljava/lang/Class;", getComponentType)
	native.Register(jlClass, "isAssignableFrom", "(Ljava/lang/Class;)Z", isAssignableFrom)
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

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsPrimitive())
}

// private static native Class<?> forName0(String name, boolean initialize,
//
//	                                    ClassLoader loader,
//	                                    Class<?> caller)
//	throws ClassNotFoundException;
//
// (Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;
func forName0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jName := vars.GetRef(0)
	initialize := vars.GetBoolean(1)
	//jLoader := vars.GetRef(2)

	goName := heap.GoString(jName)
	goName = strings.Replace(goName, ".", "/", -1)
	goClass := frame.Method().Class().Loader().LoadClass(goName)
	jClass := goClass.JClass()

	if initialize && !goClass.InitStarted() {
		// undo forName0
		thread := frame.Thread()
		frame.SetNextPC(thread.PC())
		// init class
		base.InitClass(thread, goClass)
	} else {
		stack := frame.OperandStack()
		stack.PushRef(jClass)
	}
}

// public native int getModifiers();
// ()I
func getModifiers(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	modifiers := class.AccessFlags()

	stack := frame.OperandStack()
	stack.PushInt(int32(modifiers))
}

// public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	superClass := class.SuperClass()

	stack := frame.OperandStack()
	if superClass != nil {
		stack.PushRef(superClass.JClass())
	} else {
		stack.PushRef(nil)
	}
}

// private native Class<?>[] getInterfaces0();
// ()[Ljava/lang/Class;
func getInterfaces0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	interfaces := class.Interfaces()
	classArr := toClassArr(class.Loader(), interfaces)

	stack := frame.OperandStack()
	stack.PushRef(classArr)
}

// public native boolean isArray();
// ()Z
func isArray(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsArray())
}

// public native Class<?> getComponentType();
// ()Ljava/lang/Class;
func getComponentType(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	componentClass := class.ComponentClass()
	componentClassObj := componentClass.JClass()

	stack := frame.OperandStack()
	stack.PushRef(componentClassObj)
}

// public native boolean isAssignableFrom(Class<?> cls);
// (Ljava/lang/Class;)Z
func isAssignableFrom(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	cls := vars.GetRef(1)

	thisClass := this.Extra().(*heap.Class)
	clsClass := cls.Extra().(*heap.Class)
	ok := thisClass.IsAssignableFrom(clsClass)

	stack := frame.OperandStack()
	stack.PushBoolean(ok)
}
