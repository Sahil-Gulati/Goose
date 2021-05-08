# Goose
GoLang framework for building REST API. It requires a bare minimum code to start a server by offering fast development of REST based APIs.

### Usage
`go get -v -u https://github.com/Sahil-Gulati/Goose`

### Sample code
```go
func main() {
    Goose := goose.Goose{}.GetInstance()
    Goose.Route([]string{goose.GET}, "/api/v1/transactions/").Middlewares(m1, m2).Endpoint(endpoint).Register()
    Goose.Serve(":8080")
}
```

### Middleware
##### Middleware signature
```go
type GooseMiddleware func(*http.Request, *GooseMessage) (bool, *GooseResponse)
```

### Endpoint
##### Endpoint signature
```go
type GooseEndpoint func(*http.Request, *GooseMessage) interface{}
```

### Message
##### Message signature
```go
type GooseMessage struct {
	RequestId      int64
	RequestTime    int64
	PostBody       string
	Holder         interface{}
	RequestHeaders map[string]string
	GetParams      map[string]string
	UrlParams      map[string]string
}
```
