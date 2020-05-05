package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB 连接变量
	DB *gorm.DB
)

// InitMySQL 连接
func InitMySQL() (err error) {
	dsn := "mygo:mygo@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

// Close 关闭
func Close() {
	DB.Close()
}
