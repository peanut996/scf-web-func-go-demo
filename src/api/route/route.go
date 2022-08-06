package route

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

type NodeRoute struct {
	path   string
	routes []*Route
}

func (n *NodeRoute) registerRoutes(r *gin.RouterGroup) {
	group := r.Group(n.path)
	for _, route := range n.routes {
		methodMapper(group, route.method)(route.path, route.handler)
	}
}
func NewRoute(method, path string, handler gin.HandlerFunc) *Route {
	return &Route{
		method:  method,
		path:    path,
		handler: handler,
	}
}

func NewNodeRoute(path string, routers ...*Route) *NodeRoute {
	return &NodeRoute{
		path:   path,
		routes: routers,
	}
}
func methodMapper(group *gin.RouterGroup, method string) func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	switch method {
	case "GET":
		return group.GET
	case "POST":
		return group.POST
	case "PUT":
		return group.PUT
	case "DELETE":
		return group.DELETE
	case "PATCH":
		return group.PATCH
	case "HEAD":
		return group.HEAD
	default:
		return group.Any
	}
}

func mountRoutes(r *gin.Engine, routers ...*NodeRoute) {
	router := r.Group("/")
	for _, node := range routers {
		node.registerRoutes(router)
	}
}

func InitAllRouters(r *gin.Engine) {
	nodes := []*NodeRoute{
		getPingRoutes(),
	}
	mountRoutes(r, nodes...)
}
