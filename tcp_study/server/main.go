package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func eventHandle(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			return
		}
		recvStr := string(buf[0:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	// 启动tcp服务
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen tcp 127.0.0.1:8888 failed, err: ", err)
		return
	}

	for {
		// 接受客户端请
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accpet request feaild, err: ", err)
			return
		}
		go eventHandle(conn)
	}
}
