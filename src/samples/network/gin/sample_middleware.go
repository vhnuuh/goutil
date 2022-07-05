/*
 * https://www.jianshu.com/p/88dc5cd7c4c9
 */
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StatCost(c *gin.Context) {
	// 传递数据
	c.Set("name", "jack")
	start := time.Now()
	// 调用该请求的剩余处理程序
	c.Next()
	// 不调用该请求的剩余处理程序
	// c.Abort()
	// 计算耗时
	cost := time.Since(start)
	fmt.Println(cost)
}

func main() {
	router := gin.Default()
	router.GET("/", gin.Recovery(), StatCost, func(c *gin.Context) {
		// 获取中间件传递的数据
		name := c.MustGet("name").(string)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	router.Run()
}
