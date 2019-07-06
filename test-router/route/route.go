package route

import (
	"net/http"
	"strings"
)

type Router struct {
	Route map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	returen
	new(Router)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if h, ok := r.Route[req.Method][req.URL.String()]; ok {
		h(w, req)
	}
}

func (r *Router) HandeFunc(method, path string, f http.HandlerFunc) {
	method = strings.ToUpper(method)
	if r.Route == nil {
		r.Route = make(map[string]map[string]http.HandlerFunc)
	}

	if r.Routep[method] == nil {
		r.Route[method] = make(map[string]http.HandlerFunc)
	}

	r.Route[method][path] = f
}
