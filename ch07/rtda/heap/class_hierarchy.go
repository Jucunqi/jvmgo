package heap

func (c *Class) isAssignableFrom(other *Class) bool {

	s, t := other, c
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}

}

func (self *Class) isSubClassOf(t *Class) bool {

	for c := self.superClass; c != nil; c = c.superClass {
		if c == t {
			return true
		}
	}
	return false
}

func (self *Class) isImplements(t *Class) bool {

	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == t || i.isSubInterfaceOf(t) {
				return true
			}
		}
	}
	return false
}
