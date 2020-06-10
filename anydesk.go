// Package anydesk provides an API client towards the AnyDesk REST API
// that is available with professional and enterprise licenses.
//
// Debugging
//
// If you need to see specifics about the API requests made by this library
// you can enable the debug mode and request request and response information
// directly from results.
//
//   api := NewApi("license", "password")
//   r := NewAuthenticationRequest()
//
//   SetDebug(true)
//   r.Do(api)
//
//   fmt.Printf(
//       "Url: %s, Response body: %s",
//       r.GetDebug().RequestUrl,
//       r.GetDebug().ResponseBody,
//   )
//
// Pagination
//
// There are several ways to use the paginated requests:
//
//  request := anydesk.NewSessionListRequest(nil)
//
//  // assign directly with
//  request.Limit = 10
//  request.Offset = 100
//
//  // assign everything at once with
//  request.PaginationOptions = &anydesk.PaginationOptions{
//      Offset: 0,
//      Limit:  0,
//      Sort:   "",
//      Order:  "",
//  }
package anydesk
