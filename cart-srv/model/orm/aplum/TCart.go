package aplum

/**
 * 购物车
 */
type TCart struct {
	ID             int64   `gorm:"column:id" json:"id"`
	UserID         int32   `gorm:"column:user_id" json:"user_id"`
	ProductID      int32   `gorm:"column:product_id" json:"product_id"`
	CartPrice      float64 `gorm:"column:cart_price" json:"cart_price"`
	SourcePlatform string  `gorm:"column:source_platform" json:"source_platform"`
	CreateTime     int64   `gorm:"column:create_time" json:"create_time"`
}

// TableName GORMTableName
func (t *TCart) TableName() string {
	return "t_cart"
}
