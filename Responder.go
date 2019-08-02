package goose

/**
 * Sahil Gulati {sahil.gulati1991@outlook.com}
 * This file will contains functionality of preparing full and final response to be emitted out.
 **/
import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type GooseResponder struct {
	writer     http.ResponseWriter
	headers    map[string]string
	statusCode int
}

func (gr GooseResponder) GetInstance(writer http.ResponseWriter) *GooseResponder {
	gooseResponder := new(GooseResponder)
	gooseResponder.writer = writer
	gooseResponder.headers = make(map[string]string)
	return gooseResponder
}
func (gr *GooseResponder) Respond(response interface{}) error {
	value, asserted := response.(map[string]interface{})
	if !asserted {
		return gr.
			prepareHeaders(nil).writeHeaders().
			prepareStatusCode(nil).writeStatusCode().
			emitResponse(response)
	}
	return gr.
		prepareHeaders(value[HEADERS]).writeHeaders().
		prepareStatusCode(value[STATUS_CODE]).writeStatusCode().
		emitResponse(value[RESPONSE])
}

func (gr *GooseResponder) prepareHeaders(contextHeaders interface{}) *GooseResponder {
	gr.headers = make(map[string]string)
	gr.headers[USER_AGENT] = AGENT
	gr.headers[CONTENT_TYPE] = APPLICATION_JSON
	if value, isAsserted := contextHeaders.(map[string]string); isAsserted {
		for headerName, headerValue := range value {
			gr.headers[headerName] = headerValue
		}
	}
	return gr
}
func (gr *GooseResponder) writeHeaders() *GooseResponder {
	for headerName, headerValue := range gr.headers {
		gr.writer.Header().Add(headerName, headerValue)
	}
	return gr
}
func (gr *GooseResponder) prepareStatusCode(statusCode interface{}) *GooseResponder {
	if value, isAsserted := statusCode.(int); isAsserted {
		gr.statusCode = value
	}
	gr.statusCode = 200
	return gr
}
func (gr *GooseResponder) writeStatusCode() *GooseResponder {
	gr.writer.WriteHeader(gr.statusCode)
	return gr
}

func (gr *GooseResponder) emitResponse(response interface{}) error {
	switch gr.headers[CONTENT_TYPE] {
	case APPLICATION_JSON:
		return json.NewEncoder(gr.writer).Encode(response)
	case APPLICATION_XML:
		gr.writer.Write([]byte(xml.Header))
		encoder := xml.NewEncoder(gr.writer)
		encoder.Indent(" ", " ")
		return encoder.Encode(response.(string))
	default:
		return nil
	}
}
