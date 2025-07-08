package references

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
	"reflect"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (A *ATHROW) Execute(frame *rtda.Frame) {

	// 从操作数栈弹出异常对象引用
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	// 否则查看是否可以找到并跳转到异常处理代码
	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUnCaughtException(thread, ex)
	}
}

// 若找不到异常处理逻辑，执行这个方法
func handleUnCaughtException(thread *rtda.Thread, ex *heap.Object) {

	// 清空虚拟机栈
	thread.ClearStack()

	// 获取详细信息属性
	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)
	stes := reflect.ValueOf(ex.Extra())

	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface{ String() string })
		println("\tat" + ste.String())
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {

		// 获取线程栈顶栈帧
		frame := thread.CurrentFrame()

		// 获取执行指令计数
		pc := frame.NextPC() - 1

		// 获取方法的异常处理pc
		handlerPc := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPc > 0 {

			// 如果异常处理pc>0 说明当前方法存在catch当前异常的逻辑
			stack := frame.OperandStack()

			// 清空操作数栈
			stack.Clear()

			// 把异常对象放入操作数栈
			stack.PushRef(ex)

			// 设置程序计数器执行行号
			frame.SetNextPC(handlerPc)
			return true
		}

		// 弹出当前栈帧
		thread.PopFrame()

		// 循环停止条件 - 线程中没有方法栈帧
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}
