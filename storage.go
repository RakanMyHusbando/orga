package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	CreateUser(*CreateUser) error
	DeletUser(int) error
	UpdateUser(*User) error
	GetUser() ([]*User, error)
	GetUserById(int) (*User, error)
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

/* =================== Storage user handlers =================== */

// POST
func (s *SQLiteStorage) CreateUser(user *CreateUser) error {
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
	row, err := s.db.Query(`SELECT * FROM User`)
	if err != nil {
		return nil, err
	}

	userList := []*User{}
	for row.Next() {
		newUser, err := scanIntoUser(row)
		if err != nil {
			return nil, err
		}

		lolUser, err := s.GetLeagueOfLegendsById(&newUser.Id)
		if err == nil {
			newUser.LeagueOfLegends = lolUser
		} else {
			log.Println(err)
		}

		userList = append(userList, newUser)
	}

	log.Println("Storage: get user successful")

	return userList, nil
}

// GET
func (s *SQLiteStorage) GetUserById(id int) (*User, error) {
	row, err := s.db.Query(`SELECT * FROM User WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, fmt.Errorf("User not found")
	}

	newUser, err := scanIntoUser(row)
	if err != nil {
		return nil, err
	}

	lolUser, err := s.GetLeagueOfLegendsById(&newUser.Id)
	if err == nil {
		newUser.LeagueOfLegends = lolUser
	} else {
		log.Println(err)
	}

	log.Println("Storage: get user by id successful")

	return newUser, nil
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

/* =================== Storage league_of_legends handlers =================== */

// POST
func (s *SQLiteStorage) CreateUserLeagueOfLegends(user *CreateUserLeagueOfLegends) error {
	prep, err := s.db.Prepare(`
		INSERT INTO UserLeagueOfLegends (
			user_id, 
			main_role, 
			second_role, 
			champ_0, 
			champ_1, 
			champ_2
		) VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	_, err = prep.Exec(
		user.Id,
		user.MainRole,
		user.SecondRole,
		user.MainChamps[0],
		user.MainChamps[1],
		user.MainChamps[2],
	)
	if err != nil {
		return err
	}

	log.Println("Storage: create league of legends user successful")

	return nil
}

// POST
func (s *SQLiteStorage) CreateAccountLeagueOfLegends(account *CreateAccountLeagueOfLegends) error {
	prep, err := s.db.Prepare(`INSERT INTO AccountLeagueOfLegends (user_id,name) VALUES (?, ?)`)
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
func (s *SQLiteStorage) GetLeagueOfLegendsById(userId *int) (*LeagueOfLegends, error) {
	user, err := s.GetUserLeagueOfLegendsById(userId)
	if err != nil {
		return nil, err
	}

	accounts, err := s.GetAccountLeagueOfLegendsById(userId)
	if err != nil {
		log.Println(err)
		return user, nil
	}

	user.Accounts = accounts

	return user, nil
}

// GET
func (s *SQLiteStorage) GetUserLeagueOfLegendsById(userId *int) (*LeagueOfLegends, error) {
	row, err := s.db.Query(`SELECT main_role, second_role, champ_0, champ_1, champ_2  FROM UserLeagueOfLegends WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, fmt.Errorf("Storage: league of legends user with id %v not found", userId)
	}

	user, err := scanIntoLeagueOfLegends(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GET
func (s *SQLiteStorage) GetAccountLeagueOfLegendsById(userId *int) ([]*string, error) {
	row, err := s.db.Query(`SELECT name FROM AccountLeagueOfLegends WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}

	accounts := []*string{}
	for row.Next() {
		account := new(string)
		if err := row.Scan(account); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

/* =================== Storage team handlers =================== */

// TODO
