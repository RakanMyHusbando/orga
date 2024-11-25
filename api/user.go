package api

import (
	"encoding/json"
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
	var userLst []*types.User
	id, err := GetId(r)
	if err != nil {
		userLst, err = s.store.GetUser()
	} else {
		userLst, err = s.store.GetUserById(id)
	}
	if err != nil {
		return err
	}
	for i := range userLst {
		lol, err := s.store.GetLeagueOfLegendsByUserId(userLst[i].Id)
		if err != nil {
			log.Println("[api.user] cant get league_of_legends for user with id ", userLst[i].Id)
		} else {
			accs, err := s.store.GetGameAccountBy(userLst[i].Id, "league_of_legends")
			if err != nil {
				return err
			}
			lol.Accounts = accs
		}
		userLst[i].LeagueOfLegends = lol
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
	return WriteJSON(w, http.StatusOK, "user deleted")
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
