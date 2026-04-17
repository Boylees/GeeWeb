package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)
type Engine struct {
	rou *Router
}

func New() *Engine {
	return &Engine{rou: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.rou.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(address string, handler HandlerFunc) {
	engine.addRoute("GET", address, handler)
}

func (engine *Engine) POST(address string, handler HandlerFunc) {
	engine.addRoute("POST", address, handler)
}

func (engine *Engine) Run(address string) {
	http.ListenAndServe(address, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := NewContext(writer, request)
	engine.rou.handle(c)
}
