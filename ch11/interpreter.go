package main

import (
	"fmt"

	"github.com/Jucunqi/jvmgo/ch11/instructions"
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

func interpret(thread *rtda.Thread, logInst bool) {

	// 异常处理
	defer catchErr(thread)

	// 循环执行指令
	loop(thread, logInst)
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		// 获取当前栈顶栈帧
		frame := thread.CurrentFrame()

		// 获取程序计数器
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadInt8()
		inst := instructions.NewInstruction(byte(opcode))
		inst.FetchOperands(reader)

		// 根据标识判断是否打印指令日志
		if logInst {
			logInstruction(frame, inst)

		}

		// execute
		frame.SetNextPC(reader.PC())
		inst.Execute(frame)

		// 结束标识
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func catchErr(thread *rtda.Thread) {

	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
