package utils

import "github.com/labstack/echo/v4"

func SetLocal[T any](c echo.Context, key string, value T) {
	c.Set(key, value)
}

func GetLocal[T any](c echo.Context, key string) T {
	return c.Get(key).(T)
}
