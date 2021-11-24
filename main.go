package main

import (
	"fmt"
	"github.com/An5dy/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	_ = r.Run(":9000")
}
