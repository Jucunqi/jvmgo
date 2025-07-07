package lang

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch10/native"
	"github.com/Jucunqi/jvmgo/ch10/rtda"
	"github.com/Jucunqi/jvmgo/ch10/rtda/heap"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

type StackTraceElement struct {
	fileName   string // 文件
	className  string // 类名
	methodName string // 方法名
	lineNumber int    // 行号
}

// private native Throwable fillInStackTrace(int dummy);
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {

	// 计算需要跳过的栈帧（fillInStackTrace和fillInStackTrace(int)这两个方法）+ 异常对象的构造函数-具体多找帧根据异常类的继承层级
	skip := distantToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]

	// 创建栈信息数组
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {

		// 创建栈信息对象
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

func distantToObject(class *heap.Class) int {

	// 计算需要跳过的帧数
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func (s *StackTraceElement) String() string {
	return fmt.Sprintf("at %s.%s{%s:%d}",
		s.className, s.methodName, s.fileName, s.lineNumber)
}
