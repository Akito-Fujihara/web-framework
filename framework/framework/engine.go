package framework

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
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
	pathname = strings.TrimSuffix(pathname, "/")
	existedHandler := r.RoutingTable.Search(pathname)

	if existedHandler != nil {
		panic("already exists")
	}

	r.RoutingTable.Insert(pathname, handler)
	return nil
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pathname := path.Clean(r.URL.Path)
		pathname = strings.TrimSuffix(pathname, "/")
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
	fmt.Println("Server is running on localhost:8085")
	if err := http.ListenAndServe(":8085", e); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
