package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// ANY
	r.Any("/index", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(200, gin.H{"msg": "/index get"})
		case http.MethodPost:
			c.JSON(200, gin.H{"msg": "/index post"})
		}
	})
	// 默认路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "not found"})
	})

	// 路由组
	userGroup := r.Group("/user")
	{

		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "/user/index"})
		})
		userGroup.GET("/home", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "/user/home"})
		})
	}
	r.Run(":9090")
}
