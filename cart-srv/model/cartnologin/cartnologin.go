package cartnologin

import (
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/lib/mysql"
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/orm/aplum"
)

// 根据商品ID删除购物车数据
func DeleteCartByProductid(KeyID string, productID []int32) error {
	plumDB := mysql.GetAplumDB()
	return plumDB.Where("product_id in(?) AND key_id = ?", productID, KeyID).Delete(&aplum.TCartNologin{}).Error
}
