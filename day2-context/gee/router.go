package gee

import "net/http"

type Router struct {
	m map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{m: make(map[string]HandlerFunc)}
}

func (rou *Router) addRoute(address string, handler HandlerFunc) {
	rou.m[address] = handler
}

func (rou *Router) handle(c *Context) {
	address := c.Method + "-" + c.Path
	if handler, ok := rou.m[address]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
