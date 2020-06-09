package anydesk

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewApiTestServer(t *testing.T, url string, outputFile string, statusCode int) *httptest.Server {
	// Create test ts
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(statusCode)

		if req.URL.String() != url {
			assert.Equal(t, url, req.URL.String(), "wrong request url")
		}

		if outputFile != "" {
			data, err := ioutil.ReadFile(outputFile)
			assert.NoError(t, err)

			_, err = rw.Write(data)
			assert.NoError(t, err)
		}
	}))
}

func NewApiTestClient(t *testing.T, ts *httptest.Server, licenseId string, apiPassword string) (api *Api) {
	api = NewApi(licenseId, apiPassword)
	api.ApiEndpoint = ts.URL
	api.HttpClient = ts.Client()
	return
}
