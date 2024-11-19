package types

type ReqUser struct {
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
}

type ReqGameAccount struct {
	UserId int    `json:"user_id"`
	Game   string `json:"game"`
	Name   string `json:"name"`
}

type ReqLeagueOfLegends struct {
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

type ReqTeam struct {
	GuildId      int    `json:"guild_id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}
