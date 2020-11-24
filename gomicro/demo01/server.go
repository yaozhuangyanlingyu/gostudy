/*
 * @Author: yaoxf
 * @Date: 2020-11-24 01:48:19
 * @LastEditTime: 2020-11-25 00:20:54
 * @LastEditors: Please set LastEditors
 * @Description: 服务端
 * @FilePath: \gomicro\demo01\server.go
 */
package main

import (
	"context"
	"demo01/proto"
	"fmt"

	"github.com/micro/go-micro"
)

type Product struct{}

func (this *Product) GetProductByID(ctx context.Context, req *proto.GetProductByIdRequest, res *proto.GetProductByIdResponse) error {
	fmt.Pritnln("Hello")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hello"),
	)
	service.Init()
	Product.ProductServiceHandler(service.Server(), new(Product))

	if err := service.Run(); err != nil {
		fmt.Println("service err", err)
	}
}
