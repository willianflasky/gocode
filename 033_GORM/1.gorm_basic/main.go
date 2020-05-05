package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo table
type UserInfo struct {
	ID     uint
	Name   string
	Gendar string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "mygo:mygo@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	// u1 := UserInfo{1, "tom", "男", "篮球"}
	// u2 := UserInfo{2, "cat", "女", "鱼"}
	// 创建记录
	// db.Create(&u1)
	// db.Create(&u2)

	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "鱼")
	fmt.Printf("%#v\n", uu)

	// 更新
	// db.Model(&u).Update("hobby", "双色球")

	// 删除
	// db.Delete(&u)
}
