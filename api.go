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
	Error string `json:"error"`
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

func GetId(r *http.Request) (int, error) {
	var intId int
	if strId := mux.Vars(r)["id"]; strId != "" {
		res, err := strconv.Atoi(strId)
		if err != nil {
			return intId, err
		}
		intId = res
		return intId, nil
	}
	return intId, nil
}

/* =================== API User Handlers =================== */

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	var err error = fmt.Errorf("unsupported method: %s", r.Method)
	id := mux.Vars(r)["id"]

	switch r.Method {
	case "POST":
		err = s.handleCreateUser(w, r)
	case "GET":
		if id != "" {
			err = s.handleGetUserById(w, r)
		} else {
			err = s.handleGetUser(w, r)
		}
	case "DELETE":
		if id != "" {
			err = s.handleDeleteUser(w, r)
		}
	case "PUT":
		err = s.handleUpdateUser(w, r)
	}

	return err
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUser := new(CreateUser)
	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		return err
	}

	if err := s.store.CreateUser(createUser); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, createUser)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	userList, err := s.store.GetUser()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, userList)
}

func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	user, err := s.store.GetUserById(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.DeletUser(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "User deleted")
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	user := new(User)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	if err := s.store.UpdateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

/* =================== API team handlers =================== */

// TODO
