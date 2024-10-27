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

type ReqUserLeagueOfLegends struct {
	Id         int      `json:"id"`
	MainRole   string   `json:"main_role"`
	SecondRole string   `json:"second_role"`
	MainChamps []string `json:"main_champs"`
}

/* ------------------------------ response struct ------------------------------ */

type User struct {
	Id              int              `json:"id"`
	Name            string           `json:"name"`
	DiscordID       string           `json:"discord_id"`
	LeagueOfLegends *LeagueOfLegends `json:"league_of_legends"`
}

type LeagueOfLegends struct {
	MainRole   string   `json:"main_role"`
	SecondRole string   `json:"second_role"`
	MainChamps []string `json:"main_champs"`
	Accounts   []string `json:"accounts"`
}

/* ------------------------------ helper struct ------------------------------ */

type mainChamps struct {
	champ_0 string
	champ_1 string
	champ_2 string
}

/* ------------------------------ constructor ------------------------------ */

func NewUser(name string, discordId string) *User {
	return &User{
		Name:      name,
		DiscordID: discordId,
	}
}

func NewLeagueOfLegends(mainRole string, secondRole string, mainChamps []string, accounts []string) *LeagueOfLegends {
	return &LeagueOfLegends{
		MainRole:   mainRole,
		SecondRole: secondRole,
		MainChamps: mainChamps,
		Accounts:   accounts,
	}
}
