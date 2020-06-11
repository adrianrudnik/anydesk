package anydesk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// ClientNode is the common structure of the API for AnyDesk clients.
type ClientNode struct {
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
	// Only available if queried by ClientDetailRequest.
	LastSessions []SessionNode `json:"last-sessions"`
}

// ClientSlimNode is the common short representation of the API.
// It does not contain all available information and is used in list queries.
type ClientSlimNode struct {
	ClientID      int64  `json:"cid"`
	ClientVersion string `json:"string"`
	Alias         string `json:"alias"`
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
			Method:   "GET",
			Resource: fmt.Sprintf("/clients/%d", clientID),
		},
	}
}

// ClientDetailResponse contains all fields available to the client details API resource.
type ClientDetailResponse struct {
	*ClientNode
}

// OnlineSince returns the original time when the client came online.
func (r *ClientDetailResponse) OnlineSince() time.Time {
	n := time.Now()
	return n.Add(time.Second * (-1 * time.Duration(r.OnlineSinceSeconds)))
}

func newClientDetailResponse() *ClientDetailResponse {
	return &ClientDetailResponse{}
}

// ClientListRequest is used to read a list of clients from the API resource.
type ClientListRequest struct {
	*BaseRequest
	*PaginationOptions
}

// ClientListSearch configures the search  params for NewClientListRequest.
type ClientListSearch struct {
	// Limits search to online clients.
	// Setting this to false will not list only offline clients.
	Online bool
}

// Do will execute the request against the API.
func (req *ClientListRequest) Do(api *API) (r *ClientListResponse, err error) {
	r = newClientListResponse()

	body, err := api.DoPaginated(req)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return
	}

	return
}

// NewClientListRequest returns a clean API request to retrieve a list of clients from the API.
func NewClientListRequest(search *ClientListSearch) *ClientListRequest {
	var q *url.Values

	if search != nil {
		q = &url.Values{}

		if search.Online {
			q.Set("online", "true")
		}
	}
	return &ClientListRequest{
		BaseRequest: &BaseRequest{
			Method:   "GET",
			Resource: "/clients",
			Query:    q,
		},
		PaginationOptions: NewPaginationOptions(),
	}
}

// ClientListResponse contains all fields available for client lists from the API resource.
type ClientListResponse struct {
	*PaginatedResult
	Online bool         `json:"online"`
	List   []ClientNode `json:"list"`
}

func newClientListResponse() *ClientListResponse {
	return &ClientListResponse{}
}
