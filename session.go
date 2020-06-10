package anydesk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
)

// SessionDirection defines the that connection direction  on a paginated API request.
type SessionDirection string

const (
	// DirectionIn will filter by incoming sessions only.
	DirectionIn SessionDirection = "in"

	// DirectionInOut will show all sessions, regardsless of their direction.
	DirectionInOut SessionDirection = "inout"

	// DirectionOut will filter by outgoing sessions only.
	DirectionOut SessionDirection = "out"
)

// SessionCommentChangeRequest is used to patch the /session/{id} API resource.
type SessionCommentChangeRequest struct {
	*BaseRequest
	Comment *string `json:"comment"`
}

// Do will execute the "/auth" query against the given API.
func (req *SessionCommentChangeRequest) Do(api *API) (err error) {
	// Execute the request by handing it over to the given API configuration
	body, err := api.Do(req)
	if err != nil {
		return
	}

	data, err := json.Marshal(body)
	if err != nil {
		return
	}

	fmt.Println(string(data))

	_, err = api.Do(req)
	return
}

// NewSessionCommentChangeRequest will create an API request that will set the given comment to the given session ID.
// Giving an empty comment string will remove the currently set comment.
func NewSessionCommentChangeRequest(session string, comment string) *SessionCommentChangeRequest {
	var v *string = nil

	if comment != "" {
		v = &comment
	}

	return &SessionCommentChangeRequest{
		&BaseRequest{
			Method:    "PATCH",
			Resource:  fmt.Sprintf("/sessions/%s", session),
			Timestamp: time.Now().Unix(),
		},
		v,
	}
}

// SessionListRequest is used to retrieve a list of stored sessions from the /sessions API resource.
type SessionListRequest struct {
	*BaseRequest
	*PaginationOptions
}

// SessionListSearch defines all configurable search parameters for NewSessionListRequest()
type SessionListSearch struct {
	// Limit search to client ID
	ClientID string

	// Limit search to given sessiond direction, [in, out, inout]
	Direction SessionDirection

	// Limit search to sessions after the given time
	TimeFrom *time.Time

	// Limit search to sessions up to the given time
	TimeTo *time.Time
}

// Do will execute the "/sessions" query against the given API.
func (req *SessionListRequest) Do(api *API) (err error) {
	body, err := api.DoPaginated(req)

	ioutil.WriteFile("out.json", body, 0644)

	if err != nil {
		return
	}

	return
}

// NewSessionListRequest returns a new session list query.
func NewSessionListRequest(search *SessionListSearch) *SessionListRequest {
	// Handle search
	var q *url.Values
	if search != nil {
		q = &url.Values{}

		if search.ClientID != "" {
			q.Set("cid", search.ClientID)
		}

		if search.Direction == DirectionInOut ||
			search.Direction == DirectionIn ||
			search.Direction == DirectionOut {
			q.Set("direction", string(search.Direction))
		}

		if search.TimeFrom != nil {
			q.Set("from", strconv.FormatInt(search.TimeFrom.Unix(), 10))
		}

		if search.TimeTo != nil {
			q.Set("to", strconv.FormatInt(search.TimeTo.Unix(), 10))
		}

		q.Set("limit", strconv.FormatInt(10, 10))
	}

	return &SessionListRequest{
		BaseRequest: &BaseRequest{
			Method:    "GET",
			Resource:  "/sessions",
			Query:     q,
			Timestamp: time.Now().Unix(),
		},
		PaginationOptions: NewPaginationOptions(),
	}
}
