package storage

import (
	"log"

	"github.com/RakanMyHusbando/shogun/types"
)

// POST
func (s *SQLiteStorage) CreateGameAccount(account *types.ReqGameAccount) error {
	prep, err := s.db.Prepare(`INSERT INTO GameAccount (user_id, game, name) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(account.UserId, account.Game, account.Name); err != nil {
		return err
	}
	prep.Close()

	log.Printf(
		"Storage: successfully create %v account from user with id %v",
		account.Game,
		account.UserId,
	)

	return nil
}

// GET
func (s *SQLiteStorage) GetGameAccountByUserId(userId int, game string) ([]string, error) {
	rows, err := s.db.Query(`SELECT name FROM GameAccount WHERE user_id = ? AND game = ?`, userId, game)
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
func (s *SQLiteStorage) DeleteGameAccount(account *types.ReqGameAccount) error {
	prep, err := s.db.Prepare(`DELETE FROM GameAccount WHERE user_id = ? AND name = ?`)
	if err != nil {
		return err
	}
	if _, err = prep.Exec(account.UserId, account.Name); err != nil {
		return err
	}
	prep.Close()

	log.Printf("Storage: successfully deleted %v account (%v) from user with id %v", account.Game, account.Name, account.UserId)

	return nil
}

// PATCH
func (s *SQLiteStorage) UpdateGameAccount(account *types.ReqGameAccount, oldName string) error {
	prep, err := s.db.Prepare(`UPDATE GameAccount SET name = ? WHERE user_id = ? AND name = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(
		account.Name,
		account.UserId,
		oldName,
	); err != nil {
		return err
	}

	prep.Close()

	log.Printf("Storage: successfully updated %v account from user with id %v", account.Game, account.UserId)

	return nil
}
