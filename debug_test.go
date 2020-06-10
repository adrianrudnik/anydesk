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
	api := NewApi("license", "password")

	SetDebug(true)

	r := NewAuthenticationRequest()
	r.Do(api)

	fmt.Printf(
		"Url: %s, Response body: %s",
		r.GetDebug().RequestUrl,
		r.GetDebug().ResponseBody,
	)
}

func TestApi_DebugEnabled(t *testing.T) {
	SetDebug(true)

	server := NewApiTestServer(t, "/test?demo=123", "", 201)
	defer server.Close()

	client := NewApiTestClient(t, server, "", "")

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
	a.NotNil(req.GetDebug().RequestUrl)

	a.Equal("/test", req.GetDebug().RequestUrl.Path)
	a.Equal("demo=123", req.GetDebug().RequestUrl.RawQuery)
	a.Equal("127.0.0.1", req.GetDebug().RequestUrl.Hostname())

	SetDebug(false)
}
