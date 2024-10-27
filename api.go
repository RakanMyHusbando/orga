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
		return s.handleDeleteGameAccount(w, r)
	case "PUT":
		return s.handleUpdateGameAccount(w, r)
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

	return WriteJSON(w, http.StatusOK, "'user' created")
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

	return WriteJSON(w, http.StatusOK, "'user' deleted")
}

// PUT
func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	user := new(User)
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
	reqUserLol := new(ReqUserLeagueOfLegends)
	if err := json.NewDecoder(r.Body).Decode(&reqUserLol); err != nil {
		return err
	}

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

	return WriteJSON(w, http.StatusOK, "'league_of_legends' deleted from user.")
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

	reqUserLol.Id = id

	if err := s.store.UpdateUserLeagueOfLegends(reqUserLol); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, reqUserLol)
}

/* ------------------------------ handler game account ------------------------------ */

func (s *APIServer) handleCreateGameAccount(w http.ResponseWriter, r *http.Request) error {
	createGameAccount := new(ReqGameAccount)
	if err := json.NewDecoder(r.Body).Decode(&createGameAccount); err != nil {
		return err
	}

	if err := s.store.CreateGameAccount(createGameAccount); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, createGameAccount)
}

func (s *APIServer) handleDeleteGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteGameAccount(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "'game_account' deleted from user.")
}

func (s *APIServer) handleUpdateGameAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := GetId(r)
	if err != nil {
		return err
	}

	createGameAccount := new(ReqGameAccount)
	if err := json.NewDecoder(r.Body).Decode(&createGameAccount); err != nil {
		return err
	}

	createGameAccount.UserId = id

	if err := s.store.UpdateGameAccount(createGameAccount); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, createGameAccount)
}

/* ------------------------------ handler team ------------------------------ */

/* ------------------------------ handler guide ------------------------------ */
