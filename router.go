package r4r

type Router struct {
	URL    string
	routes map[string]*Route
}

func NewRouter() *Router {
	return &Router{routes: make(map[string]*Route)}
}

func (r *Router) addRoute(URL string, withMethod func(r *Route)) {
	_, has := r.routes[URL]
	if !has {
		route := NewRoute(URL)
		withMethod(route)
		r.routes[URL] = route
	} else {
		withMethod(r.routes[URL])
	}
}

func (r *Router) Get(URL string, middlewares []Middleware, handler Handler) {
	r.addRoute(URL, func(r *Route) {
		r.Get(wrapWithMiddlewares(handler, middlewares))
	})
}

func (r *Router) Put(URL string, middlewares []Middleware, handler Handler) {
	r.addRoute(URL, func(r *Route) {
		r.Put(wrapWithMiddlewares(handler, middlewares))
	})
}

func (r *Router) Post(URL string, middlewares []Middleware, handler Handler) {
	r.addRoute(URL, func(r *Route) {
		r.Post(wrapWithMiddlewares(handler, middlewares))
	})
}

func (r *Router) Delete(URL string, middlewares []Middleware, handler Handler) {
	r.addRoute(URL, func(r *Route) {
		r.Delete(wrapWithMiddlewares(handler, middlewares))
	})
}

func (r *Router) Head(URL string, middlewares []Middleware, handler Handler) {
	r.addRoute(URL, func(r *Route) {
		r.Head(wrapWithMiddlewares(handler, middlewares))
	})
}

func (r *Router) Trace(URL string, middlewares []Middleware, handler Handler) {
	r.addRoute(URL, func(r *Route) {
		r.Trace(wrapWithMiddlewares(handler, middlewares))
	})
}

func (r *Router) Option(URL string, middlewares []Middleware, handler Handler) {
	r.addRoute(URL, func(r *Route) {
		r.Option(wrapWithMiddlewares(handler, middlewares))
	})
}
