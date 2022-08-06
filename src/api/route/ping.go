package route

import (
	"scf-web-func-go-demo/api/handler"
)

func getPingRoutes() *NodeRoute {
	routers := []*Route{
		NewRoute("GET", "/ping", handler.Ping),
	}

	return NewNodeRoute("/", routers...)
}
