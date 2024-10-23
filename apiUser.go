package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetUser(w, r)
	case "POST":
		return s.handleCreateUser(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	err := WriteJSON(w, http.StatusOK, NewUser("john", "12345"))
	if err != nil {
		return err
	}

	return nil
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	log.Println("handleCreateUser_0")
	createUserReq := new(CreateUserRequest)
	err := json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		return err
	}
	user := NewUser(createUserReq.Name, createUserReq.DiscordID)
	if err := s.store.CreateUser(user); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, createUserReq)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
