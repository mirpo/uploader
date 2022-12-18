package tests

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestUpload(t *testing.T) {
	t.Run("Upload endpoint", func(t *testing.T) {
		t.Run("400 when unknown mime type", func(t *testing.T) {
			resp := UploadFile(t, "receipt.tiff", "image/tiff")

			require.Equal(t, http.StatusBadRequest, resp.StatusCode())
		})

		t.Run("201", func(t *testing.T) {
			resp := UploadFile(t, "receipt1.jpg", "image/jpeg")

			require.Equal(t, http.StatusCreated, resp.StatusCode())
		})
	})
}
