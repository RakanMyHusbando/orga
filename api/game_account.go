package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/orga/types"
)

func (s *Store) handleCreateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	account := new(types.GameAccount)
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		return err
	}
	account.UserId = id
	if err := s.CreateGameAccount(account); err != nil {
		return err
	}
	respMessage := "game account created for user with id " + r.FormValue("id")
	log.Print("[api.game_account] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *Store) handleDeleteGameAccount(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	name := r.FormValue("accountName")
	if name == "" {
		return fmt.Errorf("account name is empty")
	}
	if err := s.DeleteGameAccount(id, name); err != nil {
		return err
	}
	respMessage := "game account deleted for user with id " + r.FormValue("id")
	log.Println("[api.game_account] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}
