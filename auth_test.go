package anydesk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestNewAuthenticationRequest(t *testing.T) {
	server := NewApiTestServer(t, "/auth", "./_tests/auth_response.json", http.StatusOK)
	defer server.Close()

	client := NewApiTestClient(t, server, "", "")

	req := NewAuthenticationRequest()
	resp, err := req.Do(client)

	assert.NoError(t, err)
	assert.Equal(t, "TEST_RESULT", resp.Result, "invalid result")
	assert.Equal(t, "TEST_LICENSE", resp.LicenseId, "invalid license")
}

func ExampleNewAuthenticationRequest() {
	api := NewApi(os.Getenv("LICENSE_ID"), os.Getenv("API_PASSWORD"))

	// Optional: Decrease timeouts
	api.HttpClient = &http.Client{Timeout: 5 * time.Second}

	// Optional: Switch to Enterprise API
	api.ApiEndpoint = "https://yourinstance:8081"

	request := NewAuthenticationRequest()
	response, _ := request.Do(api)

	fmt.Printf("Status: %s, License: %s", response.Result, response.LicenseId)
}
