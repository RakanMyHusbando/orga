package types

type User struct {
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
}

type GameAccount struct {
	UserId int    `json:"user_id"`
	Game   string `json:"game"`
	Name   string `json:"name"`
}

type LeagueOfLegends struct {
	UserId     int      `json:"user_id"`
	MainRole   string   `json:"main_role"`
	SecondRole string   `json:"second_role"`
	MainChamps []string `json:"main_champs"`
}

type Guild struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
}

type GuildRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GuildMember struct {
	UserId  int `json:"user_id"`
	GuildId int `json:"guild_id"`
	RoleId  int `json:"role_id"`
}

type Team struct {
	Id           int    `json:"id"`
	GuildId      int    `json:"guild_id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type TeamRole struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TeamMember struct {
	UserId int `json:"user_id"`
	TeamId int `json:"team_id"`
	RoleId int `json:"role_id"`
}
