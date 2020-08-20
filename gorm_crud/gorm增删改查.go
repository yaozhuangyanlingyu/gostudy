package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TUser t_user表model
type TUser struct {
	ID     uint   `gorm:"primary_key"`
	Name   string `gorm:"type:varchar(255);not null"`
	Gender string `gorm:"type:char(10);not null"`
	Hobby  string `gorm:"type:varchar(255);not null"`
}

func main() {
	// 链接数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1)/shop?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&TUser{})

	u1 := TUser{1, "张飞", "男", "篮球"}
	u2 := TUser{2, "小乔", "女", "羽毛球"}

	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	fmt.Println("EXE END!")

	// 查询记录
	var u = new(TUser)
	db.First(u)
	fmt.Println(u)

	var uu TUser
	db.Find(&uu, "hobby=?", "羽毛球")
	fmt.Printf("%#v\n", uu)

	// 更新
	var u = new(TUser)
	db.First(&u)
	db.Model(&u).Update("hobby", "吉他")

	// 删除
	db.Delete(&u)
}
