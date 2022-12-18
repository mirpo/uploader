package tests

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {
	client := resty.New().
		SetBaseURL("http://127.0.0.1:3333/v1")

	t.Run("Health endpoint", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			resp, _ := client.R().
				Get("/health")

			require.Equal(t, http.StatusOK, resp.StatusCode())
			require.Equal(t, "{\"msg\":\"All good!\"}\n", string(resp.Body()))
		})
	})
}
