package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RakanMyHusbando/shogun/types"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteStorage(dbFile string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	return &SQLiteStorage{
		db: db,
	}, nil
}

func RunSQLiteStorage(storage *SQLiteStorage, schemaFile string) error {
	// read sql file with create table statements
	byteContent, err := os.ReadFile("schema.sql")
	if err != nil {
		return err
	}
	// execute sql create table statements
	queries := strings.Split(string(byteContent), ";")
	for i := range queries {
		query := strings.TrimSpace(queries[i])
		_, err := storage.db.Exec(query)
		if err != nil {
			return fmt.Errorf("%v [Table: %v]", err.Error(), strings.Split(query, " ")[5])
		}
	}
	return nil
}

type Storage interface {
	Delete(string, string, any) error

	/* ------------------------------ ./user.go ------------------------------ */
	CreateUser(*types.ReqUser) error
	GetUser() ([]*types.ResUser, error)
	GetUserById(int) (*types.ResUser, error)
	GetUserIds() ([]*int, error)
	UpdateUser(*types.ResUser) error

	/* ------------------------------ ./league_of_legends.go ------------------------------ */
	CreateLeagueOfLegends(*types.ReqLeagueOfLegends) error
	GetLeagueOfLegendsById(int) (*types.ResLeagueOfLegends, error)
	GetLeagueOfLegendsWithAccountsById(int) (*types.ResLeagueOfLegends, error)
	UpdateLeagueOfLegends(*types.ReqLeagueOfLegends) error

	/* ------------------------------ ./game_accounts.go ------------------------------ */
	CreateGameAccount(*types.ReqGameAccount) error
	GetGameAccountByUserId(int, string) ([]string, error)
	DeleteGameAccount(*types.ReqGameAccount) error
	UpdateGameAccount(*types.ReqGameAccount, string) error

	/* ------------------------------ ./guild.go ------------------------------ */
	// guild
	CreateGuild(*types.ReqGuild) error
	GetGuild() ([]*types.ResGuild, error)
	GetGuildById(int) (*types.ResGuild, error)
	UpdateGuild(*types.ResGuild) error
	// guild role
	CreateGuildRole(*types.ReqGuildRole) error
	GetGuildRole() ([]*types.ReqGuildRole, error)
	GetGuildRoleById(int) (*types.ReqGuildRole, error)
	UpdateGuildRole(*types.ReqGuildRole) error
	// guild member
	CreateGuildMember(*types.ReqGuildMember) error
	GetGuildMemberByGuildId(int) ([]*types.ReqGuildMember, error)
	GetGuildMemberMapByGuildId(int) (map[string][]string, error)
	UpdateGuildMember(*types.ReqGuildMember) error

	/* ------------------------------ ./team.go ------------------------------ */

}

type SQLiteStorage struct {
	db *sql.DB
}

// DELETE
func (s *SQLiteStorage) Delete(table string, key string, value any) error {
	prep, err := s.db.Prepare(
		fmt.Sprintf(`DELETE FROM %s WHERE %s = ?`, table, key),
	)
	if err != nil {
		return err
	}

	if _, err = prep.Exec(value); err != nil {
		return err
	}

	prep.Close()
	log.Printf("[Storage.main] delete from %s success", table)

	return nil
}
