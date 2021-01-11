package list

import (
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/proto/cart"
)

type List struct {
	Req *cart.ListRequest
	Rsp *cart.ListResponse
}

func (_this *List) Handle() {
	// 查询数据库数据

}
