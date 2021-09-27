package r4r

type Router struct {
	URL    string
	routes map[string]*Route
}

func NewRouter() *Router {
	return &Router{routes: make(map[string]*Route)}
}

func (r *Router) hasRoute(URL string) bool {
	_, has := r.routes[URL]
	return has
}

func (r *Router) Get(URL string, middlewares []Middleware, handler Handler) {
	if r.hasRoute(URL) {
		r.routes[URL].Get(wrapWithMiddlewares(handler, middlewares))
	} else {
		route := NewRoute(URL)
		route.Get(wrapWithMiddlewares(handler, middlewares))
		r.routes[URL] = route
	}
}

func (r *Router) Put(URL string, middlewares []Middleware, handler Handler) {
	if r.hasRoute(URL) {
		r.routes[URL].Get(wrapWithMiddlewares(handler, middlewares))
	} else {
		route := NewRoute(URL)
		route.Get(wrapWithMiddlewares(handler, middlewares))
		r.routes[URL] = route
	}
}

func (r *Router) Post(URL string, middlewares []Middleware, handler Handler) {
	if r.hasRoute(URL) {
		r.routes[URL].Get(wrapWithMiddlewares(handler, middlewares))
	} else {
		route := NewRoute(URL)
		route.Get(wrapWithMiddlewares(handler, middlewares))
		r.routes[URL] = route
	}
}

func (r *Router) Delete(URL string, middlewares []Middleware, handler Handler) {
	if r.hasRoute(URL) {
		r.routes[URL].Get(wrapWithMiddlewares(handler, middlewares))
	} else {
		route := NewRoute(URL)
		route.Get(wrapWithMiddlewares(handler, middlewares))
		r.routes[URL] = route
	}
}

func (r *Router) Head(URL string, middlewares []Middleware, handler Handler) {
	if r.hasRoute(URL) {
		r.routes[URL].Get(wrapWithMiddlewares(handler, middlewares))
	} else {
		route := NewRoute(URL)
		route.Get(wrapWithMiddlewares(handler, middlewares))
		r.routes[URL] = route
	}
}

func (r *Router) Trace(URL string, middlewares []Middleware, handler Handler) {
	if r.hasRoute(URL) {
		r.routes[URL].Get(wrapWithMiddlewares(handler, middlewares))
	} else {
		route := NewRoute(URL)
		route.Get(wrapWithMiddlewares(handler, middlewares))
		r.routes[URL] = route
	}
}

func (r *Router) Option(URL string, middlewares []Middleware, handler Handler) {
	if r.hasRoute(URL) {
		r.routes[URL].Get(wrapWithMiddlewares(handler, middlewares))
	} else {
		route := NewRoute(URL)
		route.Get(wrapWithMiddlewares(handler, middlewares))
		r.routes[URL] = route
	}
}
