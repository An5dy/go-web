package gee

import (
	"log"
	"net/http"
	"strings"
)

// 定义框架的请求处理函数
type HandlerFunc func(*Context)

// 路由组
type RouterGroup struct {
	prefix      string        // 路由前缀
	middlewares []HandlerFunc // 中间件组
	parent      *RouterGroup  // 父级路由组
	engine      *Engine       // Engine 处理请求引擎
}

// Engine 处理请求引擎
type Engine struct {
	*RouterGroup
	router *router        // 路由
	groups []*RouterGroup // 路由组集合
}

var _ http.Handler = (*Engine)(nil)

// Engine 构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 路由分组
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

// 注册路由
func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// 注册 GET 请求
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// 注册 POST 请求
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// 启动 http 服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 使用中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 实现 http.Handler ServeHTTP 方法
// 解析请求的路径，如果在路由映射表中存在则执行注册的 handler 方法，反之返回 404 NOT FOUND
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, r)
	c.handlers = middlewares
	engine.router.handle(c)
}
