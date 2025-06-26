package main

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch07/rtda/heap"

	"github.com/Jucunqi/jvmgo/ch07/instructions"
	"github.com/Jucunqi/jvmgo/ch07/instructions/base"
	"github.com/Jucunqi/jvmgo/ch07/rtda"
)

func interpret(method *heap.Method) {

	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
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
		fmt.Printf("pc: %2d inst: %T %v\n", pc, instruction, instruction)
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
