package rtda

import (
	"github.com/Jucunqi/jvmgo/ch11/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(maxSize uint) LocalVars {

	if maxSize > 0 {
		return make([]Slot, maxSize)
	}
	return nil
}

func (l LocalVars) SetInt(index uint, val int32) {

	l[index].num = val
}

func (l LocalVars) GetInt(index uint) int32 {

	return l[index].num
}

func (l LocalVars) SetFloat(index uint, val float32) {

	bits := math.Float32bits(val)
	l[index].num = int32(bits)
}

func (l LocalVars) GetFloat(index uint) float32 {

	bits := l[index].num
	return math.Float32frombits(uint32(bits))
}

func (l LocalVars) SetLong(index uint, val int64) {

	l[index].num = int32(val)
	l[index+1].num = int32(val >> 32)
}

func (l LocalVars) GetLong(index uint) int64 {

	low := uint32(l[index].num)
	high := uint32(l[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (l LocalVars) SetDouble(index uint, val float64) {

	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}

func (l LocalVars) GetDouble(index uint) float64 {

	long := l.GetLong(index)
	bits := uint64(long)
	return math.Float64frombits(bits)
}
func (l LocalVars) SetRef(index uint, val *heap.Object) {

	l[index].ref = val
}

func (l LocalVars) GetRef(index uint) *heap.Object {

	return l[index].ref
}

func (l LocalVars) SetSlot(index uint, slot Slot) {
	l[index] = slot
}

func (l LocalVars) GetThis() *heap.Object {
	return l.GetRef(0)
}
