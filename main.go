package main

import (
	"db-api/req_handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.().Set("Content-Type", "application/json")
	r.HandleFunc("/", req_handler.DefaultHandler)
	r.HandleFunc("/user", req_handler.UserGetHandler).Methods("GET")
	r.HandleFunc("/user", req_handler.UserPostHandler).Methods("Post")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
