package classfile

import (
	"fmt"
	"strconv"
)

type ClassFile struct {
	magic        uint32          // 魔数
	minorVersion uint16          // 小版本
	majorVersion uint16          // 大版本
	constantPool ConstantPool    // 常量池
	accessFlags  uint16          // 访问标识符
	thisClass    uint16          // 当前类
	superClass   uint16          // 父类
	interfaces   []uint16        //接口集合
	fields       []*MemberInfo   // 字段集合
	methods      []*MemberInfo   // 方法集合
	attributes   []AttributeInfo // 属性集合
}

// Parse 解析方法
func Parse(classData []byte) (cf *ClassFile, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return cf, nil
}

// 读取
func (c *ClassFile) read(reader *ClassReader) {
	c.readAndCheckMagic(reader)
	c.readAndCheckVersion(reader)
	c.constantPool = readConstantPool(reader)
	c.accessFlags = reader.readUnit16()
	c.thisClass = reader.readUnit16()
	c.superClass = reader.readUnit16()
	c.interfaces = reader.readUnit16s()
	c.fields = readMembers(reader, c.constantPool)
	c.methods = readMembers(reader, c.constantPool)
	c.attributes = readAttributes(reader, c.constantPool)
}

// 读取并检查魔数
func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {

	magic := reader.readUnit32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 读取并检查版本号
func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {
	c.minorVersion = reader.readUnit16()
	c.majorVersion = reader.readUnit16()
	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if c.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError: Current Version: " + strconv.Itoa(int(c.majorVersion)))
}

// MinorVersion 返回小版本号
func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}

// MajorVersion 返回大版本号
func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}

// ConstantPool 返回常量池对象
func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}

// AccessFlags 返回访问标识符
func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}

// Fields 返回属性
func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}

// Methods 返回方法
func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}

// ClassName 返回类名
func (c *ClassFile) ClassName() string {

	return c.constantPool.getClassName(c.thisClass)
}

// SuperClassName 返回父类名
func (c *ClassFile) SuperClassName() string {
	if c.superClass > 0 {
		return c.constantPool.getClassName(c.superClass)
	}
	return "" // 只有java.lang.Object 没有父类
}

// InterfaceNames 返回接口名
func (c *ClassFile) InterfaceNames() []string {

	interfaceNames := make([]string, len(c.interfaces))
	for i, cpIndex := range c.interfaces {
		interfaceNames[i] = c.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
