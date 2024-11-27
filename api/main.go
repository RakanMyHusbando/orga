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

var ErrNoId = fmt.Errorf("id not found")

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

	router.HandleFunc("/team_role/", makeHTTPHandleFunc(s.handleTeamRole))
	router.HandleFunc("/team_role/{id}", makeHTTPHandleFunc(s.handleTeamRole))

	router.HandleFunc("/team_member/", makeHTTPHandleFunc(s.handleTeamMember))
	router.HandleFunc("/team_member/{id}", makeHTTPHandleFunc(s.handleTeamMember))

	router.HandleFunc("/discord", makeHTTPHandleFunc(s.handleDiscord))
	router.HandleFunc("/discord/{id}", makeHTTPHandleFunc(s.handleDiscord))

	router.HandleFunc("/discord_role/", makeHTTPHandleFunc(s.handleDiscordRole))
	router.HandleFunc("/discord_role/{id}", makeHTTPHandleFunc(s.handleDiscordRole))

	router.HandleFunc("/discord_member/", makeHTTPHandleFunc(s.handleDiscordMember))
	router.HandleFunc("/discord_member/{id}", makeHTTPHandleFunc(s.handleDiscordMember))

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
	case "PATCH":
		return s.handleUpdateLeagueOfLegends(w, r)
	case "DELETE":
		return s.handleDeleteLeagueOfLegends(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleGameAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateGameAccount(w, r)
	case "PATCH":
		return s.handleUpdateGameAccount(w, r)
	case "DELETE":
		return s.handleDeleteGameAccount(w, r)
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
	case "PATCH":
		return s.handleUpdateGuild(w, r)
	case "DELETE":
		return s.handleDeleteGuild(w, r)
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
	case "PATCH":
		return s.handleUpdateGuildRole(w, r)
	case "DELETE":
		return s.handleDeleteGuildRole(w, r)
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
	case "PATCH":
		return s.handleUpdateTeam(w, r)
	case "DELETE":
		return s.handleDeleteTeam(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleTeamRole(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateTeamRole(w, r)
	case "GET":
		return s.handleGetTeamRole(w, r)
	case "PATCH":
		return s.handleUpdateTeamRole(w, r)
	case "DELETE":
		return s.handleDeleteTeamRole(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleTeamMember(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateTeamMember(w, r)
	case "DELETE":
		return s.handleDeleteTeamMember(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleDiscord(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateDiscord(w, r)
	case "GET":
		return s.handleGetDiscord(w, r)
	case "PATCH":
		return s.handleUpdateDiscord(w, r)
	case "DELETE":
		return s.handleDeleteDiscord(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleDiscordRole(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateDiscordRole(w, r)
	case "GET":
		return s.handleGetDiscordRole(w, r)
	case "PATCH":
		return s.handleUpdateDiscordRole(w, r)
	case "DELETE":
		return s.handleDeleteDiscordRole(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *APIServer) handleDiscordMember(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateDiscordMember(w, r)
	case "DELETE":
		return s.handleDeleteDiscordMember(w, r)
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
