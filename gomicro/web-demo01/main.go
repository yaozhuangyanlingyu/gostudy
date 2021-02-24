package main

import (
    "github.com/micro/go-micro/web"
    "net/http"
	"encoding/json"
)

type Person struct {
	Name   string
	Age    int64
	Weight float64
}

// 1. 启动一个简单的Server
func main() {
	/*第一种方法，新建一个Option类型的函数，作为参数传递给NewService方法
    addr := func(o *Options){
        o.Address = "8001"
    }
    server := web.NewService(addr)
    */

    /*第二种方法,使用官方提供的简易方法web.Address(":8001")
    func Address(a string) Option {
        return func(o *Options) {
            o.Address = a
        }
    } //其实是对第一种方法的简易封装
    */
	server:=web.NewService(
		web.Address(":8081"),
	)
	server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		data := Person{
			Name:   "yaoxf",
			Age:    31,
			Weight: 65,
		}
		strData, _ := json.Marshal(data)
		writer.Write([]byte(strData))
	})
	server.Run()
}