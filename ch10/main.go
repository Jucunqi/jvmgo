package main

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch10/classpath"
	"github.com/Jucunqi/jvmgo/ch10/rtda/heap"
	"strings"
)

// 程序执行入口
func main() {

	// 获取cmd信息
	cmd := parseCmd()
	//cmd := &Cmd{cpOption: "/Users/lucaju/Documents/workspace/go/jvmgo/out/production/jvmgo/", class: "example.src.main.java.jvmgo.ch09.StrTest", verboseClassFlag: false, verboseInstFlag: false}
	// 根据cmd参数决定后面的执行内容
	if cmd.versionFlag {
		fmt.Println("version 0.0.6")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	classname := strings.ReplaceAll(cmd.class, ".", "/")
	mainClass := classLoader.LoadClass(classname)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag, cmd.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
