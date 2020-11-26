package main

import (
	"context"
	"fmt"

	proto "demo01/greeter"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("asd"),
	)
	service.Init()

	// create the greeter client using the service name and client
	greeter := proto.NewGreeterClient("aplum_shop", service.Client())

	// request the Hello method on the Greeter handler
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
