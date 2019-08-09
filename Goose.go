package goose

/**
 * Sahil Gulati {sahil.gulati1991@outlook.com}
 * Global settings for an HTTP service.
 */
type Goose struct {
	routes        map[string]*GooseRoute
	regexRoutes   map[string]*GooseRoute
	listenAddress string
	contextRoute  *GooseRoute
	holder        interface{}
}

func (g Goose) GetInstance() *Goose {
	goose := new(Goose)
	goose.routes = make(map[string]*GooseRoute)
	goose.regexRoutes = make(map[string]*GooseRoute)
	return goose
}
func (g *Goose) WithHolding(holder interface{}) *Goose {
	g.holder = holder
	return g
}
func (g *Goose) AddCors() *Goose {
	g.contextRoute.hasCors = true
	return g
}
func (g *Goose) Route(methods []string, routePath string) *Goose {
	gooseRoute := &GooseRoute{
		uRL:     routePath,
		methods: methods,
	}
	g.contextRoute = gooseRoute
	g.routes[routePath] = gooseRoute
	return g
}
func (g *Goose) RegexRoute(methods []string, routePath string) *Goose {
	gooseRoute := &GooseRoute{
		uRL:         routePath,
		methods:     methods,
		dynamics:    getDynamics(routePath),
		hasDynamics: true,
		uRLRegex:    convertDyanmicURLToRegex(routePath),
	}
	g.contextRoute = gooseRoute
	g.regexRoutes[routePath] = gooseRoute
	return g
}
func (g *Goose) Middlewares(middlewares ...GooseMiddleware) *Goose {
	g.checkContextRoute()
	for _, middleware := range middlewares {
		g.contextRoute.middlewares = append(g.contextRoute.middlewares, middleware)
	}
	return g
}
func (g *Goose) Endpoint(endpoint GooseEndpoint) *Goose {
	g.checkContextRoute()
	g.contextRoute.endpoint = endpoint
	return g
}
func (g *Goose) Register() {
	g.contextRoute = nil
}
func (g *Goose) Serve(address string) *Goose {
	g.listenAddress = address
	GooseHTTP{
		routes:      g.routes,
		regexRoutes: g.regexRoutes,
		holder: 	 g.holder,
	}.GetInstance().Register().Listen(address)
	return g
}
func (g *Goose) checkContextRoute() {
	if g.contextRoute == nil {
		panic("Must register context route first!")
	}
}
