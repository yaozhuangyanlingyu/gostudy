/*
 * @Author: yaoxf
 * @Date: 2020-11-25 01:06:07
 * @LastEditTime: 2020-11-25 01:29:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gomicro\demo03\server.go
 */

package main

import (
	"context"
	"fmt"

	"demo03/proto"

	"github.com/micro/go-micro/v2"
)

type ProductService struct{}

func (h *ProductService) Hello(ctx context.Context, req *proto.ProductRequest, res *proto.ProductResponse) error {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	res.Code = 200
	res.Msg = "党静娇"
	res.Data = "党静娇 我爱你"

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("product"),
	)
	service.Init()
	proto.RegisterProductServiceHandler(service.Server(), new(ProductService))

	if err := service.Run(); err != nil {
		fmt.Println("service err", err)
	}

}
