package gee

import (
	"log"
	"net/http"
)

// 定义 router
type router struct {
	handlers map[string]HandlerFunc // 路由 -> 请求处理映射表
}

// router 构造函数
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 注册路由
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 处理请求
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}