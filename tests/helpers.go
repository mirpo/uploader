package tests

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func UploadFile(t *testing.T, image string, mimetype string) *resty.Response {
	client := resty.New().
		SetBaseURL("http://127.0.0.1:3333/v1")

	file, err := os.Open("./testdata/" + image)
	require.NoError(t, err)
	fileBytes, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	resp, _ := client.R().
		SetMultipartField("file", "./tests/testdata/"+image, mimetype, bytes.NewReader(fileBytes)).
		SetAuthToken(testToken).
		Post("/receipts")

	return resp
}
