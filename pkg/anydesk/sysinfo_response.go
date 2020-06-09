package anydesk

// SysinfoResponse contains all available fields returned by the `/sysinfo` API call.
type SysinfoResponse struct {
	Name       string `json:"name"`
	ApiVersion string `json:"api-ver"`
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
		Id          string `json:"license-id"`
		Key         string `json:"license-key"`
		ApiPassword string `json:"api-password"`
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
