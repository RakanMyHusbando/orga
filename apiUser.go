package main

import (
	"fmt"
	"net/http"
	"strings"
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
	var lol *LeagueOfLegends = nil

	queryParam := r.URL.Query().Get("property")

	if queryParam != "" {
		properties := strings.Split(queryParam, ",")
		for _, property := range properties {
			switch property {
			case "LeagueOfLegends":
				lol := NewLeagueOfLegends("top", "jungle", []string{"a", "b"}, []string{"a", "b"})
			}
		}
	}

	err := WriteJSON(w, http.StatusOK, NewUser(1, "john", lol))
	if err != nil {
		return err
	}

	return nil
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
