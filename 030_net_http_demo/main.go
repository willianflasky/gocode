package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "<h1>Hello Golang</h1>")
}

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
