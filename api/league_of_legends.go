package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RakanMyHusbando/orga/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleCreateLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	lol := new(types.LeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&lol); err != nil {
		return err
	}
	if err := s.store.CreateLeagueOfLeagends(lol, id); err != nil {
		return err
	}
	respMessage := "[api.league_of_leagends] league_of_legends added to user with id " + mux.Vars(r)["id"]
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleDeleteLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	s.store.DeleteLeagueOfLegends(id)
	respMessage := "[api.league_of_legends] league_of_legends deleted from user with id " + mux.Vars(r)["id"]
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleUpdateLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	lol := new(types.LeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&lol); err != nil {
		return err
	}
	if err := s.store.UpdateLeagueOfLegends(lol, id); err != nil {
		return err
	}
	respMessage := "[api.league_of_legends] league_of_legends updated for user with id " + mux.Vars(r)["id"]
	return WriteJSON(w, http.StatusOK, respMessage)
}
