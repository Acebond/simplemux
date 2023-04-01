package simplemux

import (
	"net/http"
)

type Mux struct {
	NotFound http.Handler
	routes   []Route
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func (mux *Mux) HandlerFunc(method, path string, handler http.HandlerFunc) {
	route := Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
	mux.routes = append(mux.routes, route)
}

// ServeHTTP makes the router implement the http.Handler interface.
func (mux *Mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range mux.routes {
		if route.Method == req.Method && route.Path == req.URL.Path {
			route.Handler.ServeHTTP(w, req)
			return
		}
	}
	mux.NotFound.ServeHTTP(w, req)
}

// Make sure the Router conforms with the http.Handler interface
var _ http.Handler = (*Mux)(nil)
