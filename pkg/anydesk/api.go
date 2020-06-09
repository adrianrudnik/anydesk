package anydesk

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
)

type Api struct {
	LicenseId   string `json:"license_id"`
	ApiPassword string `json:"api_password"`
	ApiEndpoint string `json:"api_endpoint"`
	HttpClient  *http.Client
}

func NewApi(licenseId string, apiPassword string) *Api {
	return &Api{
		LicenseId:   licenseId,
		ApiPassword: apiPassword,
		ApiEndpoint: "https://v1.api.anydesk.com:8081",
		HttpClient:  &http.Client{},
	}
}

// GetRequestToken generates the request token used for the API request
func (api *Api) GetRequestToken(request *BaseRequest) string {
	h := hmac.New(sha1.New, []byte(api.ApiPassword))
	h.Write([]byte(request.GetRequestString()))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (api *Api) Do(request ApiRequest) (body []byte, err error) {
	r, err := request.GetHttpRequest(api)
	if err != nil {
		return
	}

	resp, err := api.HttpClient.Do(r)

	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = errors.New(resp.Status)
		return
	}

	return
}
