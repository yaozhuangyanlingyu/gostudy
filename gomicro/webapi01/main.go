package main

import (
	"net/http"

	"github.com/micro/go-micro/web"
)

func main() {
	/*
		addr := func(o *web.Options) {
			o.Address = ":8001"
		}
		service := web.NewService(addr)
		service.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Hello world"))
		})
		service.Run()
	*/

	service := web.NewService(web.Address(":8001"))
	service.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world"))
	})
	service.Run()
}
