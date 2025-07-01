package rtda

import (
	"github.com/Jucunqi/jvmgo/ch08/rtda/heap"
	"math"
)

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(stackSize uint) *OperandStack {
	if stackSize > 0 {
		return &OperandStack{slots: make([]Slot, stackSize)}
	}
	return nil
}

func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

func (o *OperandStack) PopInt() int32 {

	o.size--
	return o.slots[o.size].num
}

func (o *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	o.slots[o.size].num = int32(bits)
	o.size++
}

func (o *OperandStack) PopFloat() float32 {

	o.size--
	bits := uint32(o.slots[o.size].num)
	return math.Float32frombits(bits)
}
func (o *OperandStack) PushLong(val int64) {
	low := int32(val)
	o.slots[o.size].num = low
	o.size++
	high := int32(val >> 32)
	o.slots[o.size].num = high
	o.size++
}

func (o *OperandStack) PopLong() int64 {

	o.size -= 2
	low := uint32(o.slots[o.size].num)
	high := uint32(o.slots[o.size+1].num)
	return int64(high)<<32 | int64(low)
}
func (o *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	o.PushLong(int64(bits))
}

func (o *OperandStack) PopDouble() float64 {

	bits := uint64(o.PopLong())
	return math.Float64frombits(bits)
}

func (o *OperandStack) PushRef(val *heap.Object) {

	o.slots[o.size].ref = val
	o.size++
}

func (o *OperandStack) PopRef() *heap.Object {

	o.size--
	ref := o.slots[o.size].ref
	o.slots[o.size].ref = nil
	return ref
}

func (o *OperandStack) PushSlot(slot Slot) {
	o.slots[o.size] = slot
	o.size++
}

func (o *OperandStack) PopSlot() Slot {
	o.size--
	return o.slots[o.size]
}

func (o *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return o.slots[o.size-1-n].ref
}
