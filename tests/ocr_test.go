package tests

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"uploader/pkg/handler"
)

func TestOcr(t *testing.T) {
	resp := UploadFile(t, "receipt1.jpg", "image/jpeg")
	require.Equal(t, http.StatusCreated, resp.StatusCode())

	var response handler.FileResponse
	err := json.Unmarshal(resp.Body(), &response)
	require.NoError(t, err)

	image := response.Path

	client := resty.New().
		SetBaseURL("http://127.0.0.1:3333/v1")

	t.Run("Download OCR", func(t *testing.T) {
		resp, _ := client.R().
			SetAuthToken(testToken).
			Get("/receipts/ocr/" + image)

		require.Equal(t, http.StatusOK, resp.StatusCode())

		var ocrResp handler.OCRResponse
		err := json.Unmarshal(resp.Body(), &ocrResp)
		require.NoError(t, err)

		require.Equal(t, "29.25", ocrResp.Total)
	})
}
