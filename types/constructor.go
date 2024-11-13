package types

func NewUser(name string, discordId string) *ResUser {
	return &ResUser{
		Name:      name,
		DiscordID: discordId,
	}
}

func NewLeagueOfLegends(
	mainRole string,
	secondRole string,
	mainChamps []string,
	accounts []string,
) *ResLeagueOfLegends {
	return &ResLeagueOfLegends{
		MainRole:   mainRole,
		SecondRole: secondRole,
		MainChamps: mainChamps,
		Accounts:   accounts,
	}
}

func NewReqGameAccount(userId int, name string, game string) *ReqGameAccount {
	return &ReqGameAccount{
		UserId: userId,
		Name:   name,
		Game:   game,
	}
}
