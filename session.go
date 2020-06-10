package anydesk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
)

// SessionNode is the common structure of session information produced by AnyDesk clients.
type SessionNode struct {
	// Indicatges of the session is currently active.
	Active bool `json:"active"`

	// The unique sesson ID for this connection.
	SessionID string `json:"sid"`

	// The source client responsible for this session.
	Source *ClientNode `json:"from"`

	// The connected client of the session.
	Target *ClientNode `json:"to"`

	// Connection start as unix-timestamp.
	StartTimestamp int64 `json:"start-time"`

	// Connection end as unix-timestamp.
	EndTimestamp int64 `json:"end-time"`

	// Total duration of the session in seconds.
	DurationInSeconds int64 `json:"duration"`

	// The comment left by the source client.
	Comment string `json:"comment"`
}

// StartTime returns the connection start time.
func (n *SessionNode) StartTime() time.Time {
	return time.Unix(n.StartTimestamp, 0)
}

// EndTime returns the connection end time.
func (n *SessionNode) EndTime() time.Time {
	return time.Unix(n.EndTimestamp, 0)
}

// Duration returns the total duration of the session.
func (n *SessionNode) Duration() time.Duration {
	return time.Second * time.Duration(n.DurationInSeconds)
}

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
	var v *string

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
	ClientID int64

	// Limit search to given sessiond direction, [in, out, inout]
	Direction SessionDirection

	// Limit search to sessions after the given time
	TimeFrom time.Time

	// Limit search to sessions up to the given time
	TimeTo time.Time
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

		if search.ClientID > 0 {
			q.Set("cid", strconv.FormatInt(search.ClientID, 10))
		}

		if search.Direction == DirectionInOut ||
			search.Direction == DirectionIn ||
			search.Direction == DirectionOut {
			q.Set("direction", string(search.Direction))
		}

		if !search.TimeFrom.IsZero() {
			q.Set("from", strconv.FormatInt(search.TimeFrom.Unix(), 10))
		}

		if !search.TimeTo.IsZero() {
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
