package routes

import (
	"github.com/CrYptOz007/Fusion/internal/controllers/services"
	"github.com/CrYptOz007/Fusion/internal/server/types"

	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo, groups types.Groups) {
	r := NewRouter(app, groups)

	r.Group("services", nil, func(r *Router) {
		r.GET("", services.GetServices)
		r.POST("/create", services.CreateService)

	})
}
