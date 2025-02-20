package httpHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/orga/types"
)

func (s *Store) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	reqUser := new(types.User)
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		return err
	}
	if err := s.CreateUser(reqUser); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "user created")
}

func (s *Store) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	var userLst []*types.User
	var err error
	if id := GetId(r); id == -1 {
		userLst, err = s.GetUser()
	} else {
		userLst, err = s.GetUserById(id)
	}
	if err != nil {
		return err
	}
	for i := range userLst {
		lol, err := s.GetLeagueOfLegendsByUserId(userLst[i].Id)
		if err != nil {
			log.Println("[api.user] cant get league_of_legends for user with id ", userLst[i].Id)
		} else {
			accs, err := s.GetGameAccountByUserId(userLst[i].Id, "league_of_legends")
			if err != nil {
				return err
			}
			lol.Accounts = accs
		}
		userLst[i].LeagueOfLegends = lol
	}
	return WriteJSON(w, http.StatusOK, userLst)
}

func (s *Store) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	user := new(types.User)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}
	if err := s.UpdateUser(user, id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "user updated")
}

func (s *Store) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	if err := s.DeleteUser(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "user deleted")
}
