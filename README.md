# AnyDesk REST API client

Unoffical REST API client, written in Go.

[![license](https://img.shields.io/github/license/adrianrudnik/anydesk)](https://github.com/adrianrudnik/anydesk/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/adrianrudnik/anydesk-api?status.svg)](https://godoc.org/adrianrudnik/anydesk-api)
[![go report card](https://goreportcard.com/badge/github.com/adrianrudnik/anydesk-api)](https://goreportcard.com/report/github.com/adrianrudnik/anydesk)

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
