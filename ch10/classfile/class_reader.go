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
