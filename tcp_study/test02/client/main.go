package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println("tcp connect 127.0.0.1:8848 failed, err: ", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "Hello, 张飞"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed， err：", err)
			return
		}
		conn.Write([]byte(msg))
	}
}
