package anydesk

// SessionCommentChangeRequest is used to patch the /session/{id} API resource.
type SessionCommentChangeRequest struct {
	*BaseRequest
	Comment *string `json:"string"`
}

// Do will execute the "/auth" query against the given API.
func (req *SessionCommentChangeRequest) Do(api *Api) (err error) {
	_, err = api.Do(req)
	return
}
