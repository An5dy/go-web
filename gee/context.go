package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 方便构建 JSON 类型数据
type H map[string]interface{}

// Context 上下文
type Context struct {
	W          http.ResponseWriter // 响应体
	R          *http.Request       // 请求体
	Path       string              // 请求路径
	Method     string              // 请求方法
	StatusCode int                 // http 响应码
}

// Context 构造函数
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// 获取 POST 请求数据
func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

// 获取 GET 请求数据
func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

// 设置响应状态码
func (c *Context) Status(code int) {
	c.W.WriteHeader(code)
}

// 设置响应 Header 头
func (c *Context) SetHeader(key, value string) {
	c.W.Header().Set(key, value)
}

// 输出字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

// 输出 json
func (c *Context) JSON(code int, data interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}

// 输出 data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

// 输出 html
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.W.Write([]byte(html))
}