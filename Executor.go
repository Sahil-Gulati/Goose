package goose

/**
 * Sahil Gulati {sahil.gulati1991@outlook.com}
 * This file contains functionality of execution of middlewares.
 **/
import (
	"net/http"
)

type GooseMiddlewareExecutor struct {
	toProcess bool
	route     *GooseRoute
	message   *GooseMessage
	request   *http.Request
	response  *GooseResponse
}

func (gME GooseMiddlewareExecutor) GetInstance() *GooseMiddlewareExecutor {
	gooseMiddlewareExecutor := new(GooseMiddlewareExecutor)
	gooseMiddlewareExecutor.route = gME.route
	gooseMiddlewareExecutor.message = gME.message
	gooseMiddlewareExecutor.request = gME.request
	gooseMiddlewareExecutor.response = gME.response
	gooseMiddlewareExecutor.toProcess = gME.toProcess
	return gooseMiddlewareExecutor
}

/**
 * This function will be response for execution of middleware.
 */
func (gME *GooseMiddlewareExecutor) Execute() bool {
	if gME.toProcess {
		for _, middleware := range gME.route.middlewares {
			proceed, response := middleware(gME.request, gME.message)
			if !proceed {
				gME.toProcess = proceed
				gME.response.Headers = response.Headers
				gME.response.Response = response.Response
				gME.response.StatusCode = response.StatusCode
				return false
			}
		}
	}
	return true
}

/**
 * This function will be responsible for execution of final action.
 */
func (gME *GooseMiddlewareExecutor) Action() {
	if gME.toProcess {
		response := gME.route.endpoint(gME.request, gME.message)
		if value, asserted := response.(GooseResponse); asserted {
			gME.response.Headers = value.Headers
			gME.response.Response = value.Response
			gME.response.StatusCode = value.StatusCode
		} else if value, asserted := response.(*GooseResponse); asserted {
			gME.response.Headers = value.Headers
			gME.response.Response = value.Response
			gME.response.StatusCode = value.StatusCode
		} else {
			gME.response.Response = response
		}
	}
}
