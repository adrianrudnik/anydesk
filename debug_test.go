package anydesk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func ExampleSetDebug() {
	api := NewAPI("license", "password")

	SetDebug(true)

	r := NewAuthenticationRequest()
	r.Do(api)

	fmt.Printf(
		"Url: %s, Response body: %s",
		r.GetDebug().RequestURL,
		r.GetDebug().ResponseBody,
	)
}

func TestApi_DebugEnabled(t *testing.T) {
	SetDebug(true)

	server := NewAPITestServer(t, "/test?demo=123", "", 201)
	defer server.Close()

	client := NewAPITestClient(t, server, "", "")

	q := &url.Values{}
	q.Set("demo", strconv.FormatInt(123, 10))

	req := &BaseRequest{
		Method:    "GET",
		Resource:  "/test",
		Query:     q,
		Timestamp: time.Now().Unix(),
		Content:   []byte{},
	}
	_, err := client.Do(req)

	a := assert.New(t)
	a.NoError(err)
	a.Equal(201, req.GetDebug().Response.StatusCode)
	a.Equal(true, req.GetDebug().Available)
	a.NotNil(req.GetDebug().Response)
	a.NotNil(req.GetDebug().Request)
	a.NotNil(req.GetDebug().RequestURL)

	a.Equal("/test", req.GetDebug().RequestURL.Path)
	a.Equal("demo=123", req.GetDebug().RequestURL.RawQuery)
	a.Equal("127.0.0.1", req.GetDebug().RequestURL.Hostname())

	SetDebug(false)
}
