package heap

import (
	"fmt"

	"github.com/Jucunqi/jvmgo/ch07/classfile"
	"github.com/Jucunqi/jvmgo/ch07/classpath"
)

// 类加载器
type ClassLoader struct {
	cp          *classpath.Classpath // 类路径
	verboseFlag bool                 // 输出日志标识
	classMap    map[string]*Class    // 已加载的类
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	classLoader := &ClassLoader{}
	classLoader.cp = cp
	classLoader.verboseFlag = verboseFlag
	classLoader.classMap = make(map[string]*Class)
	return classLoader
}

func (c *ClassLoader) LoadClass(name string) *Class {

	if class, ok := c.classMap[name]; ok {
		return class

	}
	return c.loadNonArrayClass(name)
}

func (c *ClassLoader) loadNonArrayClass(name string) *Class {

	data, entry := c.readClass(name)
	class := c.defineClass(data)
	link(class)
	if c.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n,", name, entry)
	}
	return class
}

func (c *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {

	class, entry, err := c.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return class, entry
}

func (c *ClassLoader) defineClass(data []byte) *Class {

	// 解析Class，生成class对象
	class := parseClass(data)
	class.loader = c
	resolveSuperClass(class)
	resolveInterfaces(class)
	c.classMap[class.name] = class
	return class
}

func resolveInterfaces(class *Class) {

	interfaceNames := class.interfaceNames
	interfaceCount := len(interfaceNames)
	interfaceClasses := make([]*Class, interfaceCount)
	for i, name := range interfaceNames {
		interfaceClasses[i] = class.loader.LoadClass(name)
	}
	class.interfaces = interfaceClasses
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {

	// 读取解析class字节数组内容，封装到classfile中
	cf, err := classfile.Parse(data)

	if err != nil {
		panic("java.lang.ClassFormataError")
	}

	// 创建class对象
	class := newClass(cf)
	return class
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo
}
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}
