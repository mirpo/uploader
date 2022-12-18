package handler

import (
	"bytes"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"uploader/pkg/auth"
)

func getContextBasic() echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	context.Set("user", &jwt.Token{
		Claims: &auth.Claims{
			UserID:         "1",
			StandardClaims: jwt.StandardClaims{},
		},
	})

	return context
}

func getContextWithFile(t *testing.T) echo.Context {
	e := echo.New()

	path := "../../tests/testdata/receipt1.jpg"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", path)
	require.NoError(t, err)
	sample, err := os.Open(path)
	require.NoError(t, err)

	_, err = io.Copy(part, sample)
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	context.Set("user", &jwt.Token{
		Claims: &auth.Claims{
			UserID:         "1",
			StandardClaims: jwt.StandardClaims{},
		},
	})

	return context
}

func TestUpload(t *testing.T) {
	t.Run("500 when user's folder doesn't exist and failed to create", func(t *testing.T) {
		res := Upload(getContextBasic())

		require.ErrorIs(t, echo.ErrInternalServerError, res)
	})

	t.Run("400 when failed to read file from the request", func(t *testing.T) {
		t.Setenv("UPLOAD_PATH", t.TempDir())

		res := Upload(getContextBasic())

		require.ErrorIs(t, echo.ErrBadRequest, res)
	})

	t.Run("400 when invalid file mime type", func(t *testing.T) {
		t.Setenv("UPLOAD_PATH", t.TempDir())
		t.Setenv("ALLOWED_MIME_TYPES", "text/plain")

		res := Upload(getContextWithFile(t))

		require.ErrorIs(t, echo.ErrBadRequest, res)
	})

	t.Run("201 when file was successfully saved to the disk", func(t *testing.T) {
		t.Setenv("UPLOAD_PATH", t.TempDir())
		t.Setenv("ALLOWED_MIME_TYPES", "image/jpeg")
		context := getContextWithFile(t)

		_ = Upload(context)

		require.Equal(t, http.StatusCreated, context.Response().Status)
	})
}
