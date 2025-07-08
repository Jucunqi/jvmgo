package heap

func LookupMethodInClass(class *Class, name string, descriptor string) *Method {

	// 以此遍历当前类及其父类
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.Name() == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}
