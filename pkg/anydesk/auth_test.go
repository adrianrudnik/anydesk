package anydesk

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
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
