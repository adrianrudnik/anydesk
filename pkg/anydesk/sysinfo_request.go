package anydesk

import (
	"encoding/json"
	"time"
)

type SysinfoRequest struct {
	*BaseRequest
}

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
