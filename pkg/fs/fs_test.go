package fs

import (
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestFs(t *testing.T) {
	t.Run("FileExists()", func(t *testing.T) {
		_, filename, _, _ := runtime.Caller(1)

		require.True(t, FileExists(filename))
		require.False(t, FileExists(filename+"RAND"))
	})
}
