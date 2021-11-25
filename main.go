package main

import (
	"github.com/An5dy/gee"
	"net/http"
)

func main() {
	r := gee.Default()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello World\n")
	})
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"andy"}
		c.String(http.StatusOK, names[100])
	})


	_ = r.Run(":9000")
}
