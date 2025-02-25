package website

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewSessionStorage() (*SessionStorage, error) {
	db, err := sql.Open("sqlite3", "sessions.db")
	if err != nil {
		return nil, err
	}
	query := `CREATE TABLE IF NOT EXISTS Session (
		discord_id TEXT NOT NULL UNIQUE,
		ip TEXT NOT NULL,
		session_token TEXT NOT NULL UNIQUE
	)`
	if _, err = db.Exec(query); err != nil {
		return nil, err
	}
	return &SessionStorage{db}, nil
}

type UserSessionStorage interface {
	Insert(user *User) error
	Select(discordId string) (*User, error)
	Delete(sessionToken, csrfToken string) error
}

type SessionStorage struct {
	*sql.DB
}

func (s *SessionStorage) Insert(user *User) error {
	prep, err := s.Prepare("INSERT INTO Session (discord_id, ip, session_token) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement")
	}
	defer prep.Close()
	if _, err = prep.Exec(&user.DiscordId, &user.Ip, &user.SessionToken); err != nil {
		return fmt.Errorf("failed to insert session: %v", err)
	}
	return nil
}

func (s *SessionStorage) Select(discordId string) (*User, error) {
	prep, err := s.Prepare("SELECT discord_id, ip, session_token FROM Session WHERE discord_id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select statement")
	}
	defer prep.Close()
	row := prep.QueryRow(discordId)
	user := &User{}
	if err = row.Scan(&user.DiscordId, &user.Ip, &user.SessionToken); err != nil {
		return nil, fmt.Errorf("failed to select session: %v", err)
	}
	return user, nil
}

func (s *SessionStorage) Delete(discord_id string) error {
	prep, err := s.Prepare("DELETE FROM Session WHERE discord_id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare delete statement")
	}
	defer prep.Close()
	if _, err = prep.Exec(discord_id); err != nil {
		return fmt.Errorf("failed to delete session: %v", err)
	}
	return nil
}
