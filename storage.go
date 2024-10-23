package main

import (
	"database/sql"
	"errors"
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

	// defer db.Close()

	return &SQLiteStorage{
		db: db,
	}, nil
}

func (s *SQLiteStorage) CreateUser(user *User) error {
	resp, err := s.db.Exec(`INSERT INTO User (name, discord_id) VALUES ('$1', '$2')`, user.Name, user.DiscordID)
	if err != nil {
		return err
	}
	log.Printf("sqlite create user response: %v", resp)
	return nil
}

func (s *SQLiteStorage) DeletUser(id int) error {
	return nil
}

func (s *SQLiteStorage) UpdateUser(user *User) error {
	return nil
}

func (s *SQLiteStorage) GetUser() ([]*User, error) {
	return nil, nil
}

func (s *SQLiteStorage) GetUserById(id int) (*User, error) {
	return nil, nil
}
