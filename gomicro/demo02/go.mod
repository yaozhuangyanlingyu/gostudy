module demo02

go 1.15

replace "proto" => "./proto"

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1 // indirect
	google.golang.org/protobuf v1.22.0
)
