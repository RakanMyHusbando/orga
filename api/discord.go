package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *APIServer) handleCreateDiscord(w http.ResponseWriter, r *http.Request) error {
	discord := new(types.Discord)
	if err := json.NewDecoder(r.Body).Decode(&discord); err != nil {
		return err
	}
	if err := s.store.CreateDiscord(discord); err != nil {
		return err
	}
	resp := "discord server created"
	log.Print("[api.discord] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleGetDiscord(w http.ResponseWriter, r *http.Request) error {
	var discord []*types.Discord
	var err error
	if id := GetId(r); id == -1 {
		discord, err = s.store.GetDiscord()
	} else {
		discord, err = s.store.GetDiscordById(id)
	}
	if err != nil {
		return err
	}
	log.Print("[api.discord] got discord servers")
	return WriteJSON(w, http.StatusOK, discord)
}

func (s *APIServer) handleUpdateDiscord(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	discord := new(types.Discord)
	if err := json.NewDecoder(r.Body).Decode(&discord); err != nil {
		return err
	}
	if err := s.store.UpdateDiscord(discord, GetId(r)); err != nil {
		return err
	}
	resp := "discord server updated"
	log.Print("[api.discord] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleDeleteDiscord(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.store.DeleteDiscord(id); err != nil {
		return err
	}
	resp := "discord server deleted"
	log.Print("[api.discord] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

/* ------------------------------ handle role ------------------------------ */

func (s *APIServer) handleCreateDiscordRole(w http.ResponseWriter, r *http.Request) error {
	role := new(types.DiscordRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.store.CreateDiscordRole(role); err != nil {
		return err
	}
	resp := "discord role created"
	log.Print("[api.discord.role] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleGetDiscordRole(w http.ResponseWriter, r *http.Request) error {
	var role []*types.DiscordRole
	var err error
	if id := GetId(r); id == -1 {
		role, err = s.store.GetDiscordRole()
	} else {
		role, err = s.store.GetDiscordRoleById(id)
	}
	if err != nil {
		return err
	}
	log.Print("[api.discord.role] got discord roles")
	return WriteJSON(w, http.StatusOK, role)
}

func (s *APIServer) handleUpdateDiscordRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	role := new(types.DiscordRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.store.UpdateDiscordRole(role, id); err != nil {
		return err
	}
	resp := "discord role updated"
	log.Print("[api.discord.role] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleDeleteDiscordRole(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.store.DeleteDiscordRole(id); err != nil {
		return err
	}
	resp := "discord role deleted"
	log.Print("[api.discord.role] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

/* ------------------------------ handle member ------------------------------ */

func (s *APIServer) handleCreateDiscordMember(w http.ResponseWriter, r *http.Request) error {
	member := new(types.DiscordMember)
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		return err
	}
	if err := s.store.CreateDiscordMember(member); err != nil {
		return err
	}
	resp := "discord member created"
	log.Print("[api.discord.member] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleDeleteDiscordMember(w http.ResponseWriter, r *http.Request) error {
	id := GetId(r)
	if id == -1 {
		return ErrNoId
	}
	if err := s.store.DeleteDiscordMember(id); err != nil {
		return err
	}
	resp := fmt.Sprintf("discord member with id %v deleted", id)
	log.Print("[api.discord.member] " + resp)
	return WriteJSON(w, http.StatusOK, resp)
}
