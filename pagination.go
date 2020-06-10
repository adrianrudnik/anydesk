package anydesk

type SortOrder string

const (
	OrderAsc  SortOrder = "asc"
	OrderDesc SortOrder = "desc"
)

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

func (po *PaginationOptions) GetPaginationOptions() *PaginationOptions {
	return po
}
