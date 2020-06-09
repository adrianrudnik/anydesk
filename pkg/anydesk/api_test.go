package anydesk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
