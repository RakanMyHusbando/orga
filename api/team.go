package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/shogun/types"
	"github.com/gorilla/mux"
)

/* --------------------------------- handler team --------------------------------- */

func (s *APIServer) handleCreateTeam(w http.ResponseWriter, r *http.Request) error {
	team := new(types.Team)
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		return err
	}
	if err := s.store.CreateTeam(team); err != nil {
		return err
	}
	resp := "guild created"
	log.Print("[api.guild] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleGetTeam(w http.ResponseWriter, r *http.Request) error {
	var team []*types.Team
	var err error
	if id := GetId(r); id == -1 {
		team, err = s.store.GetTeam()
	} else {
		team, err = s.store.GetTeamById(id)
	}
	if err != nil {
		return err
	}
	log.Print("[api.team] got teams")
	return WriteJSON(w, http.StatusOK, team)
}

func (s *APIServer) handleUpdateTeam(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	team := new(types.Team)
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		return err
	}
	if err := s.store.UpdateTeam(team, id); err != nil {
		return err
	}
	resp := "team updated with id " + mux.Vars(r)["id"]
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleDeleteTeam(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	if err := s.store.DeleteTeam(id); err != nil {
		return err
	}
	resp := "team deleted with id " + mux.Vars(r)["id"]
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

/* --------------------------------- handler team role --------------------------------- */

func (s *APIServer) handleCreateTeamRole(w http.ResponseWriter, r *http.Request) error {
	role := new(types.TeamRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.store.CreateTeamRole(role); err != nil {
		return err
	}
	resp := "role created"
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleGetTeamRole(w http.ResponseWriter, r *http.Request) error {
	role, err := s.store.GetTeamRole()
	if err != nil {
		return err
	}
	log.Print("[api.team] got roles")
	return WriteJSON(w, http.StatusOK, role)
}

func (s *APIServer) handleUpdateTeamRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	role := new(types.TeamRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.store.UpdateTeamRole(role, id); err != nil {
		return err
	}
	resp := "role updated with id " + mux.Vars(r)["id"]
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleDeleteTeamRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	if err := s.store.DeletTeamRole(id); err != nil {
		return err
	}
	resp := "role deleted with id " + mux.Vars(r)["id"]
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

/* --------------------------------- handler team member --------------------------------- */

func (s *APIServer) handleCreateTeamMember(w http.ResponseWriter, r *http.Request) error {
	member := new(types.TeamMember)
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		return err
	}
	if err := s.store.CreateTeamMember(member); err != nil {
		return err
	}
	resp := "member created"
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleGetTeamMember(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	member, err := s.store.GetTeamMemberByUserId(id)
	if err != nil {
		return err
	}
	log.Print("[api.team] got members")
	return WriteJSON(w, http.StatusOK, member)
}

func (s *APIServer) handleDeleteTeamMember(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	if err := s.store.DeleteTeamMember(id); err != nil {
		return err
	}
	resp := "member deleted with id " + mux.Vars(r)["id"]
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleUpdateTeamMember(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	member := new(types.TeamMember)
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		return err
	}
	if err := s.store.UpdateTeamMember(member, id); err != nil {
		return err

	}
	resp := "member updated with id " + mux.Vars(r)["id"]
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}