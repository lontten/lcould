package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST("/event/dir", func(c *gin.Context) {
		// 1.获取参数

		err := c.BindJSON(&dir)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "请求参数错误",
			})
			return
		}
		c.String(http.StatusOK, "hello World!")
	})
	r.Run(":9999")
}
