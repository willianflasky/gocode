package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// UserInfo 结构体
type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./*.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	err = tmpl.ExecuteTemplate(w, "ul.tmpl", user)
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", tmplDemo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
