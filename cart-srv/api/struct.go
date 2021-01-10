package api

import (
	"context"

	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/api/list"
	"github.com/yaozhuangyanlingyu/gostudy/proto/cart"
)

type Cart struct{}

func (_this *Cart) List(ctx context.Context, req *cart.ListRequest, rsp *cart.ListResponse) error {
	rsp.Code = 200
	rsp.Msg = "success"

	handle := new(list.List)
	handle.Req = req
	handle.Rsp = rsp
	handle.Handle()
	return nil
}
