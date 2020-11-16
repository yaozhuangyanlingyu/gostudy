package main

import (
	"fmt"
	"net"
	"bytes"
	"encoding/binary"
)

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println("tcp connect 127.0.0.1:8848 failed, err: ", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "Hello, 张飞"
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed， err：", err)
			return
		}
		conn.Write(data)
	}
}
