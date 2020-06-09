package anydesk

import (
	"encoding/json"
	"time"
)

// SysinfoRequest is used to query the "/sysinfo" API endpoint for information.
type SysinfoRequest struct {
	*BaseRequest
}

// Do will execute the "/sysinfo" query against the given API.
func (req *SysinfoRequest) Do(api *Api) (resp *SysinfoResponse, err error) {
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
			Method:    "GET",
			Resource:  "/sysinfo",
			Timestamp: time.Now().Unix(),
			Content:   []byte(""),
		},
	}
}
