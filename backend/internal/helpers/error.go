package helpers

import (
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/labstack/echo/v4"
)

func ReturnUnexpectedError(c echo.Context) error {
	e := []string{"an unexpected error occurred"}
	return ReturnExpectedError(c, 500, e)
}

func ReturnExpectedError(c echo.Context, status int, e []string) error {
	return c.JSON(status, types.Response{
		Error: e,
	})
}