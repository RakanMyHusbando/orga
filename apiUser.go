package main

import (
	"fmt"
	"log"
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
	var lol *LeagueOfLegends

	queryProperty := r.URL.Query().Get("property")

	if queryProperty != "" {
		properties := strings.Split(queryProperty, ",")
		for _, property := range properties {
			switch property {
			case "league_of_legends":
				lol = NewLeagueOfLegends("top", "jungle", []string{"a", "b"}, []string{"a", "b"})
			}
		}
	}

	var search map[string]string

	queryName := r.URL.Query().Get("name")
	queryDiscordId := r.URL.Query().Get("discord_id")

	if queryName != "" {
		search = map[string]string{"name": queryName}
	} else if queryDiscordId != "" {
		search = map[string]string{"discord_id": queryDiscordId}
	}

	log.Println(search)

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
