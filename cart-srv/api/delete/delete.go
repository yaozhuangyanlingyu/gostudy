package delete

import (
	"fmt"

	"github.com/yaozhuangyanlingyu/micro-srv/proto/cart"
)

type Delete struct {
	Req *cart.DeleteRequest
	Rsp *cart.DeleteResponse
}

func (_this *Delete) Handle() {
	// 查询数据库数据
	fmt.Println("党静娇我爱你")
}
