package storage

import (
	"fmt"
	"strings"

	"github.com/Kinveil/Riot-API-Golang/constants/continent"
	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateGameAccount(account *types.GameAccount) error {
	switch account.Game {
	case "league_of_legends":
		return s.CreateLeagueOfLegendsAcc(account)
	default:
		return fmt.Errorf("Game %v not supported", account.Game)
	}

}

func (s *SQLiteStorage) GetGameAccountByUserId(userId int, game string) ([]string, error) {
	puuidLst := []*string{}
	rows, err := s.db.Query("SELECT puuid FROM GameAccount WHERE user_id = ? AND game = ?", userId, game)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		puuid := new(string)
		if err := rows.Scan(&puuid); err != nil {
			return nil, err
		}
		puuidLst = append(puuidLst, puuid)
	}
	switch game {
	case "league_of_legends":
		return s.GetLeagueOfLegendsAcc(puuidLst)
	default:
		return nil, fmt.Errorf("Game %v not supported", game)
	}
}

func (s *SQLiteStorage) DeleteGameAccount(userId int, puuid string) error {
	return s.Delete("GameAccount", map[string]any{"user_id": userId, "puuid": puuid})
}

/* --------------------------------- league_of_legends --------------------------------- */

func (s *SQLiteStorage) CreateLeagueOfLegendsAcc(account *types.GameAccount) error {
	name := strings.Split(account.Name, "#")
	riotAcc, err := apiClient.GetAccountByRiotID(continent.EUROPE, name[0], name[1])
	if err != nil {
		return err
	}
	return s.Insert("GameAccount", map[string]any{
		"user_id": account.UserId,
		"game":    account.Game,
		"puuid":   riotAcc.Puuid,
	})
}

func (s *SQLiteStorage) GetLeagueOfLegendsAcc(puuids []*string) ([]string, error) {
	accounts := []string{}
	for _, puuid := range puuids {
		riotAcc, err := apiClient.GetAccountByPuuid(continent.EUROPE, *puuid)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, riotAcc.Puuid)
	}
	return accounts, nil
}
