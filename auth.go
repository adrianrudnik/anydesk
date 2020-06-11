package anydesk

import (
	"encoding/json"
)

// AuthenticationRequest is used to read the "/auth" API resource.
type AuthenticationRequest struct {
	*BaseRequest
}

// Do will execute the "/auth" query against the given API.
func (req *AuthenticationRequest) Do(api *API) (r *AuthenticationResponse, err error) {
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
			Method:   "GET",
			Resource: "/auth",
		},
	}
}

// AuthenticationResponse contains all available fields returned by the `/auth` API call.
type AuthenticationResponse struct {
	// Status result, should be "success".
	Result string `json:"result"`

	// The humen readable error message.
	Error string `json:"error"`

	// Specific error code as string, i.e. "invalid_token".
	Code string `json:"code"`

	// Echoing the failed request method.
	Method string `json:"method"`

	// Echoing the failed request resource.
	Resource string `json:"resource"`

	// Echoping the failed request timestamp.
	RequestTimestamp string `json:"request-time"`

	// Echoing the failed request content hash.
	ContentHash string `json:"content-hash"`

	// Echoing the API license ID on success.
	LicenseID string `json:"license-id"`
}

func newAuthenticationResponse() *AuthenticationResponse {
	return &AuthenticationResponse{}
}
