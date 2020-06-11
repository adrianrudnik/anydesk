package anydesk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// Infinite is used by various calls to indicate unlimited results/items
	Infinite = int64(-1)
)

// APIRequest is the basic interface used for all compatible API requests.
type APIRequest interface {
	// Returns debug information
	GetDebug() *DebugInfo

	// Returns base information about the http request. This information
	// is used to compose the API request signature.
	GetRequestDetails() *BaseRequest

	// Returns the http.Request that is composed, signed and ready to execute
	GetHTTPRequest(api *API) (req *http.Request, err error)
}

// PaginatedAPIRequest is the base interface for all api requests that support pagination.
type PaginatedAPIRequest interface {
	APIRequest

	// Returns the pagination options for the current request
	GetPaginationOptions() *PaginationOptions
}

// API contains all information about the AnyDesk API endpoint and configurable options.
type API struct {
	// API license ID as provided by AnyDesk support
	LicenseID string `json:"license_id"`

	// API password as provided by AnyDesk support
	APIPassword string `json:"api_password"`

	// API endpoint to be used
	APIEndpoint string `json:"api_endpoint"`

	// The http client used for API requests.
	// Can be used or overwritten for timeout and transport layer configuration.
	HTTPClient *http.Client
}

// NewAPI returns an initialized AnyDesk API configuration used with a Professional license.
func NewAPI(licenseID string, apiPassword string) *API {
	return &API{
		LicenseID:   licenseID,
		APIPassword: apiPassword,
		APIEndpoint: "https://v1.api.anydesk.com:8081",
		HTTPClient:  &http.Client{},
	}
}

// GetRequestToken generates the request token used for the API request.
func (api *API) GetRequestToken(request *BaseRequest) string {
	h := hmac.New(sha1.New, []byte(api.APIPassword))
	h.Write([]byte(request.GetRequestString()))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Do will execute a given AnyDesk API request and return the plain json as string.
func (api *API) Do(request APIRequest) (body []byte, err error) {
	// Insert current timestamp so we can sign the request
	request.GetRequestDetails().Timestamp = time.Now().Unix()

	base := request.GetRequestDetails()

	// Check if the request supports the

	// Ensure we encode the optional query parameters into the BaseRequest.Resource
	// otherwise the signature will not match
	if base.Query != nil {
		base.Resource = fmt.Sprintf("%s?%s", base.Resource, base.Query.Encode())
	}

	// Ensure we encode possible request content into json
	content, err := json.Marshal(request)
	if err != nil {
		return
	}

	base.Content = content

	// Collect request body for debug
	if isDebug {
		d := request.GetDebug()
		d.Available = true
		d.RequestBody = content
	}

	// Create a clean request
	r, err := request.GetHTTPRequest(api)
	if err != nil {
		return
	}

	// Collect http request for debug
	if isDebug {
		d := request.GetDebug()
		d.Request = r
		d.RequestURL = r.URL
	}

	resp, err := api.HTTPClient.Do(r)

	if err != nil {
		return
	}

	// Collect http response for debug
	if isDebug {
		d := request.GetDebug()
		d.Response = resp
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Collect response body for debug
	if isDebug {
		d := request.GetDebug()
		d.ResponseBody = body
	}

	if resp.StatusCode == http.StatusNotFound {
		err = &APINotFoundError{}
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = errors.New(resp.Status)
		return
	}

	return
}

// DoPaginated will execute a given AnyDesk API request and return the plain json as string.
// In addition to the simple API.Do it will engrave pagination options into the request.
func (api *API) DoPaginated(request PaginatedAPIRequest) (body []byte, err error) {
	// Copy the pagination into the base query details
	p := request.GetPaginationOptions()
	base := request.GetRequestDetails()

	// Ensure the query parameters are available
	if base.Query == nil {
		base.Query = &url.Values{}
	}

	base.Query.Set("offset", strconv.FormatInt(p.Offset, 10))
	base.Query.Set("limit", strconv.FormatInt(p.Limit, 10))

	if p.Sort != "" {
		base.Query.Set("sort", p.Sort)
	}

	if p.Order != "" {
		base.Query.Set("order", string(p.Order))
	}

	body, err = api.Do(request)

	pagination := &PaginatedResult{}

	// Parse pagination info from request
	err = json.Unmarshal(body, pagination)
	if err != nil {
		return
	}

	if pagination.Selected == 0 {
		err = &APINoResultsError{}
		return
	}

	return
}

// BaseRequest contains the base information required to work against the API.
type BaseRequest struct {
	Method    string      `json:"-"`
	Resource  string      `json:"-"`
	Query     *url.Values `json:"-"`
	Timestamp int64       `json:"-"`
	Content   []byte      `json:"-"`

	debug *DebugInfo
}

// GetRequestDetails will return the base request details.
// Required by the APIRequest interface.
func (r *BaseRequest) GetRequestDetails() *BaseRequest {
	return r
}

// GetContentHash generates the content hash required for the API request string generated by GetRequestString().
func (r *BaseRequest) GetContentHash() string {
	h := sha1.New()
	h.Write(r.Content)

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// GetRequestString generates the request string required for the API token generated by GetRequestToken().
func (r *BaseRequest) GetRequestString() string {
	return fmt.Sprintf("%s\n%s\n%d\n%s", strings.ToUpper(r.Method), r.Resource, r.Timestamp, r.GetContentHash())
}

// GetHTTPRequest will return the prepared HTTP request that can be used by a http.Client
func (r *BaseRequest) GetHTTPRequest(api *API) (req *http.Request, err error) {
	endpoint := fmt.Sprintf("%s%s", api.APIEndpoint, r.Resource)

	req, err = http.NewRequest(r.Method, endpoint, bytes.NewBuffer(r.Content))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("AD %s:%d:%s", api.LicenseID, r.Timestamp, api.GetRequestToken(r)))

	return
}
