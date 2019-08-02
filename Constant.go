package goose

type GooseMessage struct {
	requestId   string
	requestTime string
	holder      interface{}
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
