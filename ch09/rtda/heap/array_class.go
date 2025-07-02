package heap

func (c *Class) NewArray(count uint) *Object {

	if !c.IsArray() {
		panic("Not array class: " + c.name)
	}
	switch c.Name() {
	case "[Z":
		return &Object{c, make([]int8, count)}
	case "[B":
		return &Object{c, make([]int8, count)}
	case "[C":
		return &Object{c, make([]uint16, count)}
	case "[S":
		return &Object{c, make([]int16, count)}
	case "[I":
		return &Object{c, make([]int32, count)}
	case "[J":
		return &Object{c, make([]int64, count)}
	case "[F":
		return &Object{c, make([]float32, count)}
	case "[D":
		return &Object{c, make([]int64, count)}
	default:
		return &Object{c, make([]*Object, count)}

	}
}

func (c *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(c.name)
	return c.loader.LoadClass(componentClassName)
}
