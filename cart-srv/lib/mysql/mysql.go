package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/yaozhuangyanlingyu/micro-srv/loader"
	"github.com/yaozhuangyanlingyu/micro-srv/mysql"
)

// 初始化mysql
func InitDB() {
	mysql.AddDB("aplum", loader.Config.GetStringMapString("aplumDB"))
}

// 获取数据库实例
func GetAplumDB() *gorm.DB {
	return mysql.GetDB("aplum")
}
