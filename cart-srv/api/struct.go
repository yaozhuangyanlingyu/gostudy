package api

import (
	"context"
	"fmt"

	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/api/delete"
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/api/list"
	"github.com/yaozhuangyanlingyu/micro-srv/proto/cart"
)

type Cart struct{}

// 购物车列表
func (_this *Cart) List(ctx context.Context, req *cart.ListRequest, rsp *cart.ListResponse) error {
	rsp.Code = 200
	rsp.Msg = "success"

	handle := new(list.List)
	handle.Req = req
	handle.Rsp = rsp
	handle.Handle()
	return nil
}

// 删除购物车数据
func (_this *Cart) Delete(ctx context.Context, req *cart.DeleteRequest, rsp *cart.DeleteResponse) error {
	rsp.Code = 200
	rsp.Msg = "success"

	handle := new(delete.Delete)
	handle.Req = req
	handle.Rsp = rsp
	handle.Handle()

	fmt.Println(rsp.Code, rsp.Msg)

	return nil
}
