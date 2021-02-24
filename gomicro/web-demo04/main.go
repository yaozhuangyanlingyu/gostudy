package main

import(
	"fmt"
    "github.com/micro/go-micro/client/selector"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-plugins/registry/consul"
    "log"
)

// 4.(服务发现1)获取consul服务列表，selector随机选择
// 来源：https://www.cnblogs.com/hualou/p/12097256.html
func main() {
    consulReg := consul.NewRegistry( //新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
        registry.Addrs("dev03.aplum-inc.com:8500"),
    )
    getService, err := consulReg.GetService("test-srv-yaoxf") //使用服务名获取服务
    if err != nil {
        log.Fatal(err)
    }
    next := selector.Random(getService) //如果getService有多个服务,随机取一个出来localhost:8080/v1/user,localhost:8081/v1/user
    node, err := next()                 //type Next func() (*registry.Node, error)
    if err != nil {
        log.Fatal()
    }
    fmt.Println(node.Id, node.Address, node.Metadata) //可以看到我们的id address还有metadata
}