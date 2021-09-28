package r4r

import "net/http"

func Use(URL string, middlewares []Middleware, router *Router) {
	router.URL = URL
	for _, route := range router.routes {
		(func(route *Route) {
			http.HandleFunc(router.URL+route.URL, wrapWithMiddlewares(route.ToHandler(), middlewares))
		})(route)
	}
}
