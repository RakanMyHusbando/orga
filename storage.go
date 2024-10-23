package main

import (
	"database/sql"
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
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{
		db: db,
	}, nil
}

func CreateSQLiteTable(sqlite *SQLiteStorage) error {
	querys := []string{
		`CREATE TABLE IF NOT EXISTS User (
			id INTEGER PRIMARY KEY,
			name TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS AccountLeagueOfLegends (
			user_id INTEGER FOREIGN KEY REFERENCES User(id),
			name TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS UserLeagueOfLegends (
			user_id INTEGER FOREIGN KEY REFERENCES User(id),
			main_role TEXT,
			second_role TEXT,
			champ_0 TEXT,
			champ_1 TEXT,
			champ_2 TEXT
		)`,
	}
	for i := range querys {
		_, err := sqlite.db.Exec(querys[i])
		if err != nil {
			return err
		}
	}
	return nil
}
