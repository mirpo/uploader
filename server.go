package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"os"
	_ "uploader/docs"
	"uploader/pkg/fs"
	"uploader/pkg/handler"
	uploaderMiddleware "uploader/pkg/middleware"
)

// @title uploader
// @version 1.0
// @description upload receipt REST API

// @contact.name mirpo
// @contact.url https://github.com/mirpo

// @host 127.0.0.1
// @BasePath /v1
func main() {
	// load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// check that required folders exist and are writable
	for _, folder := range []string{os.Getenv("UPLOAD_PATH"), os.Getenv("CACHE_PATH")} {
		if err := fs.CheckFolder(folder); err != nil {
			log.Fatalf("issues with folder: %s, %v", folder, err)
		}
	}

	// echo instance
	e := echo.New()
	// uncomment to see debug messages
	// e.Debug = true

	// middleware
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit(os.Getenv("BODY_LIMIT")))
	e.Use(uploaderMiddleware.LoggerConfig())
	e.Use(uploaderMiddleware.CorsConfig())
	e.Use(uploaderMiddleware.JWTConfig())

	// swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes in group v1
	g := e.Group("/v1")

	g.GET("/health", handler.Health)
	g.POST("/receipts", handler.Upload)
	g.GET("/receipts", handler.List)
	g.GET("/receipts/:path", handler.Download)
	g.GET("/receipts/ocr/:path", handler.Ocr)

	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))))
}
