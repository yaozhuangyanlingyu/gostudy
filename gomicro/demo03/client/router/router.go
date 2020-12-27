package router

import(
	"github.com/gin-gonic/gin"
)

/**
 * 获取路由
 */
func GetGinRouter() {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/prods", func(c *gin.Context) {
			// 验证参数
			var rquestParams proto.GetProductListRequest
			size, err := strconv.ParseInt(c.DefaultQuery("size", "10"), 10, 32)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			rquestParams.Size = int32(size)

			// 请求product服务
			rsp, err := productSrv.GetProductList(context.Background(), &rquestParams)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			// 获取数据
			c.JSON(200, gin.H{
				"size": rquestParams.Size,
				"data": rsp.Data,
			})
		})
	}
	return ginRouter
}


