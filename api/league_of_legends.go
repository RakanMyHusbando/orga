package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RakanMyHusbando/shogun/types"
)

// POST
func (s *APIServer) handleCreateLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	lol := new(types.LeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&lol); err != nil {
		return err
	}

	lol.UserId = id

	if err := s.store.CreateLeagueOfLeagends(lol); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, lol)
}

// DELETE
func (s *APIServer) handleDeleteLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.Delete("UserLeagueOfLegends", map[string]any{"user_id": id}); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "league_of_legends deleted from user with id "+strconv.Itoa(id))
}

// PATCH
func (s *APIServer) handleUpdateLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	reqUserLol := new(types.ReqLeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&reqUserLol); err != nil {
		return err
	}

	reqUserLol.UserId = id

	if err := s.store.UpdateLeagueOfLegends(reqUserLol); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqUserLol)
}
