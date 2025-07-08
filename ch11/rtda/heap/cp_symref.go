package heap

// 符号引用
type SymRef struct {
	cp        *ConstantPool // 常量池
	className string        // 类名
	class     *Class        // 类
}

func (s *SymRef) ResolveClass() *Class {

	if s.class == nil {
		s.resolveClassRef()
	}
	return s.class
}

func (s *SymRef) resolveClassRef() {
	d := s.cp.class
	c := d.loader.LoadClass(s.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	s.class = c
}
