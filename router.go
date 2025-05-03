package gohttprouter

import (
	"log"
	"net/http"
	"strings"
)

type Method int

const (
	GET Method = iota
	POST
	DELETE
	PATCH
	PUT
)

type Request struct {
	HttpReq *http.Request
	Prams   map[string]string
}

type HandlerFunc func(http.ResponseWriter, *Request)

type Route struct {
	method  Method
	handler HandlerFunc
}

type Router struct {
	notFound        http.HandlerFunc
	methodNotAllowd http.HandlerFunc

	routes map[string]Route
}

func New() *Router {
	return &Router{
		routes: make(map[string]Route),
		notFound: func(w http.ResponseWriter, r *http.Request) {
			log.Println("Path not found 404 ", r.URL.Path)
		},
	}
}

func (r *Router) NotFoundHandler(handler http.HandlerFunc) {
	r.notFound = handler
}

func (r *Router) Get(path string, handler HandlerFunc) {
	r.routes[path] = Route{
		method:  GET,
		handler: handler,
	}
}

func (r *Router) Post(path string, handler HandlerFunc) {
	r.routes[path] = Route{
		method:  POST,
		handler: handler,
	}
}

func (r *Router) Delete(path string, handler HandlerFunc) {
	r.routes[path] = Route{
		method:  DELETE,
		handler: handler,
	}
}

func (r *Router) Patch(path string, handler HandlerFunc) {
	r.routes[path] = Route{
		method:  PATCH,
		handler: handler,
	}
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	found := false
	urlSegs := strings.Split(req.URL.Path, "/")

	for path, route := range r.routes {
		pathSegs := strings.Split(path, "/")

		if len(pathSegs) != len(urlSegs) {
			continue
		}

		prams := map[string]string{}
		match := true
		for index, value := range pathSegs {
			if len(value) > 2 && value[0] == '{' && value[len(value)-1] == '}' {
				prams[value[1:len(value)-1]] = urlSegs[index]
				continue
			}
			if pathSegs[index] != urlSegs[index] {
				match = false
				break
			}
		}

		if match {
			route.handler(res, &Request{req, prams})
			found = true
		}
	}

	if !found {
		r.notFound(res, req)
	}
}
