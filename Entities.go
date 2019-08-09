package goose

import "net/http"

type GooseMiddleware func(*http.Request, *GooseMessage) (bool, *GooseResponse)
type GooseEndpoint func(*http.Request, *GooseMessage) interface{}

type GooseResponse struct {
	StatusCode int
	Response   interface{}
	Headers    map[string]string
}

/**
 * Goose route will hold all route related details. Like method, Url, Middlewares and Endpoint
 */
type GooseRoute struct {
	methods     []string
	uRL         string
	uRLRegex    string
	contentType string
	cors        map[string]string
	hasCors     bool
	middlewares []GooseMiddleware
	endpoint    GooseEndpoint
	hasDynamics bool
	dynamics    []string
}

type GooseMessage struct {
	RequestId      int64
	RequestTime    int64
	PostBody       string
	Holder         interface{}
	RequestHeaders map[string]string
	GetParams      map[string]string
	UrlParams      map[string]string
}

const (
	GET     = "GET"
	POST    = "POST"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"
	DELETE  = "DELETE"
)
const (
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_XML  = "application/xml"
	APPLICATION_JSON = "application/json"
	PLAIN_TEXT       = "text/plain"
	USER_AGENT       = "User-Agent"
	AGENT            = "golang Goose1.1"
)

const (
	METHOD_NOT_ALLOWED = "Method Not Allowed."
	PAGE_NOT_FOUND     = "Page not found."
)
