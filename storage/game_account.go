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

func (s *SQLiteStorage) DeleteGameAccount(account *types.GameAccount) error {
	return s.Delete("GameAccount", map[string]any{"name": account.Name, "user_id": account.UserId})
}

func (s *SQLiteStorage) UpdateGameAccount(account *types.GameAccount, newAccountName string) error {
	var values map[string]any
	bytes, err := json.Marshal(account)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Patch("GameAccount", map[string]any{"name": newAccountName}, values)
}
