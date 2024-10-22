package main

import (
	"db-api/req_handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", req_handler.UserGetHandler).Methods("GET")
	// router.HandleFunc("/user", req_handler.UserPostHandler).Methods("Post")

	http.Handle("/", router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
