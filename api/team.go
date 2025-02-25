package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/orga/types"
)

/* --------------------------------- Team --------------------------------- */

func (s *Store) handleCreateTeam(w http.ResponseWriter, r *http.Request) error {
	team := new(types.Team)
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		return err
	}
	if err := s.CreateTeam(team); err != nil {
		return err
	}
	resp := "team created"
	log.Print("[api.guild] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleGetTeam(w http.ResponseWriter, r *http.Request) error {
	var team []*types.Team
	var err error
	if id := GetId(r); id == -1 {
		team, err = s.GetTeam()
	} else {
		team, err = s.GetTeamById(id)
	}
	if err != nil {
		return err
	}
	for _, t := range team {
		member, err := s.GetTeamMemberByTeamId(t.Id)
		if err != nil {
			log.Println("[api.team] no team members found for team with id ", t.Id)
		} else {
			for _, m := range member {
				role, err := s.GetTeamRoleByUserId(m.UserId)
				if err != nil {
					log.Println("[api.team] no role found for team member with user_id ", m.UserId)
				} else {
					user, err := s.GetUserById(m.UserId)
					if err != nil {
						log.Println("[api.team] no user found for team member with user_id ", m.UserId)
					} else {
						t.Member[role[0].Name] = append(t.Member[role[0].Name], user[0])
					}
				}
			}
		}
	}
	log.Print("[api.team] got teams")
	return WriteJSON(w, http.StatusOK, team)
}

func (s *Store) handleUpdateTeam(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	team := new(types.Team)
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		return err
	}
	if err := s.UpdateTeam(team, id); err != nil {
		return err
	}
	resp := fmt.Sprintf("team with id %v updated ", r.FormValue("id"))
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleDeleteTeam(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	if err := s.DeleteTeam(id); err != nil {
		return err
	}
	resp := "team deleted with id " + r.FormValue("id")
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

/* -------------------------------- Role --------------------------------- */

func (s *Store) handleCreateTeamRole(w http.ResponseWriter, r *http.Request) error {
	role := new(types.TeamRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.CreateTeamRole(role); err != nil {
		return err
	}
	resp := "role created"
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleGetTeamRole(w http.ResponseWriter, r *http.Request) error {
	role, err := s.GetTeamRole()
	if err != nil {
		return err
	}
	log.Print("[api.team] got roles")
	return WriteJSON(w, http.StatusOK, role)
}

func (s *Store) handleDeleteTeamRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	if err := s.DeletTeamRole(id); err != nil {
		return err
	}
	resp := "role deleted with id " + r.FormValue("id")
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleUpdateTeamRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	role := new(types.TeamRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.UpdateTeamRole(role, id); err != nil {
		return err
	}
	resp := "role updated with id " + r.FormValue("id")
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

/* --------------------------------- Member --------------------------------- */

func (s *Store) handleCreateTeamMember(w http.ResponseWriter, r *http.Request) error {
	member := new(types.TeamMember)
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		return err
	}
	if err := s.CreateTeamMember(member); err != nil {
		return err
	}
	resp := "member created"
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleDeleteTeamMember(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return fmt.Errorf("id not found")
	}
	if err := s.DeleteTeamMember(id); err != nil {
		return err
	}
	resp := "member deleted with id " + r.FormValue("id")
	log.Print("[api.team] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}
