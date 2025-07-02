package heap

func (c *Class) isAssignableFrom(other *Class) bool {
	s, t := other, c

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is an interface
				return s.IsImplements(t)
			}
		} else {
			// s is an interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is an interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is an array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is an interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is an array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
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

// IsSubInterfaceOf c extends iface
func (c *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range c.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// IsSuperClassOf other extends c
func (c *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(c)
}

// iface extends self
func (c *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(c)
}
