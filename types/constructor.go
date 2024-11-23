package types

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
