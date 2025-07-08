package classpath

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// 通配符模式也是返回CompositeEntry对象
func newWildcardEntry(path string) CompositeEntry {

	// 创建返回值对象
	var compositeEntry []Entry
	// remove *
	baseDir := path[:len(path)-1]
	// 获取到目录下每个文件后，回调这个方法
	walkFn := func(subPath string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}
		// 过滤子目录
		if info.IsDir() && subPath != baseDir {
			return filepath.SkipDir
		}
		// 获取jar包，并解析
		if strings.HasSuffix(subPath, ".jar") || strings.HasSuffix(subPath, ".JAR") {

			entry := newZipEntry(subPath)
			compositeEntry = append(compositeEntry, entry)
		}
		return nil
	}
	// 获取目录下的所有信息
	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		return nil
	}
	return compositeEntry
}
