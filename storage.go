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
	CreateUser(*User) error
	DeletUser(int) error
	UpdateUser(*User) error
	GetUser() ([]*User, error)
	GetUserById(int) ([]*User, error)
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
			return nil, fmt.Errorf("$1 [Table: $2]", err.Error(), strings.Split(query, " ")[5])
		}
	}

	return &SQLiteStorage{
		db: db,
	}, nil
}

/* =================== Storage user handlers =================== */

func (s *SQLiteStorage) CreateUser(user *User) error {
	if _, err := s.db.Exec(
		"INSERT INTO User (name, discord_id) VALUES ($1, $2)",
		user.Name,
		user.DiscordID,
	); err != nil {
		return err
	}
	log.Println("SQLite create user successful")
	return nil
}

func (s *SQLiteStorage) DeletUser(id int) error {
	return nil
}

func (s *SQLiteStorage) UpdateUser(user *User) error {
	return nil
}

func (s *SQLiteStorage) GetUser() ([]*User, error) {
	userList := []*User{}

	row, err := s.db.Query(`SELECT * FROM User`)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		newUser, err := scanIntoUser(row)
		if err != nil {
			return nil, err
		}
		userList = append(userList, newUser)
	}

	return userList, nil
}

func (s *SQLiteStorage) GetUserById(id int) ([]*User, error) {
	userList := []*User{}

	row, err := s.db.Query(`SELECT * FROM User WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	newUser, err := scanIntoUser(row)
	if err != nil {
		return nil, err
	}

	newUser.LeagueOfLegends, err = s.GetLeagueOfLegendsUserById(id)
	if err != nil {
		return nil, err
	}

	userList = append(userList, newUser)

	return userList, nil
}

/* =================== Storage league handlers =================== */

func (s *SQLiteStorage) GetLeagueOfLegendsUserById(id int) (*LeagueOfLegends, error) {
	userLolQuery := `SELECT main_role, second_role, champ_0, champ_1, champ_2  FROM UserLeagueOfLegends WHERE user_id =` + string(id)
	AccLolQuery := `SELECT name  FROM AccountLeagueOfLegends WHERE user_id =` + string(id)

	row, err := s.db.Query(userLolQuery)
	if err != nil {
		return nil, err
	}

	lol, err := scanIntoLeagueOfLegends(row)
	if err != nil {
		return nil, err
	}

	row, err = s.db.Query(AccLolQuery)
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

	return lol, nil
}

/* =================== Storage team handlers =================== */

// TODO
