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

// POST
func (s *APIServer) handleCreateGuild(w http.ResponseWriter, r *http.Request) error {
	reqGuild := new(types.ReqGuild)
	if err := json.NewDecoder(r.Body).Decode(&reqGuild); err != nil {
		log.Println("[api.guild.(s)handleCreateGuild(w,r)] error while executing 'json.NewDecoder(r.Body).Decode(&reqGuild)': ", err)
		return err
	}

	if err := s.store.CreateGuild(reqGuild); err != nil {
		log.Println("[api.guild.(s)handleCreateGuild(w,r)] error while executing 's.store.CreateGuild(reqGuild)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, "guild created")
}

// GET
func (s *APIServer) handleGetGuild(w http.ResponseWriter, r *http.Request) error {
	guildList, err := s.store.GetGuild()
	if err != nil {
		log.Println("[api.guild.(s)handleGetGuild(w,r)] error while executing 's.store.GetGuild()': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, guildList)
}

// GET
func (s *APIServer) handleGetGuildById(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(mux.Vars(r))
	id, err := GetId(r)
	if err != nil {
		log.Println("[api.guild.(s)handleGetGuildById(w,r)] error while executing 'GetId(r)': ", err)
		return err
	}

	guild, err := s.store.GetGuildById(id)
	if err != nil {
		log.Println("[api.guild.(s)handleGetGuildById(w,r)] error while executing 'GetId(r)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, guild)
}

// DELETE
func (s *APIServer) handleDeleteGuild(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(mux.Vars(r))
	id, err := GetId(r)
	if err != nil {
		log.Println("[api.guild.handleDeleteGuild(w,r)] error while executing 'GetId(r)': ", err)
		return err
	}

	if err := s.store.DeleteGuild(id); err != nil {
		log.Println("[api.guild.(s)handleDeleteGuild(w,r)] error while executing 's.store.DeleteGuild(id)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintln("deleted guild with id ", id))
}

// PATCH
func (s *APIServer) handleUpdateGuild(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(mux.Vars(r))
	id, err := GetId(r)
	if err != nil {
		log.Println("[api.guild.(s)handleUpdateGuild(w,r)] error while executing 'GetId(r)': ", err)
		return err
	}

	resGuild := new(types.ResGuild)
	if err := json.NewDecoder(r.Body).Decode(&resGuild); err != nil {
		log.Println("[api.guild.(s)handleUpdateGuild(w,r)] error while executing 'json.NewDecoder(r.Body).Decode(&resGuild)': ", err)
		return err
	}

	resGuild.Id = id

	if err := s.store.UpdateGuild(resGuild); err != nil {
		log.Println("[api.guild.(s)handleUpdateGuild(w,r)] error while executing 's.store.UpdateGuild(resGuild)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, resGuild)
}

/* --------------------------------- handler guild role --------------------------------- */

// POST
func (s *APIServer) handleCreateGuildRole(w http.ResponseWriter, r *http.Request) error {

	reqGuildRole := new(types.ReqGuildRole)
	if err := json.NewDecoder(r.Body).Decode(&reqGuildRole); err != nil {
		log.Println("[api.guild.(s)handleCreateGuildRole(w,r)] error while executing 'json.NewDecoder(r.Body).Decode(&reqGuildRole)': ", err)
		return err
	}

	if err := s.store.CreateGuildRole(reqGuildRole); err != nil {
		log.Println("[api.guild.(s)handleCreateGuildRole(w,r)] error while executing 's.store.CreateGuildRole(reqGuildRole)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, reqGuildRole)
}

// GET
func (s *APIServer) handleGetGuildRole(w http.ResponseWriter, r *http.Request) error {
	guildRoleList, err := s.store.GetGuildRole()
	if err != nil {
		log.Println("[api.guild.(s)handleGetGuildRole(w,r)] error while executing 's.store.GetGuildRole()': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, guildRoleList)
}

// DELETE
func (s *APIServer) handleDeleteGuildRole(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(mux.Vars(r))
	id, err := GetId(r)
	if err != nil {
		log.Println("[api.guild.(s)handleDeleteGuildRole(w,r)] error while executing 'GetId(r)': ", err)
		return err
	}

	if err := s.store.DeleteGuildRole(id); err != nil {
		log.Println("[api.guild.(s)handleDeleteGuildRole(w,r)] error while executing 's.store.DeleteGuildRole(id)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintln("deleted guild role with id ", id))
}

// PATCH
func (s *APIServer) handleUpdateGuildRole(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(mux.Vars(r))
	id, err := GetId(r)
	if err != nil {
		log.Println("[api.guild.(s)handleUpdateGuildRole(w,r)] error while executing 'GetId(r)': ", err)
		return err
	}

	reqGuildRole := new(types.ReqGuildRole)
	if err := json.NewDecoder(r.Body).Decode(&reqGuildRole); err != nil {
		log.Println("[api.guild.(s)handleUpdateGuildRole(w,r)] error while executing 'json.NewDecoder(r.Body).Decode(&reqGuildRole)': ", err)
		return err
	}

	fmt.Println(reqGuildRole)
	fmt.Println(id)
	reqGuildRole.Id = id

	if err := s.store.UpdateGuildRole(reqGuildRole); err != nil {
		log.Println("[api.guild.(s)handleUpdateGuildRole(w,r)] error while executing 's.store.UpdateGuildRole(reqGuildRole)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, reqGuildRole)
}

/* --------------------------------- handler guild member --------------------------------- */

// POST
func (s *APIServer) handleCreateGuildMember(w http.ResponseWriter, r *http.Request) error {
	guildMember := new(types.ReqGuildMember)

	if err := json.NewDecoder(r.Body).Decode(&guildMember); err != nil {
		log.Println("[api.guild.(s)handleCreateGuildMember(w,r)] error while executing 'json.NewDecoder(r.Body).Decode(&guildMember)': ", err)
		return err
	}

	if err := s.store.CreateGuildMember(guildMember); err != nil {
		log.Println("[api.guild.(s)handleCreateGuildMember(w,r)] error while executing 's.store.CreateGuildMember(guildMember)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, guildMember)
}

// DELETE
func (s *APIServer) handleDeleteGuildMember(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		log.Println("[api.guild.(s)handleDeleteGuildMember(w,r)] error while executing 'GetId(r)': ", err)
		return err
	}

	if err := s.store.DeleteGuildMember(id); err != nil {
		log.Println("[api.guild.(s)handleDeleteGuildMember(w,r)] error while executing 's.store.DeleteGuildMember(id)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintln("deleted guild member with id ", id))
}

// PATCH
func (s *APIServer) handleUpdateGuildMember(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		log.Println("[api.guild.(s)handleUpdateGuildMember(w,r)] error while executing 'GetId(r)': ", err)
		return err
	}

	guildMember := new(types.ReqGuildMember)

	if err := json.NewDecoder(r.Body).Decode(&guildMember); err != nil {
		log.Println("[api.guild.(s)handleUpdateGuildMember(w,r)] error while executing 'json.NewDecoder(r.Body).Decode(&guildMember)': ", err)
		return err
	}

	guildMember.UserId = id

	if err := s.store.UpdateGuildMember(guildMember); err != nil {
		log.Println("[api.guild.(s)handleUpdateGuildMember(w,r)] error while executing 's.store.UpdateGuildMember(guildMember)': ", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, guildMember)
}
