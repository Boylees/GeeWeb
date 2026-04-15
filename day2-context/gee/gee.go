package gee

import "net/http"

type HandlerFunc func(c *Context)
type Engine struct {
	rou *Router
}

func New() *Engine {
	return &Engine{rou: newRouter()}
}

func (engine *Engine) GET(address string, handler HandlerFunc) {
	engine.rou.addRoute("GET-"+address, handler)
}

func (engine *Engine) POST(address string, handler HandlerFunc) {
	engine.rou.addRoute("POST-"+address, handler)
}

func (engine *Engine) Run(address string) {
	http.ListenAndServe(address, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := NewContext(writer, request)
	engine.rou.handle(c)
}
