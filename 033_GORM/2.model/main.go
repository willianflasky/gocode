package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User table
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}


type userInfo struct {
  ID        uint      // column name is `id`
  Name      string    // column name is `name`
  CreatedAt time.Time
}
// TableName  指定表名
func (u userInfo) TableName() string {
	return "userinfo"
}


func main() {
	// 为表增加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "eeo_" + defaultTableName
	}
	db, err := gorm.Open("mysql", "mygo:mygo@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 禁用表名复数形式;
	// db.SingularTable(true)

	// 使用User结构体创建名为`deleted_users`的表
	// db.Table("deleted_users").CreateTable(&User{})

	// 迁移表
	db.AutoMigrate(&userInfo{})

	// 新增数据
	// u1 := userInfo{1, "tom", time.Now()}
	// u2 := userInfo{2, "CAT", time.Now()}
	// db.Create(&u1)
	// db.Create(&u2)

	// 更新
	db.Model(&u1).Update("CreatedAt", time.Now())
	db.Model(&u1).Update("name", "litle tom")
}
