package storage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateLeagueOfLegends(lol *types.LeagueOfLegends) error {
	var values map[string]any
	bytes, err := json.Marshal(lol)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("UserLeagueOfLegends", values)
}

func (s *SQLiteStorage) GetLeagueOfLegendsById(userId int) (*map[string]any, error) {
	return s.SelectUnique("UserLeagueOfLegends", nil, "user_id", userId)
}

// PATCH
func (s *SQLiteStorage) UpdateLeagueOfLegends(lol *types.ReqLeagueOfLegends) error {
	var champs string
	if lol.MainChamps != nil {
		for i := range lol.MainChamps {
			champs += fmt.Sprintf(", champ_%d = '%s'", i, lol.MainChamps[i])
		}
	}
	query := fmt.Sprintf(
		`UPDATE UserLeagueOfLegends SET main_role = '%s', second_role = '%s' %s WHERE user_id = %d`,
		lol.MainRole, lol.SecondRole, champs, lol.UserId,
	)

	prep, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(); err != nil {
		return err
	}

	prep.Close()

	log.Printf("Storage: successfully update league of legends user with id %v", lol.UserId)

	return nil
}
