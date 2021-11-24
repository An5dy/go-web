package gee

import (
	"net/http"
)

// 定义框架的请求处理函数
type HandlerFunc func(*Context)

// Engine 处理请求引擎
type Engine struct {
	router *router // 路由
}

var _ http.Handler = (*Engine)(nil)

// Engine 构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 注册路由
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// 注册 GET 请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 注册 POST 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 启动 http 服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现 http.Handler ServeHTTP 方法
// 解析请求的路径，如果在路由映射表中存在则执行注册的 handler 方法，反之返回 404 NOT FOUND
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
