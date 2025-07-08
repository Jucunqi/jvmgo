package io

import (
	"github.com/Jucunqi/jvmgo/ch11/native"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

const fis = "java/io/FileInputStream"

func init() {
	native.Register(fis, "initIDs", "()V", initIDs)
}

func initIDs(frame *rtda.Frame) {
	thread := frame.Thread()
	thread.PopFrame()
}
