package main

import (
    "fmt"
    "github.com/micro/go-micro/client/selector"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-plugins/registry/consul"
    "io/ioutil"
    "log"
    "net/http"
)

func callAPI(addr string, path string, method string) (string, error) { //封装http请求函数
    req, _ := http.NewRequest(method, "http://"+addr+path, nil)
    client := http.DefaultClient
    res, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()
    buf, _ := ioutil.ReadAll(res.Body)
    return string(buf), nil
}

// 7.基本方式调用Api(http api)
// 当服务很简单的时候可以用这种方法来做基本调用
func main() {
	// 新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
    consulReg := consul.NewRegistry(
        registry.Addrs("dev03.aplum-inc.com:8500"),
    )

    for {
		// 使用服务名获取服务
        getService, err := consulReg.GetService("test-srv-yaoxf")
        if err != nil {
            log.Fatal(err)
        }

		// go-micro很智能当服务列表中一个服务出现问题后，他会自动帮我们从轮询列表中删除调，我们的轮询只会访问有效的服务,如果getService有多个服务,从第一个开始轮询 如localhost:8080/v1/user,localhost:8081/v1/user
        next := selector.RoundRobin(getService)
        node, err := next()                     //type Next func() (*registry.Node, error)
        if err != nil {
            log.Fatal()
        }
        callRes, err := callAPI(node.Address, "/user", "GET")
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(callRes)
    }
}