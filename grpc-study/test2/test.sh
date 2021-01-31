protoc -I proto/ proto/product.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/product_type.proto --go_out=plugins=grpc:proto
