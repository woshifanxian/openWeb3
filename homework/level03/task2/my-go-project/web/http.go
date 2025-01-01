package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	// 定义 POST 请求的路由
	r.POST("/submit", func(c *gin.Context) {
		// 从 POST 表单中获取数据
		name := c.PostForm("name")
		age := c.PostForm("age")

		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"status": "received",
			"name":   name,
			"age":    age,
		})
	})

	r.PUT("/xxxput")
	//监听端口默认为8080
	r.Run(":8000")
}
