package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/user/{id}", makeHTTPHandleFunc(s.handleUser))

	log.Println("API server running on ", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}

/* =================== API User Handlers =================== */

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	var err error = fmt.Errorf("unsupported method: %s", r.Method)
	id := mux.Vars(r)["id"]

	switch r.Method {
	case "GET":
		if id != "" {
			err = s.handleGetUserById(w, r)
		} else {
			err = s.handleGetUser(w, r)
		}
	case "POST":
		err = s.handleCreateUser(w, r)
	case "DELETE":
		if id != "" {
			err = s.handleDeleteUser(w, r)
		}
	case "PUT":
		err = s.handleUpdateUser(w, r)
	}

	return err
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	userList, err := s.store.GetUser()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, userList)
}

func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	stringId := mux.Vars(r)["id"]
	if stringId == "" {
		return fmt.Errorf("id is required")
	}

	intId, err := strconv.Atoi(stringId)
	if err != nil {
		return err
	}

	userList, err := s.store.GetUserById(intId)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, userList)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserReq := new(CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&createUserReq); err != nil {
		return err
	}

	if err := s.store.CreateUser(NewUser(createUserReq.Name, createUserReq.DiscordID)); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, createUserReq)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

/* =================== API team handlers =================== */

// TODO
