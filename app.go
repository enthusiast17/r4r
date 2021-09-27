package r4r

import "net/http"

func Use(URL string, middlewares []Middleware, router *Router) {
	for routeURL, route := range router.routes {
		(func(routeURL string, route *Route) {
			http.HandleFunc(URL+routeURL, wrapWithMiddlewares(route.ToHandler(), middlewares))
		})(routeURL, route)
	}
}
