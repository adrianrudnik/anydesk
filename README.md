# AnyDesk REST API client

Unoffical REST API client, written in Go.

[![license](https://img.shields.io/github/license/adrianrudnik/anydesk)](https://github.com/adrianrudnik/anydesk/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/adrianrudnik/anydesk?status.svg)](https://pkg.go.dev/github.com/adrianrudnik/anydesk?tab=doc)
[![go report card](https://goreportcard.com/badge/github.com/adrianrudnik/anydesk)](https://goreportcard.com/report/github.com/adrianrudnik/anydesk)

- [Installation](#installation)
- [Usage](#usage)
- [Requests](#requests)
  - [Authentication request](#authentication-request)
  - [System information](#system-information)
  - [Client list](#client-list)
  - [Client details](#client-details)

## Installation

To use the package inside your own project, do:

```bash
go get github.com/github.com/adrianrudnik/anydesk
```

To confirm everything runs fine:

```shell script
go test github.com/adrianrudnik/anydesk
go vet github.com/adrianrudnik/anydesk
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/adrianrudnik/anydesk"
	"net/http"
	"os"
	"time"
)

func main() {
	api := anydesk.NewAPI(os.Getenv("LICENSE_ID"), os.Getenv("API_PASSWORD"))

	// Optional: Decrease timeouts
	api.HTTPClient = &http.Client{Timeout: 5 * time.Second}

	// Optional: Switch to Enterprise API
	api.APIEndpoint = "https://yourinstance:8081"

	request := anydesk.NewAuthenticationRequest()
	response, _ := request.Do(api)

	fmt.Printf("Status: %s, License: %s", response.Result, response.LicenseID)
}
```

## Requests

The following requests are avaible with this package.

### Authentication request

To test the service authentication with the given credentials you can do a request to
the `/auth` REST API resource:

```go
api := NewAPI("license", "password")
request := NewAuthenticationRequest()

response, err := request.Do(api)
if err != nil {
    log.Fatal(err)
}

if response.Result != "success" {
    log.Fatalf("response failed: %s %s", response.Code, response.Error)
}
```

### System information

Request general license and system information from the `/sysinfo` REST-API resource:

```go
api := NewAPI("license", "password")

request := NewSysinfoRequest()
response, _ := request.Do(api)

fmt.Printf("API: %s, Max Session: %d, Active Sessions: %d",
    response.APIVersion,
    response.License.MaxSessions,
    response.Sessions.Active,
    // ...
)
```

### Client list

A list of all clients associated to the license can be requested through `/clients`.

Note: The list itself does not contain information about the last sessions each client had.

```go
api := NewAPI("license", "password")

// Define optional search parameters
search := &ClientListSearch{
    Online: true,
}

request := NewClientListRequest(search)

// Define optional pagination settings
request.Offset = 10
request.Limit = 5

response, _ := request.Do(api)

// Access the resulting pagination information
fmt.Printf(
    "Found: %d, Offset: %d, Selected: %d",
    response.Count,
    response.Offset,
    response.Selected,
)

// Iterate through the result set
if len(response.List) == 0 {
    return
}

for _, client := range response.List {
    fmt.Printf(
        "ID: %d, Alias: %s, Version: %s",
        client.ClientID,
        client.Alias,
        client.ClientVersion,
    )
}

// If you used result pagination, you could prepare the next request
if options, hasMore := response.HasMore(request); hasMore {
    // prepare next request
    request = NewClientListRequest(search)

    // options will contain the previous pagination options with shifted offse
    // to help you retrieve the next result set with the same settings.
    request.PaginationOptions = options
    response, _ = request.Do(api)
}
```

### Client details

To retrieve more detailed information about a given specific client ID:

```go
api := NewAPI("license", "password")

request := NewClientDetailRequest(123456789)
response, _ := request.Do(api)

fmt.Printf(
    "ID: %s, Alias: %s, Version: %s",
    response.ClientID,
    response.Alias,
    response.ClientVersion,
)

for _, client := range response.LastSessions {
    fmt.Printf(
        "Start: %s, %d seconds total, Comment: %s",
        client.StartTime().Format(time.RFC3339),
        client.DurationInSeconds,
        client.Comment,
    )
}
```
