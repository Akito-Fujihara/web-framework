package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "users")
// 	})

// 	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "list")
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/user" {
		fmt.Fprintf(w, "users")
		return
	}
	if r.URL.Path == "/list" {
		fmt.Fprintf(w, "list")
		return
	}
	fmt.Fprintf(w, "Not found")
}

func main() {
	handler := &Handler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	server.ListenAndServe()
}
