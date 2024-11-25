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

func (s *SQLiteStorage) GetGameAccountBy(userId int, game string) ([]string, error) {
	accs := []string{}
	rows, err := s.db.Query("SELECT name FROM User WHERE user_id = ? AND game = ?", userId, game)
	if err != nil {
		return accs, err
	}
	defer rows.Close()
	for rows.Next() {
		var acc string
		if err := rows.Scan(&acc); err != nil {
			return accs, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}

func (s *SQLiteStorage) UpdateGameAccount(userId int, oldName string, newName string) error {
	set := map[string]any{"name": newName}
	where := map[string]any{"name": oldName, "user_id": userId}
	return s.Update("GameAccount", set, where)
}

func (s *SQLiteStorage) DeleteGameAccount(userId int, name string) error {
	return s.Delete("GameAccount", map[string]any{"name": name, "user_id": userId})
}
