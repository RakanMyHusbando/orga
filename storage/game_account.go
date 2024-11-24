package storage

import (
	"encoding/json"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateGameAccount(account *types.GameAccount) error {
	var values map[string]any
	bytes, err := json.Marshal(account)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("GameAccount", values)
}

func (s *SQLiteStorage) GetGameAccountBy(account *types.GameAccount) ([]*map[string]any, error) {
	var searcParams map[string]any
	bytes, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &searcParams)
	if searcParams["userId"] == 0 {
		searcParams["userId"] = nil
	}
	return s.Select("GameAccount", []string{"name"}, searcParams)
}

func (s *SQLiteStorage) UpdateGameAccount(userId int, oldName string, newName string) error {
	set := map[string]any{"name": newName}
	where := map[string]any{"name": oldName, "user_id": userId}
	return s.Update("GameAccount", set, where)
}

func (s *SQLiteStorage) DeleteGameAccount(userId int, name string) error {
	return s.Delete("GameAccount", map[string]any{"name": name, "user_id": userId})
}
