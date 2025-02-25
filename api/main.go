package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RakanMyHusbando/orga/storage"
	"github.com/RakanMyHusbando/orga/types"
)

var ErrNoId = fmt.Errorf("id not found")

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type Store struct {
	storage.Storage
}

func NewStore(store storage.Storage) *Store {
	return &Store{store}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(types.NewJSONResponse(status, v))
}

func MakeHTTPHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errChan := make(chan error)
		go func() {
			errChan <- f(w, r)
		}()
		if err := <-errChan; err != nil {
			WriteJSON(w, http.StatusBadRequest, err)
		}
	}
}

func (s *Store) Routes() {
	http.HandleFunc("/api/user", MakeHTTPHandleFunc(s.handleUser))
	http.HandleFunc("/api/user/{id}", MakeHTTPHandleFunc(s.handleUser))

	http.HandleFunc("/api/user/{id}/league_of_legends", MakeHTTPHandleFunc(s.handleLeagueOfLegends))

	http.HandleFunc("/api/user/{id}/game_account", MakeHTTPHandleFunc(s.handleGameAccount))
	http.HandleFunc("/api/user/{id}/game_account/{accountName}", MakeHTTPHandleFunc(s.handleGameAccount))

	http.HandleFunc("/api/guild", MakeHTTPHandleFunc(s.handlerGuild))
	http.HandleFunc("/api/guild/{id}", MakeHTTPHandleFunc(s.handlerGuild))

	http.HandleFunc("/api/guild_member", MakeHTTPHandleFunc(s.handlerGuildMember))
	http.HandleFunc("/api/guild_member/{id}", MakeHTTPHandleFunc(s.handlerGuildMember))

	http.HandleFunc("/api/guild_role", MakeHTTPHandleFunc(s.handlerGuildRole))
	http.HandleFunc("/api/guild_role/{id}", MakeHTTPHandleFunc(s.handlerGuildRole))

	http.HandleFunc("/api/team", MakeHTTPHandleFunc(s.handleTeam))
	http.HandleFunc("/api/team/{id}", MakeHTTPHandleFunc(s.handleTeam))

	http.HandleFunc("/api/team_role/", MakeHTTPHandleFunc(s.handleTeamRole))
	http.HandleFunc("/api/team_role/{id}", MakeHTTPHandleFunc(s.handleTeamRole))

	http.HandleFunc("/api/team_member/", MakeHTTPHandleFunc(s.handleTeamMember))
	http.HandleFunc("/api/team_member/{id}", MakeHTTPHandleFunc(s.handleTeamMember))

	http.HandleFunc("/api/discord", MakeHTTPHandleFunc(s.handleDiscord))
	http.HandleFunc("/api/discord/{id}", MakeHTTPHandleFunc(s.handleDiscord))

	http.HandleFunc("/api/discord_role/", MakeHTTPHandleFunc(s.handleDiscordRole))
	http.HandleFunc("/api/discord_role/{id}", MakeHTTPHandleFunc(s.handleDiscordRole))

	http.HandleFunc("/api/discord_member/", MakeHTTPHandleFunc(s.handleDiscordMember))
	http.HandleFunc("/api/discord_member/{id}", MakeHTTPHandleFunc(s.handleDiscordMember))
}

/* ------------------------------ method handler ------------------------------ */

func (s *Store) handleUser(w http.ResponseWriter, r *http.Request) error {
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

func (s *Store) handleLeagueOfLegends(w http.ResponseWriter, r *http.Request) error {
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

func (s *Store) handleGameAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateGameAccount(w, r)
	case "DELETE":
		return s.handleDeleteGameAccount(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *Store) handlerGuild(w http.ResponseWriter, r *http.Request) error {
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

func (s *Store) handlerGuildRole(w http.ResponseWriter, r *http.Request) error {
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

func (s *Store) handlerGuildMember(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateGuildMember(w, r)
	case "DELETE":
		return s.handleDeleteGuildMember(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *Store) handleTeam(w http.ResponseWriter, r *http.Request) error {
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

func (s *Store) handleTeamRole(w http.ResponseWriter, r *http.Request) error {
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

func (s *Store) handleTeamMember(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateTeamMember(w, r)
	case "DELETE":
		return s.handleDeleteTeamMember(w, r)
	default:
		return fmt.Errorf("unsupported method: %s", r.Method)
	}
}

func (s *Store) handleDiscord(w http.ResponseWriter, r *http.Request) error {
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

func (s *Store) handleDiscordRole(w http.ResponseWriter, r *http.Request) error {
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

func (s *Store) handleDiscordMember(w http.ResponseWriter, r *http.Request) error {
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
	if strId := r.FormValue("id"); strId != "" {
		id, err := strconv.Atoi(strId)
		if err != nil {
			return -1
		}
		return id
	}
	return -1
}
