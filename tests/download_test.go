package tests

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"uploader/pkg/handler"
)

func TestDownload(t *testing.T) {
	resp := UploadFile(t, "receipt1.jpg", "image/jpeg")
	require.Equal(t, http.StatusCreated, resp.StatusCode())

	var response handler.FileResponse
	err := json.Unmarshal(resp.Body(), &response)
	require.NoError(t, err)

	image := response.Path

	client := resty.New().
		SetBaseURL("http://127.0.0.1:3333/v1")

	t.Run("Download", func(t *testing.T) {
		t.Run("401 when no JWT token", func(t *testing.T) {
			resp, _ := client.R().
				SetAuthToken("QWERTY").
				Get("/receipts/" + image)

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode())
		})

		t.Run("200 when downloading original receipt", func(t *testing.T) {
			resp, _ := client.R().
				SetAuthToken(testToken).
				Get("/receipts/" + image)

			require.Equal(t, http.StatusOK, resp.StatusCode())
		})

		t.Run("200 when downloading resized receipt", func(t *testing.T) {
			resp, _ := client.R().
				SetAuthToken(testToken).
				Get("/receipts/" + image + "?width=200")

			require.Equal(t, http.StatusOK, resp.StatusCode())
		})
	})
}
