package types

import "github.com/labstack/echo/v4"

type Response struct {
	Error []string `json:"error"`
	Data  Data     `json:"data"`
}

type Groups map[string]*echo.Group

type Data interface{}
