package gee

import (
	"fmt"
	"net/http"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	mapping map[string]handlerFunc
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.mapping[key]; ok {
		handler(writer, request)
	} else {
		fmt.Fprint(writer, "404 Not Found\n")
	}
}

func New() *Engine {
	return &Engine{make(map[string]handlerFunc)}
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) addRoute(method string, pattern string, handler handlerFunc) {
	key := method + "-" + pattern
	engine.mapping[key] = handler
}

func (engine *Engine) GET(address string, handler handlerFunc) {
	engine.addRoute("GET", address, handler)
}

func (engine *Engine) POST(address string, handler handlerFunc) {
	engine.addRoute("POST", address, handler)
}
