package io

import (
	"github.com/Jucunqi/jvmgo/ch11/native"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
	"os"
	"unsafe"
)

const fos = "java/io/FileOutputStream"

func init() {
	native.Register(fos, "writeBytes", "([BIIZ)V", writeBytes)
	native.Register(fos, "initIDs", "()V", fosInitIDs)

}

func fosInitIDs(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func writeBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// this := vars.GetRef(0)
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)
	//append := vars.GetInt(4)
	jBytes := b.Data().([]int8)
	goBytes := castInt8sToUint8s(jBytes)
	goBytes = goBytes[off : off+len]
	os.Stdout.Write(goBytes)
}

func castInt8sToUint8s(jBytes []int8) (goBytes []byte) {
	ptr := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(ptr))
	return
}
