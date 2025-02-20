package httpHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/orga/types"
)

func (s *Store) handleCreateDiscord(w http.ResponseWriter, r *http.Request) error {
	discord := new(types.Discord)
	if err := json.NewDecoder(r.Body).Decode(&discord); err != nil {
		return err
	}
	if err := s.CreateDiscord(discord); err != nil {
		return err
	}
	resp := "discord server created"
	log.Print("[api.discord] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleGetDiscord(w http.ResponseWriter, r *http.Request) error {
	var discord []*types.Discord
	var err error
	if id := GetId(r); id == -1 {
		discord, err = s.GetDiscord()
	} else {
		discord, err = s.GetDiscordById(id)
	}
	if err != nil {
		return err
	}
	for _, d := range discord {
		member, err := s.GetDiscordMemberByServerId(d.Id)
		if err != nil {
			log.Println("[api.discord] no member found for discord with id ", d.Id)
		} else {
			for _, m := range member {
				role, err := s.GetDiscordRoleById(m.RoleId)
				if err != nil {
					log.Println("[api.discord] no role found with id ", m.RoleId)
				} else {
					user, err := s.GetUserById(m.UserId)
					if err != nil {
						log.Println("[api.discord] no user found with id ", m.UserId)
					}
					d.Member[role[0].Name] = append(d.Member[role[0].Name], user[0])
				}
			}
		}
	}
	log.Print("[api.discord] got discord servers")
	return WriteJSON(w, http.StatusOK, discord)
}

func (s *Store) handleUpdateDiscord(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	discord := new(types.Discord)
	if err := json.NewDecoder(r.Body).Decode(&discord); err != nil {
		return err
	}
	if err := s.UpdateDiscord(discord, GetId(r)); err != nil {
		return err
	}
	resp := "discord server updated"
	log.Print("[api.discord] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleDeleteDiscord(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.DeleteDiscord(id); err != nil {
		return err
	}
	resp := "discord server deleted"
	log.Print("[api.discord] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

/* ------------------------------ handle role ------------------------------ */

func (s *Store) handleCreateDiscordRole(w http.ResponseWriter, r *http.Request) error {
	role := new(types.DiscordRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.CreateDiscordRole(role); err != nil {
		return err
	}
	resp := "discord role created"
	log.Print("[api.discord.role] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleGetDiscordRole(w http.ResponseWriter, r *http.Request) error {
	var role []*types.DiscordRole
	var err error
	if id := GetId(r); id == -1 {
		role, err = s.GetDiscordRole()
	} else {
		role, err = s.GetDiscordRoleById(id)
	}
	if err != nil {
		return err
	}
	log.Print("[api.discord.role] got discord roles")
	return WriteJSON(w, http.StatusOK, role)
}

func (s *Store) handleUpdateDiscordRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	role := new(types.DiscordRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.UpdateDiscordRole(role, id); err != nil {
		return err
	}
	resp := "discord role updated"
	log.Print("[api.discord.role] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleDeleteDiscordRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.DeleteDiscordRole(id); err != nil {
		return err
	}
	resp := "discord role deleted"
	log.Print("[api.discord.role] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

/* ------------------------------ handle member ------------------------------ */

func (s *Store) handleCreateDiscordMember(w http.ResponseWriter, r *http.Request) error {
	member := new(types.DiscordMember)
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		return err
	}
	if err := s.CreateDiscordMember(member); err != nil {
		return err
	}
	resp := "discord member created"
	log.Print("[api.discord.member] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *Store) handleDeleteDiscordMember(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.DeleteDiscordMember(id); err != nil {
		return err
	}
	resp := fmt.Sprintf("discord member with id %v deleted", id)
	log.Print("[api.discord.member] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}
