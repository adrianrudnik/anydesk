package anydesk

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewSysinfoRequest(t *testing.T) {
	server := NewApiTestServer(t, "/sysinfo", "./_tests/sysinfo_response.json", http.StatusOK)
	defer server.Close()

	client := NewApiTestClient(t, server, "", "")

	req := NewSysinfoRequest()
	resp, err := req.Do(client)

	a := assert.New(t)
	a.NoError(err)
	a.Equal("1.1", resp.ApiVersion)
	a.Equal(4, resp.Clients.Online)
	a.Equal(8, resp.Clients.Total)
	a.Equal("TEST_APIPASS", resp.License.ApiPassword)
	a.Equal(int64(1623920819), resp.License.ExpiresTimestamp)
	a.Equal(true, resp.License.HasExpired)
	a.Equal("TEST_LICENSE_ID", resp.License.Id)
	a.Equal("TEST_LICENSE_KEY", resp.License.Key)
	a.Equal(-2, resp.License.MaxClients)
	a.Equal(-3, resp.License.MaxSessionTime)
	a.Equal(4, resp.License.MaxSessions)
	a.Equal("TEST_LICENSE_NAME", resp.License.Name)
	// @todo namespace tests
	a.Equal(true, resp.License.PowerUser)
	a.Equal("AnyDesk REST", resp.Name)
	a.Equal(2, resp.Sessions.Active)
	a.Equal(599, resp.Sessions.Total)
	a.Equal(true, resp.Standalone)
}
