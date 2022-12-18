package handler

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"uploader/pkg/auth"
	"uploader/pkg/helper"

	"github.com/labstack/echo/v4"
)

type ListFileResponse struct {
	Files  []FileResponse `json:"files"`
	Widths []string       `json:"widths"`
}

// List godoc
// @Summary      List uploaded receipts
// @Produce      json
// @Success      200  {object}  ListFileResponse
// @Failure      500  {object}  string
// @Router       /receipts [get]
func List(c echo.Context) error {
	rootUserPath := helper.GetUserFolder(auth.GetUserID(c))
	userGlobPattern := fmt.Sprintf("%s/*/*/*/*.*", rootUserPath)
	c.Logger().Debugf("glob files using path: %s", userGlobPattern)

	// glob files
	files, err := filepath.Glob(userGlobPattern)
	if err != nil {
		return echo.ErrInternalServerError
	}

	c.Logger().Debugf("found files: %d\n", len(files))

	// create response
	allowedWidths := helper.GetAllowedWidth()
	list := ListFileResponse{
		Widths: allowedWidths,
	}

	for _, f := range files {
		resp := FileResponse{
			File: path.Base(f),
			Path: strings.TrimLeft(f, rootUserPath),
		}
		list.Files = append(list.Files, resp)
	}

	return c.JSON(http.StatusOK, list)
}
