package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
)

// 商品模型
type ProModel struct {
	Pid   int    // 商品ID
	PName string // 商品名称
}

//重点在这里
func callAPI(s selector.Selector) []ProModel {
	myCli := myhttp.NewClient(
		client.Selector(s),                     //通过Selector方法并且按照我们之前设置的策略返回当前service的一个节点
		client.ContentType("application/json"), //因为我们要调用的/v1/prods这个api返回的是json，所以这里要先规定好数据格式json(很重要,如果是其他的会报错)
	)

	//下面这句代码封装了一个请求(未发送)
	req := myCli.NewRequest("product-service", "/v1/prods", map[string]string{}) //这里最后一个参数是body，一般用于post请求的参数，因为我们这个api没有参数，所以我随便写了map[string]string{}

	//从之前的图可以看到我们的返回值是一个json，key-val为"data":[],所以我们的response要构建一个相同结构的，方便go-micro帮我们返回响应结构体map[string]int{"Size": 4}，这里的key大小写都可以，只要和request结构体该字段字母一样就行入 abcd:3 Abcd:3 ABCD:3都可以
	var resp []ProModel
	err := myCli.Call(context.Background(), req, &resp) //发起请求调用，并返回结果
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.95.129:8510"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin), //设置查询策略，这里是轮询
	)
	resp := callAPI(mySelector)
	fmt.Println(resp)
	for _, v := range resp {
		fmt.Println("productID:", v.Pid, " productName: ", v.PName)
	}
}
