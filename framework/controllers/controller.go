package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StudentResponse struct {
	Name string `json:"name"`
}

func StudentsController(rw http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	name := queries.Get("name")

	sutudentResponse := StudentResponse{Name: name}

	responseData, err := json.Marshal(sutudentResponse)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(responseData)
}

func ListsController(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "lists")
}

func UsersController(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "users")
}
