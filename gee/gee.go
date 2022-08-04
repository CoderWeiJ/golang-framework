package gee

import (
	"fmt"
	"net/http"
)

// HandleFunc 定义了路由处理函数类型
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine 实现了ServerHTTP接口
type Engine struct {
	// 路由映射表，key位请求方法+静态路由地址，比如：GET-/hello
	router map[string]HandleFunc
}

// New 创建一个gee.Engine的实例
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// addRoute 内部添加路由到router
func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 将路由和处理方法注册到路由映射表router当中
func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加POST请求
func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run gee开启一个http服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 解析请求的路径，查找路由映射表，有就执行注册的处理方法，没有就返回404
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
