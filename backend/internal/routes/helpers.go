package routes

import (
	"fmt"
	"net/http"

	"github.com/CrYptOz007/Fusion/internal/server/types"

	"github.com/labstack/echo/v4"
)

type Router struct {
	app    *echo.Echo
	route  string
	group  *echo.Group
	groups types.Groups
}

func NewRouter(app *echo.Echo, groups types.Groups) *Router {
	if groups == nil {
		groups = make(types.Groups)
	}

	return &Router{
		app:    app,
		groups: groups,
	}
}

func (router *Router) Group(name string, middleware []echo.MiddlewareFunc, groupFunc func(router *Router)) {
	newRouterRoute := router.route
	var newGroup *echo.Group
	if name != "" {
		newRouterRoute += name
	}

	if router.group != nil {
		newGroup = router.group.Group(name, middleware...)
	} else {
		newGroup = router.app.Group(name, middleware...)
	}

	if _, ok := router.groups[newRouterRoute]; ok {
		i := 1

		for _, ok := router.groups[newRouterRoute+fmt.Sprintf("%v", i)]; ok; {
			router.groups[newRouterRoute+fmt.Sprintf("%v", i)] = newGroup
		}
	}

	newRouter := &Router{
		app:   router.app,
		route: newRouterRoute,
		group: newGroup,
	}

	groupFunc(newRouter)
}

func (router *Router) Handle(method string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	if router.group != nil {
		router.group.Add(method, path, handler, middleware...)
	} else {
		router.app.Add(method, path, handler, middleware...)
	}
}

func (router *Router) GET(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	router.Handle(http.MethodGet, path, handler, middleware...)
}

func (router *Router) POST(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	router.Handle(http.MethodPost, path, handler, middleware...)
}

func (router *Router) PATCH(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	router.Handle(http.MethodPatch, path, handler, middleware...)
}

func (router *Router) PUT(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	router.Handle(http.MethodPut, path, handler, middleware...)
}

func (router *Router) DELETE(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	router.Handle(http.MethodDelete, path, handler, middleware...)
}
