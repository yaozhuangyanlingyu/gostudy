package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
继续在 gin-demo 目录下新建一个 examples 目录用于存放示例代码，然后在该目录下运行 go mod init gin-demo/examples 初始化 go.mod 文件。
使用 AsciiJSON 方法可以生成只包含 ASCII 字符的 JSON 格式数据，对于非 ASCII 字符会进行转义：
*/

func main() {
	r := gin.Default()

	r.GET("/asciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Gin框架",
			"tag":  "<br>",
		}

		// 输出: {"lang":"Gin\u6846\u67b6","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
