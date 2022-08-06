package gee

import (
	"net/http"
)

/**
@date 2022-08-05
@author CoderWeiJ
@description 框架入口
*/

// HandlerFunc 定义了路由处理函数类型
type HandlerFunc func(*Context)

// Engine 实现了ServerHTTP接口
type Engine struct {
	// 路由映射表，key位请求方法+静态路由地址，比如：GET-/hello
	router *router
}

// New 创建一个gee.Engine的实例
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 内部添加路由到router
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET 将路由和处理方法注册到路由映射表router当中
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加POST请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run gee开启一个http服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 解析请求的路径，查找路由映射表，有就执行注册的处理方法，没有就返回404
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
