package heap

import (
	"strings"

	"github.com/Jucunqi/jvmgo/ch08/classfile"
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
	initStarted       bool          // 是否已经初始化
}

func newClass(cf *classfile.ClassFile) *Class {

	// 数据封装
	class := &Class{}

	// 访问标识
	class.accessFlags = cf.AccessFlags()

	// 类名
	class.name = cf.ClassName()

	// 父类名
	class.superClassName = cf.SuperClassName()

	// 接口名称集合
	class.interfaceNames = cf.InterfaceNames()

	// 常量池
	class.constantPool = newConstantPool(class, cf.ConstantPool())

	// 根据常量池中的字段表，封装字段信息
	class.fields = newFields(class, cf.Fields())

	// 根据常量池中的方法表，封装方法信息
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
func (c *Class) GetPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}

func (c *Class) SuperClass() *Class {
	return c.superClass
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) InitStarted() bool {
	return c.initStarted
}

func (c *Class) StartInit() {
	c.initStarted = true
}

func (c *Class) GetClinitMethod() *Method {

	return c.getStaticMethod("<clinit>", "()V")
}
