package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc 定义了 gee 使用的请求处理
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine实现ServeHTTP的接口
type Engine struct {
	router map[string]HandlerFunc
}

// new 是 gee.Engine 的构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
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
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s", req.URL)
	}
}