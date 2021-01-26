package main

import (
	"context"
	"fmt"
	"net"

	"github.com/yaozhuangyanlingyu/gostudy/grpc-study/test2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetProduct(ctx context.Context, req *proto.ProductRequest) (*proto.ProductResponse, error) {
	return &proto.ProductResponse{
		Pid:    "100",
		Status: "onsale",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterProductServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
