package anydesk

import (
	"net/http"
	"net/url"
)

var (
	isDebug = false
)

// SetDebug switches the API request / response debug information collection to the given value.
// After enabling the debug mode, raw details about requests and their response can be queried:
func SetDebug(enable bool) {
	isDebug = enable
}

// DebugInfo contains a set of raw details about the http transaction as sent
// and received by the AnyDesk API.
// If unsure, take a look at DebugInfo.Available, which indicates if any
// debug information was collected.
type DebugInfo struct {
	// Is "true" when debug info was populated, "false" when no info was collected.
	Available bool

	// The full request URL sent by the http request.
	RequestURL *url.URL

	// The http.Request used for the request.
	Request *http.Request

	// The plain request body sent to the API.
	RequestBody []byte

	// The http.Response received by the http request.
	Response *http.Response

	// The plain response body received by the API.
	ResponseBody []byte
}

func newDebugInfo() *DebugInfo {
	return &DebugInfo{
		Available:    false,
		RequestURL:   nil,
		RequestBody:  nil,
		Response:     nil,
		ResponseBody: nil,
	}
}

// GetDebug returns the the collected information of the request.
// The debug mode must be enabled prior:
//   anydesk.SetDebug(true)
func (r *BaseRequest) GetDebug() *DebugInfo {
	if r.debug == nil {
		r.debug = newDebugInfo()
	}
	return r.debug
}
