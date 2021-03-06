package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

/**
 * 通过consul发现商品服务地址
 */
func getProductSrvAddr() string {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.95.129:8510"),
	)
	servs, err := consulReg.GetService("product-service")
	if err != nil {
		log.Fatal(err)
	}

	// 随机选择一个服务
	/*
		next := selector.Random(servs)
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}*/

	// 轮询方式
	next := selector.RoundRobin(servs)
	node, err := next()
	if err != nil {
		log.Fatal("cannot get services")
	}
	fmt.Println(node.Address)
	return node.Address
}

type ProModel struct {
	Pid   int    `json:"pid"`
	PName string `json:"pname"`
}

/**
 * 主站API
 */
func main() {
	// 使用gin路由
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/index", func(c *gin.Context) {
			// 获取商品数据
			resp, err := http.Get("http://" + getProductSrvAddr() + "/v1/prods")
			if err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  err.Error(),
					"data": nil,
				})
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  err.Error(),
					"data": nil,
				})
				return
			}

			// 将json转换为结构体
			var list []*ProModel
			err = json.Unmarshal(body, &list)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  err.Error(),
					"data": nil,
				})
				return
			}
			fmt.Println(list)

			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "success",
				"data": list,
			})
		})
	}

	// 启动go micro服务
	server := web.NewService(
		web.Address(":8888"),   // 监听IP:端口
		web.Handler(ginRouter), // 使用gin路由
	)
	server.Run()
}
