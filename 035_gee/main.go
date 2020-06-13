package main

import (
	"gee/gee"
	"net/http"
)

// index 根
func index(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
}

// hello 业务
func hello(c *gee.Context) {
	// expect /hello?name=geektutu
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

// login 登录
func login(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}

// 正则匹配 ":name"
func reHello(c *gee.Context) {
	// expect /hello/geektutu
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
}

// 正则匹配 "*filepath"
func reFilePath(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
}

func v1root(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
}

func v1world(c *gee.Context) {
	// expect /hello?name=geektutu
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func main() {
	r := gee.New()
	// r.GET("/", index)
	// r.GET("/hello", hello)
	// r.POST("/login", login)
	r.GET("/hello/:name", reHello)
	r.GET("/assets/*filepath", reFilePath)
	v1 := r.Group("v1")
	{
		v1.GET("/", v1root)
		v1.GET("/world", v1world)
	}

	r.Run(":9999")
}
