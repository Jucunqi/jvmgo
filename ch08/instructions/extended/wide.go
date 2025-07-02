package extended

import (
	"github.com/Jucunqi/jvmgo/ch08/instructions/base"
	"github.com/Jucunqi/jvmgo/ch08/instructions/loads"
	"github.com/Jucunqi/jvmgo/ch08/instructions/math"
	"github.com/Jucunqi/jvmgo/ch08/instructions/stores"
	"github.com/Jucunqi/jvmgo/ch08/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (W *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUInt8()
	switch opcode {
	case 0x15: // iload
		inst := &loads.ILOAD{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x16: // lload
		inst := &loads.LLOAD{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x17: // fload
		inst := &loads.FLOAD{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x18: // dload
		inst := &loads.DLOAD{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x19: // aload
		inst := &loads.ALOAD{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x36: // istore
		inst := &stores.ISTORE{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x37: // lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x38: // fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x39: // dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x3a: // astore
		inst := &stores.ASTORE{}
		inst.Index = uint(int(reader.ReadInt16()))
		W.modifiedInstruction = inst
		break
	case 0x84: // iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUInt16())
		inst.Const = int32(reader.ReadInt16())
		W.modifiedInstruction = inst
		break
	case 0xa9:
		panic("Unsupported opcode: 0xa9!")
	}
}

func (W *WIDE) Execute(frame *rtda.Frame) {
	W.modifiedInstruction.Execute(frame)
}
