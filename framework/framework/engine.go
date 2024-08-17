package framework

import (
	"net/http"

	"github.com/Akito-Fujihara/framework/controllers"
)

type Engine struct{}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/students" {
			controllers.StudentsController(rw, r)
			return
		}

		if r.URL.Path == "/lists" {
			controllers.ListsController(rw, r)
			return
		}

		if r.URL.Path == "/users" {
			controllers.UsersController(rw, r)
			return
		}
	}
}

func (e *Engine) Run() {
	http.ListenAndServe(":8080", e)
}
