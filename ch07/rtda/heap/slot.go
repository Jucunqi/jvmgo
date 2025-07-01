package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	return make([]Slot, slotCount)
}

func (self Slots) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

func (self Slots) SetFloat(index uint, val float32) {
	self[index].num = int32(val)
}

func (self Slots) GetFloat(index uint) float32 {
	return float32(self[index].num)
}

func (self Slots) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

func (self Slots) GetLong(index uint) int64 {
	return int64(uint64(self[index].num) | (uint64(self[index+1].num) << 32))
}

func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self Slots) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
}
