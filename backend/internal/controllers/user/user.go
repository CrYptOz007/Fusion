package user

import (
	"net/http"
	"strings"

	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/user"
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/CrYptOz007/Fusion/internal/server/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(c echo.Context) error {
	var e []string
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	u := new(user.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, types.Response{Error: []string{err.Error()}})
	}

	u.Username = strings.TrimSpace(u.Username)
	u.Password = strings.TrimSpace(u.Password)


	if u.Username == "" || u.Password == "" {
		return c.JSON(http.StatusBadRequest, types.Response{Error: []string{"username and password are required"}})
	}

	if database.Where("username = ?", u.Username).First(&user.User{}).Error == nil {
		return c.JSON(http.StatusConflict, types.Response{Error: []string{"username already exists"}})
	}

	hashedPassword, err := helpers.HashPassword(u.Password)
	if err != nil {
		return helpers.ReturnUnexpectedError(c)
	}
	u.Password = hashedPassword

	u.Salt, err = helpers.GenerateSalt()
	if err != nil {
		return helpers.ReturnUnexpectedError(c)
	}

	database.Create(u)

	return c.JSON(http.StatusCreated, types.Response{Error: e})
}
