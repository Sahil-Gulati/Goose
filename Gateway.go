package goose

import (
	// "fmt"
	"net/http"
	"time"
)

type GooseGateway struct {
	routes 		map[string]*GooseRoute
	regexRoutes map[string]*GooseRoute
	route  		*GooseRoute
	holder 		interface{}
}

/**
 * This function will be responsible for:
 * @1 Getting default variables
 * @1 Execution of middlewares
 * @2 Execution of final endpoint
 * @3 Preparing and emitting final response.
 */
func (gateway GooseGateway) Pipeline(w http.ResponseWriter, request *http.Request) {
	/** @1 **/
	message, defaultResponse := gateway.getDefaults(request)

	/** @2 **/
	toProcess := GooseValidator{
		route:    gateway.route,
		request:  request,
		response: defaultResponse,
	}.GetInstance().validate()

	/** @3 **/
	toProcess = GooseMiddlewareExecutor{
		toProcess: toProcess,
		route:     gateway.route,
		request:   request,
		message:   message,
		response:  defaultResponse,
	}.GetInstance().Execute()

	/** @4 **/
	GooseMiddlewareExecutor{
		toProcess: toProcess,
		route:     gateway.route,
		request:   request,
		message:   message,
		response:  defaultResponse,
	}.GetInstance().Action()

	/** @5 **/
	GooseResponder{
		writer:   w,
		route:    gateway.route,
		response: defaultResponse,
	}.GetInstance().Respond()
}
func (gateway GooseGateway) RegexPipeline(w http.ResponseWriter, request *http.Request) {
	/** @1 **/
	message, defaultResponse := gateway.getDefaults(request)
	/** @2 **/
	toProcess, route := GooseValidator{
		regexRoutes: gateway.regexRoutes,
		request:  request,
		message: message,
		response: defaultResponse,
	}.GetInstance().regexValidate()

	/** @3 **/
	toProcess = GooseMiddlewareExecutor{
		toProcess: toProcess,
		route:     route,
		request:   request,
		message:   message,
		response:  defaultResponse,
	}.GetInstance().Execute()

	/** @4 **/
	GooseMiddlewareExecutor{
		toProcess: toProcess,
		route:     route,
		request:   request,
		message:   message,
		response:  defaultResponse,
	}.GetInstance().Action()
	
	/** @5 **/
	GooseResponder{
		writer:   w,
		route:    route,
		response: defaultResponse,
	}.GetInstance().Respond()
}

/**
 * This function will create a new gooseMessage it
 *@1 Adding requestId
 *@2 Adding requestTime
 *@3 Previous provided user holder
 */

func (gateway GooseGateway) getDefaults(request *http.Request) (*GooseMessage, *GooseResponse) {
	requestTime := time.Now().UnixNano()
	return &GooseMessage{
			RequestId:   requestTime,
			RequestTime: requestTime,
			Holder:      gateway.holder,
			RequestHeaders: getHeaders(request),
			GetParams:		getParams(request),
		},
		&GooseResponse{
			Headers: map[string]string{
				CONTENT_TYPE: PLAIN_TEXT,
			},
			StatusCode: 200,
		}
}
