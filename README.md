# anydesk-api

Unofficial go based API client library

## Usage

```go
package main

import (
	"fmt"
	"github.com/adrianrudnik/anydesk-api/pkg/anydesk"
	"net/http"
	"os"
	"time"
)

func main() {
    api := anydesk.NewApi(os.Getenv("LICENSE_ID"), os.Getenv("API_PASSWORD"))
	api.HttpClient = &http.Client{Timeout: 5 * time.Second}

	request := anydesk.NewAuthenticationRequest()
	response, _ := request.Do(api)

	fmt.Printf("Status: %s, License: %s", response.Result, response.LicenseId)
}
```
