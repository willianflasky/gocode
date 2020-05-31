package main

import (
	"fmt"
	"gee/gee"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func hello(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func main() {
	r := gee.New()
	r.GET("/", root)
	r.GET("/hello", hello)
	r.Run(":9999")
}
