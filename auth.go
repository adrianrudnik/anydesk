package anydesk

import (
	"encoding/json"
	"time"
)

// AuthenticationRequest is used to read the "/auth" API resource.
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

// AuthenticationResponse contains all available fields returned by the `/auth` API call.
type AuthenticationResponse struct {
	Result           string `json:"result"`
	Error            string `json:"error"`
	Code             string `json:"code"`
	Method           string `json:"method"`
	Resource         string `json:"method"`
	RequestTimestamp string `json:"request-time"`
	ContentHash      string `json:"content-hash"`
	LicenseId        string `json:"license-id"`
}

func newAuthenticationResponse() *AuthenticationResponse {
	return &AuthenticationResponse{}
}
