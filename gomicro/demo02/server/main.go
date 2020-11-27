package main

import (
	"context"
	"demo02/proto"
	"fmt"
	"log"

	"github.com/micro/go-micro"
)

type Product struct{}

func (_this *Product) GetProductByID(ctx context.Context, req *proto.GetProductByIDRequest, res *proto.GetProductByIDResponse) error {
	fmt.Println(req.Pid)
	res.Code = 200
	res.Msg = "Hello world"
	res.Data = &proto.GetProductByIDResponseData{
		Pid:  req.Pid,
		Name: "玫瑰色连衣裙",
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("product-srv"),
		micro.Version("product-v0.0.1"),
	)
	service.Init()
	proto.RegisterProductHandler(service.Server(), new(Product))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
