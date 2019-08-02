package goose

/**
 * Sahil Gulati {sahil.gulati1991@outlook.com}
 * This file contains functionality of execution of middlewares.
 **/
import (
	"fmt"
	"net/http"
)

type GooseMiddleware func(*http.Request, *map[string]interface{}) (bool, error)

type GooseMiddlewareExecutor struct {
	middlewares []GooseMiddleware
}

func (gME GooseMiddlewareExecutor) GetInstance(middlewares []GooseMiddleware) *GooseMiddlewareExecutor {
	gooseMiddlewareExecutor := new(GooseMiddlewareExecutor)
	gooseMiddlewareExecutor.middlewares = middlewares
	return gooseMiddlewareExecutor
}
/**
 * This function will be response for execution of middleware.
 */
func (gME *GooseMiddlewareExecutor) Execute(req *http.Request) (bool, interface{}) {
	var err error
	var proceed bool
	var custom map[string]interface{} = make(map[string]interface{})
	for _, middleware := range gME.middlewares {
		proceed, err = middleware(req, &custom)
		if !proceed {
			fmt.Println(err)
		}
	}
	return proceed, &custom
}
/**
 * 
 */