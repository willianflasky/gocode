package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// name := c.Query("query")  // 返回字符串
		// name := c.DefaultQuery("query", "somebody")  // 取不到值，指定默认值。
		name, ok := c.GetQuery("query") //  return string, bool
		if !ok {
			name = "somebody"
		}
		c.JSON(200, gin.H{"name": name})
	})
	r.Run(":9090")
}
