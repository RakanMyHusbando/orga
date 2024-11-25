package storage

import (
	"encoding/json"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateLeagueOfLeagends(lol *types.LeagueOfLegends, userId int) error {
	return s.Insert("UserLeagueOfLegends", map[string]any{
		"user_id":     userId,
		"main_role":   lol.MainRole,
		"second_role": lol.SecondRole,
		"champ_0":     lol.MainChamps[0],
		"champ_1":     lol.MainChamps[1],
		"champ_2":     lol.MainChamps[2],
	})
}

func (s *SQLiteStorage) GetLeagueOfLegendsByUserId(userId int) (*types.LeagueOfLegends, error) {
	row := s.db.QueryRow("SELECT main_role, second_role, champ_0, champ_1, champ_2 FROM UserLeagueOfLegends WHERE user_id = ?", userId)
	lol := new(types.LeagueOfLegends)
	if err := row.Scan(
		&lol.MainChamps,
		&lol.SecondRole,
		&lol.MainChamps[0],
		&lol.MainChamps[1],
		&lol.MainChamps[2],
	); err != nil {
		return nil, err
	}
	return lol, nil
}

func (s *SQLiteStorage) DeleteLeagueOfLegends(userId int) error {
	return s.Delete("UserLeagueOfLegends", map[string]any{"user_id": userId})
}

func (s *SQLiteStorage) UpdateLeagueOfLegends(lol *types.LeagueOfLegends, userId int) error {
	var values map[string]any
	bytes, err := json.Marshal(lol)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Update("UserLeagueOfLegends", values, map[string]any{"user_id": userId})
}
