package helpers

import (
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/labstack/echo/v4"
)

func ReturnUnexpectedError(c echo.Context, e []string) error {
	e = append(e, "an unexpected error occurred")
	return c.JSON(500, types.Response{
		Error: e,
	})
}
