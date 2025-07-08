package extended

import (
	"github.com/Jucunqi/jvmgo/ch11/instructions/base"
	"github.com/Jucunqi/jvmgo/ch11/rtda"
)

type GOTO_W struct {
	offset int
}

func (G *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	G.offset = int(reader.ReadInt32())
}

func (G *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, G.offset)
}
