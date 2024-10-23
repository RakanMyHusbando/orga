package main

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	CreateUser(*User) error
	DeletUser(int) error
	UpdateUser(*User) error
	GetUserById(int) (*User, error)
	GetUserByName(int) (*User, error)
}

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dbFile string) (*SQLiteStorage, error) {
	querys := []string{
		`CREATE TABLE IF NOT EXISTS User (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS AccountLeagueOfLegends (
			user_id INTEGER NOT NULL,
			name TEXT,
			FOREIGN KEY (user_id) REFERENCES User(id)
		)`,
		`CREATE TABLE IF NOT EXISTS UserLeagueOfLegends (
			user_id INTEGER NOT NULL UNIQUE,
			main_role TEXT,
			second_role TEXT,
			champ_0 TEXT,
			champ_1 TEXT,
			champ_2 TEXT,
			FOREIGN KEY (user_id) REFERENCES User(id)
		)`,
	}
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	for i := range querys {
		_, err := db.Exec(querys[i])
		if err != nil {
			newErr := "Error: " + err.Error() + "\nQuery:" + querys[i]
			return nil, errors.New(newErr)
		}
	}
	defer db.Close()

	return &SQLiteStorage{
		db: db,
	}, nil
}
