package heap

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
