package storage

import (
	"encoding/json"
	"log"

	"github.com/RakanMyHusbando/shogun/types"
)

// POST
func (s *SQLiteStorage) CreateUser(user *types.ReqUser) error {
	var values map[string]any
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &values)
	if err = s.Insert("User", values); err != nil {
		return err
	}
	return nil
}

// GET
func (s *SQLiteStorage) GetUser() ([]*types.ResUser, error) {
	rows, err := s.db.Query(`SELECT * FROM User`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := []*types.ResUser{}
	for rows.Next() {
		user := new(types.ResUser)
		if err := rows.Scan(&user.Id, &user.Name, &user.DiscordID); err != nil {
			return nil, err
		}

		lolUser, err := s.GetLeagueOfLegendsWithAccountsById(user.Id)
		if err == nil {
			user.LeagueOfLegends = lolUser
		} else {
			log.Println(err)
		}

		userList = append(userList, user)
	}

	log.Println("Storage: successfully get user")

	return userList, nil
}

// GET
func (s *SQLiteStorage) GetUserById(id int) (*types.ResUser, error) {
	row := s.db.QueryRow(`SELECT * FROM User WHERE id = ?`, id)

	user := new(types.ResUser)
	if err := row.Scan(&user.Id, &user.Name, &user.DiscordID); err != nil {
		return nil, err
	}

	lolUser, err := s.GetLeagueOfLegendsWithAccountsById(user.Id)
	if err == nil {
		user.LeagueOfLegends = lolUser
	} else {
		log.Println(err)
	}

	log.Printf("Storage: successfully get user with id %v", id)

	return user, nil
}

// GET
func (s *SQLiteStorage) GetUserIds() ([]*int, error) {
	rows, err := s.db.Query(`SELECT id FROM User`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := []*int{}
	for rows.Next() {
		user := new(int)
		if err := rows.Scan(&user); err != nil {
			return nil, err
		}

		userList = append(userList, user)
	}

	log.Println("Storage: successfully get user ids")

	return userList, nil
}

// PATCH
func (s *SQLiteStorage) UpdateUser(user *types.ResUser) error {
	if user.Name == "" && user.DiscordID == "" {
		oldUser, err := s.GetUserById(user.Id)
		if err != nil {
			return err
		}
		if user.Name == "" {
			user.Name = oldUser.Name
		}
		if user.DiscordID == "" {
			user.DiscordID = oldUser.DiscordID
		}
	}

	prep, err := s.db.Prepare(`UPDATE User SET name = ?, discord_id = ? WHERE id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(user.Name, user.DiscordID, user.Id); err != nil {
		return err
	}

	prep.Close()

	log.Printf("Storage: successfully update user with id %v", user.Id)

	return nil
}
