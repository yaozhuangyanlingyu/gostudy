module github.com/yaozhuangyanlingyu/gostudy/cart-srv

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/yaozhuangyanlingyu/gostudy v0.0.0-20210111072803-e52f251e9bb6
	github.com/yaozhuangyanlingyu/micro-srv v0.0.0-20210111035334-261d0ff20b4a
	google.golang.org/protobuf v1.23.0
)

replace github.com/yaozhuangyanlingyu/micro-srv => /home/yaoxiaofeng/test/micro-srv
