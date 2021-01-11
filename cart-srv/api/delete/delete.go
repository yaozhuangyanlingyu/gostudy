package delete

import (
	"fmt"

	cartModel "github.com/yaozhuangyanlingyu/gostudy/cart-srv/model/cart"
	"github.com/yaozhuangyanlingyu/micro-srv/proto/cart"
)

type Delete struct {
	Req *cart.DeleteRequest
	Rsp *cart.DeleteResponse
}

func (_this *Delete) Handle() {
	userID := _this.Req.GetUserID()
	keyID := _this.Req.GetKeyID()
	productID := _this.Req.GetProductID()
	sourcePlatform := _this.Req.GetSourcePlatform()

	// 查询数据库数据
	if userID > 0 {
		err := cartModel.DeleteCartByPid(userID, productID)
		if err != nil {
			fmt.Println(err)
		}
	} else {

	}
	fmt.Println(userID, keyID, productID, sourcePlatform)
}
