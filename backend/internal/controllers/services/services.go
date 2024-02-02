package services

import (
	"net/http"

	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/service"
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/CrYptOz007/Fusion/internal/server/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dtoToService(s *service.ServiceDTO) *service.Service {
	return &service.Service{
		Name:        s.Name,
		Type:        s.Type,
		Description: s.Description,
		Hostname:    s.Hostname,
		Port:        s.Port,
		ApiKey:      s.ApiKey,
		Username:    s.Username,
		Password:    s.Password,
		UserID:      s.UserID,
		User:        s.User,
		Icon:        s.Icon,
	}
}

func CreateService(c echo.Context) error {
	var e []string

	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	s := new(service.ServiceDTO)

	if err := c.Bind(s); err != nil {
		return helpers.ReturnExpectedError(c, http.StatusBadRequest, []string{err.Error()})
	}

	err := s.BeforeCreate(database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c)
	}

	serviceModel := dtoToService(s)

	if err := database.Create(serviceModel).Error; err != nil {
		return helpers.ReturnExpectedError(c, http.StatusConflict, []string{"unable to create service"})
	}

	return c.JSON(http.StatusCreated, types.Response{Error: e})
}

func GetServices(c echo.Context) error {

	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	var services []service.Service

	if err := database.Find(&services).Error; err != nil {
		return helpers.ReturnUnexpectedError(c)
	}

	return c.JSON(http.StatusOK, types.Response{Data: services})
}
