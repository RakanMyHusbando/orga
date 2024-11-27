package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/shogun/types"
	"github.com/gorilla/mux"
)

/* --------------------------------- handler guild --------------------------------- */

func (s *APIServer) handleCreateGuild(w http.ResponseWriter, r *http.Request) error {
	guild := new(types.Guild)
	if err := json.NewDecoder(r.Body).Decode(&guild); err != nil {
		return err
	}
	if err := s.store.CreateGuild(guild); err != nil {
		return err
	}
	respMessage := "guild created"
	log.Print("[api.guild] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleGetGuild(w http.ResponseWriter, r *http.Request) error {
	var guild []*types.Guild
	var err error
	if id := GetId(r); id == -1 {
		guild, err = s.store.GetGuild()
	} else {
		guild, err = s.store.GetGuildById(id)
	}
	if err != nil {
		return err
	}
	for _, g := range guild {
		member, err := s.store.GetGuildMemberByGuildId(g.Id)
		if err != nil {
			log.Println("[api.guild] no team members found for guild with id ", g.Id)
		} else {
			for _, m := range member {
				role, err := s.store.GetGuildRoleById(m.RoleId)
				if err != nil {
					log.Println("[api.guild] no role found for guild member with user_id ", m.UserId)
				} else {
					user, err := s.store.GetUserById(m.UserId)
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

func (s *APIServer) handleUpdateGuild(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	guild := new(types.Guild)
	if err := json.NewDecoder(r.Body).Decode(&guild); err != nil {
		return err
	}
	if err := s.store.UpdateGuild(guild, id); err != nil {
		return err
	}
	respMessage := "guild updated with id " + mux.Vars(r)["id"]
	log.Print("[api.guild] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleDeleteGuild(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.store.DeleteGuild(id); err != nil {
		return err
	}
	respMessage := "guild deleted with id " + mux.Vars(r)["id"]
	log.Print("[api.guild] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

/* --------------------------------- handler guild role --------------------------------- */

func (s *APIServer) handleCreateGuildRole(w http.ResponseWriter, r *http.Request) error {
	role := new(types.GuildRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.store.CreateGuildRole(role); err != nil {
		return err
	}
	respMessage := "guild role created"
	log.Print("[api.guild.role] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleGetGuildRole(w http.ResponseWriter, r *http.Request) error {
	role, err := s.store.GetGuildRole()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, role)
}

func (s *APIServer) handleUpdateGuildRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	role := new(types.GuildRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.store.UpdateGuildRole(role, id); err != nil {
		return err
	}
	respMessage := fmt.Sprintf("guild role with id %v updated", id)
	log.Print("[api.guild.role] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleDeleteGuildRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.store.DeleteGuildRole(id); err != nil {
		return err
	}
	respMessage := fmt.Sprintf("guild role with id %v deleted", id)
	log.Print("[api.guild.role] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

/* --------------------------------- handler guild member --------------------------------- */

func (s *APIServer) handleCreateGuildMember(w http.ResponseWriter, r *http.Request) error {
	member := new(types.GuildMember)
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		return err
	}
	if err := s.store.CreateGuildMember(member); err != nil {
		return err
	}
	respMessage := "guild member created"
	log.Print("[api.guild.member] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}

func (s *APIServer) handleDeleteGuildMember(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.store.DeleteGuildMember(id); err != nil {
		return err
	}
	respMessage := fmt.Sprintf("guild member with id %v deleted", id)
	log.Print("[api.guild.member] " + respMessage)
	return WriteJSON(w, http.StatusOK, respMessage)
}
