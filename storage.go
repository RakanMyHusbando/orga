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
	// user
	CreateUser(*ReqUser) error
	GetUser() ([]*ResUser, error)
	GetUserById(int) (*ResUser, error)
	GetUserIds() ([]*int, error)
	DeletUser(int) error
	UpdateUser(*ResUser) error

	// league of legends
	CreateUserLeagueOfLegends(*ReqUserLeagueOfLegends) error
	GetUserLeagueOfLegendsById(int) (*ResLeagueOfLegends, error)
	GetUserLeagueOfLegendsWithAccountsById(int) (*ResLeagueOfLegends, error)
	DeleteUserLeagueOfLegends(int) error
	UpdateUserLeagueOfLegends(*ReqUserLeagueOfLegends) error

	// game account
	CreateGameAccount(*ReqGameAccount) error
	GetGameAccountByUserId(int, string) ([]string, error)
	DeleteGameAccount(*ReqGameAccount) error
	UpdateGameAccount(*ReqUpdateGameAccount) error

	// guild
	CreateGuild(*ReqGuild) error
	GetGuild() ([]*ResGuild, error)
	GetGuildById(int) (*ResGuild, error)
	DeleteGuild(int) error
	UpdateGuild(*ResGuild) error

	// guild user
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

/* ------------------------------ user ------------------------------ */

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
func (s *SQLiteStorage) GetUser() ([]*ResUser, error) {
	rows, err := s.db.Query(`SELECT * FROM User`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := []*ResUser{}
	for rows.Next() {
		user := new(ResUser)
		if err := rows.Scan(&user.Id, &user.Name, &user.DiscordID); err != nil {
			return nil, err
		}

		lolUser, err := s.GetUserLeagueOfLegendsWithAccountsById(user.Id)
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
func (s *SQLiteStorage) GetUserById(id int) (*ResUser, error) {
	row := s.db.QueryRow(`SELECT * FROM User WHERE id = ?`, id)

	user := new(ResUser)
	if err := row.Scan(&user.Id, &user.Name, &user.DiscordID); err != nil {
		return nil, err
	}

	lolUser, err := s.GetUserLeagueOfLegendsWithAccountsById(user.Id)
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
func (s *SQLiteStorage) UpdateUser(user *ResUser) error {
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

/* ------------------------------ league of legends ------------------------------ */

// POST
func (s *SQLiteStorage) CreateUserLeagueOfLegends(user *ReqUserLeagueOfLegends) error {
	insertKeys := "user_id, main_role, second_role"
	insertValues := strconv.Itoa(user.UserId) + ", '" + user.MainRole + "', '" + user.SecondRole + "'"

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

	log.Printf("Storage: successfully added league_of_legends to user with id %v", user.UserId)

	return nil
}

// GET
func (s *SQLiteStorage) GetUserLeagueOfLegendsById(userId int) (*ResLeagueOfLegends, error) {
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

	userLol := new(ResLeagueOfLegends)
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

	return userLol, nil
}

// GET
func (s *SQLiteStorage) GetUserLeagueOfLegendsWithAccountsById(userId int) (*ResLeagueOfLegends, error) {
	userLol, err := s.GetUserLeagueOfLegendsById(userId)
	if err != nil {
		return nil, err
	}

	accounts, err := s.GetGameAccountByUserId(userId, "league_of_legends")
	if err != nil {
		log.Println(err)
		userLol.Accounts = []string{}
	} else {
		userLol.Accounts = accounts
	}

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
		user.UserId,
	); err != nil {
		return err
	}

	log.Printf("Storage: successfully update league of legends user with id %v", user.UserId)

	return nil
}

/* ------------------------------ game account ------------------------------ */

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
		reqUpdateAccount.NameNew,
		reqUpdateAccount.UserId,
		reqUpdateAccount.NameOld,
	); err != nil {
		return err
	}

	log.Println("Storage: update league of legends account successful")

	return nil
}

/* ------------------------------ guild ------------------------------ */

// POST
func (s *SQLiteStorage) CreateGuild(Guild *ReqGuild) error {
	prep, err := s.db.Prepare(`INSERT INTO Guild (name,abbreviation,description) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(
		Guild.Name,
		Guild.Abbreviation,
		Guild.Description,
	); err != nil {
		return err
	}

	return nil
}

// GET
func (s *SQLiteStorage) GetGuild() ([]*ResGuild, error) {
	rows, err := s.db.Query(`SELECT * FROM Guild`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	guilds := []*ResGuild{}
	for rows.Next() {
		guild := new(ResGuild)
		if err := rows.Scan(
			&guild.Id,
			&guild.Name,
			&guild.Abbreviation,
			&guild.Description,
		); err != nil {
			return nil, err
		}
		roleUserList, err := s.GetGuildRoleNameUserNameByGuildId(guild.Id)
		if err != nil {
			return nil, err
		}
		guild.Member = roleUserList

		guilds = append(guilds, guild)
	}

	log.Println("Storage: successfully get guilds")

	return guilds, nil
}

// GET
func (s *SQLiteStorage) GetGuildById(id int) (*ResGuild, error) {
	row := s.db.QueryRow(`SELECT * FROM Guild WHERE id = ?`, id)

	guild := new(ResGuild)
	if err := row.Scan(
		&guild.Id,
		&guild.Name,
		&guild.Abbreviation,
		&guild.Description,
	); err != nil {
		return nil, err
	}

	roleUserList, err := s.GetGuildRoleNameUserNameByGuildId(guild.Id)
	if err != nil {
		return nil, err
	}
	guild.Member = roleUserList

	return guild, nil
}

// DELETE
func (s *SQLiteStorage) DeleteGuild(id int) error {
	prep, err := s.db.Prepare(`DELETE FROM Guild WHERE id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(id); err != nil {
		return err
	}

	return nil
}

// PUT
func (s *SQLiteStorage) UpdateGuild(guild *ResGuild) error {
	prep, err := s.db.Prepare(`
		UPDATE 
			Guild 
		SET 
			name = ?, 
			abbreviation = ?, 
			description = ? 
		WHERE 
			id = ?
	`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(
		guild.Name,
		guild.Abbreviation,
		guild.Description,
		guild.Id,
	); err != nil {
		return err
	}

	return nil
}

/* ------------------------------ guild role ------------------------------ */

// POST
func (s *SQLiteStorage) CreateGuildRole(guildRole *ReqGuildRole) error {
	prep, err := s.db.Prepare(`INSERT INTO GuildRole (name, description) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(guildRole.Name, guildRole.Description); err != nil {
		return err
	}

	return nil
}

// GET
func (s *SQLiteStorage) GetGuildRoleById(id int) (*ReqGuildRole, error) {
	row := s.db.QueryRow(`SELECT name, description FROM GuildRole WHERE id = ?`, id)

	guildRole := new(ReqGuildRole)
	if err := row.Scan(&guildRole.Name, &guildRole.Description); err != nil {
		return nil, err
	}

	return guildRole, nil
}

// DELETE
func (s *SQLiteStorage) DeleteGuildRole(id int) error {
	prep, err := s.db.Prepare(`DELETE FROM GuildRole WHERE id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(id); err != nil {
		return err
	}

	return nil
}

// PUT
func (s *SQLiteStorage) UpdateGuildRole(guildRole *ReqGuildRole) error {
	prep, err := s.db.Prepare(`
		UPDATE 
			GuildRole 
		SET 
			name = ?, 
			description = ? 
		WHERE 
			id = ?
	`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(guildRole.Name, guildRole.Description, guildRole.Id); err != nil {
		return err
	}

	return nil
}

/* ------------------------------ guild user ------------------------------ */

// POST
func (s *SQLiteStorage) CreateGuildUser(guildUser *ReqGuildUser) error {
	prep, err := s.db.Prepare(`INSERT INTO GuildMember (user_id, guild_id, role_id) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(guildUser.UserId, guildUser.GuildId, guildUser.RoleId); err != nil {
		return err
	}

	return nil
}

// GET
func (s *SQLiteStorage) GetGuildUserByGuildId(guildId int) ([]*ReqGuildUser, error) {
	rows, err := s.db.Query(`SELECT user_id, role_id FROM GuildMember WHERE guild_id = ?`, guildId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	guildUsers := []*ReqGuildUser{}
	for rows.Next() {
		guildUser := new(ReqGuildUser)
		if err := rows.Scan(&guildUser.UserId, &guildUser.RoleId); err != nil {
			return nil, err
		}
		guildUsers = append(guildUsers, guildUser)
	}

	return guildUsers, nil
}

// GET
func (s *SQLiteStorage) GetGuildRoleNameUserNameByGuildId(guildId int) (map[string][]string, error) {
	guildUser, err := s.GetGuildUserByGuildId(guildId)
	if err != nil {
		return nil, err
	}

	var roleUserMap map[string][]string

	for _, guildUser := range guildUser {
		role, err := s.GetGuildRoleById(guildUser.RoleId)
		if err != nil {
			return nil, err
		}

		user, err := s.GetUserById(guildUser.UserId)
		if err != nil {
			return nil, err
		}

		if roleUserMap[role.Name] == nil {
			roleUserMap[role.Name] = []string{user.Name}
		} else {
			roleUserMap[role.Name] = append(roleUserMap[role.Name], user.Name)
		}
	}

	return roleUserMap, nil
}

// DELETE
func (s *SQLiteStorage) DeleteGuildUser(user_id int) error {
	prep, err := s.db.Prepare(`DELETE FROM GuildMember WHERE user_id = ? `)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(user_id); err != nil {
		return err
	}

	return nil
}

// PUT
func (s *SQLiteStorage) UpdateGuildUser(guildUser *ReqGuildUser) error {
	prep, err := s.db.Prepare(`UPDATE GuildMember SET role_id = ?, guild_id = ? WHERE user_id = ?`)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(guildUser.RoleId, guildUser.GuildId, guildUser.UserId); err != nil {
		return err
	}

	return nil
}
