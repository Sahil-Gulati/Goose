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
    Goose := goose.Goose{}.GetInstance()
    
    Goose.Route([]string{goose.GET}, "/health").Endpoint(endpoint).Register()
    Goose.RegexRoute([]string{goose.GET}, "/api/v1/transaction/{transactionId}").Endpoint(endpoint).Register()
    
    Goose.Serve(":8080")
}
var endpoint goose.GooseEndpoint = func(request *http.Request, gooseMessage *goose.GooseMessage) (interface{}, error) {
    fmt.Println("Executing actions --->" + gooseMessage.UrlParams["transactionId"])
    return map[string]string{"message": "Hello World!"}, nil
}


```
