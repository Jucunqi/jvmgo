package main

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch05/classfile"
	"github.com/Jucunqi/jvmgo/ch05/classpath"
	"strings"
)

// 程序执行入口
func main() {

	// 获取cmd信息
	//cmd := parseCmd()
	cmd := &Cmd{cpOption: "/Users/lucaju/Documents/workspace/go/jvmgo/out/production/jvmgo/example/src/main/java/jvmgo/ch05", class: "GuessTest"}
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
	cf := loadClass(classname, cp)
	mainMethod := getMainMethod(cf)
	interpret(mainMethod)
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {

	methods := cf.Methods()
	for _, method := range methods {
		methodName := method.Name()
		if methodName == "main" && method.Descriptor() == "([Ljava/lang/String;)V" {
			return method
		}
	}
	return nil
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
