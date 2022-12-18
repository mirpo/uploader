package handler

import (
	"github.com/otiai10/gosseract/v2"
	"net/http"
	"strings"
	"uploader/pkg/auth"
	"uploader/pkg/fs"
	"uploader/pkg/helper"

	"github.com/labstack/echo/v4"
)

type OCRResponse struct {
	Total string   `json:"total,omitempty"`
	Raw   []string `json:"raw"`
}

func Ocr(c echo.Context) error {
	path := c.Param("path")
	userId := auth.GetUserID(c)
	c.Logger().Debugf(`userId: "%s" trying to ocr: "%s"`, userId, path)

	// check that original file exists
	originalFilePath := helper.GetUserFolder(userId) + path
	if !fs.FileExists(originalFilePath) {
		c.Logger().Errorf("file doesn't exist: %s", originalFilePath)
		return echo.ErrNotFound
	}

	// try to OCR image
	client := gosseract.NewClient()
	defer client.Close()

	if err := client.SetImage(originalFilePath); err != nil {
		return echo.ErrInternalServerError
	}

	text, _ := client.Text()
	lines := strings.Split(text, "\n")

	var total string
	for _, line := range lines {
		lineToLower := strings.ToLower(line)
		// try to find total :)
		if strings.HasPrefix(lineToLower, "total ") || strings.HasPrefix(lineToLower, "subtotal ") {
			total = strings.Split(line, " ")[1]
		}
	}

	return c.JSON(http.StatusOK, OCRResponse{
		Total: total,
		Raw:   lines,
	})
}
