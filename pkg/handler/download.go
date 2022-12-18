package handler

import (
	"github.com/disintegration/imaging"
	"golang.org/x/exp/slices"
	"strconv"
	"uploader/pkg/auth"
	"uploader/pkg/fs"
	"uploader/pkg/helper"

	"github.com/labstack/echo/v4"
)

func Download(c echo.Context) error {
	width := c.QueryParam("width")
	path := c.Param("path")
	userID := auth.GetUserID(c)
	c.Logger().Debugf(`userID: "%s" trying to download: "%s", width: "%s"`, userID, path, width)

	if width != "" && !slices.Contains(helper.GetAllowedWidth(), width) {
		c.Logger().Errorf("got unsupported width: %s", width)
		return echo.ErrBadRequest
	}

	// check that original file exists
	originalFilePath := helper.GetUserFolder(userID) + path
	if !fs.FileExists(originalFilePath) {
		c.Logger().Errorf("file doesn't exist: %s", originalFilePath)
		return echo.ErrNotFound
	}

	// original file
	if width == "" {
		c.Logger().Debugf("returning original file: %s", originalFilePath)
		return c.File(originalFilePath)
	}

	// if other size, check that cached file exist
	cacheFolder, cachedFilename := helper.GetCachePaths(path, helper.GetUserCacheFolder(userID), width)
	cacheFullPath := cacheFolder + cachedFilename
	if fs.FileExists(cacheFullPath) {
		c.Logger().Debugf("returning file from cache: %s", cacheFullPath)
		return c.File(cacheFullPath)
	}

	c.Logger().Debugf("cached file doesn't exist, resizing file: %s", cacheFullPath)

	// if cache folder doesn't exist try to create
	if err := fs.CheckFolder(cacheFolder); err != nil {
		if err := fs.CreateFolder(cacheFolder); err != nil {
			c.Logger().Errorf("failed to create cache folder: %s", cacheFolder)
			return echo.ErrInternalServerError
		}
	}

	// read original file
	src, err := imaging.Open(originalFilePath)
	if err != nil {
		c.Logger().Errorf("failed to open image: %v", err)
		return echo.ErrInternalServerError
	}

	// resize file and write to disk
	intWidth, _ := strconv.Atoi(width)
	dstImage := imaging.Resize(src, intWidth, 0, imaging.Lanczos)
	err = imaging.Save(dstImage, cacheFullPath)
	if err != nil {
		c.Logger().Errorf("failed to save image: %v", err)
		return echo.ErrInternalServerError
	}

	return c.File(cacheFullPath)
}
