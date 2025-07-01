package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {

	// 遍历每一个entry，然后读取文件内容
	for _, entry := range self {

		// 调用每个entry的readClass方法
		bytes, from, err := entry.readClass(className)
		if err == nil {
			return bytes, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}

func newCompositeEntry(pathList string) CompositeEntry {

	compositeEntry := []Entry{}
	split := strings.Split(pathList, pathListSeparator)
	for _, path := range split {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
