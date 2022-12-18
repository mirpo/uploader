package helper

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func GetAllowedWidth() []string {
	return strings.Split(os.Getenv("ALLOWED_WIDTHS"), ",")
}

func GetAllowedMimeTypes() []string {
	return strings.Split(os.Getenv("ALLOWED_MIME_TYPES"), ",")
}

func GetUserFolder(userID string) string {
	return fmt.Sprintf("%s/%s/", os.Getenv("UPLOAD_PATH"), userID)
}

func GetUserCacheFolder(userID string) string {
	return fmt.Sprintf("%s/%s/", os.Getenv("CACHE_PATH"), userID)
}

func GetFilename(file *multipart.FileHeader) string {
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
}

// GetCachePaths returns path to user's cache folder and filename with width inside filename
func GetCachePaths(uploadedFilePath string, userCacheFolder string, width string) (string, string) {
	filename := path.Base(uploadedFilePath)
	folder := path.Dir(uploadedFilePath) + "/"
	extension := filepath.Ext(filename)
	filename = strings.TrimSuffix(filename, extension)

	return userCacheFolder + folder, fmt.Sprintf("%s_%spx%s", filename, width, extension)
}

var ErrFailedToGetMimeType = errors.New("failed to get mime type")

// DetectAndValidateMimeType helper to detect mime type based on first 512 bytes
// for simplicity return one error type, which will result ErrBadRequest in the handler
//result nil means that mime type was detected and is in allowed types
func DetectAndValidateMimeType(src multipart.File) error {
	// detect MIME type
	buff := make([]byte, 512)
	if _, err := src.Read(buff); err != nil {
		return ErrFailedToGetMimeType
	}

	filetype := http.DetectContentType(buff)
	if filetype == "" || !slices.Contains(GetAllowedMimeTypes(), filetype) {
		return ErrFailedToGetMimeType
	}

	// go back to the 0 position to write correct file to the disk
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return ErrFailedToGetMimeType
	}

	return nil
}
