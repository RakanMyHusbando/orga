package types

func NewUser(name, discordId string, id int, leagueOfLegends *LeagueOfLegends) *User {
	return &User{
		Id:              id,
		Name:            name,
		DiscordId:       discordId,
		LeagueOfLegends: leagueOfLegends,
	}
}

func NewLeagueOfLegends(mainRole, secondRole string, mainChamps, accounts []string) *LeagueOfLegends {
	resp := &LeagueOfLegends{
		MainRole:   mainRole,
		SecondRole: secondRole,
		Accounts:   accounts,
	}
	copy(resp.MainChamps[:], mainChamps)
	return resp
}

func NewGameAccount(userId int, name string, game string) *GameAccount {
	return &GameAccount{
		UserId: userId,
		Name:   name,
		Game:   game,
	}
}
