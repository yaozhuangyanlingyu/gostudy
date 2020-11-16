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
		conn.Write([]byte("hello world"))
	}
}
