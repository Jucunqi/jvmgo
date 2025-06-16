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