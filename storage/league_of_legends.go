package storage

import (
	"encoding/json"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateLeagueOfLeagends(lol *types.LeagueOfLegends) error {
	var values map[string]any
	bytes, err := json.Marshal(lol)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("LeagueOfLegends", values)
}

func (s *SQLiteStorage) GetLeagueOfLegends() ([]*map[string]any, error) {
	return s.Select("LeagueOfLegends", nil, nil)
}

func (s *SQLiteStorage) GetLeagueOfLegendsByUserId(userId int) ([]*map[string]any, error) {
	return s.SelectUnique("LeagueOfLegends", nil, "user_id", userId)
}

func (s *SQLiteStorage) DeleteLeagueOfLegends(userId int) error {
	return s.Delete("LeagueOfLegends", map[string]any{"user_id": userId})
}

func (s *SQLiteStorage) UpdateLeagueOfLegends(lol *types.LeagueOfLegends, userId int) error {
	var values map[string]any
	bytes, err := json.Marshal(lol)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Update("LeagueOfLegends", values, map[string]any{"user_id": userId})
}
