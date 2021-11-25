package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// 开始事件
		t := time.Now()
		// 处理下一步请求
		c.Next()
		// 计算响应时间
		log.Printf("[%d] %s in %v", c.StatusCode, c.R.RequestURI, time.Since(t))
	}
}
