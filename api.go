package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/user/{id}", makeHTTPHandleFunc(s.handleUser))

	router.HandleFunc("/user/{id}/league_of_legends", makeHTTPHandleFunc(s.handleLeagueOfLegends))

	router.HandleFunc("/user/{id}/game_account", makeHTTPHandleFunc(s.handleGameAccount))
	router.HandleFunc("/user/{id}/game_account/{accountName}", makeHTTPHandleFunc(s.handleGameAccount))

	router.HandleFunc("/user/guild", makeHTTPHandleFunc(s.handlerGuildMember))
	router.HandleFunc("/user/guild/{id}", makeHTTPHandleFunc(s.handlerGuildMember))

	router.HandleFunc("/guild", makeHTTPHandleFunc(s.handlerGuild))
	router.HandleFunc("/guild/{id}", makeHTTPHandleFunc(s.handlerGuild))

	router.HandleFunc("/guild_role", makeHTTPHandleFunc(s.handlerGuildRole))
	router.HandleFunc("/guild_role/{id}", makeHTTPHandleFunc(s.handlerGuildRole))

	log.Println("API server running on ", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}

func GetId(r *http.Request) (int, error) {
	var intId int
	if strId := mux.Vars(r)["id"]; strId != "" {
		res, err := strconv.Atoi(strId)
		if err != nil {
			return intId, err
		}
		intId = res
		return intId, nil
	}
	return intId, nil
}

/* ============================== method handler ============================== */

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		if mux.Vars(r)["id"] != "" {
			return s.handleGetUserById(w, r)
		}
		return s.handleGetUser(w, r)
	case "POST":
		return s.handleCreateUser(w, r)
	case "PUT":
		return s.handleUpdateUser(w, r)
	case "DELETE":
		return s.handleDeleteUser(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateLeagueOfLegends(w, r)
	case "DELETE":
		return s.handleDeleteLeagueOfLegends(w, r)
	case "PUT":
		return s.handleUpdateLeagueOfLegends(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleGameAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateGameAccount(w, r)
	case "DELETE":
		return s.handleDeleteGameAccount(w, r)
	case "PUT":
		return s.handleUpdateGameAccount(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handlerGuild(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateGuild(w, r)
	case "GET":
		if mux.Vars(r)["id"] != "" {
			return s.handleGetGuildById(w, r)
		}
		return s.handleGetGuild(w, r)
	case "DELETE":
		return s.handleDeleteGuild(w, r)
	case "PUT":
		return s.handleUpdateGuild(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handlerGuildRole(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateGuildRole(w, r)
	case "DELETE":
		return s.handleDeleteGuildRole(w, r)
	case "PUT":
		return s.handleUpdateGuildRole(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handlerGuildMember(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateGuildMember(w, r)
	case "DELETE":
		return s.handleDeleteGuildMember(w, r)
	case "PUT":
		return s.handleUpdateGuildMember(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

/* ------------------------------ handler user ------------------------------ */

// POST
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	reqUser := new(ReqUser)
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		return err
	}

	if err := s.store.CreateUser(reqUser); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "user created")
}

// GET
func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	userList, err := s.store.GetUser()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, userList)
}

// GET
func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	user, err := s.store.GetUserById(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

// DELETE
func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.DeletUser(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "user with id "+strconv.Itoa(id)+" deleted")
}

// PUT
func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	user := new(ResUser)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	user.Id = id

	if err := s.store.UpdateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

/* ------------------------------ handler league of legends ------------------------------ */

// POST
func (s *APIServer) handleCreateLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	reqUserLol := new(ReqUserLeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&reqUserLol); err != nil {
		return err
	}

	reqUserLol.UserId = id

	if err := s.store.CreateUserLeagueOfLegends(reqUserLol); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqUserLol)
}

// DELETE
func (s *APIServer) handleDeleteLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteUserLeagueOfLegends(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "league_of_legends deleted from user with id "+strconv.Itoa(id))
}

// PUT
func (s *APIServer) handleUpdateLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	reqUserLol := new(ReqUserLeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&reqUserLol); err != nil {
		return err
	}

	reqUserLol.UserId = id

	if err := s.store.UpdateUserLeagueOfLegends(reqUserLol); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqUserLol)
}

/* ------------------------------ handler game account ------------------------------ */

// POST
func (s *APIServer) handleCreateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	reqGameAcc := new(ReqGameAccount)
	if err := json.NewDecoder(r.Body).Decode(&reqGameAcc); err != nil {
		return err
	}

	reqGameAcc.UserId = id

	if err := s.store.CreateGameAccount(reqGameAcc); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqGameAcc)
}

// DELETE
func (s *APIServer) handleDeleteGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	reqGameAcc := new(ReqGameAccount)
	if err := json.NewDecoder(r.Body).Decode(&reqGameAcc); err != nil {
		return err
	}

	reqGameAcc.UserId = id

	if err := s.store.DeleteGameAccount(reqGameAcc); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "game account deleted from user with id "+strconv.Itoa(id))
}

// PUT
func (s *APIServer) handleUpdateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	updateGameAccount := new(ReqUpdateGameAccount)
	if err := json.NewDecoder(r.Body).Decode(&updateGameAccount); err != nil {
		return err
	}

	updateGameAccount.UserId = id

	if err := s.store.UpdateGameAccount(updateGameAccount); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, updateGameAccount)
}

/* --------------------------------- handler guild --------------------------------- */

// POST
func (s *APIServer) handleCreateGuild(w http.ResponseWriter, r *http.Request) error {
	reqGuild := new(ReqGuild)
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

// PUT
func (s *APIServer) handleUpdateGuild(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	resGuild := new(ResGuild)
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
	reqGuildRole := new(ReqGuildRole)
	if err := json.NewDecoder(r.Body).Decode(&reqGuildRole); err != nil {
		return err
	}

	if err := s.store.CreateGuildRole(reqGuildRole); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqGuildRole)
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

// PUT
func (s *APIServer) handleUpdateGuildRole(w http.ResponseWriter, r *http.Request) error {
	reqUpdateGuildRole := new(ReqUpdateGuildRole)
	if err := json.NewDecoder(r.Body).Decode(&reqUpdateGuildRole); err != nil {
		return err
	}

	if err := s.store.UpdateGuildRole(reqUpdateGuildRole); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqUpdateGuildRole)
}

/* --------------------------------- handler guild member --------------------------------- */

// POST
func (s *APIServer) handleCreateGuildMember(w http.ResponseWriter, r *http.Request) error {
	guildMember := new(ReqGuildMember)

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

// PUT
func (s *APIServer) handleUpdateGuildMember(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	guildMember := new(ReqGuildMember)

	if err := json.NewDecoder(r.Body).Decode(&guildMember); err != nil {
		return err
	}

	guildMember.UserId = id

	if err := s.store.UpdateGuildMember(guildMember); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, guildMember)
}
