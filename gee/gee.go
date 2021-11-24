package gee

import (
	"fmt"
	"net/http"
)

// 定义框架的请求处理函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 处理请求引擎
type Engine struct {
	router map[string]HandlerFunc // 路由映射表
}

var _ http.Handler = (*Engine)(nil)

// Engine 构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 注册路由
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
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
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
