package anydesk

import (
	"encoding/json"
	"time"
)

// AuthenticationRequest is used to query the "/auth" API endpoint for information.
type AuthenticationRequest struct {
	*BaseRequest
}

// Do will execute the "/auth" query against the given API.
func (req *AuthenticationRequest) Do(api *Api) (r *AuthenticationResponse, err error) {
	r = newAuthenticationResponse()

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

// NewAuthenticationRequest will a clean API request against "/auth".
func NewAuthenticationRequest() *AuthenticationRequest {
	return &AuthenticationRequest{
		&BaseRequest{
			Method:    "GET",
			Resource:  "/auth",
			Timestamp: time.Now().Unix(),
			Content:   []byte(""),
		},
	}
}
