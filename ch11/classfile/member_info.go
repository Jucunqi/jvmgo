package classfile

type MemberInfo struct {
	cp              ConstantPool    // 常量池
	accessFlags     uint16          // 访问标识符
	nameIndex       uint16          // 名称索引
	descriptorIndex uint16          // 描述符索引
	attributes      []AttributeInfo // 属性集合
}

// 读取Member集合
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUnit16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取一个Member
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {

	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUnit16(),
		nameIndex:       reader.readUnit16(),
		descriptorIndex: reader.readUnit16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}
func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}
func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}

func (m *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attr := range m.attributes {
		switch attr.(type) {
		case *CodeAttribute:
			return attr.(*CodeAttribute)
		}
	}
	return nil
}

func (m *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attr := range m.attributes {
		switch attr.(type) {
		case *ConstantValueAttribute:
			return attr.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) ExceptionsAttribute() *ExceptionAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ExceptionAttribute:
			return attrInfo.(*ExceptionAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) RuntimeVisibleAnnotationsAttributeData() []byte {
	return self.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (self *MemberInfo) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return self.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (self *MemberInfo) AnnotationDefaultAttributeData() []byte {
	return self.getUnparsedAttributeData("AnnotationDefault")
}

func (self *MemberInfo) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttribute:
			unparsedAttr := attrInfo.(*UnparsedAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
