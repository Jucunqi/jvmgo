package main

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch05/classfile"
	"github.com/Jucunqi/jvmgo/ch05/instructions"
	"github.com/Jucunqi/jvmgo/ch05/instructions/base"
	"github.com/Jucunqi/jvmgo/ch05/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {

	// 获取方法中的code属性
	codeAttr := methodInfo.CodeAttribute()

	// 从属性中解析最大局部变量表，最大栈，code属性
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	// 创建一个线程
	thread := rtda.NewThread()

	// 创建一个栈帧
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))

	// 栈帧压入栈
	thread.PushFrame(frame)

	// 异常处理
	defer catchErr(frame)

	// 循环解析指令并执行
	loop(thread, bytecode)

}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadInt8()
		instruction := instructions.NewInstruction(byte(opcode))
		instruction.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		instruction.Execute(frame)
	}
}

func catchErr(frame *rtda.Frame) {

	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
