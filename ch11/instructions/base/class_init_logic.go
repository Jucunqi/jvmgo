package base

import (
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {

	// 设置初始化标记位
	class.StartInit()

	// 计划执行clinit方法
	scheduleClinit(thread, class)

	// 初始化父类
	initSuperClass(thread, class)
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {

	clinit := class.GetClinitMethod()

	if clinit != nil {

		// 创建栈帧
		frame := thread.NewFrame(clinit)
		// 压入栈
		thread.PushFrame(frame)
	}
}
