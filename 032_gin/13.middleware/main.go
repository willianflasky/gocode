package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	r := gin.New()
	// 1. 全局使用中间健
	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	// 2. 为某个路由设置中间键
	// r.GET("/test2", StatCost(), func(c *gin.Context) {
	// 	name := c.MustGet("name").(string) // 从上下文取值
	// 	log.Println(name)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Hello world!",
	// 	})
	// })

	// 3.为路由组设置中间键
	// shopGroup := r.Group("/shop")
	// shopGroup.Use(StatCost())

	r.Run(":9090")

}
