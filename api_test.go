package anydesk

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// NewApiTestServer will create a AnyDesk API stub to work tests against.
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

// NewApiTestClient will create a AnyDesk API client that accepts an API stub server.
func NewApiTestClient(t *testing.T, ts *httptest.Server, licenseId string, apiPassword string) (api *Api) {
	api = NewApi(licenseId, apiPassword)
	api.ApiEndpoint = ts.URL
	api.HttpClient = ts.Client()
	return
}


func TestApi_GetRequestToken(t *testing.T) {
	a := NewApi("1438129266231705", "UYETICGU2CT3KES")

	r := &BaseRequest{
		Method:    "GET",
		Resource:  "/auth",
		Timestamp: 1445440997,
		Content:   []byte(""),
	}

	assert.Equal(t, "T2YsCOj2o3Rb79nLPUgx3Gl+nnw=", a.GetRequestToken(r))
}

func TestApi_InvalidHttpStatusCode(t *testing.T) {
	server := NewApiTestServer(t, "/", "", 400)
	defer server.Close()

	client := NewApiTestClient(t, server, "", "")

	req := &BaseRequest{
		Method:    "GET",
		Resource:  "/",
		Timestamp: time.Now().Unix(),
		Content:   []byte{},
	}
	_, err := client.Do(req)
	assert.Error(t, err)
}

func TestRequest_GetContentHash(t *testing.T) {
	r := &BaseRequest{
		Method:    "GET",
		Resource:  "/auth",
		Timestamp: 1445440997,
		Content:   []byte(""),
	}

	assert.Equal(t, "2jmj7l5rSw0yVb/vlWAYkK/YBwk=", r.GetContentHash())
}

func TestRequest_GetRequestString(t *testing.T) {
	r := &BaseRequest{
		Method:    "GET",
		Resource:  "/auth",
		Timestamp: 1445440997,
		Content:   []byte(""),
	}

	assert.Equal(t, "GET\n/auth\n1445440997\n2jmj7l5rSw0yVb/vlWAYkK/YBwk=", r.GetRequestString())
}