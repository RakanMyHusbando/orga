package main

import (
	"database/sql"
	"errors"
	"os"
	"strings"

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
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	byteContent, err := os.ReadFile("schema.sql")
	if err != nil {
		return nil, err
	}

	queries := strings.Split(string(byteContent), ";")
	for i := range queries {
		query := strings.TrimSpace(queries[i])
		_, err := db.Exec(query)
		if err != nil {
			return nil, errors.New(err.Error() + " [Table:" + strings.Split(query, " ")[5] + "]")
		}
	}

	defer db.Close()

	return &SQLiteStorage{
		db: db,
	}, nil
}
