package handler

import (
	"io"
	"net/http"
	"os"
	"time"
	"uploader/pkg/auth"
	"uploader/pkg/fs"
	"uploader/pkg/helper"

	"github.com/labstack/echo/v4"
)

type FileResponse struct {
	File   string   `json:"file"`
	Path   string   `json:"path"`
	Widths []string `json:"widths,omitempty"`
}

func CreateFileResponse(filename string, currentDayPath string) FileResponse {
	return FileResponse{
		File:   filename,
		Path:   currentDayPath + filename,
		Widths: helper.GetAllowedWidth(),
	}
}

func Upload(c echo.Context) error {
	userID := auth.GetUserID(c)

	// check user's folder exist, if not => try create
	currentDayPath := time.Now().UTC().Format("2006/01/02/")
	userFolder := helper.GetUserFolder(userID) + currentDayPath
	if err := fs.CheckFolder(userFolder); err != nil {
		if err := fs.CreateFolder(userFolder); err != nil {
			c.Logger().Errorf("failed to create folder: %s", userFolder)
			return echo.ErrInternalServerError
		}
	}

	// read file from the request
	file, err := c.FormFile("file")
	if err != nil {
		c.Logger().Error("failed to read file from the request")
		return echo.ErrBadRequest
	}

	src, err := file.Open()
	if err != nil {
		c.Logger().Error("failed to read file")
		return echo.ErrInternalServerError
	}
	defer src.Close()

	// detect and validate mime type
	if err = helper.DetectAndValidateMimeType(src); err != nil {
		c.Logger().Error("failed to detect and validate mime type")
		return echo.ErrBadRequest
	}

	// write image to the disk
	newFilename := helper.GetFilename(file)
	fullPath := userFolder + newFilename
	dst, err := os.Create(fullPath)
	if err != nil {
		c.Logger().Errorf("failed to create file: %s", fullPath)
		return echo.ErrInternalServerError
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		c.Logger().Error("failed to copy file content from the request to the disk")
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, CreateFileResponse(newFilename, currentDayPath))
}
