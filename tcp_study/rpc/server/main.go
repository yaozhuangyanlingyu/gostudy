package main

import (
	"io"
	"net/http"
)

type Panda int

func (this *Panda) Getinfo(argType int, replyType *int) error {

}

func main() {
	// 注册1个页面请求
	http.HandleFunc("/panda", pandatext)

	// new一个对象
}

// 用来实现web函数
func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "panda")
}
