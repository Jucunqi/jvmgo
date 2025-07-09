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



# 笔记部分

## 目录

[第一章 命令行工具](# 第一章 命令行工具)
	[一、环境准备](# 一、环境准备)
		[1、JDK 1.8](# 1、JDK 1.8)
		[2、go 1.23.10](# 2、go 1.23.10)
	[二、开发命令行工具代码](# 二、开发命令行工具代码)	
		[1、创建cmd.go](# 1、创建cmd.go)
		[2、创建main.go作为程序主入口](# 2、创建main.go作为程序主入口)
		[3、编译本章代码](# 3、编译本章代码)
​	

[第二章 搜索class文件](# 第二章 搜索class文件)
	[一、通过-Xjre命令获取jre路径，用于解析java类](# 一、通过-Xjre命令获取jre路径，用于解析java类)
	[二、定义Entry接口](# 二、定义Entry接口)
	[三、定义Classpath类](# 三、定义Classpath类)
	[四、修改main.go用于测试](# 四、修改main.go用于测试)
	[五、项目install并执行命令测试](# 五、项目install并执行命令测试)

[第三章 解析class文件](# 第三章 解析class文件)
	[一、解析class一般信息](# 一、解析class一般信息)
		[1、封装ClassReader类，实现对class文件字节数组的读取方法](# 1、封装ClassReader类，实现对class文件字节数组的读取方法)
		[2、定义ClassFile类型，存储class数据](# 2、定义ClassFile类型，存储class数据)
		[2、定义ClassFile类型，存储class数据](# 2、定义ClassFile类型，存储class数据)
	[二、解析常量池](# 二、解析常量池)
		[1、常量池结构](# 1、常量池结构)
		[2、定义常量池结构体](# 2、定义常量池结构体)
	[三、解析字段、方法](# 三、解析字段、方法)
		[1、字段结构](# 1、字段结构)
		[2、方法结构](# 2、方法结构)
		[3、定义结构体，封装字段、方法数据](# 3、定义结构体，封装字段、方法数据)
	[四、解析属性](# 四、解析属性)
		[1、属性结构](# 1、属性结构)
			[1）通用属性定义](# 1 通用属性定义)
			[2) Code属性](# 2 Code属性)
			[3) ConstantValue 属性](# 3 ConstantValue 属性)
			[4) Exceptions 属性](# 4 Exceptions 属性)
			[5) LineNumberTable 属性](# 5 LineNumberTable 属性)
			[6) LocalVariableTable 属性](# 6 LocalVariableTable 属性)
			[7) SourceFile 属性](# 7 SourceFile 属性)
		[2、定义结构体封装属性信息](# 2、定义结构体封装属性信息)
	[五、测试本章代码](# 五、测试本章代码)
		[1、修改main.go中startJvm函数](# 1、修改main.go中startJvm函数)
		[2、项目install并执行命令测试](# 2、项目install并执行命令测试)
[第四章 运行时数据区](# 第四章 运行时数据区)
	[一、示意图](# 一、示意图)
	[二、数据类型](# 二、数据类型)
	[三、实现运行时数据区](# 三、实现运行时数据区)
		[1、线程Thread](# 1、线程Thread)
		[2、栈](# 2、栈)
		[3、栈帧](# 3、栈帧)
		[4、局部变量表](# 4、局部变量表)
		[5、操作数栈](# 5、操作数栈)
	[四、测试本章代码](# 四、测试本章代码)
		[1、修改main.go中startJvm函数](# 1、修改main.go中startJvm函数)
		[2、项目install并执行命令测试](# 2、项目install并执行命令测试)
[第五章 指令集和解释器](# 第五章 指令集和解释器)
	[一、常量池和tag对应关系](# 一、常量池和tag对应关系)
	[二、AccessFlags结构定义](# 二、AccessFlags结构定义)
	[三、字段表定义](# 三、字段表定义)
	[四、方法表定义](# 四、方法表定义)
	[五、属性表定义](# 五、属性表定义)
	[六、解释指令](# 六、解释指令)
		[1、nop指令](# 1、nop指令)
		[2、const指令](# 2、const指令)
		[3、bipush和sipush命令](# 3、bipush和sipush命令)
		[4、加载指令](# 4、加载指令)
		[5、存储指令](# 5、存储指令)
		[6、栈指令](# 6、栈指令)
		[7、pop和pop2指令](# 7、pop和pop2指令)
		[8、dup指令](# 8、dup指令)
		[9、swap指令](# 9、swap指令)
		[10、数字指令](# 10、数字指令)
			[1、算数指令](# 1、算数指令)
			[2、位移指令](# 2、位移指令)
			[3、布尔运算指令](# 3、布尔运算指令)
			[4、iinc指令](# 4、iinc指令)
		[11、类型转换指令](# 11、类型转换指令)		
		[12、比较指令](# 12、比较指令)		
			[1、lcmp](# 1、lcmp)		
			[2、fcmp<op>](# 2、fcmp<op>)		
			[3、if<cond>指令](# 3、if<cond>指令)		
		[13、控制指令](# 13、控制指令)		
			[1、goto指令](# 1、goto指令)		
			[2、tableswitch指令](# 2、tableswitch指令)		
			[3、lookupswitch指令](# 3、lookupswitch指令)		
		[14、wide指令](# 14、wide指令)		
	[七、解释器](# 七、解释器)		
		[1、创建工厂类，根据不同tag获取不同的指令对象](# 1、创建工厂类，根据不同tag获取不同的指令对象)		
		[2、interpret方法](# 2、interpret方法)		
	[八、测试](# 八、测试)		
		[1、编写java代码](# 1、编写java代码)		
		[2、使用go测试](# 2、使用go测试)		
[第六章 类和对象](# 第六章 类和对象)		
	[一、方法区](# 一、方法区)		
		[1、类信息](# 1、类信息)		
		[2、字段信息](# 2、字段信息)		
		[3、方法信息](# 3、方法信息)		
		[4、其他信息](# 4、其他信息)		
	[二、运行时常量池](# 二、运行时常量池)		
		[1、类符号引用](# 1、类符号引用)		
		[2、字段符号引用](# 2、字段符号引用)		
		[3、方法符号引用](# 3、方法符号引用)		
		[4、接口方法符号引用](# 4、接口方法符号引用)		
	[三、类加载器](# 三、类加载器)		
	[四、对象、实例变量和类变量](# 四、对象、实例变量和类变量)		
	[五、类和字段符号引用解析](# 五、类和字段符号引用解析)		
	[六、类和字段相关指令](# 六、类和字段相关指令)		
		[1、NEW](# 1、NEW)		
		[2、putstatic](# 2、putstatic)		
	[七、功能测试](# 七、功能测试)		
		[1、install](# 1、install)		
		[2、执行](# 2、执行)		
		[3、结果如图](# 3、结果如图)		
[第七章 方法调用和返回](# 第七章 方法调用和返回)		
	[一、概述](# 一、概述)		
	[二、解析方法符号引用](# 二、解析方法符号引用)		
	[三、方法调用和参数传递](# 三、方法调用和参数传递)		
	[四、返回指令](# 四、返回指令)		
	[五、方法调用指令](# 五、方法调用指令)		
	[六、改进解释器](# 六、改进解释器)		
	[七、类初始化](# 七、类初始化)		
[第八章 数组和字符串](# 第八章 数组和字符串)		
	[一、概述](# 一、概述)		
	[二、数组实现](# 二、数组实现)		
		[1、数组对象](# 1、数组对象)		
		[2、数组类](# 2、数组类)		
		[3、加载数组类](# 3、加载数组类)		
	[三、数组相关指令](# 三、数组相关指令)	
		[1、newarray指令](# 1、newarray指令)		
		[2、anewarray指令](# 2、anewarray指令)		
		[3、arraylength指令](# 3、arraylength指令)		
		[4、<t>aload指令](# 4、<t>aload指令)		
		[5、 <t>astore指令](# 5、 <t>astore指令)		
		[6、multianewarray指令](# 6、multianewarray指令)		
	[四、测试数组](# 四、测试数组)	
	[五、字符串](# 五、字符串)	
		[1、字符串池](# 1、字符串池)		
		[2、测试字符串](# 2、测试字符串)		
[第九章 本地方法调用](# 第九章 本地方法调用)		
	[一、注册和使用本地方法](# 一、注册和使用本地方法)		
	[二、调用本地方法](# 二、调用本地方法)		
	[三、反射](# 三、反射)		
		[1、类和对象之间的关系](# 1、类和对象之间的关系)		
		[2、修改类加载加载逻辑，加载类的时候，绑定类对象信息](# 2、修改类加载加载逻辑，加载类的时候，绑定类对象信息)		
		[3、基本类型的类](# 3、基本类型的类)		
		[4、修改ldc指令](# 4、修改ldc指令)		
		[5、通过反射获取类名](# 5、通过反射获取类名)		
		[6、测试本节代码](# 6、测试本节代码)		
	[四、字符串拼接和String.intern()方法](# 四、字符串拼接和String.intern()方法)		
		[1、字符串拼接涉及到的Java类库](# 1、字符串拼接涉及到的Java类库)		
		[2、String.intern()方法设计到的类库](# 2、String.intern()方法设计到的类库)		
		[3、测试本节代码](# 3、测试本节代码)		
[第十章 异常处理](# 第十章 异常处理)		
	[一、概述](# 一、概述)		
	[二、异常抛出](# 二、异常抛出)		
	[三、异常处理表](# 三、异常处理表)		
	[四、实现athrow指令](# 四、实现athrow指令)		
	[五、Java虚拟机栈信息](# 五、Java虚拟机栈信息)		
	[六、测试代码](# 六、测试代码)


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



## 一、通过-Xjre命令获取jre路径，用于解析java类

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



## 二、定义Entry接口

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



## 三、定义Classpath类

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



## 四、修改main.go用于测试

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



## 五、项目install并执行命令测试

1. `go install ./ch02/`
2. `ch02 java.lang.Object`

3. 效果如图

![ch02test.png](https://s2.loli.net/2025/06/13/Wut54loJMEVZ61a.png)



# 第三章 解析class文件

## 一、解析class一般信息

class文件的结构定义

```
ClassFile {
    u4             magic;                  // 魔数，固定值 0xCAFEBABE
    u2             minor_version;          // 次版本号
    u2             major_version;          // 主版本号
    u2             constant_pool_count;    // 常量池大小
    cp_info        constant_pool[constant_pool_count-1]; // 常量池
    u2             access_flags;           // 类访问标志
    u2             this_class;             // 类本身的常量池索引
    u2             super_class;            // 父类的常量池索引
    u2             interfaces_count;       // 实现的接口数量
    u2             interfaces[interfaces_count]; // 接口列表
    u2             fields_count;           // 字段数量
    field_info     fields[fields_count];   // 字段表
    u2             methods_count;          // 方法数量
    method_info    methods[methods_count]; // 方法表
    u2             attributes_count;       // 属性数量
    attribute_info attributes[attributes_count]; // 属性表
}
```



### 1、封装ClassReader类，实现对class文件字节数组的读取方法

```go 
package classfile

import "encoding/binary"

// ClassReader 封装读取class字节流的方法
type ClassReader struct {
	data []byte
}

// 读取u1
func (c *ClassReader) readUnit8() uint8 {
	val := c.data[0]
	c.data = c.data[1:]
	return val
}

// 读取u2
func (c *ClassReader) readUnit16() uint16 {
	val := binary.BigEndian.Uint16(c.data)
	c.data = c.data[2:]
	return val
}

// 读取u4
func (c *ClassReader) readUnit32() uint32 {
	val := binary.BigEndian.Uint32(c.data)
	c.data = c.data[4:]
	return val
}

// 读取u8
func (c *ClassReader) readUnit64() uint64 {
	val := binary.BigEndian.Uint64(c.data)
	c.data = c.data[8:]
	return val
}

// 读取u2集合,集合的大小由开头的uint16数据指出
func (c *ClassReader) readUnit16s() []uint16 {
	length := c.readUnit16()
	uint16s := make([]uint16, length)
	for i := range uint16s {
		uint16s[i] = c.readUnit16()
	}
	return uint16s
}

// 读取指定长度的字节数组
func (c *ClassReader) readBytes(length uint32) []byte {
	bytes := c.data[:length]
	c.data = c.data[length:]
	return bytes
}
```



### 2、定义ClassFile类型，存储class数据

```go
package classfile

import "fmt"

type ClassFile struct {
	magic        uint32          // 魔数
	minorVersion uint16          // 小版本
	majorVersion uint16          // 大版本
	constantPool ConstantPool    // 常量池
	accessFlags  uint16          // 访问标识符
	thisClass    uint16          // 当前类
	superClass   uint16          // 父类
	interfaces   []uint16        //接口集合
	fields       []*MemberInfo   // 字段集合
	methods      []*MemberInfo   // 方法集合
	attributes   []AttributeInfo // 属性集合
}

// Parse 解析方法
func Parse(classData []byte) (cf *ClassFile, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
  // 关键方法
	cf.read(cr)
	return
}

// 读取
func (c *ClassFile) read(reader *ClassReader) {
	c.readAndCheckMagic(reader)
	c.readAndCheckVersion(reader)
  // 常量池
	c.constantPool = readConstantPool(reader)
	c.accessFlags = reader.readUnit16()
	c.thisClass = reader.readUnit16()
	c.superClass = reader.readUnit16()
	c.interfaces = reader.readUnit16s()
	c.fields = readMembers(reader, c.constantPool)
	c.methods = readMembers(reader, c.constantPool)
  // 属性
	c.attributes = readAttributes(reader, c.constantPool)
}

// 读取并检查魔数
func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {

	magic := reader.readUnit32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 读取并检查版本号
func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {
	c.minorVersion = reader.readUnit16()
	c.majorVersion = reader.readUnit16()
	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if c.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// MinorVersion 返回小版本号
func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}

// MajorVersion 返回大版本号
func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}

// ConstantPool 返回常量池对象
func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}

// AccessFlags 返回访问标识符
func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}

// Fields 返回属性
func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}

// Methods 返回方法
func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}

// ClassName 返回类名
func (c *ClassFile) ClassName() string {

	return c.constantPool.getClassName(c.thisClass)
}

// SuperClassName 返回父类名
func (c *ClassFile) SuperClassName() string {
	if c.superClass > 0 {
		return c.constantPool.getClassName(c.superClass)
	}
	return "" // 只有java.lang.Object 没有父类
}

// InterfaceNames 返回接口名
func (c *ClassFile) InterfaceNames() []string {

	interfaceNames := make([]string, len(c.interfaces))
	for i, cpIndex := range c.interfaces {
		interfaceNames[i] = c.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
```



## 二、解析常量池

### 1、常量池结构

```txt
cp_info {
    u1 tag;       // 常量类型标识（1-21 之间的值）
    u1 info[];    // 具体内容，格式取决于 tag
}
```

### 2、定义常量池结构体

存放常量池结构信息接口数组

根据tag决定具体常量池的接口实现

```go
package classfile

type ConstantPool []ConstantInfo

// 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {

	cpCount := int(reader.readUnit16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { // 索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 如果是Long 或者 Double 占两个位置
		}
	}
	return cp
}

// 获取指定常量信息
func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {

	if cpInfo := c[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 读取名称和类型 (字段、方法)
func (c ConstantPool) getNameAndType(index uint16) (string, string) {

	ntInfo := c.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := c.getUtf8(ntInfo.nameIndex)
	_type := c.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 读取类名
func (c ConstantPool) getClassName(index uint16) string {
	classInfo := c.getConstantInfo(index).(*ConstantClassInfo)
	return c.getUtf8(classInfo.nameIndex)
}

// 读取utf8字符
func (c ConstantPool) getUtf8(index uint16) string {
	utf8Info := c.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
```



## 三、解析字段、方法

### 1、字段结构

```
field_info {
    u2             access_flags;       // 字段访问标志
    u2             name_index;         // 字段名的常量池索引
    u2             descriptor_index;   // 字段类型描述符的常量池索引
    u2             attributes_count;   // 属性数量
    attribute_info attributes[attributes_count]; // 属性表
}
```

### 2、方法结构

```
method_info {
    u2             access_flags;       // 通常包含 ACC_PUBLIC 和 ACC_ABSTRACT
    u2             name_index;         // 方法名的常量池索引
    u2             descriptor_index;   // 方法描述符的常量池索引
    u2             attributes_count;   // 属性数量（抽象方法通常为 0）
    attribute_info attributes[attributes_count]; // 属性表
}
```

### 3、定义结构体，封装字段、方法数据

```go
package classfile

type MemberInfo struct {
	cp              ConstantPool    // 常量池
	accessFlags     uint16          // 访问标识符
	nameIndex       uint16          // 名称索引
	descriptorIndex uint16          // 描述符索引
	attributes      []AttributeInfo // 属性集合
}

// 读取Member集合
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUnit16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取一个Member
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {

	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUnit16(),
		nameIndex:       reader.readUnit16(),
		descriptorIndex: reader.readUnit16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}
func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}
func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}
```



## 四、解析属性

### 1、属性结构

#### 1 通用属性定义

```
attributes_count {
    u2 attributes_count;    // 属性数量
    attribute_info attributes[attributes_count]; // 属性数组
}
attribute_info {
    u2 attribute_name_index;  // 指向常量池中的 UTF-8 字符串，表示属性名
    u4 attribute_length;      // 属性长度（不包含前 6 字节）
    u1 info[attribute_length]; // 属性内容，格式取决于属性名
}
```

#### 2 Code属性

```
Code_attribute {
    u2 attribute_name_index;    // 固定为 "Code"
    u4 attribute_length;
    u2 max_stack;               // 操作数栈最大深度
    u2 max_locals;              // 局部变量表大小
    u4 code_length;
    u1 code[code_length];       // 字节码指令
    u2 exception_table_length;
    {
        u2 start_pc;            // 异常处理开始位置
        u2 end_pc;              // 异常处理结束位置
        u2 handler_pc;          // 异常处理程序位置
        u2 catch_type;          // 捕获的异常类型
    } exception_table[exception_table_length];
    u2 attributes_count;        // Code 属性的子属性
    attribute_info attributes[attributes_count];
}
```

#### 3 ConstantValue 属性

```
ConstantValue_attribute {
    u2 attribute_name_index;    // 固定为 "ConstantValue"
    u4 attribute_length;        // 固定为 2
    u2 constantvalue_index;     // 指向常量池中的常量值
}
```

#### 4 Exceptions 属性

```
Exceptions_attribute {
    u2 attribute_name_index;    // 固定为 "Exceptions"
    u4 attribute_length;
    u2 number_of_exceptions;    // 异常数量
    u2 exception_index_table[number_of_exceptions]; // 指向常量池中的异常类
}
```

#### 5 LineNumberTable 属性

```
LineNumberTable_attribute {
    u2 attribute_name_index;    // 固定为 "LineNumberTable"
    u4 attribute_length;
    u2 line_number_table_length;
    {
        u2 start_pc;            // 字节码起始位置
        u2 line_number;         // 对应的源代码行号
    } line_number_table[line_number_table_length];
}
```

#### 6 LocalVariableTable 属性

```
LocalVariableTable_attribute {
    u2 attribute_name_index;    // 固定为 "LocalVariableTable"
    u4 attribute_length;
    u2 local_variable_table_length;
    {
        u2 start_pc;            // 局部变量作用域起始位置
        u2 length;              // 作用域长度
        u2 name_index;          // 变量名在常量池中的索引
        u2 descriptor_index;    // 类型描述符在常量池中的索引
        u2 index;               // 局部变量表中的 Slot 索引
    } local_variable_table[local_variable_table_length];
}
```

#### 7 SourceFile 属性

```
SourceFile_attribute {
    u2 attribute_name_index;    // 固定为 "SourceFile"
    u4 attribute_length;        // 固定为 2
    u2 sourcefile_index;        // 指向常量池中的源文件名
}
```

### 2、定义结构体封装属性信息

```
package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {

	attributeCount := reader.readUnit16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {

	attrNameIndex := reader.readUnit16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUnit32()
	attribute := newAttribute(attrName, attrLen, cp)
	attribute.readInfo(reader)
	return attribute
}

func newAttribute(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {

	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	//case "Depreated":
	//	return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
```

## 五、测试本章代码

### 1、修改main.go中startJvm函数

```go
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
```



### 2、项目install并执行命令测试

1. `go install ./ch03/`
2. `ch02 java.lang.String`

3. 效果如图

![ch03test.png](https://s2.loli.net/2025/06/15/JBvAEgSq6ZNmknT.png)



# 第四章 运行时数据区

### 一、示意图

![ch04rtdata.png](https://s2.loli.net/2025/06/16/xGZ468yRrdO3UCY.png)

### 二、数据类型

1. 基本数据类型
   1. 存放数据本身
2. 引用数据类型
   1. 存放引用指针

![ch04datatype.png](https://s2.loli.net/2025/06/16/l7cwapX8zktA9fU.png)



### 三、实现运行时数据区

#### 1、线程Thread

```go
package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{stack: newStack(1024)}
}

func (t *Thread) PC() int {
	return t.pc
}
func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}
func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}

```



#### 2、栈

```go
package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func (s *Stack) push(frame *Frame) {

	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {

	if s._top == nil {
		panic("jvm stack is empty")
	}
	top := s._top
	s._top = top.lower
	s.size--
	return top

}

func (s *Stack) top() *Frame {

	if s.top() == nil {
		panic("jvm stack is empty!")
	}
	return s._top
}

func newStack(size uint) *Stack {
	return &Stack{maxSize: size}
}
```

#### 3、栈帧

```go
package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack)}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}
```

#### 4、局部变量表

```go
package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxSize uint) LocalVars {

	if maxSize > 0 {
		return make([]Slot, maxSize)
	}
	return nil
}

func (l LocalVars) SetInt(index uint, val int32) {

	l[index].num = val
}

func (l LocalVars) GetInt(index uint) int32 {

	return l[index].num
}

func (l LocalVars) SetFloat(index uint, val float32) {

	bits := math.Float32bits(val)
	l[index].num = int32(bits)
}

func (l LocalVars) GetFloat(index uint) float32 {

	bits := l[index].num
	return math.Float32frombits(uint32(bits))
}

func (l LocalVars) SetLong(index uint, val int64) {

	l[index].num = int32(val)
	l[index+1].num = int32(val >> 32)
}

func (l LocalVars) GetLong(index uint) int64 {

	low := uint32(l[index].num)
	high := uint32(l[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (l LocalVars) SetDouble(index uint, val float64) {

	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}

func (l LocalVars) GetDouble(index uint) float64 {

	long := l.GetLong(index)
	bits := uint64(long)
	return math.Float64frombits(bits)
}
func (l LocalVars) SetRef(index uint, val *Object) {

	l[index].ref = val
}

func (l LocalVars) GetRef(index uint) *Object {

	return l[index].ref
}
```

#### 5、操作数栈

```go
package rtda

import "math"

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(stackSize uint) *OperandStack {
	if stackSize > 0 {
		return &OperandStack{slots: make([]Slot, stackSize)}
	}
	return nil
}

func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

func (o *OperandStack) PopInt() int32 {

	o.size--
	return o.slots[o.size].num
}

func (o *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	o.slots[o.size].num = int32(bits)
	o.size++
}

func (o *OperandStack) PopFloat() float32 {

	o.size--
	bits := uint32(o.slots[o.size].num)
	return math.Float32frombits(bits)
}
func (o *OperandStack) PushLong(val int64) {
	low := int32(val)
	o.slots[o.size].num = low
	o.size++
	high := int32(val >> 32)
	o.slots[o.size].num = high
	o.size++
}

func (o *OperandStack) PopLong() int64 {

	o.size -= 2
	low := uint32(o.slots[o.size].num)
	high := uint32(o.slots[o.size+1].num)
	return int64(high)<<32 | int64(low)
}
func (o *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	o.PushLong(int64(bits))
}

func (o *OperandStack) PopDouble() float64 {

	bits := uint64(o.PopLong())
	return math.Float64frombits(bits)
}

func (o *OperandStack) PushRef(val *Object) {

	o.slots[o.size].ref = val
	o.size++
}

func (o *OperandStack) PopRef() *Object {

	o.size--
	ref := o.slots[o.size].ref
	o.slots[o.size].ref = nil
	return ref
}
```

## 四、测试本章代码

### 1、修改main.go中startJvm函数

```go
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
```

### 2、项目install并执行命令测试

1. `go install ./ch04/`
2. `ch04`

3. 效果如图

![ch04test.png](https://s2.loli.net/2025/06/16/XaMuLFisnpkEbZf.png)



# 第五章 指令集和解释器



## 一、常量池和tag对应关系

| Tag值（十进制） | Tag 值（十六进制） | 助记符                      | 说明                                  |
| --------------- | ------------------ | --------------------------- | ------------------------------------- |
| 1               | 0x01               | CONSTANT_Utf8               | UTF-8 编码的字符串常量                |
| 3               | 0x03               | CONSTANT_Integer            | 整型常量                              |
| 4               | 0x04               | CONSTANT_Float              | 浮点型常量                            |
| 5               | 0x05               | CONSTANT_Long               | 长整型常量（占两个条目）              |
| 6               | 0x06               | CONSTANT_Double             | 双精度浮点型常量（占两个条目）        |
| 7               | 0x07               | CONSTANT_Class              | 类或接口的符号引用                    |
| 8               | 0x08               | CONSTANT_String             | 字符串类型常量的引用                  |
| 9               | 0x09               | CONSTANT_Fieldref           | 字段（类变量或实例变量）的符号引用    |
| 10              | 0x0a               | CONSTANT_Methodref          | 类方法的符号引用                      |
| 11              | 0x0b               | CONSTANT_InterfaceMethodref | 接口方法的符号引用                    |
| 12              | 0x0c               | CONSTANT_NameAndType        | 字段或方法的名称和描述符的符号引用    |
| 15              | 0x0f               | CONSTANT_MethodHandle       | 方法句柄（Java 7+）                   |
| 16              | 0x10               | CONSTANT_MethodType         | 方法类型（Java 7+）                   |
| 18              | 0x12               | CONSTANT_InvokeDynamic      | 动态调用点（Java 7+，用于 Lambda 等） |
| 19              | 0x13               | CONSTANT_Module             | 模块（Java 9+）                       |
| 20              | 0x14               | CONSTANT_Package            | 包（Java 9+）                         |



### 结构定义

- `CONSTANT_Methodref_info` 的结构由三部分组成：

```txt
CONSTANT_Methodref_info {
    u1 tag;                // 标记位，固定值 0x0A (十进制10)
    u2 class_index;        // 指向 CONSTANT_Class_info 的索引，表示方法所属的类
    u2 name_and_type_index; // 指向 CONSTANT_NameAndType_info 的索引，表示方法的名称和描述符
}
```

- `CONSTANT_Class_info` 的结构由两部分组成

```txt
CONSTANT_Class_info {
    u1 tag;             // 标记位，固定值 0x07 (十进制7)
    u2 name_index;      // 指向 CONSTANT_Utf8_info 的索引，表示类或接口的全限定名
}
```

- `CONSTANT_Utf8_info` 的结构由三部分组成

```txt
CONSTANT_Utf8_info {
    u1 tag;             // 标记位，固定值 0x01 (十进制1)
    u2 length;          // UTF-8 编码的字节长度
    u1 bytes[length];   // UTF-8 编码的字节数据
}
```

- `CONSTANT_NameAndType` 的结构如下

```txt
CONSTANT_NameAndType_info {
    u1 tag;          // 标记值，固定为 0x0C
    u2 name_index;   // 指向常量池中的 UTF-8 字符串（方法/字段名）
    u2 descriptor_index;  // 指向常量池中的 UTF-8 字符串（描述符）
}
```



## 二、AccessFlags结构定义

![](https://img2018.cnblogs.com/i-beta/1911569/202001/1911569-20200105193407734-486225375.png)



## 三、字段表定义

```txt
field_info {
    u2             access_flags;       // 字段访问标志
    u2             name_index;         // 字段名索引
    u2             descriptor_index;   // 字段描述符索引
    u2             attributes_count;   // 属性数量
    attribute_info attributes[attributes_count]; // 属性表
}
```



## 四、方法表定义

```txt
method_info {
    u2             access_flags;       // 方法访问标志
    u2             name_index;         // 方法名索引
    u2             descriptor_index;   // 方法描述符索引
    u2             attributes_count;   // 属性数量
    attribute_info attributes[attributes_count]; // 属性表
}
```



## 五、属性表定义

```txt
attribute_info {
    u2 attribute_name_index; // 属性名索引
    u4 attribute_length;     // 属性长度
    u1 info[attribute_length]; // 属性内容（不同属性结构不同）
}

// 常见属性示例：Code 属性（方法字节码）
Code_attribute {
    u2 attribute_name_index; // 固定为 "Code"
    u4 attribute_length;
    u2 max_stack;            // 操作数栈最大深度
    u2 max_locals;           // 局部变量表大小
    u4 code_length;          // 字节码长度
    u1 code[code_length];    // 字节码
    u2 exception_table_length; // 异常表长度
    exception_table_entry exception_table[exception_table_length]; // 异常表
    u2 attributes_count;     // 子属性数量
    attribute_info attributes[attributes_count]; // 子属性表
}
```



## 六、解释指令



### 1、nop指令

```go
func (n *Nop) Execute(frame *rtda.Frame) {
	// noting to do
}
```



### 2、const指令

```go
type ACONST_NULL struct {
	base.NoOperandsInstruction
}
type DCONST_0 struct {
	base.NoOperandsInstruction
}
type DCONST_1 struct {
	base.NoOperandsInstruction
}
....
```



### 3、bipush和sipush命令

```go
// 常量放入操作数栈
type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

```



### 4、加载指令

```go
type ILOAD struct {
	base.Index8Instruction
}
type ILOAD_0 struct {
	base.NoOperandsInstruction
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}
type ILOAD_2 struct {
	base.NoOperandsInstruction
}
type ILOAD_3 struct {
	base.NoOperandsInstruction
}
```



### 5、存储指令

```go
type ISTORE struct {
	base.Index8Instruction
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}
type ISTORE_1 struct {
	base.NoOperandsInstruction
}
type ISTORE_2 struct {
	base.NoOperandsInstruction
}
type ISTORE_3 struct {
	base.NoOperandsInstruction
}
```



### 6、栈指令

```go
func (o *OperandStack) PushSlot(slot Slot) {
	o.slots[o.size] = slot
	o.size++
}

func (o *OperandStack) PopSlot() Slot {
	o.size--
	return o.slots[o.size]
}
```



### 7、pop和pop2指令

```go
type POP struct {
	base.NoOperandsInstruction
}

type POP2 struct {
	base.NoOperandsInstruction
}
```



### 8、dup指令

```go
// DUP
/*
bottom -> top
[...][c][b][a]

	\_
	  |
	  V

[...][c][b][a][a]
*/
type DUP struct {
	base.NoOperandsInstruction
}
```



### 9、swap指令

```go
// SWAP the top two operand stack values
type SWAP struct {
	base.NoOperandsInstruction
}

// Execute
/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/
func (S *SWAP) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

```



### 10、数字指令

#### 1、算数指令

```go
// Add double
type DADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }
type IADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }
```



#### 2、位移指令

```go
// ISHL int左移
type ISHL struct {
	base.NoOperandsInstruction
}

// ISHR int算数右移
type ISHR struct {
	base.NoOperandsInstruction
}
```



#### 3、布尔运算指令

```go
type IAND struct {
	base.NoOperandsInstruction
}

type LAND struct {
	base.NoOperandsInstruction
}
```



#### 4、iinc指令

```go
// IINC 从局部变量表中的int变量增加常量值，局部变量表索引和常量值都由指令的操作数提供
type IINC struct {
	Index uint
	Const int32
}
```



### 11、类型转换指令

```go
type D2F struct {
	base.NoOperandsInstruction
}
type D2I struct {
	base.NoOperandsInstruction
}
type D2L struct {
	base.NoOperandsInstruction
}
```



### 12、比较指令

#### 1、lcmp

```go
type LCMP struct {
	base.NoOperandsInstruction
}

func (L *LCMP) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := 0
	if v1 > v2 {
		result = 1
	} else if v1 < v2 {
		result = -1
	}

	stack.PushInt(int32(result))
}
```



#### 2、fcmp<op>

```go
// Compare float
type FCMPG struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
```



#### 3、if<cond>指令

```go
type IF_ACMPEQ struct {
	base.BranchInstruction
}
type IF_ACMPNE struct {
	base.BranchInstruction
}

func (I *IF_ACMPEQ) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v1 == v2 {
		base.Branch(frame, I.Offset)
	}
}

func (I *IF_ACMPNE) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v1 != v2 {
		base.Branch(frame, I.Offset)
	}
}
```



### 13、控制指令

#### 1、goto指令

```go
type GOTO struct {
	base.BranchInstruction
}

func (G *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, G.Offset)
}
```



#### 2、tableswitch指令

```go
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (T *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {

	reader.SkipPadding()
	T.defaultOffset = reader.ReadInt32()
	T.low = reader.ReadInt32()
	T.high = reader.ReadInt32()
	count := T.high - T.low + 1
	T.jumpOffsets = reader.ReadInt32s(count)
}

func (T *TABLE_SWITCH) Execute(frame *rtda.Frame) {

	stack := frame.OperandStack()
	i := stack.PopInt()
	if i >= T.low && i <= T.high {
		base.Branch(frame, int(T.jumpOffsets[i-T.low]))
	} else {
		base.Branch(frame, int(T.defaultOffset))
	}
}
```



#### 3、lookupswitch指令

```go
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (L *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	L.defaultOffset = reader.ReadInt32()
	L.npairs = reader.ReadInt32()
	L.matchOffsets = reader.ReadInt32s(L.npairs * 2)
}

func (L *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	key := stack.PopInt()
	for i := int32(0); i < L.npairs*2; i += 2 {
		value := L.matchOffsets[i]
		if value == key {
			base.Branch(frame, int(L.matchOffsets[i+1]))
			return
		}
	}
	base.Branch(frame, int(L.defaultOffset))
}
```



### 14、wide指令

这是一个扩展指令，加载类指令、存储类指令、ref指令和iinc指令需要按索引访问局部变量表，索引一u1存储在字节码中，如果出现超过u1能存储的256，那么就需要扩展，原来一个字节的扩展为2字节



## 七、解释器

### 1、创建工厂类，根据不同tag获取不同的指令对象

```go
func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
    .......
```

### 2、interpret方法

```
func interpret(methodInfo *classfile.MemberInfo) {

    // 获取方法中的code属性
    codeAttr := methodInfo.CodeAttribute()

    // 从属性中解析最大局部变量表，最大栈，code属性
    maxLocals := codeAttr.MaxLocals()
    maxStack := codeAttr.MaxStack()
    bytecode := codeAttr.Code()

    // 创建一个线程
    thread := rtda.NewThread()

    // 创建一个栈帧
    frame := thread.NewFrame(uint(maxLocals), uint(maxStack))

    // 栈帧压入栈
    thread.PushFrame(frame)

    // 异常处理
    defer catchErr(frame)

    // 循环解析指令并执行
    loop(thread, bytecode)
}
```



## 八、测试

### 1、编写java代码

```java
/**
 * 提供java代码  - 1-100 求和
 *
 * @author : jucunqi
 * @since : 2025/6/20
 */
public class GuessTest {

    public static void main(String[] args) {

        int result = 0;
        for (int i = 1; i <= 100 ; i++) {
            result += i;
        }
    }
}
```



手动解析了一下字节码

![ch05_mannualanalyze.png](https://s2.loli.net/2025/06/20/slR6HYp4CwEQJDT.png)



### 2、使用go测试

因为没有实现return指令，所以会报错，但是局部变量表中已经可以看到5050我们想要的答案了，结果如图

![ch05_test.png](https://s2.loli.net/2025/06/20/3kzrIRVoiEft6yC.png)



# 第六章 类和对象

## 一、方法区

> 一块可以被多个线程共享的逻辑区域
>
> 主要存放从class文件获取的类信息、类变量，加载一次并缓存



### 1、类信息

```go
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
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethod(class, cf.Methods())
	return class
}
```



### 2、字段信息

```go
type Field struct {
	ClassMember          //继承自ClassMember
	constValueIndex uint //常量值索引
	slotId          uint //槽位ID
}

type ClassMember struct {
	accessFlags uint16 // 访问标识符
	name        string // 名称
	descriptor  string // 描述符
	class       *Class // 所属类
}
```



### 3、方法信息

```go
type Method struct {
	ClassMember        //继承自ClassMember
	maxStack    uint   //最大栈深度
	maxLocals   uint   //最大局部变量表大小
	code        []byte //字节码
}

type ClassMember struct {
	accessFlags uint16 // 访问标识符
	name        string // 名称
	descriptor  string // 描述符
	class       *Class // 所属类
}
```

### 4、其他信息

>instanceSlotCount uint          //实例字段数量
>staticSlotCount   uint             //静态字段数量
>staticVars        Slots                //静态变量



## 二、运行时常量池

> 主要存放两类信息：字面量和符号引用
>
> 字面量包括：整数、浮点数、字符串字面量
>
> 符号引用包括：类符号引用、字段符号引用、方法符号引用、接口方法符号引用

```go
// 常量接口
type Constant interface {
}

// 常量池
type ConstantPool struct {
	class  *Class // 所属类
	consts []Constant
}
```



### 1、类符号引用

```go
// 类符号引用
type ClassRef struct {
	SymRef //继承自SymRef
}

// 符号引用
type SymRef struct {
	cp        *ConstantPool // 常量池
	className string        // 类名
	class     *Class        // 类
}
```

### 2、字段符号引用

```go
// 字段符号引用
type FieldRef struct {
	MemberRef        //继承自MemberRef
	field     *Field // 字段
}
// 成员符号引用
type MemberRef struct {
	SymRef            //继承自SymRef
	name       string // 名称
	descriptor string
}
```



### 3、方法符号引用

```go
// 方法符号引用
type MethodRef struct {
	MemberRef         //继承自MemberRef
	method    *Method // 方法
}
```



### 4、接口方法符号引用

```go
// 接口方法符号引用
type InterfaceMethodref struct {
	MemberRef         //继承自MemberRef
	method    *Method // 方法
}
```

## 三、类加载器

>简单分为步骤：
>
>1. 读取class文件
>2. load加载
>3. 链接(验证、准备)



```go
// 类加载器
type ClassLoader struct {
	cp       *classpath.Classpath // 类路径
	classMap map[string]*Class    // 已加载的类
}

// 加载
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
	fmt.Printf("[Loaded %s from %s]\n,", name, entry)
	return class
}
```



## 四、对象、实例变量和类变量

> 使用slots来存放每个属性，数组结构
>
> 链接阶段，为变量赋默认值，为常量赋值
>
> 如果常量是引用类型(排除String)或者需要执行代码的，将通过<clinit> 方法进行复制，不会直接在常量值赋值



## 五、类和字段符号引用解析

```go
// 类解析
func (s *SymRef) ResolveClass() *Class {

	if s.class == nil {
		s.resolveClassRef()
	}
	return s.class
}

func (s *SymRef) resolveClassRef() {
	d := s.cp.class
	c := d.loader.LoadClass(s.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	s.class = c
}
```



```go
// 字段解析
func (r *FieldRef) ResolvedField() *Field {

	if r.field == nil {
		r.resolveFieldRef()
	}
	return r.field
}

func (r *FieldRef) resolveFieldRef() {

	d := r.cp.class
	c := r.ResolveClass()
	field := lookupField(c, r.name, r.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.Lang.IllegalAccessError")
	}
	r.field = field
}

func lookupField(c *Class, name string, descriptor string) *Field {

	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}
```



## 六、类和字段相关指令

> 实现了 new、putstatic、getstatic、putfield、getfield、instanceof、checkcast、ldc系列

列举两个指令的代码

### 1、NEW

```go
type NEW struct {
	base.Index16Instruction
}

func (n *NEW) Execute(frame *rtda.Frame) {

	cp := frame.Method().Class().ConstantPool()
	constant := cp.GetConstant(n.Index)
	classRef := constant.(*heap.ClassRef)
	class := classRef.ResolveClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

func (c *Class) NewObject() *Object {

	return newObject(c)
}

func newObject(c *Class) *Object {
	return &Object{
		class:  c,
		fields: newSlots(c.instanceSlotCount),
	}
}
```



### 2、putstatic

```go
func (p *PUT_STATIC) Execute(frame *rtda.Frame) {

	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(p.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())

	}
}
```

## 七、功能测试

> 修改staratJVM函数，通过类加载加载指定的类，实现方法区、运行时常量池、类和对象关联、以及部分指令

### 1、install

> go install ./ch06



### 2、执行

> sh ch06 Myobject

### 3、结果如图

![ch06test.png](https://s2.loli.net/2025/06/24/9GX2KVle7svauk5.png)



# 第七章 方法调用和返回

## 一、概述

> 调用角度，类可以分为：静态方法、实例方法。
>
> 实现角度，类可以分位：抽象方法、java方法(或jvm上的其他语言：Groovy、Scala等)
>
> 静态方法和抽象方法是互斥的
>
> jdk7之前，java提供了4个方法调用指令：invoke_static, invoke_special, invoke_vritual, invoke_interface
>
> jdk8提供了：invoke_dynamic（暂不实现）



**方法调用指令大致流程:** 

1. 解析常量池中方法 
2. 根据方法参数从操作数栈中弹出变量 
3. 新建栈帧压入虚拟机栈 
4. 将之前弹出的变量放入当前栈帧的局部变量表中(参数传递) 
5. 方法执行完毕后将返回值推入前一帧的操作数栈顶 
6. 弹出当前帧



## 二、解析方法符号引用

> 不实现接口的静态方法和默认方法



**解析方法符号引用大致流程：**

1. 获取当前类的常量池
2. 从常量池中获取方法引用的信息
3. 解析方法所在类、若未初始化需要进行初始化
4. 遍历类中的方法，根据方法名匹配 (动态匹配需要加一些额外逻辑，详见代码)
5. 方法访问权限验证(private、static等)
6. 创建方法栈帧，压入操作数栈



**以非接口方法符号引用为例，实例代码**

```go
func (r *MethodRef) ResolveMethod() *Method {

	if r.method == nil {
		r.resolveMethodRef()
	}
	return r.method
}

func (r *MethodRef) resolveMethodRef() {

	// 获取当前类
	currentClass := r.cp.class

	// 获取方法所在类
	methodClass := r.ResolveClass()

	// 如果方法所在类是个接口，抛出异常
	if methodClass.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 遍历类中的方法，获取当前方法
	method := lookupMethod(methodClass, r.name, r.descriptor)

	// 如果方法为空，抛出异常
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	// 验证访问权限
	if !method.isAccessibleTo(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	r.method = method
}

func lookupMethod(class *Class, name string, descriptor string) *Method {

	// 封装通过方法，从类中根据名称和描述符匹配方法
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterface(class.interfaces, name, descriptor)
	}
	return method
}

func lookupMethodInInterface(ifaces []*Class, name string, descriptor string) *Method {

	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterface(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
```



## 三、方法调用和参数传递

> 大致流畅，前面已经有提过了，提供一下实例代码

```go
// InvokeMethod 方法调用逻辑实现
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {

	// 获取当前栈帧所在的线程
	thread := invokerFrame.Thread()

	// 因为方法调用需要向当前线程的栈中压入一个栈帧，所以创建一个方法的栈帧
	newFrame := thread.NewFrame(method)

	// 压入虚拟机栈
	thread.PushFrame(newFrame)

	// 参数传递：1. 获取参数数量
	argSlotCount := int(method.ArgSlotCount())

	// 参数传递：2. 从调用方的操作数栈中参数变量，放入方法栈帧的局部变量表中
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack：因为还未实现Native方法，所以直接跳过
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n", method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
```



## 四、返回指令

> 做的事情比较简单：弹出当前帧，结果压入上一帧的操作数栈

**示例代码**

```go
func (i *IRETURN) Execute(frame *rtda.Frame) {

	// 获取当前线程
	thread := frame.Thread()

	// 获取操作数栈中的变量
	currentFrame := thread.PopFrame()
	result := currentFrame.OperandStack().PopInt()

	// 把变量压入调用线程栈帧中的操作数栈
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushInt(result)
}

```



## 五、方法调用指令

>Invoke_static 		 静态调用
>
>invoke_special	       构造器、super、private方法
>
>invoke_vritual                动态调用
>
>Invoke_interface           接口调用

**示例代码：以invoke_vritual为例**

```go
func (i *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {

	// 获取当前类的常量池
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()

	// 获取方法的符号引用
	methodRef := cp.GetConstant(i.Index).(*heap.MethodRef)

	// 方法解析
	resolvedMethod := methodRef.ResolveMethod()

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 从操作数栈中获取this对象引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)

	if ref == nil {
		// hack！
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPinterException")
	}
	// 校验protect权限
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() && ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
```



## 六、改进解释器

> 命令行新增 verbose:inst 和 verbose:class参数，用于输出类解析和指令解析详细日志

**loop函数支持多方法解析,示例代码**

```go
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		// 获取当前栈顶栈帧
		frame := thread.CurrentFrame()

		// 获取程序计数器
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadInt8()
		inst := instructions.NewInstruction(byte(opcode))
		inst.FetchOperands(reader)

		// 根据标识判断是否打印指令日志
		if logInst {
			logInstruction(frame, inst)

		}

		// execute
		frame.SetNextPC(reader.PC())
		inst.Execute(frame)

		// 结束标识
		if thread.IsStackEmpty() {
			break
		}
	}
}

```



## 七、类初始化

> 实际上就是执行类的<clinit> 方法、一下情况会触发
>
> - 执行new指令创建对象，类还未初始化
> - 执行putstatic、getstatis指令，存储静态变量，声明该字段的类还未初始化
> - 执行invoke_static指令调用类的静态方法，声明该方法的类还未初始化
> - 初始化类时，父类还未初始化
> - 反射操作

**示例代码**

```go
func InitClass(thread *rtda.Thread, class *heap.Class) {

	// 设置初始化标记位
	class.StartInit()

	// 计划执行clinit方法
	scheduleClinit(thread, class)

	// 初始化父类
	initSuperClass(thread, class)
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {

	clinit := class.GetClinitMethod()

	if clinit != nil {

		// 创建栈帧
		frame := thread.NewFrame(clinit)
		// 压入栈
		thread.PushFrame(frame)
	}
}
```



# 第八章 数组和字符串

## 一、概述

> 数组类和普通类不同，普通类从class文件加载，数组类由jvm虚拟机在运行时生成
>
> 数组对象创建方式不同，普通对象通过new指令，然后由构造器初始化、数组有newarray、anewarray、multianewarray指令创建



## 二、数组实现

### 1、数组对象

> 依然使用Object封装，不过字段信息类型修改为interface{}
>
> go语言中interface{}表示任意类型

```go
type Object struct {
	class *Class
	data  interface{} // 标识任何类型
}
```



### 2、数组类

> 不需要修改class结构体，增加NewArray函数

```go
func (c *Class) NewArray(count uint) *Object {

	if !c.IsArray() {
		panic("Not array class: " + c.name)
	}
	switch c.Name() {
	case "[Z":
		return &Object{c, make([]int8, count)}
	case "[B":
		return &Object{c, make([]int8, count)}
	case "[C":
		return &Object{c, make([]uint16, count)}
	case "[S":
		return &Object{c, make([]int16, count)}
	case "[I":
		return &Object{c, make([]int32, count)}
	case "[J":
		return &Object{c, make([]int64, count)}
	case "[F":
		return &Object{c, make([]float32, count)}
	case "[D":
		return &Object{c, make([]int64, count)}
	default:
		return &Object{c, make([]*Object, count)}

	}
}
```

### 3、加载数组类

> 修改ClassLoader类中的LoadClass方法

```go
func (c *ClassLoader) LoadClass(name string) *Class {

	// 判断类是否已经加载
	if class, ok := c.classMap[name]; ok {
		return class

	}

	// 判断类是否属于数组类
	if name[0] == '[' {
		return c.loadArrayClass(name)
	}

	// 加载非数组类
	return c.loadNonArrayClass(name)
}

func (c *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        name,       // 类名
		loader:      c,          // 加载器
		initStarted: true,       // 数组不需要初始化
		superClass:  c.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			c.LoadClass("java/lang/Cloneable"),
			c.LoadClass("java/io/Serializable"),
		},
	}
	c.classMap[name] = class
	return class
}
```



## 三、数组相关指令

### 1、newarray指令

> 创建基本数据类型的数组，有两个操作数
>
> 1. 紧跟在指令后面 u1，代表基本数据类型
> 2. 从操作数栈中获取，数组长度

```go
// 基本数据类型和操作数对应关系
const (
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

func (n *NEW_ARRAY) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出元素，代表数组长度
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	// 获取类加载器对象
	classLoder := frame.Method().Class().Loader()

	// 解析数组类
	arrClass := getPrimitiveArrayClass(classLoder, n.atype)

	// 创建数组对象
	arr := arrClass.NewArray(uint(count))

	// 压入操作数栈
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loder *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loder.LoadClass("[Z")
	case AT_BYTE:
		return loder.LoadClass("[B")
	case AT_CHAR:
		return loder.LoadClass("[C")
	case AT_SHORT:
		return loder.LoadClass("[S")
	case AT_INT:
		return loder.LoadClass("[I")
	case AT_LONG:
		return loder.LoadClass("[J")
	case AT_FLOAT:
		return loder.LoadClass("[F")
	case AT_DOUBLE:
		return loder.LoadClass("[D")
	default:
		panic("Invalid atype!")

	}
}
```



### 2、anewarray指令

> 创建引用数据类型的数组，有两个操作数
>
> 1. u2的操作数，执行常量池中类符号引用索引
> 2. 从操作数栈中弹出数组长度

```go
func (a *ANEW_ARRAY) Execute(frame *rtda.Frame) {

	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()

	// 获取类符号引用，并解析
	classRef := cp.GetConstant(a.Index).(*heap.ClassRef)
	componentClass := classRef.ResolveClass()

	// 操作数栈中获取数组长度
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

```



### 3、arraylength指令

> 获取数组的长度

```go

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (a *ARRAY_LENGTH) Execute(frame *rtda.Frame) {

	// 从栈顶获取数组引用
	stack := frame.OperandStack()
	arrRef := stack.PopRef()

	if arrRef == nil {
		panic("java.lang.NullPointException")
	}
	// 获取数组长度
	length := arrRef.ArrayLength()

	// 压入操作数栈
	stack.PushInt(length)
}

func (o *Object) ArrayLength() int32 {

	switch o.data.(type) {
	case []int8:
		return int32(len(o.data.([]int8)))
	case []int16:
		return int32(len(o.data.([]int16)))
	case []int32:
		return int32(len(o.data.([]int32)))
	case []int64:
		return int32(len(o.data.([]int64)))
	case []uint16:
		return int32(len(o.data.([]uint16)))
	case []float32:
		return int32(len(o.data.([]float32)))
	case []float64:
		return int32(len(o.data.([]float64)))
	case []*Object:
		return int32(len(o.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
```

### 4、<t>aload指令

> 从数组中获取元素，两个操作数
>
> 1. 操作数栈中获取数组索引
> 2. 操作数栈中获取数组引用

执行逻辑

1. 从操作数栈中弹出数组索引
2. 从操作数栈中弹出数组引用
3. 将数组指定引用的值压入操作数栈顶

以aaload为例，示例代码

```go
func (a *AALOAD) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()

	// 非空验证
	checkNotNil(arrRef)

	// 索引越界验证
	refs := arrRef.Refs()
	checkIndex(len(refs), index)

	// 结果压入栈顶
	stack.PushRef(refs[index])
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
```



### 5、 <t>astore指令

> 将指定值，放入数组中指定索引位置，三个操作数
>
> 1. 赋给数组的值
> 2. 数组索引
> 3. 数组引用

大致逻辑

1. 从操作数栈中弹出值
2. 从操作数栈弹出数组索引
3. 从操作数栈中弹出数组引用
4. 数组指定索引赋值

以iastore为例，示例代码

```go
func (i *IASTORE) Execute(frame *rtda.Frame) {

	// 获取操作数栈
	stack := frame.OperandStack()

	// 栈顶弹出目标值
	val := stack.PopInt()

	// 栈顶弹出数组索引
	index := stack.PopInt()

	// 栈顶弹出数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	// 校验数组是否越界 ， 将值放入数组
	ints := arrRef.Ints()
	checkIndex(len(ints), index)

	// 数组赋值
	ints[index] = int32(val)
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

```



### 6、multianewarray指令

> 一共有2+n个操作数
>
> 前两个操作数紧跟操作码后面，第一个代表常量池索引，可以查询到数组类型的符号引用
>
> 第二个操作数表示数组纬度
>
> n个操作数表示每个纬度的数组长度



## 四、测试数组

测试冒泡排序

```java
public class BubbleSortTest {

    public static void main(String[] args) {
        int[] arr = {22, 84, 77, 56, 10, 43, 59};
        int[] ints = bubbleSort(arr);
        for (int anInt : ints) {
            System.out.println(anInt);
        }
    }

    /**
     * 冒泡排序
     *
     * @param arr 数组
     * @return 排序后的数组
     */
    public static int[] bubbleSort(int[] arr) {

        boolean swapped = true;
        int j = 0;
        int tmp;
        while (swapped) {
            swapped = false;
            j++;
            for (int i = 0; i < arr.length - j; i++) {
                if (arr[i] > arr[i + 1]) {
                    tmp = arr[i];
                    arr[i] = arr[i + 1];
                    arr[i + 1] = tmp;
                    swapped = true;
                }
            }
        }
        return arr;
    }
}
```

测试结果，如图

![ch08-array-test.png](https://s2.loli.net/2025/07/02/kpm4Suzcoaj1QYi.png)



## 五、字符串

> java中字符串是以java/lang/String类型存储
>
> String类有两个重要变量value，类型是字符数组，jdk17中类型是字符数组
>
> 另一个是hash值，缓存字符串的hash码



### 1、字符串池

用map表示

```go
// 用map表示字符串池，key是go字符串，value是Java字符串
var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {

	// 从字符串池中获取
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	internedStrings[goStr] = jStr
	return jStr
}
```



### 2、测试字符串

> 久违的**hello world**即将到来

```java
package example.src.main.java.jvmgo.ch08;

public class HelloWorld {

    public static void main(String[] args) {
        System.out.println("Hello World");
    }
}
```

结果如图

![ch08-string-test.png](https://s2.loli.net/2025/07/02/7JtUClXg16FVyAo.png)



# 第九章 本地方法调用

> java中有很多使用native修饰的本地方法，使用c语言编写的
>
> 下面将使用go语言实现本地方法



## 一、注册和使用本地方法

> 定义map结构的registry
>
> key：类名+方法名+描述符
>
> value：本地方法函数

```go
type NativeMethod func(frame *rtda.Frame)
// 存储本地方法实现
var registry = map[string]NativeMethod{}

// Register 本地方法注册逻辑
func Register(className string, methodName string, methodDescriptor string, method NativeMethod) {

	// 定义缓存key
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}
```



## 二、调用本地方法

> java虚拟机规范留了两个指令代表本地方法分别是，0xFE、0xFF，下面使用0xFE标识
>
> 由于本地方法没有code属性，所以修改方法解析逻辑，如果是native方法，则创建code属性，两个字节，一个字节固定0xFE，另一个固定为方法返回指令
>
> 本地方法操作数栈深度最大值默认4，局部变量表最大槽数与方法参数一致
>
> 执行逻辑：
>
> 通过本地方法注册表找到方法函数，然后执行

```go
// 根据方法返回类型，构建本地方法的code属性，本地方法的第一个指令是0xfe
func (m *Method) injectCodeAttribute(returnType string) {

	// 最大栈深，默认4
	m.maxStack = 4

	// 布局变量表最大 = 参数个数
	m.maxLocals = m.argSlotCount

	// 根据返回类型，构建code字节码
	switch returnType[0] {
	case 'V':
		m.code = []byte{0xfe, 0xb1} // return
	case 'D':
		m.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		m.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		m.code = []byte{0xfe, 0xad} // lreturn
	case 'L', '[':
		m.code = []byte{0xfe, 0xb0} // areturn
	default:
		m.code = []byte{0xfe, 0xac} // ireturn
	}
}

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (i *INVOKE_NATIVE) Execute(frame *rtda.Frame) {

	// 获取方法名，类名，描述符
	method := frame.Method()
	methodName := method.Name()
	className := method.Class().Name()
	descriptor := method.Descriptor()

	// 根据上述参数，去本地方法注册表中匹配本地方法
	nativeMethod := native.FindNativeMethod(className, methodName, descriptor)

	// 非空校验
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + "." + descriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	// 执行本地方法
	nativeMethod(frame)
}
```

## 三、反射

### 1、类和对象之间的关系

> Java中每个类可以有多个实例对象，但是这能有一个Class对象，也就是类对象
>
> 类对象中存储这类的相关信息，在反射中很有用处

> 修改代码，class类也可以关联上类对象信息

```go
type Class struct {
	// ...
  
	jClass            *Object       //java.lang.Class实例 - 类对象、JVM 自动创建（类加载时）、用于获取类的元数据，动态操作类、每个类在 JVM 中只有一个 Class 对象
}
```

### 2、修改类加载加载逻辑，加载类的时候，绑定类对象信息

```go 
func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	classLoader := &ClassLoader{}
	classLoader.cp = cp
	classLoader.verboseFlag = verboseFlag
	classLoader.classMap = make(map[string]*Class)

	// 类与类对象绑定关联
	classLoader.loadBasicClasses()

	// 加载void和基本数据类型
	classLoader.loadPrimitiveClasses()
	return classLoader
}

func (c *ClassLoader) LoadClass(name string) *Class {

	// 判断类是否已经加载
	if class, ok := c.classMap[name]; ok {
		return class

	}

	var class *Class
	if name[0] == '[' {
		// 加载数组类
		class = c.loadArrayClass(name)
	} else {
		// 加载非数组类
		class = c.loadNonArrayClass(name)
	}

	if jlClassClass, ok := c.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}
	return class
}
```



### 3、基本类型的类

> void和基本数据类型也有对应的类对象，通过字面量来访问
>
> 有jvm运行时创建，并且都获取类对象的时候会调用Class.getPrimitiveClass()这个本地方法

###### 

### 4、修改ldc指令

> 类对象字面值是通过ldc指令加载的

```go
func _ldc(frame *rtda.Frame, index uint) {
  //....
	case *heap.ClassRef:

		// 转成类符号引用
		classRef := c.(*heap.ClassRef)
		// 解析类符号引用， 加载类并获取类的类对象
		classObj := classRef.ResolveClass().JClass()
		stack.PushRef(classObj)
		break
	default:
		panic("todo: ldc!")
	}
}
```

### 5、通过反射获取类名

> 为了支持通过反射获取类名，需要实现下面四个本地方法
>
> java.lang.Object.getClass()
>
> Java.lang.Class.getPrimitiveClass()
>
> java.lang.Class.getName0()
>
> java.lang.Class.desiredAssertionStatus0()

go语言中的init函数是一个特殊函数只要包被引用， 那么init方法就会执行

通过init函数注册本地方法

```go
func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

func getClass(frame *rtda.Frame) {

	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
```

以getName0方法为例

```go
func init() {
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
}
func getName0(frame *rtda.Frame) {

	// 从局部变量表中获取当前对象
	this := frame.LocalVars().GetThis()

	// 通过类对象的extra属性可以获取类的信息
	class := this.Extra().(*heap.Class)

	// 名称转换java类名，/ -> .
	name := class.JavaName()

	// 转换为String对象
	jString := heap.JString(class.Loader(), name)

	// 压入操作数栈
	frame.OperandStack().PushRef(jString)
}
```

### 6、测试本节代码

```java
public class ClassTest {

    public static void main(String[] args) {
        System.out.println(void.class.getName());
        System.out.println(boolean.class.getName());
        System.out.println(byte.class.getName());
        System.out.println(char.class.getName());
        System.out.println(short.class.getName());
        System.out.println(int.class.getName());
        System.out.println(long.class.getName());
        System.out.println(float.class.getName());
        System.out.println(double.class.getName());
        System.out.println(Object.class.getName());
        System.out.println(int[].class.getName());
        System.out.println(int[][].class.getName());
        System.out.println(Object[].class.getName());
        System.out.println(Object[][].class.getName());
        System.out.println(Runnable.class.getName());
        System.out.println("abc".getClass().getName());
        System.out.println(new double[0].getClass().getName());
        System.out.println(new String[0].getClass().getName());
    }
}

```

执行结果如图

![ch09-reflecttest.png](https://s2.loli.net/2025/07/03/MXWkqjRIw1JV2Y8.png)



## 四、字符串拼接和String.intern()方法

### 1、字符串拼接涉及到的Java类库

> 我们在Java中使用字符串拼接时，会被优化成使用StringBuilder的append方法拼接
>
> 为了执行append方法，需要实现下面四个本地方法
>
> 1. System.arraycopy()
> 2. Float.floatToRawIntBits()
> 3. Double.doubleToRawLongBits()
> 4. Double.longBitsToDouble()



```go
func init() {
	native.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

// 对应Java本地方法   public static native void arraycopy(Object src,  int  srcPos, Object dest, int destPos, int length);
func arraycopy(frame *rtda.Frame) {

	// 从局部变量表中拿到5个参数
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	// 非空校验
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	// 源数组和目标数组必须兼容
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}

	// 检查索引位置
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	// 数组拷贝
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src *heap.Object, dest *heap.Object) bool {

	srcClass := src.Class()
	descClass := dest.Class()

	// 必须都是数组
	if !srcClass.IsArray() || !descClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() ||
		descClass.ComponentClass().IsPrimitive() {
		return srcClass == descClass
	}
	return true
}
```

```go
func init() {
	native.Register("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
}

func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}

func init() {
	native.Register("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)
}

func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}

func longBitsToDouble(frame *rtda.Frame) {
	value := frame.LocalVars().GetLong(0)
	frame.OperandStack().PushDouble(float64(value))
}
```

### 2、String.intern()方法设计到的类库

```go
func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
```

### 3、测试本节代码

```java
public class StrTest {

    public static void main(String[] args) {
        String s1 = "abc1";
        String s2 = "abc1";
        System.out.println(s1 == s2);   // true
        int x = 1;
        String s3 = "abc" + x;
        System.out.println(s1 == s3);   // false
        s3 = s3.intern();
        System.out.println(s1 == s3);   // true
    }
}
```

测试结果如图

![ch09-strtest.png](https://s2.loli.net/2025/07/04/36MUgxN2AKiwCOa.png)



# 第十章 异常处理

## 一、概述

> Java中异常分为两类：Checked 和 Unchecked
>
> Unchecked包括：java.lang.RuntimeException、java.lang.Error以及他们的子类，其他的异常通常都是checked，所有异常最终都继承子java.lang.Throwable



## 二、异常抛出

> Java 代码中通过 throw关键字抛出
>
> 对应字节码指令：athrow



## 三、异常处理表

> 异常处理通过try-catch实现
>
> 根据异常处理表找到对应执行的逻辑
>
> 异常处理表的每一项都包含3个信息：处理哪部分代码抛出的异常，哪类异常，异常处理代码在哪里

```go
type ExceptionHandler struct {
	startPc   int       // 其实pc
	endPc     int       // 结束pc
	handlerPc int       // 跳转pc
	catchType *ClassRef // 捕获类型
}

type ExceptionTable []*ExceptionHandler

func (t ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {

	for _, handler := range t {
		// 判断pc是否处于当前异常类的起始和结束区间内
		if pc > handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler // catch - all
			}
			class := handler.catchType.ResolveClass()
			if class == exClass || exClass.IsSubClassOf(class) {
				// 并且处理类型是exClass的子类
				return handler
			}
		}
	}
	return nil
}
```

## 四、实现athrow指令

> 从操作数栈中，弹出异常对象引用
>
> 查看是否可以找到并跳转到异常处理代码
>
> 操作数栈清空，把异常对象引用放入栈顶

```go
func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {

		// 获取线程栈顶栈帧
		frame := thread.CurrentFrame()

		// 获取执行指令计数
		pc := frame.NextPC() - 1

		// 获取方法的异常处理pc
		handlerPc := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPc > 0 {

			// 如果异常处理pc>0 说明当前方法存在catch当前异常的逻辑
			stack := frame.OperandStack()

			// 清空操作数栈
			stack.Clear()

			// 把异常对象放入操作数栈
			stack.PushRef(ex)

			// 设置程序计数器执行行号
			frame.SetNextPC(handlerPc)
			return true
		}

		// 弹出当前栈帧
		thread.PopFrame()

		// 循环停止条件 - 线程中没有方法栈帧
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}
```

## 五、Java虚拟机栈信息

> 实现fillInStackTrace本地方法，具体逻辑不在给出

```go
func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

type StackTraceElement struct {
	fileName   string // 文件
	className  string // 类名
	methodName string // 方法名
	lineNumber int    // 行号
}
```

## 六、测试代码

```java
public class ParseIntTest {

    public static void main(String[] args) {
        foo(args);
    }

    private static void foo(String[] args) {
        try {
            bar(args);
        } catch (NumberFormatException e) {
            System.out.println(e.getMessage());
        }
    }

    private static void bar(String[] args) {
        if (args.length == 0) {
            throw new IndexOutOfBoundsException("no args!");
        }
        int x = Integer.parseInt(args[0]);
        System.out.println(x);
    }
}

```

测试结果，如图

![ch10-test.png](https://s2.loli.net/2025/07/07/6cwruQqNiEoG5Vn.png)
