package anydesk

// APINotFoundError will be thrown when a API request could not find any specifc data
type APINotFoundError struct {
}

func (e *APINotFoundError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return "not found"
}

// APINoResultsError will be thrown when a list API request delivered no results.
type APINoResultsError struct{}

func (e *APINoResultsError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return "not results"
}
