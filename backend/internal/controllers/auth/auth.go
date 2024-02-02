package auth

import (
	"net/http"
	"strings"

	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/user"
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/CrYptOz007/Fusion/internal/server/utils"
	jwt "github.com/CrYptOz007/Fusion/internal/utils"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	var user, dbUser user.User

	if err := c.Bind(&user); err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)

	if user.Username == "" || user.Password == "" {
		return helpers.ReturnUnexpectedError(c, []string{"username and password are required"})
	}

	if err := database.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		return helpers.ReturnUnexpectedError(c, []string{"username does not exist"})
	}

	if !helpers.ComparePassword(dbUser.Password, user.Password) {
    return c.JSON(http.StatusUnauthorized, types.Response{Error: []string{"invalid password"}})
	}

	authToken, refreshToken, err := jwt.GenerateTokenPair(dbUser.ID, dbUser.Username)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	c.SetCookie(&http.Cookie{Name: "refreshToken", Value: refreshToken, HttpOnly: true, Path: "/"})
	return c.JSON(http.StatusOK, map[string]string{"token": authToken})
}

func Refresh(c echo.Context) error {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, types.Response{Error: []string{"refresh token is missing"}})
	}

	claims, err := jwt.ValidateRefreshToken(refreshToken.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, types.Response{Error: []string{"refresh token is invalid"}})
	}

	username := claims["username"].(string)
	id := uint(claims["id"].(float64))

	authToken, err := jwt.GenerateAuthToken(id, username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, types.Response{Error: []string{"refresh token is invalid"}})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"token": authToken})
}

func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "refreshToken", Value: "", HttpOnly: true, Path: "/", MaxAge: -1})
	return c.JSON(200, map[string]string{"message": "logged out"})
}
