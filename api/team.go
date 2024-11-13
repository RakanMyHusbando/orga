package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RakanMyHusbando/shogun/types"
	"github.com/gorilla/mux"
)

// POST
func (s *APIServer) handleCreateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	reqGameAcc := new(types.ReqGameAccount)
	if err := json.NewDecoder(r.Body).Decode(&reqGameAcc); err != nil {
		return err
	}

	reqGameAcc.UserId = id

	if err := s.store.CreateGameAccount(reqGameAcc); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqGameAcc)
}

// DELETE
func (s *APIServer) handleDeleteGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	accountName := mux.Vars(r)["accountName"]
	if accountName == "" {
		return fmt.Errorf("account name is empty")
	}

	reqGameAcc := types.NewReqGameAccount(id, accountName, "")

	if err := s.store.DeleteGameAccount(reqGameAcc); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "game account deleted from user with id "+strconv.Itoa(id))
}

// PATCH
func (s *APIServer) handleUpdateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	reqGameAcc := new(types.ReqGameAccount)
	if err := json.NewDecoder(r.Body).Decode(&reqGameAcc); err != nil {
		return err
	}

	reqGameAcc.UserId = id

	accountName := mux.Vars(r)["accountName"]
	if accountName == "" {
		return fmt.Errorf("account name is empty")
	}

	if err := s.store.UpdateGameAccount(reqGameAcc, accountName); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqGameAcc)
}
