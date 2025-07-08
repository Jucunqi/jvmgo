package classfile

type ConstantPool []ConstantInfo

// 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {

	cpCount := int(reader.readUnit16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { // 索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 如果是Long 或者 Double 占两个位置
		}
	}
	return cp
}

// 获取指定常量信息
func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {

	if cpInfo := c[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 读取名称和类型 (字段、方法)
func (c ConstantPool) getNameAndType(index uint16) (string, string) {

	ntInfo := c.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := c.getUtf8(ntInfo.nameIndex)
	_type := c.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 读取类名
func (c ConstantPool) getClassName(index uint16) string {
	classInfo := c.getConstantInfo(index).(*ConstantClassInfo)
	return c.getUtf8(classInfo.nameIndex)
}

// 读取utf8字符
func (c ConstantPool) getUtf8(index uint16) string {
	utf8Info := c.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
