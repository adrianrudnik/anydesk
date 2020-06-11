package anydesk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestNewClientDetailRequest(t *testing.T) {
	server := NewAPITestServer(t, "/clients/100000000", "./_tests/client_detail.json", 200)
	defer server.Close()

	client := NewAPITestClient(t, server, "", "")

	req := NewClientDetailRequest(100000000)
	resp, err := req.Do(client)

	a := assert.New(t)
	a.NoError(err)
	a.Equal("xyz", resp.Alias)
	a.Equal(int64(100000000), resp.ClientID)
	a.Equal("1.2.3", resp.ClientVersion)
	a.Equal("TEST-COMMENTA", resp.Comment)
	a.True(resp.Online)
	a.Equal(int64(456), resp.OnlineSinceSeconds)
	a.Equal(time.Now().Unix(), resp.OnlineSince().Add(time.Second*time.Duration(resp.OnlineSinceSeconds)).Unix())

	a.Len(resp.LastSessions, 2)

	s1 := resp.LastSessions[0]
	a.False(s1.Active)
	a.Equal("SESSIONA", s1.SessionID)
	a.Equal("", s1.Comment)
	a.Equal(int64(123), s1.DurationInSeconds)
	a.Equal(s1.DurationInSeconds, int64(s1.Duration()/time.Second))
	a.Equal(int64(1590504637), s1.EndTimestamp)
	a.Equal(s1.EndTimestamp, s1.EndTime().Unix())
	a.Equal(int64(1590504626), s1.StartTimestamp)
	a.Equal(s1.StartTimestamp, s1.StartTime().Unix())
	a.Equal("TEST_ALIAS1", s1.Source.Alias)
	a.Equal(int64(100000000), s1.Source.ClientID)
	a.Equal("", s1.Target.Alias)
	a.Equal(int64(100000001), s1.Target.ClientID)

	s2 := resp.LastSessions[1]
	a.True(s2.Active)
	a.Equal("SESSIONB", s2.SessionID)
	a.Equal("TEST_COMMENTB", s2.Comment)
	a.Equal(int64(321), s2.DurationInSeconds)
	a.Equal(s2.DurationInSeconds, int64(s2.Duration()/time.Second))
	a.Equal(int64(1587473931), s2.EndTimestamp)
	a.Equal(s2.EndTimestamp, s2.EndTime().Unix())
	a.Equal(int64(1587473919), s2.StartTimestamp)
	a.Equal(s2.StartTimestamp, s2.StartTime().Unix())
	a.Equal("", s2.Source.Alias)
	a.Equal(int64(100000010), s2.Source.ClientID)
	a.Equal("TEST_ALIAS2", s2.Target.Alias)
	a.Equal(int64(100000000), s2.Target.ClientID)
}

func ExampleNewClientDetailRequest() {
	api := NewAPI(os.Getenv("LICENSE_ID"), os.Getenv("API_PASSWORD"))

	request := NewClientDetailRequest(123456789)
	response, _ := request.Do(api)

	fmt.Printf(
		"Version: %s, Started %s (%d seconds)",
		response.ClientVersion,
		response.LastSessions[0].StartTime(),
		response.LastSessions[0].DurationInSeconds,
	)
}

func TestNewClientListRequest(t *testing.T) {
	server := NewAPITestServer(t, "/clients?limit=-1&offset=0&order=desc", "./_tests/client_list_all.json", 200)
	defer server.Close()

	client := NewAPITestClient(t, server, "", "")

	req := NewClientListRequest(nil)
	resp, err := req.Do(client)

	a := assert.New(t)
	a.NoError(err)

	a.Equal(int64(321), resp.Count)
	a.Equal(int64(654), resp.Limit)
	a.Equal(int64(528), resp.Offset)
	a.Equal(int64(671), resp.Selected)
	a.True(resp.Online)

	a.Len(resp.List, 8)

	n2 := resp.List[1]
	a.Equal("", n2.Alias)
	a.Equal(int64(122), n2.ClientID)
	a.Equal("5.4.0", n2.ClientVersion)
	a.Equal("TEST-02", n2.Comment)
	a.True(n2.Online)
	a.Equal(int64(49215), n2.OnlineSinceSeconds)

	n8 := resp.List[7]
	a.Equal("demo@ad", n8.Alias)
	a.Equal(int64(128), n8.ClientID)
	a.Equal("5.5.3", n8.ClientVersion)
	a.Equal("", n8.Comment)
	a.False(n8.Online)
	a.Equal(int64(0), n8.OnlineSinceSeconds)
}
