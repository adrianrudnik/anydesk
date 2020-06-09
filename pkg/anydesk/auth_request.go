package anydesk

import (
	"encoding/json"
	"time"
)

type AuthenticationRequest struct {
	*BaseRequest
}

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
