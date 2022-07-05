package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// path

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/account/*name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// query
	router.GET("/index", func(c *gin.Context) {
		name := c.Query("name")
		id := c.DefaultQuery("id", "1")
		c.String(http.StatusOK, "Hello, name:%s, id:%v", name, id)
	})
	// http://localhost:8080/user?id=10&id=20&id=40
	router.GET("/user", func(c *gin.Context) {
		ids := c.QueryArray("id")
		c.JSON(http.StatusOK, gin.H{
			"ids": ids,
		})
	})
	// http://localhost:8080/article?articles[tittle]=golang
	router.GET("/article", func(c *gin.Context) {
		article := c.QueryMap("articles")
		c.JSON(http.StatusOK, article)
	})

	router.Run()
}
