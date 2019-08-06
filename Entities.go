package goose

import "net/http"

type GooseMiddleware func(*http.Request, *GooseMessage) (bool, error)
type GooseEndpoint func(*http.Request, *GooseMessage) (interface{}, error)

/**
 * Goose route will hold all route related details. Like method, Url, Middlewares and Endpoint
 */
type GooseRoute struct {
	methods     []string
	uRL         string
	uRLRegex    string
	contentType string
	cors        map[string]string
	haveCors    bool
	middlewares []GooseMiddleware
	endpoint    GooseEndpoint
	holder      interface{}
	hasDynamics bool
	dynamics    []string
	urlParams   map[string]string
}

type GooseMessage struct {
	RequestId   int64
	RequestTime int64
	GetParams   map[string]string
	Holder      interface{}
}

const (
	HOLDER        = "XXX_HOLDER"
	LOGGER        = "XXX_LOGGER"
	HEADERS       = "XXX_HEADERS"
	RESPONSE      = "XXX_RESPONSE"
	STATUS_CODE   = "XXX_STATUS_CODE"
	RESPONSE_TYPE = "XXX_RESPONSE_TYPE"

	GET              = "GET"
	POST             = "POST"
	PATCH            = "PATCH"
	OPTIONS          = "OPTIONS"
	DELETE           = "DELETE"
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_XML  = "application/xml"
	APPLICATION_JSON = "application/json"
	PLAIN_TEXT       = "text/plain"
	USER_AGENT       = "User-Agent"
	AGENT            = "golang Goose1.1"

	METHOD_NOT_ALLOWED = "Method Not Allowed."
)
