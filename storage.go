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

		lolUser, err := s.GetLeagueOfLegendsUserById(newUser.Id)
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

	lolUser, err := s.GetLeagueOfLegendsUserById(newUser.Id)
	if err == nil {
		newUser.LeagueOfLegends = lolUser
	} else {
		log.Println(err)
	}

	log.Println("Storage: get user by id successful")

	return newUser, nil
}

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

func (s *SQLiteStorage) GetLeagueOfLegendsUserById(id int) (*LeagueOfLegends, error) {
	row, err := s.db.Query(`SELECT main_role, second_role, champ_0, champ_1, champ_2  FROM UserLeagueOfLegends WHERE user_id = ?`, id)
	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, fmt.Errorf("Storage: no league of legends user found")
	}

	lol, err := scanIntoLeagueOfLegends(row)
	if err != nil {
		return nil, err
	}

	row, err = s.db.Query(`SELECT name  FROM AccountLeagueOfLegends WHERE user_id = ?`, id)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var name string
		if err := row.Scan(&name); err != nil {
			return nil, err
		}
		lol.Accounts = append(lol.Accounts, name)
	}

	log.Println("Storage: get league of legends user by id successful")

	return lol, nil
}

/* =================== Storage team handlers =================== */

// TODO
