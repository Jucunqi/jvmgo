package classpath

import (
	"os"
	"strings"
)

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
