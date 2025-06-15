package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string // 目录的绝对路径
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {

	// 获取class文件的绝对路径
	fullPath := filepath.Join(self.absDir, className)

	// 读取文件流
	bytes, err := os.ReadFile(fullPath)

	return bytes, self, err
}

func (self *DirEntry) String() string {

	return self.absDir
}

func newDirEntry(path string) *DirEntry {

	// 将path转换为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir: absDir}
}
