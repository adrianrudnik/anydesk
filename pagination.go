package anydesk

// SortOrder defines that ordering you want on a paginated API request.
type SortOrder string

const (
	// OrderAsc indicates that you want to receive the results ascending order.
	OrderAsc SortOrder = "asc"

	// OrderDesc indicates that you want to receive the results in descencing order.
	OrderDesc SortOrder = "desc"
)

// PaginationOptions contain all configurable settings for the pagination of API requests.
type PaginationOptions struct {
	// Result offset, starting at 0
	Offset int64 `json:"-"`

	// Result limit, use anydesk.Infinite for unlimited results
	Limit int64 `json:"-"`

	// Result sort by property name
	Sort string `json:"-"`

	// Result sort order, use anydesk.OrderAsc or anydesk.OrderDesc
	Order SortOrder `json:"-"`
}

// NewPaginationOptions returns the default pagination options used by the AnyDesk API.
func NewPaginationOptions() *PaginationOptions {
	return &PaginationOptions{
		Offset: 0,
		Limit:  Infinite,
		Sort:   "",
		Order:  OrderDesc,
	}
}

// GetPaginationOptions returns the currently configured pagination settings.
func (po *PaginationOptions) GetPaginationOptions() *PaginationOptions {
	return po
}

// PaginatedResult contains the current pagination settings as returned by the API.
type PaginatedResult struct {
	// Total result count
	Count int64 `json:"count"`

	// Unknown
	Selected int64 `json:"selected"`

	// Applied offset of the current result
	Offset int64 `json:"offset"`

	// Applied limit of the current result
	Limit int64 `json:"limit"`
}

// HasMore will indicate if more results could be fetched and also return
// a possible version of the next pages page options.
func (pr *PaginatedResult) HasMore(request PaginatedAPIRequest) (*PaginationOptions, bool) {
	if pr.Count <= pr.Offset+pr.Selected {
		// no more results expected
		return nil, false
	}

	prev := request.GetPaginationOptions()

	return &PaginationOptions{
		Offset: pr.Offset + pr.Limit,
		Limit:  pr.Limit,
		Sort:   prev.Sort,
		Order:  prev.Order,
	}, true
}
