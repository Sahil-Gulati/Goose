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
	writer   http.ResponseWriter
	route    *GooseRoute
	response *GooseResponse
}

func (gr GooseResponder) GetInstance() *GooseResponder {
	gooseResponder := new(GooseResponder)
	gooseResponder.route = gr.route
	gooseResponder.writer = gr.writer
	gooseResponder.response = gr.response
	return gooseResponder
}
func (gr *GooseResponder) Respond() error {
	gr.writeHeaders(
		gr.prepareHeaders(),
	)
	return gr.writeResponse(
		gr.prepareResponse(),
	)
}

func (gr *GooseResponder) prepareHeaders() (map[string]string, int) {
	gr.response.Headers = ifElse(len(gr.response.Headers) == 0, make(map[string]string), gr.response.Headers).(map[string]string)
	gr.response.StatusCode = ifElse(gr.response.StatusCode == 0, 200, gr.response.StatusCode).(int)
	gr.response.Headers[CONTENT_TYPE] = ifElse(gr.response.Headers[CONTENT_TYPE] != "", gr.response.Headers[CONTENT_TYPE], PLAIN_TEXT).(string)
	if gr.route.hasCors {
		gr.response.Headers["Access-Control-Allow-Origin"] = "*"
		gr.response.Headers["Access-Control-Allow-Methods"] = "GET, POST, PUT"
		gr.response.Headers["Access-Control-Allow-Headers"] = "Content-Type"
	}
	return gr.response.Headers, gr.response.StatusCode
}
func (gr *GooseResponder) writeHeaders(headers map[string]string, statusCode int) *GooseResponder {
	for headerName, headerValue := range headers {
		gr.writer.Header().Add(headerName, headerValue)
	}
	gr.writer.WriteHeader(statusCode)
	return gr
}
func (gr *GooseResponder) prepareResponse() *GooseResponse {
	return gr.response
}

func (gr *GooseResponder) writeResponse(response *GooseResponse) error {
	switch response.Headers[CONTENT_TYPE] {
	case APPLICATION_JSON:
		return json.NewEncoder(gr.writer).Encode(response.Response)
	case APPLICATION_XML:
		gr.writer.Write([]byte(xml.Header))
		encoder := xml.NewEncoder(gr.writer)
		encoder.Indent(" ", " ")
		return encoder.Encode(response.Response)
	default:
		gr.writer.Write([]byte(response.Response.(string)))
		return nil
	}
}
