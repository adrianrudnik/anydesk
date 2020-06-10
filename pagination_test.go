package anydesk

func ExampleNewPaginationOptions() {
	api := NewApi("license", "password")
	request := NewSessionListRequest(nil)

	// change default values
	request.Offset = 100
	request.Limit = 10

	// change by clean default values
	options := NewPaginationOptions()
	options.Offset = 20
	options.Order = OrderAsc
	request.PaginationOptions = options

	// change by struct
	request.PaginationOptions = &PaginationOptions{
		Offset: 0,
		Limit:  0,
		Sort:   "",
		Order:  "",
	}

	request.Do(api)
}
