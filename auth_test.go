package anydesk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestNewAuthenticationRequest(t *testing.T) {
	server := NewAPITestServer(t, "/auth", "./_tests/auth_response.json", http.StatusOK)
	defer server.Close()

	client := NewAPITestClient(t, server, "", "")

	req := NewAuthenticationRequest()
	resp, err := req.Do(client)

	assert.NoError(t, err)
	assert.Equal(t, "TEST_RESULT", resp.Result, "invalid result")
	assert.Equal(t, "TEST_LICENSE", resp.LicenseID, "invalid license")
}

func ExampleNewAuthenticationRequest() {
	api := NewAPI(os.Getenv("LICENSE_ID"), os.Getenv("API_PASSWORD"))

	request := NewAuthenticationRequest()
	response, err := request.Do(api)
	if err != nil {
		log.Fatal(err)
	}

	if response.Result != "success" {
		log.Fatalf("response failed: %s %s", response.Code, response.Error)
	}

	fmt.Printf("Status: %s, License: %s", response.Result, response.LicenseID)
}
