package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	reqUser := new(types.User)
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		return err
	}
	if err := s.store.CreateUser(reqUser); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "user created")
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	var getUser []*map[string]any
	id, err := GetId(r)
	if err != nil {
		getUser, err = s.store.GetUser(nil)
	} else {
		getUser, err = s.store.GetUserById(id)
	}
	if err != nil {
		return err
	}
	var userLst []*types.User
	for _, user := range getUser {
		newUser := types.NewUser(
			(*user)["name"].(string),
			(*user)["discord_id"].(string),
			(*user)["id"].(int),
			nil,
		)
		getLol, err := s.store.GetLeagueOfLegendsByUserId((*user)["id"].(int))
		if err != nil {
			log.Println(err)
		} else {
			getAccs, err := s.store.GetGameAccountBy(
				types.NewGameAccount((*user)["id"].(int), "", "league_of_legends"),
			)
			if err != nil {
				return err
			}
			LeagueOfLegendsMapToStruct(getLol[0], getAccs)
		}
		userLst = append(userLst, newUser)
	}
	return WriteJSON(w, http.StatusOK, userLst)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteUser(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, fmt.Sprintf("[api.user] user deleted"))
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	user := new(types.User)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}
	user.Id = id
	if err := s.store.UpdateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}
