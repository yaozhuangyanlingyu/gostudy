package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/spf13/viper"
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/api"
	"github.com/yaozhuangyanlingyu/gostudy/proto/cart"
)

// 读取配置文件
func LoadConfig(path string) {
	bindEnv()

	if viper.GetString("env") == "" {
		log.Fatal("没有指定[MICRO_SERVICE_ENV]环境变量,usage export MICRO_SERVICE_ENV=dev or usage export MICRO_SERVICE_ENV=prod")
	}
	log.Printf("当前运行环境 %s", viper.GetString("env"))

	LoadEnvConfig(path)
}

func bindEnv() {
	viper.SetEnvPrefix("micro_service")
	viper.BindEnv("env") // 区分线上线下的环境变量 dev | prod
}

func main() {
	LoadConfig("")

	// 将服务注册到consul
	consulReg := consul.NewRegistry(
		registry.Addrs(Config.GetString("consul.address")),
	)

	// 创建micro服务
	service := micro.NewService(
		micro.Name("cart-srv-test"),
		micro.Registry(consulReg),
		micro.Address(Config.GetString("server.address")),
	)

	// 解析命令参数
	service.Init()

	// 注册处理函数
	cart.RegisterCartHandler(service.Server(), new(api.Cart))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
