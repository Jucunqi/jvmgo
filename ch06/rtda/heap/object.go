package heap

type Object struct {
	class  *Class
	fields Slots
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}

func (o *Object) Fields() Slots {
	return o.fields
}
