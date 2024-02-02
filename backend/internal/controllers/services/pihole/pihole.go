package pihole

import (
	"strconv"

	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/service"
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/CrYptOz007/Fusion/internal/server/utils"
	"github.com/CrYptOz007/Fusion/internal/services/pihole"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetSummary(c echo.Context) error {
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	id := c.QueryParam("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get service from database
	service, err := service.FetchService(parsedId, database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get summary from pihole
	summary, err := pihole.Summary(service)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	return c.JSON(200, types.Response{Data: summary})
}

func GetStatus(c echo.Context) error {
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	id := c.QueryParam("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get service from database
	service, err := service.FetchService(parsedId, database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get status from pihole
	status, err := pihole.Status(service)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	return c.JSON(200, types.Response{Data: status})
}

func Disable(c echo.Context) error {
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	id := c.QueryParam("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get service from database
	service, err := service.FetchService(parsedId, database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Disable pihole
	err = pihole.Disable(service)

	return c.JSON(200, types.Response{Data: err})
}

func Enable(c echo.Context) error {
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	id := c.QueryParam("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get service from database
	service, err := service.FetchService(parsedId, database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Enable pihole
	err = pihole.Enable(service)

	return c.JSON(200, types.Response{Data: err})
}
