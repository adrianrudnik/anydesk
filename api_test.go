package anydesk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// NewAPITestServer will create a AnyDesk API stub to work tests against.
func NewAPITestServer(t *testing.T, url string, outputFile string, statusCode int) *httptest.Server {
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

// NewAPITestClient will create a AnyDesk API client that accepts an API stub server.
func NewAPITestClient(t *testing.T, ts *httptest.Server, licenseID string, apiPassword string) (api *API) {
	api = NewAPI(licenseID, apiPassword)
	api.APIEndpoint = ts.URL
	api.HTTPClient = ts.Client()
	return
}

func TestApi_GetRequestToken(t *testing.T) {
	a := NewAPI("1438129266231705", "UYETICGU2CT3KES")

	r := &BaseRequest{
		Method:    "GET",
		Resource:  "/auth",
		Timestamp: 1445440997,
	}

	assert.Equal(t, "T2YsCOj2o3Rb79nLPUgx3Gl+nnw=", a.GetRequestToken(r))
}

func TestApi_InvalidHttpStatusCode(t *testing.T) {
	server := NewAPITestServer(t, "/", "", 400)
	defer server.Close()

	client := NewAPITestClient(t, server, "", "")

	req := &BaseRequest{
		Method:    "GET",
		Resource:  "/",
		Timestamp: time.Now().Unix(),
	}
	_, err := client.Do(req)
	assert.Error(t, err)
}

func TestApi_NotFoundStatusCode(t *testing.T) {
	server := NewAPITestServer(t, "/", "", 404)
	defer server.Close()

	client := NewAPITestClient(t, server, "", "")

	req := &BaseRequest{
		Method:    "GET",
		Resource:  "/",
		Timestamp: time.Now().Unix(),
	}
	_, err := client.Do(req)
	assert.Error(t, err)
	assert.IsType(t, &APINotFoundError{}, err)
}

func TestRequest_GetContentHash(t *testing.T) {
	r := &BaseRequest{
		Method:    "GET",
		Resource:  "/auth",
		Timestamp: 1445440997,
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

func ExampleNewApi() {
	api := NewAPI(os.Getenv("LICENSE_ID"), os.Getenv("API_PASSWORD"))

	// Optional: Decrease timeouts
	api.HTTPClient = &http.Client{Timeout: 5 * time.Second}

	// Optional: Switch to Enterprise API
	api.APIEndpoint = "https://yourinstance:8081"

	request := NewAuthenticationRequest()
	response, _ := request.Do(api)

	fmt.Printf("Status: %s, License: %s", response.Result, response.LicenseID)
}
