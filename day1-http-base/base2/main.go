package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(res, "hello\n")
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(res, "Header[%q]: %q\n", k, v)
		}
	default:
		fmt.Fprintf(res, "404 page not found\n")
	}
}

func main() {

	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
