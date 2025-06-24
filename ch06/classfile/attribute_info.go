package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {

	attributeCount := reader.readUnit16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {

	attrNameIndex := reader.readUnit16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUnit32()
	attribute := newAttribute(attrName, attrLen, cp)
	attribute.readInfo(reader)
	return attribute
}

func newAttribute(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {

	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	//case "Depreated":
	//	return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
