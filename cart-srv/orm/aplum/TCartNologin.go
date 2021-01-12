package aplum

/**
 * 未登录购物车
 */
type TCartNologin struct {
	ID             int64   `gorm:"column:id" json:"id"`
	KeyID          int32   `gorm:"column:key_id" json:"key_id"`
	ProductID      int32   `gorm:"column:product_id" json:"product_id"`
	CartPrice      float64 `gorm:"column:cart_price" json:"cart_price"`
	SourcePlatform string  `gorm:"column:source_platform" json:"source_platform"`
	CreateTime     int64   `gorm:"column:create_time" json:"create_time"`
}

// TableName GORMTableName
func (t *TCartNologin) TableName() string {
	return "t_cart_nologin"
}
