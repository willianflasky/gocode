package gee

import (
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee.  HandlerFunc -> Context -> (request, writer, method...)
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP. Engine -> router -> handlers
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

// RouterGroup 路由组
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
}

// New is the constructor of gee. Engine 初始化函数
func New() *Engine {
	engine := &Engine{router: newRouter()}             // new engine
	engine.RouterGroup = &RouterGroup{engine: engine}  // 这里存上new routergroup, 再把engine存到new routergroup
	engine.groups = []*RouterGroup{engine.RouterGroup} // 将routergroup存到routergroup切片;
	return engine
}

// Group 绑定group方法
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request.  在使用GET方法时，就将路由映射就增加到 router.handlers
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request  在使用POST方法时，就将路由映射就增加到 router.handlers
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
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
