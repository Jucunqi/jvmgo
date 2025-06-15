package main

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch03/classfile"
	"github.com/Jucunqi/jvmgo/ch03/classpath"
	"strings"
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
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	classname := strings.ReplaceAll(cmd.class, ".", "/")
	fmt.Println(classname)
	class := loadClass(classname, cp)
	printClass(class)
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
