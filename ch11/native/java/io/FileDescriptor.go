package io

import (
	"github.com/Jucunqi/jvmgo/ch11/native"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

const fd = "java/io/FileDescriptor"

func init() {
	native.Register(fd, "set", "(I)J", set)
	native.Register(fd, "initIDs", "()V", descInitIDs)
}

func descInitIDs(frame *rtda.Frame) {

	// todo:
	frame.Thread().PopFrame()
}

// private static native long set(int d);
// (I)J
func set(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
