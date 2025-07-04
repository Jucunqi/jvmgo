package base

import "github.com/Jucunqi/jvmgo/ch10/rtda"

func Branch(frame *rtda.Frame, offect int) {
	pc := frame.Thread().PC()
	nextPC := pc + offect
	frame.SetNextPC(nextPC)
}
