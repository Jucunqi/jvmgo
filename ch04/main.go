package main

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch04/classfile"
	"github.com/Jucunqi/jvmgo/ch04/classpath"
	"github.com/Jucunqi/jvmgo/ch04/rtda"
)

// 程序执行入口
func main() {

	// 获取cmd信息
	//cmd := parseCmd()
	cmd := &Cmd{class: "java.lang.String"}
	// 根据cmd参数决定后面的执行内容
	if cmd.versionFlag {
		fmt.Println("version 0.0.2")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	//fmt.Printf("classpath:%v class:%v args:%v\n",
	//	cp, cmd.class, cmd.args)
	//
	//classname := strings.ReplaceAll(cmd.class, ".", "/")
	//fmt.Println(classname)
	//class := loadClass(classname, cp)
	//printClass(class)
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopInt())
	println(ops.PopInt())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopFloat())
	println(ops.PopDouble())
	println(ops.PopRef())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))

}

func printClass(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: *s", len(cf.ConstantPool()))
	fmt.Printf("access flag: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %s\n", cf.ClassName())
	fmt.Printf("super class: %s\n", cf.SuperClassName())
	fmt.Printf("interfaces name: %s\n", cf.InterfaceNames())
	fmt.Printf("filed count: %s\n", len(cf.Fields()))
	for _, field := range cf.Fields() {
		fmt.Printf("%s\n", field.Name())
	}
	fmt.Printf("method count: %s\n", len(cf.Methods()))
	for _, method := range cf.Methods() {
		fmt.Printf("%s\n", method.Name())
	}
}

func loadClass(classname string, cp *classpath.Classpath) *classfile.ClassFile {
	data, _, err := cp.ReadClass(classname)
	if err != nil {
		fmt.Printf("Cound not find or load main class %s\n", classname)
		return nil
	}
	cf, err := classfile.Parse(data)
	if err != nil {
		panic(err)
	}
	return cf
}
