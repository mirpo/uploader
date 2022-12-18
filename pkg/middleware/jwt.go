package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"uploader/pkg/auth"
)

func JWTConfig() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    auth.SkipperFn,
		Claims:     &auth.Claims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	})
}
