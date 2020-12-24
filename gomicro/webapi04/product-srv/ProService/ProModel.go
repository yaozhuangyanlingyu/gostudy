package ProService

import "fmt"

// 商品模型
type ProModel struct {
	Pid   int32  // 商品ID
	Pname string // 商品名称
}

/**
 * 创建商品模型对象
 */
func NewProModel(pid int32, pname string) *ProModel {
	return &ProModel{
		Pid:   pid,
		Pname: pname,
	}
}

/**
 * 获取商品列表数据
 */
func GetProductList(n int32) []*ProModel {
	list := make([]*ProModel, 0)
	var i int32 = 0
	for i = 0; i < n; i++ {
		list = append(list, NewProModel(i, fmt.Sprintf("华为手机-Note%d", i+1)))
	}
	return list
}
