package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RakanMyHusbando/shogun/types"
)

/* --------------------------------- handler guild --------------------------------- */

// POST
func (s *APIServer) handleCreateGuild(w http.ResponseWriter, r *http.Request) error {
	reqGuild := new(types.ReqGuild)
	if err := json.NewDecoder(r.Body).Decode(&reqGuild); err != nil {
		return err
	}

	if err := s.store.CreateGuild(reqGuild); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "guild created")
}

// GET
func (s *APIServer) handleGetGuild(w http.ResponseWriter, r *http.Request) error {
	guildList, err := s.store.GetGuild()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, guildList)
}

// GET
func (s *APIServer) handleGetGuildById(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	guild, err := s.store.GetGuildById(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, guild)
}

// DELETE
func (s *APIServer) handleDeleteGuild(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteGuild(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "deleted guild with id "+strconv.Itoa(id))
}

// PATCH
func (s *APIServer) handleUpdateGuild(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	resGuild := new(types.ResGuild)
	if err := json.NewDecoder(r.Body).Decode(&resGuild); err != nil {
		return err
	}

	resGuild.Id = id

	if err := s.store.UpdateGuild(resGuild); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, resGuild)
}

/* --------------------------------- handler guild role --------------------------------- */

// POST
func (s *APIServer) handleCreateGuildRole(w http.ResponseWriter, r *http.Request) error {
	reqGuildRole := new(types.ReqGuildRole)
	if err := json.NewDecoder(r.Body).Decode(&reqGuildRole); err != nil {
		return err
	}

	if err := s.store.CreateGuildRole(reqGuildRole); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqGuildRole)
}

// GET
func (s *APIServer) handleGetGuildRole(w http.ResponseWriter, r *http.Request) error {
	guildRoleList, err := s.store.GetGuildRole()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, guildRoleList)
}

// DELETE
func (s *APIServer) handleDeleteGuildRole(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteGuildRole(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "deleted guild role with id "+strconv.Itoa(id))
}

// PATCH
func (s *APIServer) handleUpdateGuildRole(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	reqGuildRole := new(types.ReqGuildRole)
	if err := json.NewDecoder(r.Body).Decode(&reqGuildRole); err != nil {
		return err
	}

	reqGuildRole.Id = id

	if err := s.store.UpdateGuildRole(reqGuildRole); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqGuildRole)
}

/* --------------------------------- handler guild member --------------------------------- */

// POST
func (s *APIServer) handleCreateGuildMember(w http.ResponseWriter, r *http.Request) error {
	guildMember := new(types.ReqGuildMember)

	if err := json.NewDecoder(r.Body).Decode(&guildMember); err != nil {
		return err
	}

	if err := s.store.CreateGuildMember(guildMember); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, guildMember)
}

// DELETE
func (s *APIServer) handleDeleteGuildMember(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteGuildMember(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "deleted guild member with id "+strconv.Itoa(id))
}

// PATCH
func (s *APIServer) handleUpdateGuildMember(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	guildMember := new(types.ReqGuildMember)

	if err := json.NewDecoder(r.Body).Decode(&guildMember); err != nil {
		return err
	}

	guildMember.UserId = id

	if err := s.store.UpdateGuildMember(guildMember); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, guildMember)
}
