package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	UserName string   `form:"userName" json:"userName" uri:"userName" binding:"required"`
	Password string   `form:"password" json:"password" uri:"password" binding:"required"`
	Hobbys   []string `form:"hobbys" json:"bobbys" uri:"hobbys" binding:"required"`
}

func main() {
	router := gin.Default()

	// Path
	router.GET("/user/:userName/:password", func(c *gin.Context) {
		var user User
		c.ShouldBindUri(&user)
		c.JSON(http.StatusOK, user)
	})
	// Query
	router.GET("/index", func(c *gin.Context) {
		var user User
		c.ShouldBind(&user)
		c.JSON(http.StatusOK, user)
	})
	// Body
	router.POST("/register", func(c *gin.Context) {
		var user User
		c.ShouldBind(&user)
		c.JSON(http.StatusOK, user)
	})
	router.Run()
}
