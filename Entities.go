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
	contentType string
	cors        map[string]string
	haveCors    bool
	middlewares []GooseMiddleware
	endpoint    GooseEndpoint
	holder      interface{}
}

type GooseMessage struct {
	RequestId   int64
	RequestTime int64
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
	OPTIONS          = "OPTIONS"
	DELETE           = "DELETE"
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_XML  = "application/xml"
	APPLICATION_JSON = "application/json"
	USER_AGENT       = "User-Agent"
	AGENT            = "golang Goose1.1"
)
