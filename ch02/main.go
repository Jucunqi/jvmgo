package main

import (
	"fmt"
	"github.com/Jucunqi/jvmgo/ch02/classpath"
	"strings"
)

// 程序执行入口
func main() {

	// 获取cmd信息
	cmd := parseCmd()
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
	data, _, err := cp.ReadClass(classname)
	if err != nil {
		fmt.Printf("Cound not find or load main class %s\n", cmd.class)
		return
	}
	// 输出类的内容
	fmt.Printf("class data:%v\n", data)
}
