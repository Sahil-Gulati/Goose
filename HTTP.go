package goose

import (
	"fmt"
	"net/http"
)

/**
 * Sahil Gulati {sahil.gulati1991@outlook.com}
 * This structure will be responsible for registering routes with HTTP handlers.
 */
type GooseHTTP struct {
	routes map[string]GooseRoute
}

func (gHTTP GooseHTTP) GetInstance(routes map[string]GooseRoute) *GooseHTTP {
	gh := new(GooseHTTP)
	gh.routes = routes
	return gh
}

/**
 * This function will register all GooseRoutes to net/http in which
 * each request will be first forwarded to GooseHTTPHandler.Pipeline
 */
func (gHTTP *GooseHTTP) Register() *GooseHTTP {
	for routeURL, gooseRoute := range gHTTP.routes {
		fmt.Println(routeURL, gooseRoute)
		http.HandleFunc(
			routeURL,
			GooseHTTPHandler{
				route: gooseRoute,
			}.Pipeline,
		)
	}
	return gHTTP
}
func (gHTTP *GooseHTTP) Listen(address string) {
	fmt.Println(fmt.Sprintf("Listening at %s.", address))
	http.ListenAndServe(address, nil)
}

type GooseHTTPHandler struct {
	route GooseRoute
}

/**
 * This function will be responsible for:
 * @1 Execution of middlewares
 * @2 Execution of final endpoint
 * @3 Preparing and emitting final response.
 */
func (gHTTPHandler GooseHTTPHandler) Pipeline(w http.ResponseWriter, req *http.Request) {
	proceed, middlewareMessage := GooseMiddlewareExecutor{}.GetInstance(gHTTPHandler.route).Execute(req)
	if proceed {
		response, err := gHTTPHandler.route.endpoint(req, middlewareMessage.(*GooseMessage))
		if err == nil {
			GooseResponder{}.GetInstance(w).Respond(response)
		}
	}
}
