package types

func NewUser(name, discordId string) *User {
	return &User{
		Name:      name,
		DiscordID: discordId,
	}
}

func NewLeagueOfLegends(mainRole, secondRole string, mainChamps, accounts []string) *LeagueOfLegends {
	return &LeagueOfLegends{
		MainRole:   mainRole,
		SecondRole: secondRole,
		MainChamps: mainChamps,
		Accounts:   accounts,
	}
}

func NewGameAccount(userId int, name string, game string) *GameAccount {
	return &GameAccount{
		UserId: userId,
		Name:   name,
		Game:   game,
	}
}
