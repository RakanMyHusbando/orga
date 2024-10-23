package main

type User struct {
	Id    int
	Name  string
	Games *Games
}

type Games struct {
	LeagueOfLegends *LeagueOfLegends
}

type LeagueOfLegends struct {
	MainRole   string
	SecondRole string
	MainChamps []string
	Accounts   []string
}

func NewLeagueOfLegends(mainRole string, secondRole string, mainChamps []string, accounts []string) *LeagueOfLegends {
	return &LeagueOfLegends{
		MainRole:   mainRole,
		SecondRole: secondRole,
		MainChamps: mainChamps,
		Accounts:   accounts,
	}
}

func NewUser(id int, name string, lol *LeagueOfLegends) *User {
	return &User{
		Id:   id,
		Name: name,
		Games: &Games{
			LeagueOfLegends: lol,
		},
	}
}
