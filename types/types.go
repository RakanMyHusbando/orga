package types

import (
	"reflect"
)

type User struct {
	Id              int              `json:"id"`
	Name            string           `json:"name"`
	DiscordId       string           `json:"discord_id"`
	LeagueOfLegends *LeagueOfLegends `json:"league_of_legends"`
}

type GameAccount struct {
	UserId int    `json:"user_id"`
	Game   string `json:"game"`
	Name   string `json:"name"`
	Puuid  string `json:"puuid"`
}

type LeagueOfLegends struct {
	MainRole   string    `json:"main_role"`
	SecondRole string    `json:"second_role"`
	MainChamps [3]string `json:"main_champs"`
	Accounts   []string  `json:"accounts"`
}

type Guild struct {
	Id           int                `json:"id"`
	Name         string             `json:"name"`
	Abbreviation string             `json:"abbreviation"`
	Description  string             `json:"description"`
	Member       map[string][]*User `json:"member"`
}

type GuildRole struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Member      map[string][]*User `json:"member"`
}

type GuildMember struct {
	UserId  int `json:"user_id"`
	GuildId int `json:"guild_id"`
	RoleId  int `json:"role_id"`
}

type Team struct {
	Id           int                `json:"id"`
	GuildId      int                `json:"guild_id"`
	Name         string             `json:"name"`
	Abbreviation string             `json:"abbreviation"`
	Member       map[string][]*User `json:"member"`
}

type TeamRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TeamMember struct {
	UserId int `json:"user_id"`
	TeamId int `json:"team_id"`
	RoleId int `json:"role_id"`
}

type Discord struct {
	Id          int                `json:"id"`
	Discord_id  string             `json:"discord_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	GuildId     int                `json:"guild_id"`
	Member      map[string][]*User `json:"member"`
}

type DiscordRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DiscordMember struct {
	UserId   int `json:"user_id"`
	RoleId   int `json:"role_id"`
	ServerId int `json:"server_id"`
}

type JSONResponse struct {
	Status   int    `json:"status"`
	Response any    `json:"response"`
	Error    string `json:"error"`
}

func NewJSONResponse(status int, v any) *JSONResponse {
	result := &JSONResponse{Status: status}
	if reflect.TypeOf(v).String() == "error" {
		result.Error = v.(error).Error()
	} else {
		result.Response = v
	}
	return result
}
