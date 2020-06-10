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
