package storage

import (
	"encoding/json"

	"github.com/RakanMyHusbando/orga/types"
)

func (s *SQLiteStorage) CreateUser(user *types.User) error {
	return s.Insert("User", map[string]any{
		"name":       user.Name,
		"discord_id": user.DiscordId,
	})
}

func (s *SQLiteStorage) GetUser() ([]*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var userLst []*types.User
	for rows.Next() {
		user := new(types.User)
		if err := rows.Scan(
			&user.Id, &user.Name, &user.DiscordId, &user.SessionCoockie, &user.CSRFToken,
		); err != nil {
			return nil, err
		}
		userLst = append(userLst, user)
	}
	return userLst, nil
}

func (s *SQLiteStorage) GetUserById(id int) ([]*types.User, error) {
	row := s.db.QueryRow("SELECT * FROM User WHERE id = ?", id)
	user := new(types.User)
	if err := row.Scan(
		&user.Id, &user.Name, &user.DiscordId, &user.SessionCoockie, &user.CSRFToken,
	); err != nil {
		return nil, err
	}
	return []*types.User{user}, nil
}

func (s *SQLiteStorage) UpdateUser(user *types.User, id int) error {
	var values map[string]any
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	values["id"] = nil
	return s.Update("User", values, map[string]any{"id": user.Id})
}

func (s *SQLiteStorage) DeleteUser(id int) error {
	return s.Delete("User", map[string]any{"id": id})
}
