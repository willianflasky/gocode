package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dns := "root:root@tcp(localhost:3306)/db1?parseTime=true"
	db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(16)
	return nil
}

func createUser(username, password string) error {
	sql := "insert into userinfo(username, password) values(?, ?)"
	_, err := db.Exec(sql, username, password)
	if err != nil {
		fmt.Println("insert err: ", err)
	}
	return nil
}

func queryUser(username, password string) (bool, error) {
	sql := "select id from userinfo where username=? and password=? limit 1"
	var id int64
	err := db.Get(&id, sql, username, password)
	if err != nil {
		return false, err
	}
	if id > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
