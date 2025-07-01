package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absPath string // 绝对路径
}

func (self ZipEntry) readClass(className string) ([]byte, Entry, error) {

	// 解压zip包
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	// 关闭资源(defer定义的方法不会立即执行，会在return前执行)
	defer deferClose(r)

	// 读取文件
	for _, f := range r.File {

		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			bytes, err := io.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return bytes, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func deferClose(r *zip.ReadCloser) {

	err := r.Close()
	if err != nil {
		panic(err)
	}
}

func (self ZipEntry) String() string {

	return self.absPath
}

func newZipEntry(path string) *ZipEntry {

	// 将path转换为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}
