package main

/* ------------------------------ request struct ------------------------------ */

type ReqUser struct {
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
}

type ReqGameAccount struct {
	UserId int    `json:"user_id"`
	Game   string `json:"game"`
	Name   string `json:"name"`
}

type ReqUpdateGameAccount struct {
	UserId  int    `json:"user_id"`
	Game    string `json:"game"`
	NameOld string `json:"name_old"`
	NameNew string `json:"name_new"`
}

type ReqUserLeagueOfLegends struct {
	UserId     int      `json:"user_id"`
	MainRole   string   `json:"main_role"`
	SecondRole string   `json:"second_role"`
	MainChamps []string `json:"main_champs"`
}

type ReqGuild struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
}

type ReqGuildRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ReqGuildMember struct {
	UserId  int `json:"user_id"`
	GuildId int `json:"guild_id"`
	RoleId  int `json:"role_id"`
}

type ReqUpdateGuildRole struct {
	NameOld        string `json:"name_old"`
	NameNew        string `json:"name_new"`
	DescriptionNew string `json:"description_new"`
}

/* ------------------------------ response struct ------------------------------ */

type ResUser struct {
	Id              int                 `json:"id"`
	Name            string              `json:"name"`
	DiscordID       string              `json:"discord_id"`
	LeagueOfLegends *ResLeagueOfLegends `json:"league_of_legends"`
}

type ResLeagueOfLegends struct {
	MainRole   string   `json:"main_role"`
	SecondRole string   `json:"second_role"`
	MainChamps []string `json:"main_champs"`
	Accounts   []string `json:"accounts"`
}

type ResGuild struct {
	Id           int                 `json:"id"`
	Name         string              `json:"name"`
	Abbreviation string              `json:"abbreviation"`
	Description  string              `json:"description"`
	Member       map[string][]string `json:"member"`
}

/* ------------------------------ constructor ------------------------------ */

func NewUser(name string, discordId string) *ResUser {
	return &ResUser{
		Name:      name,
		DiscordID: discordId,
	}
}

func NewLeagueOfLegends(mainRole string, secondRole string, mainChamps []string, accounts []string) *ResLeagueOfLegends {
	return &ResLeagueOfLegends{
		MainRole:   mainRole,
		SecondRole: secondRole,
		MainChamps: mainChamps,
		Accounts:   accounts,
	}
}
