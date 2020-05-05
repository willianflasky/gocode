package routers

import (
	"todolist/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由配置
func SetupRouter() (DB *gin.Engine) {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.CreateToDo)
		v1Group.GET("/todo", controller.GetToDoList)
		v1Group.PUT("/todo/:id", controller.UpdateAToDo)
		v1Group.DELETE("/todo/:id", controller.DeleteAToDo)
	}
	return r
}
