package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RakanMyHusbando/shogun/storage"
	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type APIServer struct {
	listenAddr string
	store      storage.Storage
}

func NewAPIServer(listenAddr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errChan := make(chan error)
		go func() {
			errChan <- f(w, r)
		}()
		if err := <-errChan; err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/user/{id}", makeHTTPHandleFunc(s.handleUser))

	router.HandleFunc("/user/{id}/league_of_legends", makeHTTPHandleFunc(s.handleLeagueOfLegends))

	router.HandleFunc("/user/{id}/game_account", makeHTTPHandleFunc(s.handleGameAccount))
	router.HandleFunc("/user/{id}/game_account/{accountName}", makeHTTPHandleFunc(s.handleGameAccount))

	router.HandleFunc("/guild", makeHTTPHandleFunc(s.handlerGuild))
	router.HandleFunc("/guild/{id}", makeHTTPHandleFunc(s.handlerGuild))

	router.HandleFunc("/guild_user", makeHTTPHandleFunc(s.handlerGuildMember))
	router.HandleFunc("/guild_user/{id}", makeHTTPHandleFunc(s.handlerGuildMember))

	router.HandleFunc("/guild_role", makeHTTPHandleFunc(s.handlerGuildRole))
	router.HandleFunc("/guild_role/{id}", makeHTTPHandleFunc(s.handlerGuildRole))

	router.HandleFunc("/team", makeHTTPHandleFunc(s.handleTeam))
	router.HandleFunc("/team/{id}", makeHTTPHandleFunc(s.handleTeam))

	log.Println("API server running on ", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}

/* ------------------------------ method handler ------------------------------ */

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetUser(w, r)
	case "POST":
		return s.handleCreateUser(w, r)
	case "PATCH":
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
	case "PATCH":
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
	case "PATCH":
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
		return s.handleGetGuild(w, r)
	case "DELETE":
		return s.handleDeleteGuild(w, r)
	case "PATCH":
		return s.handleUpdateGuild(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handlerGuildRole(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateGuildRole(w, r)
	case "GET":
		return s.handleGetGuildRole(w, r)
	case "DELETE":
		return s.handleDeleteGuildRole(w, r)
	case "PATCH":
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
	case "PATCH":
		return s.handleUpdateGuildMember(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleTeam(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateTeam(w, r)
	case "GET":
		return s.handleGetTeam(w, r)
	case "DELETE":
		return s.handleDeleteTeam(w, r)
	case "PATCH":
		return s.handleUpdateTeam(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

/* ------------------------------ helper functions ------------------------------ */

func GetId(r *http.Request) int {
	if strId := mux.Vars(r)["id"]; strId != "" {
		id, err := strconv.Atoi(strId)
		if err != nil {
			return -1
		}
		return id
	}
	return -1
}