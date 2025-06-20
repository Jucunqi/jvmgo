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

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {

	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

// 获取jre目录
func getJreDir(jreOption string) string {

	// 如果指定了jre path，并且目录存在，直接返回
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	// 如果未指定，查看当前目录下是否存在jre目录
	if exists("./jre") {
		return ".jre"
	}

	// 如果当前目录下也没有jre目录，获取系统环境变量的jre
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not found jre folder!")
}

// 验证目录是否存在
func exists(jreOption string) bool {

	//_, err := os.Stat(jreOption)
	//if err != nil {
	//	if os.IsNotExist(err) {
	//		return false
	//	}
	//}
	// 这里是两行代码的结合简写，分开写法如上
	if _, err := os.Stat(jreOption); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 解析用户classpath
func (c *Classpath) parseUserClasspath(cpOption string) {

	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}

// Parse 解析classpath
func Parse(jreOption, cpOption string) *Classpath {

	cp := &Classpath{}

	// 完善启动类和扩展类路径
	cp.parseBootAndExtClasspath(jreOption)

	// 完善用户类路径
	cp.parseUserClasspath(cpOption)
	return cp
}

// ReadClass 以此从启动类、扩展类、用户类路径下寻找
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {

	// 拼接.class后缀
	className = className + ".class"
	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := c.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return c.userClasspath.readClass(className)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}
