package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BasicResponse struct {
	Msg string `json:"msg"`
}

// Health godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, BasicResponse{Msg: "All good!"})
}
