package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc  函数类型, 存放函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine struct
type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	// key:  GET-/hello
	// value: 函数
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// New method
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// Run xx
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
