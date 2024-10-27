package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	// handler user
	CreateUser(*ReqUser) error
	GetUser() ([]*User, error)
	GetUserById(int) (*User, error)
	DeletUser(int) error
	UpdateUser(*User) error
	// handler league of legends
	CreateUserLeagueOfLegends(*ReqUserLeagueOfLegends) error
	GetUserLeagueOfLegendsById(int) (*LeagueOfLegends, error)
	DeleteUserLeagueOfLegends(int) error
	UpdateUserLeagueOfLegends(*ReqUserLeagueOfLegends) error
	// handlergame account
	CreateGameAccount(*ReqGameAccount) error
	GetGameAccountByUserId(int) ([]string, error)
	DeleteGameAccount(int) error
	UpdateGameAccount(*ReqGameAccount) error
}

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dbFile string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	// read sql file with create table statements
	byteContent, err := os.ReadFile("schema.sql")
	if err != nil {
		return nil, err
	}
	// execute sql create table statements
	queries := strings.Split(string(byteContent), ";")
	for i := range queries {
		query := strings.TrimSpace(queries[i])
		_, err := db.Exec(query)
		if err != nil {
			return nil, fmt.Errorf("%v [Table: %v]", err.Error(), strings.Split(query, " ")[5])
		}
	}

	return &SQLiteStorage{
		db: db,
	}, nil
}

/* ------------------------------ handler user ------------------------------ */

// POST
func (s *SQLiteStorage) CreateUser(user *ReqUser) error {
	prep, err := s.db.Prepare(`INSERT INTO User (name, discord_id) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	_, err = prep.Exec(user.Name, user.DiscordID)
	if err != nil {
		return err
	}

	log.Println("Storage: create user successful")

	return nil
}

// GET
func (s *SQLiteStorage) GetUser() ([]*User, error) {
	rows, err := s.db.Query(`SELECT * FROM User`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := []*User{}
	for rows.Next() {
		user := new(User)
		if err := rows.Scan(&user.Id, &user.Name, &user.DiscordID); err != nil {
			return nil, err
		}

		lolUser, err := s.GetUserLeagueOfLegendsById(user.Id)
		if err == nil {
			user.LeagueOfLegends = lolUser
		} else {
			log.Println(err)
		}

		userList = append(userList, user)
	}

	log.Println("Storage: get user successful")

	return userList, nil
}

// GET
func (s *SQLiteStorage) GetUserById(id int) (*User, error) {
	row := s.db.QueryRow(`SELECT * FROM User WHERE id = ?`, id)

	user := new(User)
	if err := row.Scan(&user.Id, &user.Name, &user.DiscordID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	lolUser, err := s.GetUserLeagueOfLegendsById(user.Id)
	if err == nil {
		user.LeagueOfLegends = lolUser
	} else {
		log.Println(err)
	}

	log.Println("Storage: get user by id successful")

	return user, nil
}

// DELETE
func (s *SQLiteStorage) DeletUser(id int) error {
	prep, err := s.db.Prepare(`DELETE FROM User WHERE id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(id); err != nil {
		return err
	}

	log.Println("Storage: delete user successful")

	return nil
}

// PUT
func (s *SQLiteStorage) UpdateUser(user *User) error {
	prep, err := s.db.Prepare(`UPDATE User SET name = ?, discord_id = ? WHERE id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(user.Name, user.DiscordID, user.Id); err != nil {
		return err
	}

	log.Println("Storage: update user successful")

	return nil
}

/* ------------------------------ handler league of legends ------------------------------ */

// POST
func (s *SQLiteStorage) CreateUserLeagueOfLegends(user *ReqUserLeagueOfLegends) error {
	prep, err := s.db.Prepare(`INSERT INTO UserLeagueOfLegends (user_id, main_role, second_role, champ_0, champ_1, champ_2) VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(user.Id, user.MainRole, user.SecondRole, user.MainChamps[0], user.MainChamps[1], user.MainChamps[2]); err != nil {
		return err
	}

	log.Println("Storage: create league of legends user successful")

	return nil
}

// GET
func (s *SQLiteStorage) GetUserLeagueOfLegendsById(userId int) (*LeagueOfLegends, error) {
	row := s.db.QueryRow(`SELECT main_role, second_role, champ_0, champ_1, champ_2 FROM UserLeagueOfLegends WHERE user_id = ?`, userId)

	userLol := new(LeagueOfLegends)
	mainChamps := []string{"", "", ""}

	if err := row.Scan(&userLol.MainRole, &userLol.SecondRole, &mainChamps[0], &mainChamps[1], &mainChamps[2]); err != nil {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, os.ErrNotExist
		}
		return nil, err
	}

	for i := range mainChamps {
		if mainChamps[i] != "" {
			userLol.MainChamps = append(userLol.MainChamps, mainChamps[i])
		}
	}

	log.Println("Storage: get league of legends user successful")

	accounts, err := s.GetGameAccountByUserId(userId)
	if err != nil {
		log.Println(err)
		return userLol, nil
	}
	userLol.Accounts = accounts

	return userLol, nil
}

// DELETE
func (s *SQLiteStorage) DeleteUserLeagueOfLegends(userId int) error {
	prep, err := s.db.Prepare(`DELETE FROM UserLeagueOfLegends WHERE user_id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(userId); err != nil {
		return err
	}

	log.Println("Storage: delete league of legends user successful")

	return nil
}

// PUT
func (s *SQLiteStorage) UpdateUserLeagueOfLegends(user *ReqUserLeagueOfLegends) error {
	prep, err := s.db.Prepare(`UPDATE UserLeagueOfLegends SET main_role = ?, second_role = ?, champ_0 = ?, champ_1 = ?, champ_2 = ? WHERE user_id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(user.MainRole, user.SecondRole, user.MainChamps[0], user.MainChamps[1], user.MainChamps[2], user.Id); err != nil {
		return err
	}

	log.Println("Storage: update league of legends user successful")

	return nil
}

/* ------------------------------ handlergame account ------------------------------ */

// POST
func (s *SQLiteStorage) CreateGameAccount(account *ReqGameAccount) error {
	prep, err := s.db.Prepare(`INSERT INTO GameAccount (user_id,name) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	_, err = prep.Exec(account.UserId, account.Name)
	if err != nil {
		return err
	}

	log.Println("Storage: create league of legends account successful")

	return nil
}

// GET
func (s *SQLiteStorage) GetGameAccountByUserId(userId int) ([]string, error) {
	rows, err := s.db.Query(`SELECT name FROM GameAccount WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []string{}
	for rows.Next() {
		var account string
		if err := rows.Scan(&account); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// DELETE
func (s *SQLiteStorage) DeleteGameAccount(userId int) error {
	prep, err := s.db.Prepare(`DELETE FROM GameAccount WHERE user_id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(userId); err != nil {
		return err
	}

	log.Println("Storage: delete league of legends account successful")

	return nil
}

// PUT
func (s *SQLiteStorage) UpdateGameAccount(account *ReqGameAccount) error {
	prep, err := s.db.Prepare(`UPDATE AccountLeagueOfLegends SET name = ? WHERE user_id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(account.Name, account.UserId); err != nil {
		return err
	}

	log.Println("Storage: update league of legends account successful")

	return nil
}

/* ------------------------------ handler team ------------------------------ */

/* ------------------------------ handler guide ------------------------------ */
