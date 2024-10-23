package main

import (
	"fmt"
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
	lol := NewLeagueOfLegends("top", "jungle", []string{"a", "b"}, []string{"a", "b"})
	user := NewUser(1, "john", lol)
	WriteJSON(w, http.StatusOK, user)
	return nil
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
