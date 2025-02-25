package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/orga/types"
)

/* --------------------------------- handler guild --------------------------------- */

func (s *Store) handleCreateGuild(w http.ResponseWriter, r *http.Request) error {
	guild := new(types.Guild)
	if err := json.NewDecoder(r.Body).Decode(&guild); err != nil {
		return err
	}
	if err := s.CreateGuild(guild); err != nil {
		return err
	}
	respMessage := "guild created"
	log.Print("[api.guild] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *Store) handleGetGuild(w http.ResponseWriter, r *http.Request) error {
	var guild []*types.Guild
	var err error
	if id := GetId(r); id == -1 {
		guild, err = s.GetGuild()
	} else {
		guild, err = s.GetGuildById(id)
	}
	if err != nil {
		return err
	}
	for _, g := range guild {
		member, err := s.GetGuildMemberByGuildId(g.Id)
		if err != nil {
			log.Println("[api.guild] no team members found for guild with id ", g.Id)
		} else {
			for _, m := range member {
				role, err := s.GetGuildRoleById(m.RoleId)
				if err != nil {
					log.Println("[api.guild] no role found for guild member with user_id ", m.UserId)
				} else {
					user, err := s.GetUserById(m.UserId)
					if err != nil {
						log.Println("[api.guild] no user found for guild member with user_id ", m.UserId)
					} else {
						g.Member[role[0].Name] = append(g.Member[role[0].Name], user[0])
					}
				}
			}
		}
	}
	log.Print("[api.guild] got guilds")
	return WriteJSON(w, http.StatusOK, guild)
}

func (s *Store) handleUpdateGuild(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	guild := new(types.Guild)
	if err := json.NewDecoder(r.Body).Decode(&guild); err != nil {
		return err
	}
	if err := s.UpdateGuild(guild, id); err != nil {
		return err
	}
	respMessage := "guild updated with id " + r.FormValue("id")
	log.Print("[api.guild] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *Store) handleDeleteGuild(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.DeleteGuild(id); err != nil {
		return err
	}
	respMessage := "guild deleted with id " + r.FormValue("id")
	log.Print("[api.guild] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

/* --------------------------------- handler guild role --------------------------------- */

func (s *Store) handleCreateGuildRole(w http.ResponseWriter, r *http.Request) error {
	role := new(types.GuildRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.CreateGuildRole(role); err != nil {
		return err
	}
	respMessage := "guild role created"
	log.Print("[api.guild.role] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *Store) handleGetGuildRole(w http.ResponseWriter, r *http.Request) error {
	role, err := s.GetGuildRole()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, role)
}

func (s *Store) handleUpdateGuildRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	role := new(types.GuildRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.UpdateGuildRole(role, id); err != nil {
		return err
	}
	respMessage := fmt.Sprintf("guild role with id %v updated", id)
	log.Print("[api.guild.role] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *Store) handleDeleteGuildRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.DeleteGuildRole(id); err != nil {
		return err
	}
	respMessage := fmt.Sprintf("guild role with id %v deleted", id)
	log.Print("[api.guild.role] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

/* --------------------------------- handler guild member --------------------------------- */

func (s *Store) handleCreateGuildMember(w http.ResponseWriter, r *http.Request) error {
	member := new(types.GuildMember)
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		return err
	}
	if err := s.CreateGuildMember(member); err != nil {
		return err
	}
	respMessage := "guild member created"
	log.Print("[api.guild.member] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *Store) handleDeleteGuildMember(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.DeleteGuildMember(id); err != nil {
		return err
	}
	respMessage := fmt.Sprintf("guild member with id %v deleted", id)
	log.Print("[api.guild.member] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}
