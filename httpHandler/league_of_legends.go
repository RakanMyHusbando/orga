package httpHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RakanMyHusbando/orga/types"
	"github.com/gorilla/mux"
)

func (s *Store) handleCreateLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	lol := new(types.LeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&lol); err != nil {
		return err
	}
	if err := s.CreateLeagueOfLeagends(lol, id); err != nil {
		return err
	}
	respMessage := "[api.league_of_leagends] league_of_legends added to user with id " + mux.Vars(r)["id"]
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *Store) handleDeleteLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	s.DeleteLeagueOfLegends(id)
	respMessage := "[api.league_of_legends] league_of_legends deleted from user with id " + mux.Vars(r)["id"]
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *Store) handleUpdateLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	lol := new(types.LeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&lol); err != nil {
		return err
	}
	if err := s.UpdateLeagueOfLegends(lol, id); err != nil {
		return err
	}
	respMessage := "[api.league_of_legends] league_of_legends updated for user with id " + mux.Vars(r)["id"]
	return WriteJSON(w, http.StatusOK, respMessage)
}
