package main

import (
	"fmt"
)

// 程序执行入口
func main() {

	// 获取cmd信息
	//cmd := parseCmd()
	cmd := &Cmd{cpOption: "/Users/lucaju/Documents/workspace/go/jvmgo/out/production/jvmgo/", class: "example.src.main.java.jvmgo.ch11.HelloWorld", verboseClassFlag: false, verboseInstFlag: false}
	// 根据cmd参数决定后面的执行内容
	if cmd.versionFlag {
		fmt.Println("version 0.0.11")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		newJVM(cmd).start()
	}
}
