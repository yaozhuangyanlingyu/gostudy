package delete

import (
	cartModel "github.com/yaozhuangyanlingyu/gostudy/cart-srv/model/cart"
	cartNologinModel "github.com/yaozhuangyanlingyu/gostudy/cart-srv/model/cartnologin"
	"github.com/yaozhuangyanlingyu/micro-srv/proto/cart"
)

type Delete struct {
	Req *cart.DeleteRequest
	Rsp *cart.DeleteResponse
}

func (_this *Delete) Handle() {
	var err error
	userID := _this.Req.GetUserID()
	keyID := _this.Req.GetKeyID()
	productID := _this.Req.GetProductID()

	// 删除数据库数据
	if userID > 0 {
		err = cartModel.DeleteCartByProductid(userID, productID)
	} else {
		err = cartNologinModel.DeleteCartByProductid(keyID, productID)
	}
	if err != nil {
		_this.Rsp.Code = 50000
		_this.Rsp.Msg = err.Error()
	}
}
