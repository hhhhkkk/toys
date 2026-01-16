package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Named interface {
	Name() string
}

type IRouter interface {
	Named
	Method() string
	Path() string
	Handler() gin.HandlerFunc
}

type IRouterGroup interface {
	Named
	Routers() []IRouter
	Middlewares() []gin.HandlerFunc
	AddRouter(r ...IRouter)
	AddMiddleware(m ...gin.HandlerFunc)
}

type router struct {
	method string
	path   string
	name   string
	f      gin.HandlerFunc
}

func (r *router) Method() string {
	return r.method
}

func (r *router) Path() string {
	return r.path
}

func (r *router) Name() string {
	if r.name == "" {
		return r.path
	}
	return r.name
}

func (r *router) Handler() gin.HandlerFunc {
	return r.f
}

func NewRouter(m, p, n string, f gin.HandlerFunc) IRouter {
	return &router{
		method: m,
		path:   p,
		name:   n,
		f:      f,
	}
}

type routerGroup struct {
	name        string
	routers     []IRouter
	middlewares []gin.HandlerFunc
}

func NewRouterGroup(name string, routers ...IRouter) IRouterGroup {
	return &routerGroup{
		name:        name,
		routers:     routers,
		middlewares: make([]gin.HandlerFunc, 0),
	}
}

func (rg *routerGroup) Name() string {
	return rg.name
}

func (rg *routerGroup) AddRouter(r ...IRouter) {
	rg.routers = append(rg.routers, r...)
}

func (rg *routerGroup) AddMiddleware(m ...gin.HandlerFunc) {
	rg.middlewares = append(rg.middlewares, m...)
}

func (rg *routerGroup) Routers() []IRouter {
	return rg.routers
}

func (rg *routerGroup) Middlewares() []gin.HandlerFunc {
	return rg.middlewares
}

func NewRouterWithoutName(m, p string, f gin.HandlerFunc) IRouter {
	return NewRouter(m, p, "", f)
}

func NewGetRouter(p string, f gin.HandlerFunc) IRouter {
	return NewRouterWithoutName(http.MethodGet, p, f)
}
func NewPostRouter(p string, f gin.HandlerFunc) IRouter {
	return NewRouterWithoutName(http.MethodPost, p, f)
}
func NewPutRouter(p string, f gin.HandlerFunc) IRouter {
	return NewRouterWithoutName(http.MethodPut, p, f)
}
func NewDeletetRouter(p string, f gin.HandlerFunc) IRouter {
	return NewRouterWithoutName(http.MethodDelete, p, f)
}
