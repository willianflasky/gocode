package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 静态路径
	r.Static("/static", "./static")
	// 为模板增加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(200, "posts/index.tmpl", gin.H{"title": "posts pages"})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(200, "users/index.tmpl", gin.H{"title": "<a href='https://www.baidu.com'>百度</a>"})
	})
	r.Run(":9090")
}
