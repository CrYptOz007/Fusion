package middlewares

import (
	"net/http"
	"strings"

	"github.com/CrYptOz007/Fusion/internal/server/types"
	jwt "github.com/CrYptOz007/Fusion/internal/utils"
	"github.com/labstack/echo/v4"
)

var Protected = []echo.MiddlewareFunc{func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return authMiddleware(c, next)
	}
}}

func authMiddleware(c echo.Context, next echo.HandlerFunc) error {
	authToken := c.Request().Header.Get("Authorization")

	token := strings.TrimSpace(strings.Replace(authToken, "Bearer", "", 1))
	if token == "" {
		return c.JSON(http.StatusUnauthorized, types.Response{
			Error: []string{"no token provided"},
		})
	}

	_, err := jwt.ValidateAuthToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, types.Response{
			Error: []string{"invalid token"},
		})
	}
	return next(c)

}
