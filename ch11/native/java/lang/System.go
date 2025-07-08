package lang

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/native"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
	"runtime"
	"time"
)

const jlSystem = "java/lang/System"

func init() {
	native.Register(jlSystem, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register(jlSystem, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	native.Register(jlSystem, "identityHashCode", "(Ljava/lang/Object;)I", identityHashCode)
	native.Register(jlSystem, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	native.Register(jlSystem, "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	native.Register(jlSystem, "setErr0", "(Ljava/io/PrintStream;)V", setErr0)
	native.Register(jlSystem, "currentTimeMillis", "()J", currentTimeMillis)

}

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

func setOut0(frame *rtda.Frame) {
	out := frame.LocalVars().GetRef(0)
	sysClass := frame.Method().Class()
	sysClass.SetRefVar("out", "Ljava/io/PrintStream;", out)

}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
	vars := frame.LocalVars()
	props := vars.GetRef(0)

	stack := frame.OperandStack()
	stack.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread()
	for key, val := range _sysProps() {
		jKey := heap.JString(frame.Method().Class().Loader(), key)
		jVal := heap.JString(frame.Method().Class().Loader(), val)
		ops := rtda.NewOperandStack(3)
		ops.PushRef(props)
		ops.PushRef(jKey)
		ops.PushRef(jVal)
		shimFrame := rtda.NewShimFrame(thread, ops)
		thread.PushFrame(shimFrame)

		base.InvokeMethod(shimFrame, setPropMethod)
	}
}
func _sysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https://github.com/zxh0/jvm.go",
		"java.home":            "/Library/Java/JavaVirtualMachines/jdk-1.8.jdk/Contents/Home",
		"java.class.version":   "52.0",
		"java.class.path":      ".",
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "1.1",          // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	in := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("in", "Ljava/io/InputStream;", in)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	err := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("err", "Ljava/io/PrintStream;", err)
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *rtda.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	stack := frame.OperandStack()
	stack.PushLong(millis)
}
