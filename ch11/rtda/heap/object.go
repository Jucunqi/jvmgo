package heap

import "unsafe"

type Object struct {
	class *Class
	data  interface{} // 标识任何类型
	extra interface{} // 作为类对象时，用于记录额外信息
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}

func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) SetRefVar(name string, descriptor string, ref *Object) {

	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (o *Object) GetRefVar(name string, descriptor string) *Object {

	// 根据名称描述附匹配属性
	field := o.class.getField(name, descriptor, false)

	// 获取对象所有属性
	slots := o.data.(Slots)

	// 获取属性值
	return slots.GetRef(field.slotId)

}

func (o *Object) Extra() interface{} {
	return o.extra
}

func (o *Object) SetExtra(extra interface{}) {

	o.extra = extra
}

// HashCode 返回对象的hashcode
// 使用对象的地址作为hashcode，这符合Java中System.identityHashCode()的语义
func (o *Object) HashCode() int32 {
	// 使用unsafe.Pointer获取对象的地址，然后转换为int32
	ptr := unsafe.Pointer(o)
	hash := uintptr(ptr)
	// 将64位地址转换为32位hashcode
	return int32(hash ^ (hash >> 32))
}
