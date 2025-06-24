package heap

import (
	"strings"

	"github.com/Jucunqi/jvmgo/ch06/classfile"
)

type Class struct {
	accessFlags       uint16        //访问权限
	name              string        //类名
	superClassName    string        //父类名
	interfaceNames    []string      //接口名
	constantPool      *ConstantPool //常量池
	fields            []*Field      //字段
	methods           []*Method     //方法
	loader            *ClassLoader  //类加载器
	superClass        *Class        //父类
	interfaces        []*Class      //接口
	instanceSlotCount uint          //实例字段数量
	staticSlotCount   uint          //静态字段数量
	staticVars        Slots         //静态变量
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethod(class, cf.Methods())
	return class
}

func (c *Class) IsPublic() bool {
	return c.accessFlags&ACC_PUBLIC != 0
}

func (c *Class) IsFinal() bool {
	return c.accessFlags&ACC_FINAL != 0
}
func (c *Class) IsSuper() bool {
	return c.accessFlags&ACC_SUPER != 0
}

func (c *Class) IsInterface() bool {
	return c.accessFlags&ACC_INTERFACE != 0
}

func (c *Class) IsAbstract() bool {
	return c.accessFlags&ACC_ABSTRACT != 0
}

func (c *Class) IsSynthetic() bool {
	return c.accessFlags&ACC_SYNTHETIC != 0
}

func (c *Class) IsAnnotation() bool {
	return c.accessFlags&ACC_ANNOTATION != 0
}

func (c *Class) IsEnum() bool {
	return c.accessFlags&ACC_ENUM != 0
}

func (c *Class) isAccessibleTo(other *Class) bool {
	return c.IsPublic() || c.getPackageName() == other.getPackageName()
}

func (c *Class) getPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}

func (c *Class) NewObject() *Object {

	return newObject(c)
}

func (c *Class) isSubInterfaceOf(t *Class) bool {
	for _, superInterface := range c.interfaces {
		if superInterface == t || superInterface.isSubInterfaceOf(t) {
			return true
		}
	}
	return false
}

func newObject(c *Class) *Object {
	return &Object{
		class:  c,
		fields: newSlots(c.instanceSlotCount),
	}
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}

func (c *Class) GetMainMethod() *Method {

	return c.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (c *Class) getStaticMethod(name string, descriptor string) *Method {

	for _, method := range c.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (c *Class) StaticVars() Slots {
	return c.staticVars
}
