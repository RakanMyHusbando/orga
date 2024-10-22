package main

import (
	"db-api/req_handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", req_handler.DefaultHandler)
	r.HandleFunc("/user", req_handler.UserGetHandler).Methods("GET")
	r.HandleFunc("/user", req_handler.UserPostHandler).Methods("Post")
}
