package storage

import (
	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateGameAccount(account *types.GameAccount) error {
	return s.Insert("GameAccount", map[string]any{
		"user_id": account.UserId,
		"game":    account.Game,
		"name":    account.Name,
	})
}

func (s *SQLiteStorage) GetGameAccountBy(userId int, game string) ([]string, error) {
	accs := []string{}
	rows, err := s.db.Query("SELECT name FROM GameAccount WHERE user_id = ? AND game = ?", userId, game)
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
	return s.Update(
		"GameAccount",
		map[string]any{"name": newName},
		map[string]any{"name": oldName, "user_id": userId},
	)
}

func (s *SQLiteStorage) DeleteGameAccount(userId int, name string) error {
	return s.Delete("GameAccount", map[string]any{"name": name, "user_id": userId})
}
