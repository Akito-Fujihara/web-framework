package framework

import (
	"errors"
	"net/http"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{},
	}
}

type Router struct {
	RoutingTable map[string]func(http.ResponseWriter, *http.Request)
}

func (r *Router) Get(pathname string, handler func(http.ResponseWriter, *http.Request)) error {
	if r.RoutingTable == nil {
		r.RoutingTable = make(map[string]func(http.ResponseWriter, *http.Request))
	}

	if r.RoutingTable[pathname] != nil {
		return errors.New("pathname is already registered")
	}

	r.RoutingTable[pathname] = handler
	return nil
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handler := e.Router.RoutingTable[r.URL.Path]
		if handler == nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		handler(rw, r)
		return
	}
}

func (e *Engine) Run() {
	http.ListenAndServe(":8080", e)
}
