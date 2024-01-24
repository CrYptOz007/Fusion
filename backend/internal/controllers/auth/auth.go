package auth

import (
	"net/http"

	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/user"
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

	if user.Username == "" || user.Password == "" {
		return helpers.ReturnUnexpectedError(c, []string{"username and password are required"})
	}

	if err := database.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		return helpers.ReturnUnexpectedError(c, []string{"username does not exist"})
	}

	if err := helpers.ComparePassword(dbUser.Password, user.Password); err == false {
		return c.JSON(401, map[string]string{"error": "password is incorrect"})
	}

	authToken, refreshToken, err := jwt.GenerateTokenPair(dbUser.ID, dbUser.Username)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	c.SetCookie(&http.Cookie{Name: "refreshToken", Value: refreshToken, HttpOnly: true, Path: "/"})
	return c.JSON(200, map[string]interface{}{"token": authToken})
}

func Refresh(c echo.Context) error {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		return c.JSON(401, map[string]string{"error": "refresh token is missing"})
	}

	claims, err := jwt.ValidateRefreshToken(refreshToken.Value)
	if err != nil {
		return c.JSON(401, map[string]string{"error": "refresh token is invalid"})
	}

	username := claims["username"].(string)
	id := uint(claims["id"].(float64))

	authToken, err := jwt.GenerateAuthToken(id, username)
	if err != nil {
		return c.JSON(401, map[string]string{"error": "refresh token is invalid"})
	}

	return c.JSON(200, map[string]interface{}{"token": authToken})
}
