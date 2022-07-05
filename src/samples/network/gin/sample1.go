package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 当客户端以GET方法请求/路径时，会执行后面的匿名函数
	r.GET("/", func(c *gin.Context) {
		// 返回json格式的数据
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.Run()
}
