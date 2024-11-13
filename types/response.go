package types

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
