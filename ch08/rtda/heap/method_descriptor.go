package heap

type MethodDescriptor struct {
	parameterTypes []string // 方法形参类型列表
	returnType     string   // 返回类型
}

func (m *MethodDescriptor) addParameterType(paramType string) {

	// 创建 扩容 逻辑
	pLen := len(m.parameterTypes)
	if pLen == cap(m.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, m.parameterTypes)
		m.parameterTypes = s
	}

	// 添加元素
	m.parameterTypes = append(m.parameterTypes, paramType)
}
