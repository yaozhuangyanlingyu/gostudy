package ProService

// 商品模型
type ProModel struct {
	Pid   int    // 商品ID
	PName string // 商品名称
}

/**
 * 创建商品模型对象
 */
func NewProModel(pid int, pname string) *ProModel {
	return &ProModel{
		Pid:   pid,
		PName: pname,
	}
}

/**
 * 获取商品列表数据
 */
func GetProductList(n int) []*ProModel {
	list := make([]*ProModel, 0)
	for i := 0; i < n; i++ {
		list = append(list, NewProModel(i, "商品："+string(i)))
	}
	return list
}
