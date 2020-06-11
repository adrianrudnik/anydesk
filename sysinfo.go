package anydesk

import (
	"encoding/json"
)

// SysinfoRequest is used to read the "/sysinfo" API resource.
type SysinfoRequest struct {
	*BaseRequest
}

// Do will execute the "/sysinfo" query against the given API.
func (req *SysinfoRequest) Do(api *API) (resp *SysinfoResponse, err error) {
	resp = newSysinfoResponse()

	body, err := api.Do(req.BaseRequest)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}

	return
}

// NewSysinfoRequest will a clean API request against "/sysinfo".
func NewSysinfoRequest() *SysinfoRequest {
	return &SysinfoRequest{
		&BaseRequest{
			Method:   "GET",
			Resource: "/sysinfo",
		},
	}
}

// SysinfoResponse contains all available fields returned by the `/sysinfo` API call.
type SysinfoResponse struct {
	Name       string `json:"name"`
	APIVersion string `json:"api-ver"`
	License    struct {
		Name             string `json:"name"`
		ExpiresTimestamp int64  `json:"expires"`
		HasExpired       bool   `json:"has-expired"` // undocumented or deprecated
		MaxClients       int    `json:"max-clients"`
		MaxSessions      int    `json:"max-sessions"`
		MaxSessionTime   int    `json:"max-session-time"`
		Namespaces       []struct {
			Name string `json:"name"`
			Size int    `json:"size"`
		} `json:"namespaces"`
		ID          string `json:"license-id"`
		Key         string `json:"license-key"`
		APIPassword string `json:"api-password"`
		PowerUser   bool   `json:"power-user"` // undocumented or deprecated
	} `json:"license"`
	Clients struct {
		Total  int `json:"total"`
		Online int `json:"online"`
	} `json:"clients"`
	Sessions struct {
		Total  int `json:"total"`
		Active int `json:"active"`
	} `json:"sessions"`
	Standalone bool `json:"standalone"`
}

func newSysinfoResponse() *SysinfoResponse {
	return &SysinfoResponse{}
}
