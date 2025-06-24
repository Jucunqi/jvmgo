package heap

import "github.com/Jucunqi/jvmgo/ch06/classfile"

// 常量接口
type Constant interface {
}

// 常量池
type ConstantPool struct {
	class  *Class // 所属类
	consts []Constant
}

func newConstantPool(class *Class, cfcp classfile.ConstantPool) *ConstantPool {

	cpCount := len(cfcp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class: class}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfcp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value() // int32
			break
		case *classfile.ConstantFloatInfo:
			intInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = intInfo.Value() // float
			break
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value() // long
			i++
			break
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() // double
			i++
			break
		case *classfile.ConstantStringInfo:
			strInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = strInfo.String()
			break
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
			break
		case *classfile.ConstantFieldrefInfo:
			fieldInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldInfo)
			break
		case *classfile.ConstantMethodrefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodRefInfo)
			break
		case *classfile.ConstantInterfaceMethodrefInfo:
			interfaceRef := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, interfaceRef)
			break
		}
	}
	rtCp.consts = consts
	return rtCp
}
func (c *ConstantPool) GetConstant(index uint) Constant {
	if p := c.consts[index]; p != nil {
		return p
	}
	return nil
}
