package goose

import (
	"fmt"
	"net/http"
	"regexp"
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
		http.HandleFunc(
			routeURL,
			GooseHTTPHandler{
				route: gooseRoute,
			}.Pipeline,
		)
	}
	gHTTP.DefaultRoute()
	return gHTTP
}
func (gHTTP *GooseHTTP) DefaultRoute() *GooseHTTP {
	route := "/"
	http.HandleFunc(
		route,
		GooseHTTPHandler{
			route: GooseRoute{
				methods: []string{GET, POST, OPTIONS, PATCH, DELETE},
				uRL:     route,
			},
			routes: gHTTP.routes,
		}.RegexPipeline,
	)
	return gHTTP
}
func (gHTTP *GooseHTTP) Listen(address string) {
	fmt.Println(fmt.Sprintf("Listening at %s.", address))
	http.ListenAndServe(address, nil)
}

type GooseHTTPHandler struct {
	routes map[string]GooseRoute
	route  GooseRoute
}

/**
 * This function will be responsible for:
 * @1 Execution of middlewares
 * @2 Execution of final endpoint
 * @3 Preparing and emitting final response.
 */
func (gHTTPHandler GooseHTTPHandler) Pipeline(w http.ResponseWriter, req *http.Request) {
	validator := GooseValidator{}.GetInstance(gHTTPHandler.route, req)
	if validator.validate() {
		proceed, middlewareMessage := GooseMiddlewareExecutor{}.GetInstance(gHTTPHandler.route).Execute(req)
		if proceed {
			response, err := gHTTPHandler.route.endpoint(req, middlewareMessage.(*GooseMessage))
			if err == nil {
				GooseResponder{}.GetInstance(w).Respond(response)
			}
		}
	} else {
		GooseResponder{}.GetInstance(w).Respond(gHTTPHandler.buildCustomMessage())
	}
}
func (gHTTPHandler GooseHTTPHandler) RegexPipeline(w http.ResponseWriter, req *http.Request) {
	found, gooseRoute := gHTTPHandler.getRouteForRegexURI(req.URL.Path)
	validator := GooseValidator{}.GetInstance(gooseRoute, req)
	fmt.Println(validator, found)
}

func (gHTTPHandler GooseHTTPHandler) buildCustomMessage() interface{} {
	return map[string]interface{}{
		HEADERS: map[string]string{
			CONTENT_TYPE: PLAIN_TEXT,
		},
		RESPONSE:    METHOD_NOT_ALLOWED,
		STATUS_CODE: 405,
	}
}
func (gHTTPHandler GooseHTTPHandler) getRouteForRegexURI(url string) (bool, GooseRoute) {
	for _, route := range gHTTPHandler.routes {
		regex := regexp.MustCompile(route.uRLRegex)
		substrings := regex.FindAllStringSubmatch(url, 1)
		urlParams := map[string]string{}
		if len(substrings) >= 1 && len(substrings[0]) > 1 {
			for i := 1; i < len(substrings[0]); i++ {
				key := route.dynamics[i-1]
				value := substrings[0][i]
				urlParams[key] = value
			}
			route.urlParams = urlParams
			return true, route
		}
	}
	return false, GooseRoute{}
}
