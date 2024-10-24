package main

import "database/sql"

type CreateUserRequest struct {
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
}

/*===========================================================================*/

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
}

type UserGames struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
	Games     *Games `json:"games"`
}

type Games struct {
	LeagueOfLegends *LeagueOfLegends `json:"league_of_legends"`
}

type LeagueOfLegends struct {
	MainRole   string   `json:"main_role"`
	SecondRole string   `json:"second_role"`
	MainChamps []string `json:"main_champs"`
	Accounts   []string `json:"accounts"`
}

func NewUser(name string, discordId string) *User {
	return &User{
		Name:      name,
		DiscordID: discordId,
	}
}

func NewUserGames(name string, discordId string, lol *LeagueOfLegends) *UserGames {
	return &UserGames{
		Name: name,
		Games: &Games{
			LeagueOfLegends: lol,
		},
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

func scanIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)
	if err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.DiscordID,
	); err != nil {
		return nil, err
	}
	return user, nil
}
