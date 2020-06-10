package anydesk

import (
	"encoding/json"
	"fmt"
	"time"
)

// ClientNode is the common structure of the API for AnyDesk clients.
type ClientNode struct {
	// The unique client ID of the client AnyDesk software installation.
	ClientID int64 `json:"cid"`

	// The unique client alias of the AnyDesk software installation.
	Alias string `json:"alias"`
}

// ClientDetailRequest is used to read details about a single client from the REST API.
type ClientDetailRequest struct {
	*BaseRequest
}

// Do will execute the "/auth" query against the given API.
func (req *ClientDetailRequest) Do(api *API) (r *ClientDetailResponse, err error) {
	r = newClientDetailResponse()

	body, err := api.Do(req)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return
	}

	return
}

// NewClientDetailRequest returns a clean API request to retrieve client details from the API.
func NewClientDetailRequest(clientID int64) *ClientDetailRequest {
	return &ClientDetailRequest{
		&BaseRequest{
			Method:    "GET",
			Resource:  fmt.Sprintf("/clients/%d", clientID),
			Timestamp: time.Now().Unix(),
		},
	}
}

// ClientDetailResponse contains all available fields returned by the `/auth` API call.
type ClientDetailResponse struct {
	// ID of the client the response is about.
	ClientID int64 `json:"cid"`

	// Current version of the clients AnyDesk software.
	ClientVersion string `json:"client-version"`

	// Currently set alias of the client.
	Alias string `json:"alias"`

	// Indicates if the client is currently online.
	Online bool `json:"online"`

	// Comment for the given client, as defined in the address book.
	Comment string `json:"comment"`

	// Seconds since the client came online.
	// Will be -1 if client is Offline, but please use the .Online attribute for check.
	OnlineSinceSeconds int64 `json:"online-time"`

	// Last five sessions that this client was involved in.
	LastSessions []SessionNode `json:"last-sessions"`
}

// OnlineSince returns the original time when the client came online.
func (r *ClientDetailResponse) OnlineSince() time.Time {
	n := time.Now()
	return n.Add(time.Second * (-1 * time.Duration(r.OnlineSinceSeconds)))
}

func newClientDetailResponse() *ClientDetailResponse {
	return &ClientDetailResponse{}
}
