package heap

func (c *Class) NewArray(count uint) *Object {

	if !c.IsArray() {
		panic("Not array class: " + c.name)
	}
	switch c.Name() {
	case "[Z":
		return &Object{c, make([]int8, count), nil}
	case "[B":
		return &Object{c, make([]int8, count), nil}
	case "[C":
		return &Object{c, make([]uint16, count), nil}
	case "[S":
		return &Object{c, make([]int16, count), nil}
	case "[I":
		return &Object{c, make([]int32, count), nil}
	case "[J":
		return &Object{c, make([]int64, count), nil}
	case "[F":
		return &Object{c, make([]float32, count), nil}
	case "[D":
		return &Object{c, make([]int64, count), nil}
	default:
		return &Object{c, make([]*Object, count), nil}

	}
}

func (c *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(c.name)
	return c.loader.LoadClass(componentClassName)
}
