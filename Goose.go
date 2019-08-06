package goose

/**
 * Sahil Gulati {sahil.gulati1991@outlook.com}
 * Global settings for an HTTP service.
 */
type Goose struct {
	gooseRoutes   map[string]GooseRoute
	listenAddress string
	contextRoute  *GooseRoute
	holder        interface{}
}

func (g Goose) GetInstance() *Goose {
	goose := new(Goose)
	goose.gooseRoutes = make(map[string]GooseRoute)
	return goose
}
func (g *Goose) WithHolding(holder interface{}) *Goose {
	g.holder = holder
	return g
}
func (g *Goose) Route(methods []string, routePath string) *Goose {
	gooseRoute := GooseRoute{
		methods:  methods,
		uRL:      routePath,
		haveCors: false,
		holder:   g.holder,
		dynamics: getDynamics(routePath),
	}
	if len(gooseRoute.dynamics) > 0 {
		gooseRoute.hasDynamics = true
		gooseRoute.uRLRegex = convertDyanmicURLToRegex(gooseRoute.uRL)
	}
	g.contextRoute = &gooseRoute
	g.gooseRoutes[routePath] = gooseRoute
	return g
}
func (g *Goose) Middlewares(middlewares ...GooseMiddleware) *Goose {
	contextGooseRoute, isset := g.gooseRoutes[g.contextRoute.uRL]
	if isset {
		contextGooseRoute.middlewares = make([]GooseMiddleware, len(middlewares))
		contextGooseRoute.middlewares = middlewares
		g.gooseRoutes[g.contextRoute.uRL] = contextGooseRoute
	}
	return g
}
func (g *Goose) Endpoint(endpoint GooseEndpoint) *Goose {
	contextGooseRoute, isset := g.gooseRoutes[g.contextRoute.uRL]
	if isset {
		contextGooseRoute.endpoint = endpoint
		g.gooseRoutes[g.contextRoute.uRL] = contextGooseRoute
	}
	return g
}
func (g *Goose) Register() {
	g.contextRoute = nil
}
func (g *Goose) Serve(address string) *Goose {
	g.listenAddress = address
	GooseHTTP{}.GetInstance(g.gooseRoutes).Register().Listen(address)
	return g
}
