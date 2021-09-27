package r4r

import "net/http"

func WithEmptyMiddlewares() []Middleware {
	return []Middleware{}
}

func WithMiddlewares(middlewares ...Middleware) []Middleware {
	return middlewares
}

func everyMiddlewares(w http.ResponseWriter, r *http.Request, middlewares []Middleware) error {
	for _, middleware := range middlewares {
		if err := middleware(w, r); err != nil {
			return err
		}
	}

	return nil
}

func wrapWithMiddlewares(handler Handler, middlewares []Middleware) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := everyMiddlewares(w, r, middlewares); err != nil {
			return
		}

		handler(w, r)
	}
}
