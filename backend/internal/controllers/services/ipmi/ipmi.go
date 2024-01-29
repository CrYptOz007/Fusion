package ipmi

import (
	"strconv"

	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/service"
	"github.com/CrYptOz007/Fusion/internal/server/utils"
	"github.com/CrYptOz007/Fusion/pkg/ipmitool"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetInfo(c echo.Context) error {
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

	client, err := ipmitool.NewClient(service.Hostname, uint16(service.Port), service.Username, service.Password)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	info := client.GetInfo()

	return c.JSON(200, info)
}
