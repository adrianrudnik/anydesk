package anydesk

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
