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

#### 1）通用属性定义

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

#### 2) Code属性

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

#### 3) ConstantValue 属性

```
ConstantValue_attribute {
    u2 attribute_name_index;    // 固定为 "ConstantValue"
    u4 attribute_length;        // 固定为 2
    u2 constantvalue_index;     // 指向常量池中的常量值
}
```

#### 4) Exceptions 属性

```
Exceptions_attribute {
    u2 attribute_name_index;    // 固定为 "Exceptions"
    u4 attribute_length;
    u2 number_of_exceptions;    // 异常数量
    u2 exception_index_table[number_of_exceptions]; // 指向常量池中的异常类
}
```

#### 5) LineNumberTable 属性

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

#### 6) LocalVariableTable 属性

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

#### 7) SourceFile 属性

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