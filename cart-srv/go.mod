module github.com/yaozhuangyanlingyu/gostudy/cart-srv

go 1.13

require (
	code.aliyun.com/aplum_service/micro-srv v0.2.3 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/yaozhuangyanlingyu/gostudy v0.0.0-20210112092825-944df317ef74
	github.com/yaozhuangyanlingyu/micro-srv v0.0.0-20210111035334-261d0ff20b4a
	google.golang.org/protobuf v1.24.0
)

replace github.com/yaozhuangyanlingyu/micro-srv => /home/yaoxiaofeng/test/micro-srv

replace google.golang.org/protobuf => google.golang.org/protobuf v1.23.0
