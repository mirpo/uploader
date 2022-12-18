package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelper(t *testing.T) {
	t.Run("GetAllowedWidth()", func(t *testing.T) {
		t.Setenv("ALLOWED_WIDTHS", "1,2,3,4")

		require.Equal(t, []string{"1", "2", "3", "4"}, GetAllowedWidth())
	})

	t.Run("GetUserFolder()", func(t *testing.T) {
		t.Setenv("UPLOAD_PATH", "./super-folder")

		require.Equal(t, "./super-folder/12345/", GetUserFolder("12345"))
	})

	t.Run("GetCachePaths()", func(t *testing.T) {
		cacheFolder, cachedFilename := GetCachePaths("2022/12/15/1671118166178899000.jpg", "./cache/1/", "200")

		require.Equal(t, "./cache/1/2022/12/15/", cacheFolder)
		require.Equal(t, "1671118166178899000_200px.jpg", cachedFilename)
	})
}
