package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init mysql
func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		"root",
		"",
		"127.0.0.1",
		3306,
		"db1",
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect failed.")
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	fmt.Println("connect success.")
	return
}

func main() {
	if err := Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	type master struct {
		File            string `db:"File"`
		Postion         string `db:"Position"`
		BinlogDoDB      string `db:"Binlog_Do_DB"`
		BinlogIgnoreDB  string `db:"Binlog_Ignore_DB"`
		ExecutedGtidSet string `db:"Executed_Gtid_Set"`
	}

	var row master
	err := db.Get(&row, "show master status")
	fmt.Println(row, err)

	defer db.Close()

	// filename := "mysql-bin.000001"
	// s := strings.Split(filename, ".")
	// fmt.Printf("%T\n", s[1])
	// v, err := strconv.Atoi(s[1])
	// fmt.Println(v, err)
}
