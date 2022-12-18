package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"strings"
)

type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func SkipperFn(c echo.Context) bool {
	path := c.Request().URL.Path
	return path == "/v1/health" || strings.Contains(path, "swagger")
}

func GetUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims)
	return claims.UserID
}
