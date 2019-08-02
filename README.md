# Goose
GoLang framework for building REST API. It requires a bare minimum code to start a server by offering fast development of REST based APIs.


`HelloWorld.go`
```go
package main

import (
	"fmt"
	"net/http"
	goose "github.com/Sahil-Gulati/Goose"
)

func main() {
    goose := goose.Goose{}.GetInstance()
    goose.Route([]string{goose.GET}, "/health").Endpoint(endpoint).Register()
    goose.Serve(":8080")
}
var endpoint goose.GooseEndpoint = func(request *http.Request, gooseMessage *goose.GooseMessage) (interface{}, error) {
    fmt.Println("Executing actions --->")
    return map[string]string{"message": "Hello World!"}, nil
}


```
