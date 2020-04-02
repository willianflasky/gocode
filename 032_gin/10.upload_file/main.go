package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return

		}

		dst := path.Join("./", file.Filename)
		c.SaveUploadedFile(file, dst)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("%s uploaded!", file.Filename),
		})
	})

	r.Run(":9090")
}
