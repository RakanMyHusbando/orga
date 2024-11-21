package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RakanMyHusbando/shogun/types"
)

// POST
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	reqUser := new(types.ReqUser)
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		return err
	}

	if err := s.store.CreateUser(reqUser); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "user created")
}

// GET
func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	userList, err := s.store.GetUser()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, userList)
}

// GET
func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	user, err := s.store.GetUserById(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

// DELETE
func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.Delete("User", map[string]any{"id": id}); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintf("user with id %v deleted", id))
}

// PATCH
func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	user := new(types.ResUser)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	user.Id = id

	if err := s.store.UpdateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}
