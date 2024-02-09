package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	accessTokenMaxAge  = 15 * time.Minute
	refreshTokenMaxAge = 2 * 24 * time.Hour
)

var authKey = os.Getenv("AUTH_KEY")
var refreshKey = os.Getenv("REFRESH_KEY")

type Token struct {
	UserId   uint
	Username string
	Expiry   time.Time
	Key      string
}

func (t *Token) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       t.UserId,
		"username": t.Username,
		"exp":      t.Expiry.Unix(),
	})

	tokenString, err := token.SignedString([]byte(t.Key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAuthToken(userId uint, username string) (string, error) {
	token := &Token{
		UserId:   userId,
		Username: username,
		Expiry:   time.Now().Add(accessTokenMaxAge),
		Key:      authKey,
	}

	return token.Generate()
}

func GenerateTokenPair(userId uint, username string) (string, string, error) {
	authToken := &Token{
		UserId:   userId,
		Username: username,
		Expiry:   time.Now().Add(accessTokenMaxAge),
		Key:      authKey,
	}

	authTokenString, err := authToken.Generate()
	if err != nil {
		return "", "", err
	}

	refreshToken := &Token{
		UserId:   userId,
		Username: username,
		Expiry:   time.Now().Add(refreshTokenMaxAge),
		Key:      refreshKey,
	}

	refreshTokenString, err := refreshToken.Generate()
	if err != nil {
		return "", "", err
	}

	return authTokenString, refreshTokenString, nil
}

func ValidateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	return validateToken(tokenString, []byte(refreshKey))
}

func ValidateAuthToken(tokenString string) (jwt.MapClaims, error) {
	return validateToken(tokenString, []byte(authKey))
}

func validateToken(tokenString string, key []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return key, nil
	})

	if err != nil {
		return nil, errors.New("could not parse token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	username := claims["username"].(string)
	id := claims["id"].(float64)

	if username == "" || id < 0 {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func GetUserIdFromToken(tokenString string) (uint, error) {
	claims, err := ValidateAuthToken(tokenString)
	if err != nil {
		return 0, err
	}

	id := claims["id"].(float64)

	return uint(id), nil
}
