package gee

import (
	"log"
	"net/http"
)

// router type
type router struct {
	// GET-/hello: hello
	handlers map[string]HandlerFunc
}

// newRouter router的初始化函数
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRoute 增加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	// key = GET-/hello
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle 调度函数， 被ServeHTTP调用。
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
