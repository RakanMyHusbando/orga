package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/shogun/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleCreateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	account := new(types.GameAccount)
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		return err
	}
	account.UserId = id
	if err := s.store.CreateGameAccount(account); err != nil {
		return err
	}
	respMessage := "game account created for user with id " + mux.Vars(r)["id"]
	log.Print("[api.game_account] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleUpdateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	oldName := mux.Vars(r)["accountName"]
	if oldName == "" {
		return fmt.Errorf("account name is empty")
	}
	var newName map[string]string
	if err := json.NewDecoder(r.Body).Decode(&newName); err != nil {
		return err
	}
	if err := s.store.UpdateGameAccount(id, oldName, newName["name"]); err != nil {
		return err
	}
	respMessage := "game account updated for user with id " + mux.Vars(r)["id"]
	log.Print("[api.game_account] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleDeleteGameAccount(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	name := mux.Vars(r)["accountName"]
	if name == "" {
		return fmt.Errorf("account name is empty")
	}
	if err := s.store.DeleteGameAccount(id, name); err != nil {
		return err
	}
	respMessage := "game account deleted for user with id " + mux.Vars(r)["id"]
	log.Print("[api.game_account] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}
