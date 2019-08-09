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
	routes      map[string]*GooseRoute
	regexRoutes map[string]*GooseRoute
	holder      interface{}
}

func (gHTTP GooseHTTP) GetInstance() *GooseHTTP {
	gh := new(GooseHTTP)
	gh.routes = gHTTP.routes
	gh.holder = gHTTP.holder
	gh.regexRoutes = gHTTP.regexRoutes
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
			GooseGateway{
				route:  gooseRoute,
				holder: gHTTP.holder,
			}.Pipeline,
		)
	}
	gHTTP.DefaultRoute()
	return gHTTP
}

/**
 * This function will register all Regex GooseRoutes to net/http in which
 * each request will be first forwarded to GooseHTTPHandler.RegexPipeline
 */
func (gHTTP *GooseHTTP) DefaultRoute() *GooseHTTP {
	route := "/"
	http.HandleFunc(
		route,
		GooseGateway{
			regexRoutes: gHTTP.regexRoutes,
			holder:      gHTTP.holder,
		}.RegexPipeline,
	)
	return gHTTP
}
func (gHTTP *GooseHTTP) Listen(address string) {
	fmt.Println(fmt.Sprintf("Listening at %s.", address))
	http.ListenAndServe(address, nil)
}
