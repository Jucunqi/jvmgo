# 手写Java虚拟机



## 为什么想去自己实现一个JVM虚拟机
- 兴趣使然，通过自己动手，体会"Write once,run anywhere"的本质，不想做api调用程序员。
- 并且看网上的资料，没有找到mac平台实现的，也给大家做一个参考



## 为什么使用Go语言
- 相较于 C 和 C++，Go 语言的语法简洁且开发效率更高，能降低开发门槛，更专注于 JVM 核心功能实现
- 同时也是一个学习Go语言的机会



## 参考书籍
> 《自己动手写Java虚拟机》 —— 张秀宏



## 环境
- 操作系统 MacOS 15.5
- JDK 1.8
- Go 1.23.10



# 第一章 命令行工具

## 一、环境准备

### 1、JDK 1.8
下载安装，配置环境变量



### 2、go 1.23.10
下载安装，配置GOROOT、GOPATH环境变量



## 二、开发命令行工具代码



### 1、创建cmd.go
1. 定义cmd结构体，用于接收命令行参数
```go
package main

import (
	"flag"
	"fmt"
	"os"
)

// 命令行选项和参数 结构题
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

// 解析命令行参数，并赋值到Cmd结构体
func parseCmd() *Cmd {

	cmd := &Cmd{}
	flag.Usage = printUsage

	// 接收命令行中的参数，并给cmd中的属性赋值
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.versionFlag, "v", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
```



### 2、创建main.go作为程序主入口

1. 与cmd.go文件一样，main.go的包名也是main。
2. 在go语言中main是一个特殊的包，这个包所在的目录会被编译成一个可执行文件。
3. go的程序入口也是main()函数，但是不接受任何参数，也不能有返回值

```go
package main

import "fmt"

// 程序执行入口
func main() {

	// 获取cmd信息
	cmd := parseCmd()

	// 根据cmd参数决定后面的执行内容
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)

}

```


### 3、编译本章代码
1. 执行一下命令
    ```bash
    go install ./ch01
    ```
2. 在GOPATH/bin目录下就会生成ch01可执行文件
3. 执行命令，测试效果如图
   ![version.png](https://s2.loli.net/2025/06/13/8N29SWTqMicge6Q.png)





# 第二章 搜索class文件



### 一、通过-Xjre命令获取jre路径，用于解析java类

- 修改cmd.go接收xjre参数

```go
// 命令行选项和参数 结构题
type Cmd struct {

	...省略
	// -Xjre
	xJreOption string
	...省略
}

// 解析命令行参数，并赋值到Cmd结构体
func parseCmd() *Cmd {

	cmd := &Cmd{}
	flag.Usage = printUsage

	// 接收命令行中的参数，并给cmd中的属性赋值
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.versionFlag, "v", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
  // 添加解析Xjre参数
	flag.StringVar(&cmd.xJreOption, "Xjre", "", "path to jre")
	flag.Parse()
	....省略
	return cmd
}
```



### 二、定义Entry接口

```go
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {

	// 读取类
	readClass(name string) ([]byte, Entry, error)
	// String 方法 ，类似java中toString 方法
	String() string
}

// 获取Entry对象
func newEntry(path string) Entry {

	// 复合情况
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	// 通配符情况
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	// .jar .zip情况
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	// 目录情况
	return newDirEntry(path)
}
```



- Entry接口中定义了两个方法
  - 读取类
  - String 方法
- 有4个实现类 - go语言中的实现类与java不同不需要implement关键字，只需要实现方法就可以
  - CompositeEntry
  - WildcardEntry 本质也是 CompositeEntry，没有创建新的结构体
  - ZipEntry
  - DirEntry
  - 具体结构看代码中实现



### 三、定义Classpath类

部分代码如下

```go
package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// 解析启动类和扩展类
func (c *Classpath) parseBootAndExtClasspath(jreOption string) {

	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

// 解析用户类classpath
func (c *Classpath) parseUserClasspath(cpOption string) {

	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}
```



### 四、修改main.go用于测试

> 主要修改startJVM这个函数中的内容，实现输出class二进制内容

```go
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
```



### 五、项目install并执行命令测试

1. `go install ./ch02/`
2. `ch02 java.lang.Object`

3. 效果如图

![ch02test.png](https://s2.loli.net/2025/06/13/Wut54loJMEVZ61a.png)