package storage

import (
	"encoding/json"

	"github.com/RakanMyHusbando/shogun/types"
)

func (s *SQLiteStorage) CreateUser(user *types.User) error {
	var values map[string]any
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Insert("User", values)
}

func (s *SQLiteStorage) GetUser(selectKeys []string) ([]*map[string]any, error) {
	// TODO: add User-Game relation
	return s.Select("User", selectKeys, nil)
}

func (s *SQLiteStorage) GetUserById(id int) (*map[string]any, error) {
	// TODO: add User-Game relation
	return s.SelectUnique("User", nil, "id", id)
}

func (s *SQLiteStorage) UpdateUser(user *types.User, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	return s.Patch("User", values, map[string]any{"id": id})
}

func (s *SQLiteStorage) DeleteUser(id int) error {
	return s.Delete("User", map[string]any{"id": id})
}
