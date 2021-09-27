package r4r

import "net/http"

type Route struct {
	URL     string
	methods map[string]Handler
}

func NewRoute(URL string) *Route {
	return &Route{URL: URL, methods: make(map[string]Handler)}
}

func (r *Route) ToHandler() Handler {
	return func(writer http.ResponseWriter, request *http.Request) {
		handler, has := r.methods[request.Method]
		if has {
			handler(writer, request)
		}
	}
}

func (r *Route) FindHandlerByMethod(method string) Handler {
	return r.methods[method]
}

func (r *Route) Get(handler Handler) {
	r.methods[GET] = handler
}

func (r *Route) Put(handler Handler) {
	r.methods[PUT] = handler
}

func (r *Route) Post(handler Handler) {
	r.methods[POST] = handler
}

func (r *Route) Delete(handler Handler) {
	r.methods[DELETE] = handler
}

func (r *Route) Head(handler Handler) {
	r.methods[HEAD] = handler
}

func (r *Route) Trace(handler Handler) {
	r.methods[TRACE] = handler
}

func (r *Route) Option(handler Handler) {
	r.methods[OPTION] = handler
}
