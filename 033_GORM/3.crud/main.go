package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User table
type User struct {
	gorm.Model
	ID   int64
	Name string `gorm:"default:'nobody'"`
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "mygo:mygo@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	// 1、新增数据
	// u2 := User{Name: "panda", Age: 61}
	// db.Create(&u2)

	// 2. 一般查询
	// var user User
	//  获取第一条数据
	// db.First(&user)

	// 获取随机一条
	// db.Take(&user)

	// 获取最后一条
	// db.Last(&user)

	// 查询所有数据
	// var users []User
	// db.Find(&users)
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	// 查询主健为7的
	// db.First(&user, 7)

	// 3. 条件查询
	// 查询一条
	// db.Where("name = ?", "cat").First(&user)

	// 查询多条
	var users []User
	// db.Where("name =?", "nobody").Find(&users)

	// db.Where("name <> ?", "nobody").Find(&users)

	// db.Where("name IN (?)", []string{"cat", "fesh"}).Find(&users)

	// db.Where("name LIKE ?", "no%").Find(&users)

	// db.Where("name = ? AND age = ?", "nobody", "20").Find(&users)

	// db.Where("updated_at > ? and name = ?", "2020-04-12 00:00:00", "panda").Find(&users)

	// db.Where(&User{Name: "cat", Age: 51}).First(&users)

	// db.Where(map[string]interface{}{"name": "cat", "age": "51"}).Find(&users)

	// db.Where([]int{2, 3, 40}).Find(&users)

	// db.Where(&User{Name: "nobody", Age: 0}).Find(&users)

	// 4. Not查询
	// db.Not("name", "cat").First(&users)

	// db.Not("name", []string{"cat", "nobody"}).Find(&users)

	// db.Not([]int64{1, 2, 3}).First(&users)

	// db.Not([]int64{}).Find(&users)

	// db.Not("name = ?", "cat").First(&users)

	// db.Not(User{Name: "cat"}).First(&users)

	// 5. Or查询
	// db.Where("name = ?", "cat").Or("name = ?", "panda").Find(&users)
	// db.Where("name = ?", "cat").Or(User{Name: "panda"}).Find(&users)
	// db.Where("name = ?", "cat").Or(map[string]interface{}{"name": "padna"}).Find(&users)

	// 6. 内联查询
	// db.First(&users, 3)

	// db.First(&users, "id = ?", "string_primary_key")

	// db.Find(&users, "name = ?", "cat")

	// db.Find(&users, User{Age: 20})

	// db.Find(&users, map[string]interface{}{"name": "cat"})

	// 7. 额外查询
	// db.Set("gorm:query_option", "FOR UPDATE").First(&users, 10)

	// 8. FirstOrInit
	// db.Where(User{Name: "cat"}).FirstOrInit(&users)

	// 9. Attrs
	// db.Where(User{Name: "not_exist"}).Attrs(User{Age: 20}).FirstOrInit(&users)
	// db.Where(User{Name: "not_exist"}).Assign(User{Age: 20}).FirstOrInit(&users)

	// db.FirstOrCreate(&users, User{Name: "non_exsiting"})

	// 10. 子查询
	// db.Where("age = ?", db.Table("users").Select("AVG(age)").SubQuery()).Find(&users)

	// 11. 选择字段
	// db.Select("name").First(&users)

	// 12. order
	// db.Order("age desc").Order("name desc").Find(&users)

	// 13. limit
	// db.Limit(3).Find(&users)

	// 14. offset
	// db.Offset(9).Find(&users)

	// 15. count
	// var count int32
	// db.Table("users").Count(&count)

	// 16. having group
	// type Result struct {
	// 	Date  time.Time
	// 	Total int
	// }
	// var rets Result

	// db.Table("users").Select("sum(age)").Group("name").Scan(&rets)
	// fmt.Println(rets)

	// 17. join
	// db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

	// 18. pluck
	// var ages []int64
	// db.Find(&users).Pluck("age", &ages)
	// db.Table("users").Pluck("age", &ages)
	// db.Table("users").Select("name, age").Find(&users)

	// 19. scan
	// type Result struct {
	// 	Name string
	// 	Age  int
	// }
	// var result Result
	// var results []Result
	// db.Table("users").Select("name,age").Where("name = ?", "panda").Scan(&result)
	// db.Table("users").Select("name,age").Where("id > ?", 0).Scan(&results)
	// db.Raw("select name,age from users where name = ?", "panda").Scan(&result)

	// 20. 链式操作
	// tx := db.Where("name = ?", "nobody")
	// tx = tx.Where("age = ?", 30)
	// tx.Find(&users)

	// 21. scopes
	// db.Scopes(ageLT30).Find(&users)

	// 22. save
	// var user User

	// db.First(&user)
	// user.Age = 20
	// user.Name = "eeo1"

	// db.Debug().Save(&user)

	// 23. update updates
	var user User
	// db.Model(&user).Update("age", 50)
	// db.Model(&user).Where("name = ?", "eeo").Update("age", 60)
	// db.Model(&user).Where("name = ?", "").Update(map[string]interface{}{"age": 100})
	// db.Debug().Model(&user).Where("name = ?", "").Update(User{Name: "alex"})

	// 24. omit exclude
	// db.Debug().Model(&user).Where("name = ?", "panda").Omit("name").Update(map[string]interface{}{"name": "hello", "age": 19})

	// 25. 无钩子更新, 没有更新时间。
	// db.Debug().Model(&user).Where("name = ?", "panda").UpdateColumn("age", 96)

	// 26. 批量更新
	// db.Table("users").Where("name = ?", "panda").Updates(User{Age: 95, Name: "panda1"})
	// v := db.Table("users").Where("name = ?", "panda1").Updates(User{Age: 94, Name: "panda"}).RowsAffected
	// fmt.Println(v)

	// 27. 表达式更新
	// db.First(&user)
	// db.Debug().Model(&user).Update("age", gorm.Expr("age * ?", 2))

	// 28. 软删除和查询被删除的
	// db.Where("name = ?", "panda").Delete(&user)
	// db.Debug().Unscoped().Where("name = ?", "panda").Find(&user)

	// 29. 物理删除
	// db.Debug().Unscoped().Where("name = ?", "panda").Delete(&user)

	rows, _ := db.Table("users").Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Rows()

	fmt.Println(user)
	fmt.Println(users)
	for rows.Next() {

	}
}

func ageLT30(db *gorm.DB) *gorm.DB {
	return db.Where("age < ?", 30)
}
