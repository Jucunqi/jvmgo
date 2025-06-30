package heap

func (c *Class) isAssignableFrom(other *Class) bool {

	s, t := other, c
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}

}

func (c *Class) IsSubClassOf(t *Class) bool {

	for c := c.superClass; c != nil; c = c.superClass {
		if c == t {
			return true
		}
	}
	return false
}

func (c *Class) IsImplements(t *Class) bool {

	for c := c; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == t || i.isSubInterfaceOf(t) {
				return true
			}
		}
	}
	return false
}

// c extends iface
func (c *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range c.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// other extends c
func (c *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(c)
}
