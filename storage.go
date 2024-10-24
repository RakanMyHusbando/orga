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

	rows, err := s.db.Query(`SELECT * FROM User`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		newUser, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		userList = append(userList, newUser)
	}

	return userList, nil
}

func (s *SQLiteStorage) GetUserById(id int) ([]*User, error) {
	userList := []*User{}

	rows, err := s.db.Query(`SELECT * FROM User WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	newUser, err := scanIntoUser(rows)
	if err != nil {
		return nil, err
	}

	userList = append(userList, newUser)

	return userList, nil
}
