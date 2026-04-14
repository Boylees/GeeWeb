package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	})
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	r.POST("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world\n")
	})
	r.Run(":9999")
}
