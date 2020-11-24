package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
)

type HelloService struct{}

func (h *HelloService) Hello(ctx context.Context, req *TestHello.HelloRequest, rp *TestHello.HelloResponse) error {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	name := req.Name
	rp.Code = 200

	rp.Msg = fmt.Sprintf("%s helloworld", name)
	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("hello"),
	)
	service.Init()
	TestHello.RegisterHelloServiceHandler(service.Server(), new(HelloService))

	if err := service.Run(); err != nil {
		fmt.Println("service err", err)
	}

}
