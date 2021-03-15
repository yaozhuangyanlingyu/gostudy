package main

import (
	"context"
	"demo03/proto"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

/**
 * 创建商品模型对象
 */
func NewProModel(pid int32, pname string) *proto.ProductRow {
	return &proto.ProductRow{
		ProductID:   pid,
		ProductName: pname,
	}
}

type Product struct{}

func (_this *Product) GetProductList(ctx context.Context, req *proto.GetProductListRequest, res *proto.GetProductListResponse) error {
	list := make([]*proto.ProductRow, 0)
	var i int32 = 0
	for i = 0; i < req.Size; i++ {
		list = append(list, NewProModel(i, fmt.Sprintf("华为手机-Note%d", i+1)))
	}
	res.Data = list
	return nil
}

func main() {
	// 将服务注册到consul
	reg := consul.NewRegistry(
		registry.Addrs("dev03.aplum-inc.com:8500"),
	)

	// 将服务注册到etcd
	/*
		reg := etcdv3.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{
				"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
			}
		})*/

	// 创建micro服务
	service := micro.NewService(
		micro.Name("product-srv-yaoxf"),
		micro.Registry(reg),
	)

	// 解析命令参数
	service.Init()

	// 注册处理函数
	proto.RegisterProductHandler(service.Server(), new(Product))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
