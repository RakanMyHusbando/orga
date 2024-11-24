package api

import (
	"encoding/json"
	"net/http"

	"github.com/RakanMyHusbando/shogun/types"
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
	return WriteJSON(w, http.StatusOK, "guild created")
}

func (s *APIServer) handleGetGuild(w http.ResponseWriter, r *http.Request) error {
	var guild []*map[string]any
	id, err := GetId(r)
	if err != nil {
		guild, err = s.store.GetGuild()
	} else {
		guild, err = s.store.GetGuildById(id)
	}
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, guild)
}

func (s *APIServer) handleUpdateGuild(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	guild := new(types.Guild)
	if err := json.NewDecoder(r.Body).Decode(&guild); err != nil {
		return err
	}
	if err := s.store.UpdateGuild(guild, id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "guild updated")
}

func (s *APIServer) handleDeleteGuild(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteGuild(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "deleted guild")
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
	return WriteJSON(w, http.StatusOK, "guild role created")
}

func (s *APIServer) handleGetGuildRole(w http.ResponseWriter, r *http.Request) error {
	role, err := s.store.GetGuildRole()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, role)
}

func (s *APIServer) handleUpdateGuildRole(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	role := new(types.GuildRole)
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return err
	}
	if err := s.store.UpdateGuildRole(role, id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "guild role updated")
}

func (s *APIServer) handleDeleteGuildRole(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteGuildRole(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "deleted guild role")
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
	return WriteJSON(w, http.StatusOK, "guild member created")
}

func (s *APIServer) handleUpdateGuildMember(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	member := new(types.GuildMember)
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		return err
	}
	if err := s.store.UpdateGuildMember(member, id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "guild member updated")
}

func (s *APIServer) handleDeleteGuildMember(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteGuildMember(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, "deleted guild member")
}
