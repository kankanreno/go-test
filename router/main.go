package main

import (
	"fmt"
	"net/http"
	"strings"
)

// 路由结构体，包含一个记录方法、路径的 map
type Router struct {
	Route map[string]map[string]http.HandlerFunc
}

// 返回一个 Router 实例
func NewRouter() *Router {
	return new(Router)
}

// 实现 Handler 接口，匹配方法以及路径
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if h, ok := r.Route[req.Method][req.URL.String()]; ok {
		h(w, req)
	}
}

// 根据方法、路径将方法注册到路由
func (r *Router) HandleFunc(method, path string, f http.HandlerFunc) {
	if r.Route == nil {
		r.Route = make(map[string]map[string]http.HandlerFunc)
	}

	method = strings.ToUpper(method)
	if r.Route[method] == nil {
		r.Route[method] = make(map[string]http.HandlerFunc)
	}

	r.Route[method][path] = f
}

func main() {
	r := NewRouter()
	r.HandleFunc("GET", "/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "FOO!")
	})

	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ROOT!")
	})

	http.ListenAndServe(":8080", r)
}
