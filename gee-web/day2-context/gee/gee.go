package gee

import (
	"net/http"
)

// HandlerFunc 定义了 gee 使用的请求处理
type HandlerFunc func(ctx *Context)

// Engine实现ServeHTTP的接口
type Engine struct {
	router *router
}

// new 是 gee.Engine 的构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)

}

//GET 定义了 添加GET 请求方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 定义了添加post请求方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 定义了启动 http 服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w,req)
	engine.router.handle(c)


}


