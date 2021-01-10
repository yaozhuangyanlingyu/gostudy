package list

import (
	"github.com/yaozhuangyanlingyu/gostudy/proto/cart"
)

type List struct {
	Req *cart.ListRequest
	Rsp *cart.ListResponse
}

func (_this *List) Handle() {
}
