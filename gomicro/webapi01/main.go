package main

import (
	"github.com/micro/go-micro/web"
	"net/http"
)

func main() {

	server := web.NewService(web.Address(":8081")) // ??
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go micro"))
	})
	server.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("shop api"))
	})
	_ = server.Run()
}
