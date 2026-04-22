package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)
type Engine struct {
	*RouteGroup
	groups []*RouteGroup
	rou    *Router
}

type RouteGroup struct {
	Prefix      string
	middlewares []HandlerFunc
	engine      *Engine
	parent      *RouteGroup
}

func New() *Engine {
	engine := &Engine{
		rou: newRouter(),
	}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (group *RouteGroup) Group(prefix string) *RouteGroup {
	engine := group.engine
	newGroup := &RouteGroup{
		Prefix: group.Prefix + prefix,
		engine: engine,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouteGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	pattern = group.Prefix + pattern
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.rou.addRoute(method, pattern, handler)
}

func (group *RouteGroup) GET(address string, handler HandlerFunc) {
	group.addRoute("GET", address, handler)
}

func (group *RouteGroup) POST(address string, handler HandlerFunc) {
	group.addRoute("POST", address, handler)
}

func (engine *Engine) Run(address string) {
	http.ListenAndServe(address, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(request.URL.Path, group.Prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(writer, request)
	c.handlers = middlewares
	engine.rou.handle(c)
}

func (group *RouteGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
