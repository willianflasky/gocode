package main

import (
	"todolist/dao"
	"todolist/models"
	"todolist/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 关闭数据库
	defer dao.Close()

	// 模型绑定、同步数据库
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}
