package services

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/service"
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/CrYptOz007/Fusion/internal/server/utils"
	jwt "github.com/CrYptOz007/Fusion/internal/utils"
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

	authHeader := c.Request().Header.Get("Authorization")

	tokenString := strings.Split(authHeader, "Bearer ")[1]

	userId, err := jwt.GetUserIdFromToken(tokenString)
	if err != nil {
		return helpers.ReturnExpectedError(c, http.StatusUnauthorized, []string{err.Error()})
	}

	s.UserID = int(userId)

	if err := c.Bind(s); err != nil {
		return helpers.ReturnExpectedError(c, http.StatusBadRequest, []string{err.Error()})
	}

	err = s.BeforeCreate(database)
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

func PingService(c echo.Context) error {
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	id := c.QueryParam("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.ReturnExpectedError(c, http.StatusBadRequest, []string{"invalid query type for: id"})
	}

	// Get service from database
	service, err := service.FetchService(parsedId, database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c)
	}

	client := &http.Client{
			Timeout: 1000 * time.Millisecond,
	}

	httpURL := "http://"+service.Hostname+":"+strconv.Itoa(service.Port)
	httpsURL := "https://"+service.Hostname+":"+strconv.Itoa(service.Port)

	httpErrChan := make(chan error, 1)
	httpsErrChan := make(chan error, 1)

	go func() {
		_, err := client.Get(httpURL)
		httpErrChan <- err
	}()

	go func() {
		_, err := client.Get(httpsURL)
		httpsErrChan <- err
	}()

	var httpErr, httpsErr error
	 httpErr = <-httpErrChan
	 httpsErr = <-httpsErrChan


	if httpErr != nil && httpsErr != nil {
    urlErr, ok := httpsErr.(*url.Error)
    if ok {
        if _, ok := urlErr.Err.(*tls.CertificateVerificationError); !ok {
            return helpers.ReturnExpectedError(c, 408, []string{"unable to reach service"})
        }
    }
	}

	return c.JSON(http.StatusOK, types.Response{Data: "pong"})
}