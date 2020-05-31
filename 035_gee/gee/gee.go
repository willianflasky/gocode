package gee

import "net/http"

// HandlerFunc defines the request handler used by gee.  HandlerFunc -> Context -> (request, writer, method...)
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP. Engine -> router -> handlers
type Engine struct {
	router *router
}

// New is the constructor of gee. Engine 初始化函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute  调用router.addRoute
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request.  在使用GET方法时，就将路由映射就增加到 router.handlers
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request  在使用POST方法时，就将路由映射就增加到 router.handlers
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server 运行
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP  1. 被谁调用？
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
