package cart

import (
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/lib/mysql"
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/orm/aplum"
)

// 根据商品ID删除购物车数据
func DeleteCartByPid(userID int32, productID []int32) error {
	plumDB := mysql.GetAplumDB()
	return plumDB.Where("product_id in(?) AND user_id = ?", productID, userID).Delete(&aplum.TCart{}).Error
}
