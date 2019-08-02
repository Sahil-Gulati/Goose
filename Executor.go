package goose

/**
 * Sahil Gulati {sahil.gulati1991@outlook.com}
 * This file contains functionality of execution of middlewares.
 **/
import (
	"fmt"
	"net/http"
	"time"
)

type GooseMiddlewareExecutor struct {
	route GooseRoute
}

func (gME GooseMiddlewareExecutor) GetInstance(gooseRoute GooseRoute) *GooseMiddlewareExecutor {
	gooseMiddlewareExecutor := new(GooseMiddlewareExecutor)
	gooseMiddlewareExecutor.route = gooseRoute
	return gooseMiddlewareExecutor
}

/**
 * This function will be response for execution of middleware.
 */
func (gME *GooseMiddlewareExecutor) Execute(req *http.Request) (bool, interface{}) {
	var err error
	proceed := true
	gooseMessage := gME.getMessage()
	for _, middleware := range gME.route.middlewares {
		proceed, err = middleware(req, gooseMessage)
		if !proceed {
			fmt.Println(err)
		}
	}
	return proceed, gooseMessage
}

/**
 * This function will create a new gooseMessage it
 *@1 Adding requestId
 *@2 Adding requestTime
 *@3 Previous provided user holder
 */
func (gME *GooseMiddlewareExecutor) getMessage() *GooseMessage {
	requestTime := time.Now().UnixNano()
	return &GooseMessage{
		RequestId:   requestTime,
		RequestTime: requestTime,
		Holder:      gME.route.holder,
	}
}
