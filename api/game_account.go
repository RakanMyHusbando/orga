package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RakanMyHusbando/shogun/types"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleCreateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	account := new(types.GameAccount)
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		return err
	}
	account.UserId = id
	if err := s.store.CreateGameAccount(account); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "game account created")
}

func (s *APIServer) handleUpdateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
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
	return WriteJSON(w, http.StatusOK, "game account updated")
}

func (s *APIServer) handleDeleteGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	name := mux.Vars(r)["accountName"]
	if name == "" {
		return fmt.Errorf("account name is empty")
	}
	if err := s.store.DeleteGameAccount(id, name); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "game account deleted")
}
