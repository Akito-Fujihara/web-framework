package framework

import (
	"net/http"
	"path"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{
			RoutingTable: NewTreeNode(),
		},
	}
}

type Router struct {
	RoutingTable TreeNode
}

func (r *Router) Get(pathname string, handler func(http.ResponseWriter, *http.Request)) error {
	r.RoutingTable.Insert(pathname, handler)
	return nil
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pathname := path.Clean(r.URL.Path)
		handler := e.Router.RoutingTable.Search(pathname)

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
