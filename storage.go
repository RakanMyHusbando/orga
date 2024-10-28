package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	// handler user
	CreateUser(*ReqUser) error
	GetUser() ([]*User, error)
	GetUserById(int) (*User, error)
	GetUserIds() ([]*int, error)
	DeletUser(int) error
	UpdateUser(*User) error

	// handler league of legends
	CreateUserLeagueOfLegends(*ReqUserLeagueOfLegends) error
	GetUserLeagueOfLegendsById(int) (*LeagueOfLegends, error)
	DeleteUserLeagueOfLegends(int) error
	UpdateUserLeagueOfLegends(*ReqUserLeagueOfLegends) error

	// handlergame account
	CreateGameAccount(*ReqGameAccount) error
	GetGameAccountByUserId(int, string) ([]string, error)
	DeleteGameAccount(*ReqGameAccount) error
	UpdateGameAccount(*ReqUpdateGameAccount) error
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

/* ------------------------------ handler user ------------------------------ */

// POST
func (s *SQLiteStorage) CreateUser(user *ReqUser) error {
	prep, err := s.db.Prepare(`INSERT INTO User (name, discord_id) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	_, err = prep.Exec(user.Name, user.DiscordID)
	if err != nil {
		return err
	}

	log.Printf("Storage: successfully create user %v", user.Name)

	return nil
}

// GET
func (s *SQLiteStorage) GetUser() ([]*User, error) {
	rows, err := s.db.Query(`SELECT * FROM User`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := []*User{}
	for rows.Next() {
		user := new(User)
		if err := rows.Scan(&user.Id, &user.Name, &user.DiscordID); err != nil {
			return nil, err
		}

		lolUser, err := s.GetUserLeagueOfLegendsById(user.Id)
		if err == nil {
			user.LeagueOfLegends = lolUser
		} else {
			log.Println(err)
		}

		userList = append(userList, user)
	}

	log.Println("Storage: successfully get user")

	return userList, nil
}

// GET
func (s *SQLiteStorage) GetUserById(id int) (*User, error) {
	row := s.db.QueryRow(`SELECT * FROM User WHERE id = ?`, id)

	user := new(User)
	if err := row.Scan(&user.Id, &user.Name, &user.DiscordID); err != nil {
		return nil, err
	}

	lolUser, err := s.GetUserLeagueOfLegendsById(user.Id)
	if err == nil {
		user.LeagueOfLegends = lolUser
	} else {
		log.Println(err)
	}

	log.Printf("Storage: successfully get user with id %v", id)

	return user, nil
}

// GET
func (s *SQLiteStorage) GetUserIds() ([]*int, error) {
	rows, err := s.db.Query(`SELECT id FROM User`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := []*int{}
	for rows.Next() {
		user := new(int)
		if err := rows.Scan(&user); err != nil {
			return nil, err
		}

		userList = append(userList, user)
	}

	log.Println("Storage: successfully get user ids")

	return userList, nil
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

	log.Printf("Storage: successfully delete user with id %v", id)

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

	log.Printf("Storage: successfully update user with id %v", user.Id)

	return nil
}

/* ------------------------------ handler league of legends ------------------------------ */

// POST
func (s *SQLiteStorage) CreateUserLeagueOfLegends(user *ReqUserLeagueOfLegends) error {
	insertKeys := "user_id, main_role, second_role"
	insertValues := strconv.Itoa(user.Id) + ", '" + user.MainRole + "', '" + user.SecondRole + "'"

	if user.MainChamps != nil {
		for i := range user.MainChamps {
			insertKeys += ", champ_" + strconv.Itoa(i)
			insertValues += ", '" + user.MainChamps[i] + "'"
		}
	}

	prep, err := s.db.Prepare(
		"INSERT INTO UserLeagueOfLegends (" + insertKeys + ") VALUES (" + insertValues + ")",
	)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(); err != nil {
		return err
	}

	log.Printf("Storage: successfully added league_of_legends to user with id %v", user.Id)

	return nil
}

// GET
func (s *SQLiteStorage) GetUserLeagueOfLegendsById(userId int) (*LeagueOfLegends, error) {
	row := s.db.QueryRow(`
		SELECT 
			main_role, 
			second_role, 
			IFNULL(champ_0, ''), 
			IFNULL(champ_1, ''), 
			IFNULL(champ_2, '') 
		FROM 
			UserLeagueOfLegends 
		WHERE 
			user_id = ?`,
		userId,
	)

	userLol := new(LeagueOfLegends)
	mainChamps := []string{"", "", ""}

	if err := row.Scan(
		&userLol.MainRole,
		&userLol.SecondRole,
		&mainChamps[0], &mainChamps[1],
		&mainChamps[2],
	); err != nil {
		return nil, err
	}

	userLol.MainChamps = []string{}
	for i := range mainChamps {
		if mainChamps[i] != "" {
			userLol.MainChamps = append(userLol.MainChamps, mainChamps[i])
		}
	}

	log.Printf("Storage: successfully get league_of_legends from user with id %v", userId)

	accounts, err := s.GetGameAccountByUserId(userId, "league_of_legends")
	if err != nil {
		log.Println(err)
		return userLol, nil
	}
	userLol.Accounts = accounts

	return userLol, nil
}

// DELETE
func (s *SQLiteStorage) DeleteUserLeagueOfLegends(userId int) error {
	log.Println(userId)
	prep, err := s.db.Prepare(`DELETE FROM UserLeagueOfLegends WHERE user_id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(userId); err != nil {
		return err
	}

	log.Printf("Storage: successfully delete league_of_legends from user with id %v", userId)

	return nil
}

// PUT
func (s *SQLiteStorage) UpdateUserLeagueOfLegends(user *ReqUserLeagueOfLegends) error {
	prep, err := s.db.Prepare(`
		UPDATE 
			UserLeagueOfLegends 
		SET 
			main_role = ?, 
			second_role = ?, 
			champ_0 = ?, 
			champ_1 = ?, 
			champ_2 = ? 
		WHERE 
			user_id = ?
	`)

	if err != nil {
		return err
	}

	if _, err = prep.Exec(
		user.MainRole,
		user.SecondRole,
		user.MainChamps[0],
		user.MainChamps[1],
		user.MainChamps[2],
		user.Id,
	); err != nil {
		return err
	}

	log.Printf("Storage: successfully update league of legends user with id %v", user.Id)

	return nil
}

/* ------------------------------ handlergame account ------------------------------ */

// POST
func (s *SQLiteStorage) CreateGameAccount(account *ReqGameAccount) error {
	prep, err := s.db.Prepare(`INSERT INTO GameAccount (user_id, game, name) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = prep.Exec(account.UserId, account.Game, account.Name) // TODO
	if err != nil {
		return err
	}

	log.Printf(
		"Storage: successfully create  %v account for user with id %v",
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
func (s *SQLiteStorage) DeleteGameAccount(reqAccount *ReqGameAccount) error {
	prep, err := s.db.Prepare(`DELETE FROM GameAccount WHERE user_id = ? AND name = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(reqAccount.UserId, reqAccount.Name); err != nil {
		return err
	}

	log.Println("Storage: successfully delete game account")

	return nil
}

// PUT
func (s *SQLiteStorage) UpdateGameAccount(reqUpdateAccount *ReqUpdateGameAccount) error {
	row, err := s.db.Query(
		`SELECT name FROM GameAccount WHERE user_id = ? AND name = ?`,
		reqUpdateAccount.UserId,
		reqUpdateAccount.NameOld,
	)
	if err != nil {
		return err
	}

	if !row.Next() {
		return fmt.Errorf(
			"account '%v' from user with id %v not found",
			reqUpdateAccount.NameOld,
			reqUpdateAccount.UserId,
		)
	}

	prep, err := s.db.Prepare(`UPDATE GameAccount SET name = ? WHERE user_id = ? AND name = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(
		reqUpdateAccount.NameNew, reqUpdateAccount.UserId, reqUpdateAccount.NameOld,
	); err != nil {
		return err
	}

	log.Println("Storage: update league of legends account successful")

	return nil
}

/* ------------------------------ handler team ------------------------------ */

/* ------------------------------ handler guide ------------------------------ */
