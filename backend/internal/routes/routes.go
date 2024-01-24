package routes

import (
	"github.com/CrYptOz007/Fusion/internal/controllers/auth"
	"github.com/CrYptOz007/Fusion/internal/controllers/services"
	"github.com/CrYptOz007/Fusion/internal/controllers/services/proxmox"
	"github.com/CrYptOz007/Fusion/internal/controllers/user"
	"github.com/CrYptOz007/Fusion/internal/middlewares"
	"github.com/CrYptOz007/Fusion/internal/server/types"

	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo, groups types.Groups) {
	r := NewRouter(app, groups)

	r.Group("services", middlewares.Protected, func(r *Router) {
		r.GET("", services.GetServices)
		r.POST("/create", services.CreateService)

		r.GET("/proxmox/nodes", proxmox.GetNodes)
		r.GET("/proxmox/virtualmachines", proxmox.GetVirtualMachines)
		r.GET("/proxmox/containers", proxmox.GetContainers)
	})

	r.Group("auth", nil, func(r *Router) {
		r.POST("/login", auth.Login)
		r.GET("/refresh", auth.Refresh)
		r.GET("/logout", auth.Logout)
	})

	r.Group("user", nil, func(r *Router) {
		r.POST("/register", user.Register)
	})
}
